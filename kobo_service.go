package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"strings"
	"time"

	"github.com/pgaskin/koboutils/v2/kobo"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	highlightsEndpoint = "https://readwise.io/api/v2/highlights/"
)

type KoboService struct {
	SelectedKobo Kobo
	ConnectedDB  *gorm.DB
	Settings     *Settings
	Logger       *zap.SugaredLogger
}

type ReadwiseResponse struct {
	Highlights []Highlight `json:"highlights"`
}

type Highlight struct {
	Text          string `json:"text"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	SourceType    string `json:"source_type"`
	Category      string `json:"category"`
	Note          string `json:"note,omitempty"`
	HighlightedAt string `json:"highlighted_at"`
}

type Kobo struct {
	Name       string `json:"name"`
	Storage    int    `json:"storage"`
	DisplayPPI int    `json:"display_ppi"`
	MntPath    string `json:"mnt_path"`
	DbPath     string `json:"db_path"`
}

type Content struct {
	ContentID               string `gorm:"column:ContentID"`
	ContentType             string `gorm:"column:ContentType"`
	MimeType                string `gorm:"column:MimeType"`
	BookID                  string
	BookTitle               string `gorm:"column:BookTitle"`
	ImageId                 string
	Title                   string `gorm:"column:Title"`
	Attribution             string `gorm:"column:Attribution"`
	Description             string `gorm:"column:Description"`
	DateCreated             string `gorm:"column:DateCreated"`
	ShortCoverKey           string
	AdobeLocation           string `gorm:"column:adobe_location"`
	Publisher               string
	IsEncrypted             bool
	DateLastRead            string
	FirstTimeReading        bool
	ChapterIDBookmarked     string
	ParagraphBookmarked     int
	BookmarkWordOffset      int
	NumShortcovers          int
	VolumeIndex             int `gorm:"column:VolumeIndex"`
	NumPages                int `gorm:"column:___NumPages"`
	ReadStatus              int
	SyncTime                string `gorm:"column:___SyncTime"`
	UserID                  string `gorm:"column:___UserID"`
	PublicationId           string
	FileOffset              int    `gorm:"column:___FileOffset"`
	FileSize                int    `gorm:"column:___FileSize"`
	PercentRead             string `gorm:"column:___PercentRead"`
	ExpirationStatus        int    `gorm:"column:___ExpirationStatus"`
	FavouritesIndex         int
	Accessibility           int
	ContentURL              string
	Language                string
	BookshelfTags           string
	IsDownloaded            bool
	FeedbackType            int
	AverageRating           float64
	Depth                   int
	PageProgressDirection   string
	InWishlist              string
	ISBN                    int64
	WishlistedDate          string
	FeedbackTypeSynced      bool
	IsSocialEnabled         bool
	EpubType                string
	Monetization            string
	ExternalId              string
	Series                  string
	SeriesNumber            string
	Subtitle                string
	WordCount               string
	Fallback                string
	RestOfBookEstimate      string
	CurrentChapterEstimate  string
	CurrentChapterProgress  float32
	PocketStatus            string
	UnsyncedPocketChanges   string
	ImageUrl                string
	DateAdded               string
	WorkId                  string
	Properties              string
	RenditionSpread         string
	RatingCount             string
	ReviewsSyncDate         string
	MediaOverlay            string
	RedirectPreviewUrl      bool
	PreviewFileSize         int
	EntitlementId           string
	CrossRevisionId         string
	DownloadUrl             bool
	ReadStateSynced         bool
	TimesStartedReading     int
	TimeSpentReading        int
	LastTimeStartedReading  string
	LastTimeFinishedReading string
	ApplicableSubscriptions string
	ExternalIds             string
	PurchaseRevisionId      string
	SeriesID                string
	SeriesNumberFloat       float64
	AdobeLoanExpiration     string
	HideFromHomePage        bool
	IsInternetArchive       bool
	TitleKana               string `gorm:"column:titleKana"`
	SubtitleKana            string `gorm:"column:subtitleKana"`
	SeriesKana              string `gorm:"column:seriesKana"`
	AttributionKana         string `gorm:"column:attributionKana"`
	PublisherKana           string `gorm:"column:publisherKana"`
	IsPurchaseable          bool
	IsSupported             bool
	AnnotationsSyncToken    string
	DateModified            string
}

type Bookmark struct {
	BookmarkID               string
	VolumeID                 string
	ContentID                string
	StartContainerPath       string
	StartContainerChild      string
	StartContainerChildIndex string
	StartOffset              string
	EndContainerPath         string
	EndContainerChildIndex   string
	EndOffset                string
	Text                     string
	Annotation               string
	ExtraAnnotationData      string
	DateCreated              string
	ChapterProgress          float64
	Hidden                   string
	Version                  string
	DateModified             string
	Creator                  string
	UUID                     string
	UserID                   string
	SyncTime                 string
	Published                string
	ContextString            string
	Type                     string
}

func (Content) TableName() string {
	return "Content"
}

func (Bookmark) TableName() string {
	return "Bookmark"
}

func NewKoboService(settings *Settings, logger *zap.SugaredLogger) *KoboService {
	return &KoboService{
		Settings: settings,
		Logger:   logger,
	}
}

func (k *KoboService) DetectKobos() []Kobo {
	var kobos []Kobo
	connectedKobos, err := kobo.Find()
	if err != nil {
		k.Logger.Errorw("Failed to check for Kobos", "error", err)
		panic(err)
	}
	k.Logger.Debugw("Found %d Kobos", len(connectedKobos))
	for _, koboPath := range connectedKobos {
		_, _, deviceId, err := kobo.ParseKoboVersion(koboPath)
		k.Logger.Debugw("Found Kobo with Device ID of %s", deviceId)
		if err != nil {
			k.Logger.Errorw("Failed to parse Kobo version", "error", err)
			panic(err)
		}
		device, found := kobo.DeviceByID(deviceId)
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
		kobos = append(kobos, Kobo{
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
	_, _, deviceId, err := kobo.ParseKoboVersion(devicePath)
	if err != nil {
		panic(err)
	}
	device, found := kobo.DeviceByID(deviceId)
	foundKobo := Kobo{}
	if !found {
		fallbackKobo, err := GetKoboFallbackMetadata(deviceId, devicePath)
		if err != nil {
			panic("unknown device? unplugged?")
		}
		foundKobo = fallbackKobo
	} else {
		foundKobo = Kobo{
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

func (k *KoboService) GetSelectedKobo() Kobo {
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

func (k *KoboService) ListDeviceContent() ([]Content, error) {
	var content []Content
	k.Logger.Debugw("Retrieving content from device")
	result := k.ConnectedDB.Where(
		&Content{ContentType: "6", MimeType: "application/x-kobo-epub+zip", VolumeIndex: -1},
	).Order("___PercentRead desc, title asc").Find(&content)
	if result.Error != nil {
		k.Logger.Errorw("Failed to retrieve content from device", "error", result.Error)
		return nil, result.Error
	}
	k.Logger.Debugw(fmt.Sprintf("Successfully retrieved %d pieces of content from device DB", len(content)))
	return content, nil
}

func (k *KoboService) ListDeviceBookmarks() ([]Bookmark, error) {
	var bookmarks []Bookmark
	k.Logger.Debugw("Retrieving bookmarks from device")
	result := k.ConnectedDB.Order("VolumeID ASC, ChapterProgress ASC").Find(&bookmarks).Limit(1)
	if result.Error != nil {
		k.Logger.Errorw("Failed to retrieve bookmarks from device", "error", result.Error)
		return nil, result.Error
	}
	k.Logger.Debugw(fmt.Sprintf("Successfully retrieved %d pieces of content from device DB", len(bookmarks)))
	return bookmarks, nil
}

func (k *KoboService) BuildContentIndex(content []Content) map[string]Content {
	k.Logger.Debugw("Building an index out of device content")
	contentIndex := make(map[string]Content)
	for _, item := range content {
		contentIndex[item.ContentID] = item
	}
	k.Logger.Debugw(fmt.Sprintf("Built an index out with %d items", len(contentIndex)))
	return contentIndex
}

func (k *KoboService) CountDeviceBookmarks() int64 {
	var count int64
	result := k.ConnectedDB.Model(&Bookmark{}).Count(&count)
	if result.Error != nil {
		k.Logger.Errorw("Failed to count bookmarks on device", "error", result.Error)
	}
	return count
}

func (k *KoboService) BuildReadwisePayload() ([]Highlight, error) {
	content, err := k.ListDeviceContent()
	if err != nil {
		return nil, err
	}
	contentIndex := k.BuildContentIndex(content)
	bookmarks, err := k.ListDeviceBookmarks()
	if err != nil {
		return nil, err
	}
	var highlights []Highlight
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
			k.Logger.Infow("Bookmark has no source title so skipping to next item", "source", source, "bookmark", entry)
			fmt.Println("Found no source for ", entry.VolumeID)
			continue
		}
		highlight := Highlight{
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
	if err := k.Settings.save(); err != nil {
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

func getKoboFallbackSkus() map[string]Kobo {
	return map[string]Kobo{
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

func GetKoboFallbackMetadata(deviceId string, devicePath string) (Kobo, error) {
	fallbackSkus := getKoboFallbackSkus()
	if !deviceIdInSkuList(deviceId) {
		return Kobo{}, errors.New("no kobo found with that device id")
	}
	deviceInfo := fallbackSkus[deviceId]
	deviceInfo.MntPath = devicePath
	deviceInfo.DbPath = fmt.Sprintf("%s/.kobo/KoboReader.sqlite", devicePath)
	return deviceInfo, nil
}
