package main

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/marcus-crane/october/pkg/logger"
	"github.com/pgaskin/koboutils/v2/kobo"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/marcus-crane/october/pkg/device"
	"github.com/marcus-crane/october/pkg/settings"
)

type KoboService struct {
	SelectedKobo device.Kobo
	ConnectedDB  *gorm.DB
	Settings     *settings.Settings
}

func NewKoboService(settings *settings.Settings) *KoboService {
	return &KoboService{
		Settings: settings,
	}
}

func (k *KoboService) DetectKobos() []device.Kobo {
	var kobos []device.Kobo
	connectedKobos, err := kobo.Find()
	if err != nil {
		logger.Log.Errorw("Failed to check for Kobos", "error", err)
		panic(err)
	}
}

func (k *KoboService) SelectKobo(devicePath string) bool {
	_, _, deviceId, err := kobo.ParseKoboVersion(devicePath)
	if err != nil {
		panic(err)
	}
	deviceFound, found := kobo.DeviceByID(deviceId)
	foundKobo := device.Kobo{}
	if !found {
		fallbackKobo, err := GetKoboFallbackMetadata(deviceId, devicePath)
		if err != nil {
			panic("unknown device? unplugged?")
		}
		foundKobo = fallbackKobo
	} else {
		foundKobo = device.Kobo{
			Name:       deviceFound.Name(),
			Storage:    deviceFound.StorageGB(),
			DisplayPPI: deviceFound.DisplayPPI(),
			MntPath:    devicePath,
			DbPath:     fmt.Sprintf("%s/.device/KoboReader.sqlite", devicePath),
		}
	}
	k.SelectedKobo = foundKobo
	logger.Log.Infow(fmt.Sprintf("User has selected %s", k.SelectedKobo.Name), "device", k.SelectedKobo)

	err = k.OpenDBConnection(k.SelectedKobo.DbPath)
	if err != nil {
		logger.Log.Errorw(fmt.Sprintf("Failed to open a connection to %s", k.SelectedKobo.DbPath), "device", k.SelectedKobo)
		return false
	}
	logger.Log.Infow(fmt.Sprintf("Successfully opened connection to %s", k.SelectedKobo.DbPath))
	return true
}

func (k *KoboService) GetSelectedKobo() device.Kobo {
	return k.SelectedKobo
}

func (k *KoboService) PromptForLocalDBPath() error {
	var ctx context.Context // noop, this doesn't actually work
	logger.Log.Debugw("Asking user to provide path to local sqlite3 DB")
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

func (k *KoboService) ListDeviceContent() ([]device.Content, error) {
	var content []device.Content
	logger.Log.Debugw("Retrieving content from device")
	result := k.ConnectedDB.Where(
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
	result := k.ConnectedDB.Order("VolumeID ASC, ChapterProgress ASC").Find(&bookmarks).Limit(1)
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
	result := k.ConnectedDB.Model(&device.Bookmark{}).Count(&count)
	if result.Error != nil {
		logger.Log.Errorw("Failed to count bookmarks on device", "error", result.Error)
	}
	return count
}

func (k *KoboService) BuildReadwisePayload() ([]device.Highlight, error) {
	content, err := k.ListDeviceContent()
	if err != nil {
		return nil, err
	}
	contentIndex := k.BuildContentIndex(content)
	bookmarks, err := k.ListDeviceBookmarks()
	if err != nil {
		return nil, err
	}
	var highlights []device.Highlight
	logger.Log.Infow(fmt.Sprintf("Starting to build Readwise payload out of %d bookmarks", len(bookmarks)))
	for _, entry := range bookmarks {
		source := contentIndex[entry.VolumeID]
		t, err := time.Parse("2006-01-02T15:04:05.000", entry.DateCreated)
		if err != nil {
			logger.Log.Errorw(fmt.Sprintf("Failed to parse timestamp %s from bookmark", entry.DateCreated), "bookmark", entry)
			return nil, err
		}
		createdAt := t.Format("2006-01-02T15:04:05-07:00")
		text := k.NormaliseText(entry.Text)
		if entry.Annotation != "" && text == "" {
			text = "Placeholder for attached annotation"
		}
		if entry.Annotation == "" && text == "" {
			logger.Log.Infow("Found an entry with no annotation of text so skipping to next item", "source", source, "bookmark", entry)
			fmt.Printf("Ignoring entry from %s", source.Title)
			continue
		}
		if source.Title == "" {
			sourceFile, err := url.Parse(entry.VolumeID)
			if err != nil {
				logger.Log.Errorw("No title. Fallback of using filename failed. Not required so will send with no title.", "source", source, "bookmark", entry)
				continue
			}
			filename := path.Base(sourceFile.Path)
			logger.Log.Debugw(fmt.Sprintf("No source title. Constructing title from filename: %s", filename))
			source.Title = strings.TrimSuffix(filename, ".epub")
		}
		highlight := device.Highlight{
			Text:          text,
			Title:         source.Title,
			Author:        source.Attribution,
			SourceType:    "OctoberForKobo",
			Category:      "books",
			Note:          entry.Annotation,
			HighlightedAt: createdAt,
		}
		logger.Log.Debugw("Succesfully built highlight", "highlight", highlight)
		highlights = append(highlights, highlight)
	}
	logger.Log.Infow(fmt.Sprintf("Successfully parsed %d highlights", len(highlights)))
	return highlights, nil
}
