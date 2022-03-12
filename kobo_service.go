package main

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	pgakobo "github.com/pgaskin/koboutils/v2/kobo"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/marcus-crane/october/pkg/kobo"
	"github.com/marcus-crane/october/pkg/settings"
)

var (
	highlightsEndpoint = "https://readwise.io/api/v2/highlights/"
)

type KoboService struct {
	SelectedKobo kobo.Kobo
	ConnectedDB  *gorm.DB
	Settings     *settings.Settings
	Logger       *zap.SugaredLogger
}

type ReadwiseResponse struct {
	Highlights []kobo.Highlight `json:"highlights"`
}

func NewKoboService(settings *settings.Settings, logger *zap.SugaredLogger) *KoboService {
	return &KoboService{
		Settings: settings,
		Logger:   logger,
	}
}

func (k *KoboService) DetectKobos() []kobo.Kobo {
	var kobos []kobo.Kobo
	connectedKobos, err := pgakobo.Find()
	if err != nil {
		k.Logger.Errorw("Failed to check for Kobos", "error", err)
		panic(err)
	}
	k.Logger.Debugw("Found %d Kobos", len(connectedKobos))
	for _, koboPath := range connectedKobos {
		_, _, deviceId, err := pgakobo.ParseKoboVersion(koboPath)
		k.Logger.Debugw("Found Kobo with Device ID of %s", deviceId)
		if err != nil {
			k.Logger.Errorw("Failed to parse Kobo version", "error", err)
			panic(err)
		}
		device, found := pgakobo.DeviceByID(deviceId)
		if !found {
			fallbackKobo, err := GetKoboFallbackMetadata(deviceId, koboPath)
			if err != nil {
				continue
			}
			k.Logger.Infow(fmt.Sprintf("Found a %s through fallback method", fallbackKobo.Name))
			kobos = append(kobos, fallbackKobo)
			continue
		}
		k.Logger.Infof(fmt.Sprintf("Detected a %s", device.Name()))
		kobos = append(kobos, kobo.Kobo{
			Name:       device.Name(),
			Storage:    device.StorageGB(),
			DisplayPPI: device.DisplayPPI(),
			MntPath:    koboPath,
			DbPath:     fmt.Sprintf("%s/.kobo/KoboReader.sqlite", koboPath),
		})
	}
	return kobos
}

func (k *KoboService) SelectKobo(devicePath string) bool {
	_, _, deviceId, err := pgakobo.ParseKoboVersion(devicePath)
	if err != nil {
		panic(err)
	}
	device, found := pgakobo.DeviceByID(deviceId)
	foundKobo := kobo.Kobo{}
	if !found {
		fallbackKobo, err := GetKoboFallbackMetadata(deviceId, devicePath)
		if err != nil {
			panic("unknown device? unplugged?")
		}
		foundKobo = fallbackKobo
	} else {
		foundKobo = kobo.Kobo{
			Name:       device.Name(),
			Storage:    device.StorageGB(),
			DisplayPPI: device.DisplayPPI(),
			MntPath:    devicePath,
			DbPath:     fmt.Sprintf("%s/.kobo/KoboReader.sqlite", devicePath),
		}
	}
	k.SelectedKobo = foundKobo
	k.Logger.Infow(fmt.Sprintf("User has selected %s", k.SelectedKobo.Name), "kobo", k.SelectedKobo)

	err = k.OpenDBConnection(k.SelectedKobo.DbPath)
	if err != nil {
		k.Logger.Errorw(fmt.Sprintf("Failed to open a connection to %s", k.SelectedKobo.DbPath), "kobo", k.SelectedKobo)
		return false
	}
	k.Logger.Infow(fmt.Sprintf("Successfully opened connection to %s", k.SelectedKobo.DbPath))
	return true
}

func (k *KoboService) GetSelectedKobo() kobo.Kobo {
	return k.SelectedKobo
}

func (k *KoboService) OpenDBConnection(filepath string) error {
	if filepath == "" {
		filepath = k.SelectedKobo.DbPath
	}
	db, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})
	if err != nil {
		k.Logger.Errorw(fmt.Sprintf("Failed to open DB connection to %s", filepath), "error", err)
		panic(err)
	}
	k.ConnectedDB = db
	return nil
}

