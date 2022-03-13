package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/pgaskin/koboutils/v2/kobo"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/marcus-crane/october/pkg/db"
	"github.com/marcus-crane/october/pkg/device"
	"github.com/marcus-crane/october/pkg/logger"
	"github.com/marcus-crane/october/pkg/readwise"
	"github.com/marcus-crane/october/pkg/settings"
)

type KoboService struct {
	SelectedKobo   device.Kobo
	ConnectedKobos map[string]device.Kobo
	runtimeContext context.Context
	Settings       *settings.Settings
}

func NewKoboService(settings *settings.Settings) *KoboService {
	return &KoboService{
		Settings: settings,
	}
}

func (k *KoboService) SetContext(ctx context.Context) {
	k.runtimeContext = ctx
}

func (k *KoboService) DetectKobos() []device.Kobo {
	connectedKobos, err := kobo.Find()
	if err != nil {
		logger.Log.Errorw("Failed to check for Kobos", "error", err)
		panic(err)
	}
	kobos := device.GetKoboMetadata(connectedKobos)
	for _, kobo := range kobos {
		k.ConnectedKobos[kobo.MntPath] = kobo
	}
	return kobos
}

func (k *KoboService) SelectKobo(devicePath string) error {
	// We assume this matches a read device, instead of eg; a local sqlite db
	if strings.Contains(devicePath, ".kobo/KoboReader.sqlite") {
		k.SelectedKobo = k.ConnectedKobos[devicePath]
	} else {
		k.SelectedKobo = device.Kobo{
			Name:       "Local Database",
			Storage:    0,
			DisplayPPI: 0,
			MntPath:    devicePath,
			DbPath:     devicePath,
		}
	}
	if err := db.OpenConnection(devicePath); err != nil {
		logger.Log.Errorw("Failed to open DB connection", "error", err)
		return err
	}
	return nil
}

func (k *KoboService) GetSelectedKobo() device.Kobo {
	return k.SelectedKobo
}

func (k *KoboService) PromptForLocalDBPath() error {
	logger.Log.Debugw("Asking user to provide path to local sqlite3 DB")
	selectedFile, err := runtime.OpenFileDialog(k.runtimeContext, runtime.OpenDialogOptions{
		Title: "Select local Kobo database",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "sqlite (*.sqlite;*.sqlite3)",
				Pattern:     "*.sqlite;*.sqlite3",
			},
		},
	})
	if err != nil {
		logger.Log.Errorw("Failed to load local Kobo database", "error", err)
		return err
	}
	logger.Log.Info(fmt.Sprintf("Saw db path: %s", selectedFile))
	return k.SelectKobo(selectedFile)
}

func (k *KoboService) ListDeviceContent() ([]device.Content, error) {
	var content []device.Content
	logger.Log.Debugw("Retrieving content from device")
	result := db.Conn.Where(
		&device.Content{ContentType: "6", MimeType: "application/x-device-epub+zip", VolumeIndex: -1},
	).Order("___PercentRead desc, title asc").Find(&content)
	if result.Error != nil {
		logger.Log.Errorw("Failed to retrieve content from device", "error", result.Error)
		return nil, result.Error
	}
	logger.Log.Debugw(fmt.Sprintf("Successfully retrieved %d pieces of content from device DB", len(content)))
	return content, nil
}

func (k *KoboService) ListDeviceBookmarks() ([]device.Bookmark, error) {
	var bookmarks []device.Bookmark
	logger.Log.Debugw("Retrieving bookmarks from device")
	result := db.Conn.Order("VolumeID ASC, ChapterProgress ASC").Find(&bookmarks).Limit(1)
	if result.Error != nil {
		logger.Log.Errorw("Failed to retrieve bookmarks from device", "error", result.Error)
		return nil, result.Error
	}
	logger.Log.Debugw(fmt.Sprintf("Successfully retrieved %d pieces of content from device DB", len(bookmarks)))
	return bookmarks, nil
}

func (k *KoboService) BuildContentIndex(content []device.Content) map[string]device.Content {
	logger.Log.Debugw("Building an index out of device content")
	contentIndex := make(map[string]device.Content)
	for _, item := range content {
		contentIndex[item.ContentID] = item
	}
	logger.Log.Debugw(fmt.Sprintf("Built an index out with %d items", len(contentIndex)))
	return contentIndex
}

func (k *KoboService) CountDeviceBookmarks() int64 {
	var count int64
	result := db.Conn.Model(&device.Bookmark{}).Count(&count)
	if result.Error != nil {
		logger.Log.Errorw("Failed to count bookmarks on device", "error", result.Error)
	}
	return count
}

func (k *KoboService) CheckTokenValidity() error {
	if !k.CheckReadwiseConfig() {
		return fmt.Errorf("readwise token is empty")
	}
	return readwise.CheckTokenValidity(k.Settings.ReadwiseToken)
}

func (k *KoboService) GetReadwiseToken() string {
	return k.Settings.ReadwiseToken
}
func (k *KoboService) SetReadwiseToken(token string) error {
	return k.Settings.SetReadwiseToken(token)
}

func (k *KoboService) CheckReadwiseConfig() bool {
	return k.Settings.ReadwiseTokenExists()
}

func (k *KoboService) ForwardToReadwise() (int, error) {
	content, err := k.ListDeviceContent()
	if err != nil {
		return 0, err
	}
	contentIndex := k.BuildContentIndex(content)
	bookmarks, err := k.ListDeviceBookmarks()
	if err != nil {
		return 0, err
	}
	payload, err := readwise.BuildPayload(bookmarks, contentIndex)
	if err != nil {
		return 0, err
	}
	return readwise.SendBookmarks(payload, k.Settings.ReadwiseToken)
}