func (k *KoboService) PromptForLocalDBPath() error {
	var ctx context.Context // noop, this doesn't actually work
	k.Logger.Debugw("Asking user to provide path to local sqlite3 DB")
	selectedFile, err := runtime.OpenFileDialog(ctx, runtime.OpenDialogOptions{
		Title: "Select local Kobo database",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "sqlite (*.sqlite;*.sqlite3)",
				Pattern:     "*.sqlite;*.sqlite3",
			},
		},
	})
	if err != nil {
		return err
	}
	return k.OpenDBConnection(selectedFile)
}

func (k *KoboService) ListDeviceContent() ([]kobo.Content, error) {
	var content []kobo.Content
	k.Logger.Debugw("Retrieving content from device")
	result := k.ConnectedDB.Where(
		&kobo.Content{ContentType: "6", MimeType: "application/x-kobo-epub+zip", VolumeIndex: -1},
	).Order("___PercentRead desc, title asc").Find(&content)
	if result.Error != nil {
		k.Logger.Errorw("Failed to retrieve content from device", "error", result.Error)
		return nil, result.Error
	}
	k.Logger.Debugw(fmt.Sprintf("Successfully retrieved %d pieces of content from device DB", len(content)))
	return content, nil
}

func (k *KoboService) ListDeviceBookmarks() ([]kobo.Bookmark, error) {
	var bookmarks []kobo.Bookmark
	k.Logger.Debugw("Retrieving bookmarks from device")
	result := k.ConnectedDB.Order("VolumeID ASC, ChapterProgress ASC").Find(&bookmarks).Limit(1)
	if result.Error != nil {
		k.Logger.Errorw("Failed to retrieve bookmarks from device", "error", result.Error)
		return nil, result.Error
	}
	k.Logger.Debugw(fmt.Sprintf("Successfully retrieved %d pieces of content from device DB", len(bookmarks)))
	return bookmarks, nil
}

func (k *KoboService) BuildContentIndex(content []kobo.Content) map[string]kobo.Content {
	k.Logger.Debugw("Building an index out of device content")
	contentIndex := make(map[string]kobo.Content)
	for _, item := range content {
		contentIndex[item.ContentID] = item
	}
	k.Logger.Debugw(fmt.Sprintf("Built an index out with %d items", len(contentIndex)))
	return contentIndex
}

func (k *KoboService) CountDeviceBookmarks() int64 {
	var count int64
	result := k.ConnectedDB.Model(&kobo.Bookmark{}).Count(&count)
	if result.Error != nil {
		k.Logger.Errorw("Failed to count bookmarks on device", "error", result.Error)
	}
	return count
}

func (k *KoboService) BuildReadwisePayload() ([]kobo.Highlight, error) {
	content, err := k.ListDeviceContent()
	if err != nil {
		return nil, err
	}
	contentIndex := k.BuildContentIndex(content)
	bookmarks, err := k.ListDeviceBookmarks()
	if err != nil {
		return nil, err
	}
	var highlights []kobo.Highlight
	k.Logger.Infow(fmt.Sprintf("Starting to build Readwise payload out of %d bookmarks", len(bookmarks)))
	for _, entry := range bookmarks {
		source := contentIndex[entry.VolumeID]
		t, err := time.Parse("2006-01-02T15:04:05.000", entry.DateCreated)
		if err != nil {
			k.Logger.Errorw(fmt.Sprintf("Failed to parse timestamp %s from bookmark", entry.DateCreated), "bookmark", entry)
			return nil, err
		}
		createdAt := t.Format("2006-01-02T15:04:05-07:00")
		text := k.NormaliseText(entry.Text)
		if entry.Annotation != "" && text == "" {
			text = "Placeholder for attached annotation"
		}
		if entry.Annotation == "" && text == "" {
			k.Logger.Infow("Found an entry with no annotation of text so skipping to next item", "source", source, "bookmark", entry)
			fmt.Printf("Ignoring entry from %s", source.Title)
			continue
		}
		if source.Title == "" {
			sourceFile, err := url.Parse(entry.VolumeID)
			if err != nil {
				k.Logger.Errorw("No title. Fallback of using filename failed. Not required so will send with no title.", "source", source, "bookmark", entry)
				continue
			}
			filename := path.Base(sourceFile.Path)
			k.Logger.Debugw(fmt.Sprintf("No source title. Constructing title from filename: %s", filename))
			source.Title = strings.TrimSuffix(filename, ".epub")
		}
		highlight := kobo.Highlight{
			Text:          text,
			Title:         source.Title,
			Author:        source.Attribution,
			SourceType:    "OctoberForKobo",
			Category:      "books",
			Note:          entry.Annotation,
			HighlightedAt: createdAt,
		}
		k.Logger.Debugw("Succesfully built highlight", "highlight", highlight)
		highlights = append(highlights, highlight)
	}
	k.Logger.Infow(fmt.Sprintf("Successfully parsed %d highlights", len(highlights)))
	return highlights, nil
}

func (k *KoboService) GetReadwiseToken() string {
	return k.Settings.ReadwiseToken
}
func (k *KoboService) SetReadwiseToken(token string) error {
	k.Settings.ReadwiseToken = token
	if err := k.Settings.Save(); err != nil {
		k.Logger.Errorw("Failed to save Readwise token", "error", err)
		return err
	}
	k.Logger.Infow("Saved Readwise token to storage")
	return nil
}

func (k *KoboService) NormaliseText(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\n", " ")
	return s
}

func (k *KoboService) SendBookmarksToReadwise() (int, error) {
	bookmarks, err := k.BuildReadwisePayload()
	if err != nil {
		return 0, err
	}
	payload := ReadwiseResponse{
		Highlights: bookmarks,
	}
	if err != nil {
		k.Logger.Infow("Failed to save Readwise payload to disc for debugging but we can carry on")
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Token %s", k.Settings.ReadwiseToken)).
		SetHeader("User-Agent", "october/1.0.0 <https://github.com/marcus-crane/october>").
		SetBody(payload).
		Post(highlightsEndpoint)
	if resp.StatusCode() != 200 {
		k.Logger.Errorw("Received a non-200 response from Readwise", "status", resp.StatusCode(), "response", string(resp.Body()))
		return 0, errors.New(fmt.Sprintf("Received a non-200 status code from Readwise: code %d", resp.StatusCode()))
	}
	k.Logger.Infow(fmt.Sprintf("Successfully sent %d bookmarks to Readwise", len(bookmarks)))
	return len(bookmarks), nil
}

func getKoboFallbackSkus() map[string]kobo.Kobo {
	return map[string]kobo.Kobo{
		"00000000-0000-0000-0000-000000000383": {Name: "Kobo Sage", Storage: 32, DisplayPPI: 300},
		"00000000-0000-0000-0000-000000000387": {Name: "Kobo Elipsa", Storage: 32, DisplayPPI: 227},
		"00000000-0000-0000-0000-000000000388": {Name: "Kobo Libra 2", Storage: 32, DisplayPPI: 300},
	}
}

func deviceIdInSkuList(deviceId string) bool {
	for k, _ := range getKoboFallbackSkus() {
		if k == deviceId {
			return true
		}
	}
	return false
}

func (k *KoboService) CheckReadwiseConfig() bool {
	if k.Settings.ReadwiseToken == "" {
		return false
	}
	return true
}

func GetKoboFallbackMetadata(deviceId string, devicePath string) (kobo.Kobo, error) {
	fallbackSkus := getKoboFallbackSkus()
	if !deviceIdInSkuList(deviceId) {
		return kobo.Kobo{}, errors.New("no kobo found with that device id")
	}
	deviceInfo := fallbackSkus[deviceId]
	deviceInfo.MntPath = devicePath
	deviceInfo.DbPath = fmt.Sprintf("%s/.kobo/KoboReader.sqlite", devicePath)
	return deviceInfo, nil
}
