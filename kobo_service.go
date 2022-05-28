package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/pgaskin/koboutils/v2/kobo"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/marcus-crane/october/backend"
	"github.com/marcus-crane/october/pkg/db"
	"github.com/marcus-crane/october/pkg/device"
	"github.com/marcus-crane/october/pkg/logger"
	"github.com/marcus-crane/october/pkg/readwise"
)

type KoboService struct {
	SelectedKobo   device.Kobo
	ConnectedKobos map[string]device.Kobo
	runtimeContext context.Context
	Settings       backend.Settings
}

func NewKoboService(settings backend.Settings) *KoboService {
	return &KoboService{
		Settings:       settings,
		ConnectedKobos: map[string]device.Kobo{},
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
	for _, kb := range kobos {
		logger.Log.Infow("Found Kobo", "device", kb)
		k.ConnectedKobos[kb.MntPath] = kb
	}
	return kobos
}

func (k *KoboService) SelectKobo(devicePath string) error {
	if val, ok := k.ConnectedKobos[devicePath]; ok {
		k.SelectedKobo = val
	} else {
		logger.Log.Info("Trying to access local db")
		k.SelectedKobo = device.Kobo{
			Name:       "Local Database",
			Storage:    0,
			DisplayPPI: 0,
			MntPath:    devicePath,
			DbPath:     devicePath,
		}
	}
	if err := db.OpenConnection(k.SelectedKobo.DbPath); err != nil {
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

func (k *KoboService) FindBookOnDevice(bookID string) (device.Content, error) {
	var content device.Content
	logger.Log.Debugw("Retrieving books that have been uploaded to Readwise previously")
	result := db.Conn.Where(&device.Content{ContentType: "6", VolumeIndex: -1, ContentID: bookID}).Find(&content)
	if result.Error != nil {
		logger.Log.Errorw("Failed to retrieve content from device", "error", result.Error)
		return content, result.Error
	}
	logger.Log.Debugw(fmt.Sprintf("Successfully retrieved %s from device DB", content.Title))
	return content, nil
}

func (k *KoboService) ListDeviceContent() ([]device.Content, error) {
	var content []device.Content
	logger.Log.Debugw("Retrieving content from device")
	result := db.Conn.Where(
		&device.Content{ContentType: "6", VolumeIndex: -1},
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
	logger.Log.Debugw(fmt.Sprintf("Built an index out of %d items", len(contentIndex)))
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
	k.Settings.ReadwiseToken = token
	return k.Settings.Save()
}

func (k *KoboService) GetCoverUploadStatus() bool {
	return k.Settings.UploadCovers
}
func (k *KoboService) SetCoverUploadStatus(enabled bool) error {
	k.Settings.UploadCovers = enabled
	return k.Settings.Save()
}

func (k *KoboService) CheckReadwiseConfig() bool {
	if k.Settings.ReadwiseToken == "" {
		return false
	}
	return true
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
	numUploads, err := readwise.SendBookmarks(payload, k.Settings.ReadwiseToken)
	if err != nil {
		return 0, err
	}
	uploadedBooks, err := readwise.RetrieveUploadedBooks(k.Settings.ReadwiseToken)
	if err != nil {
		return numUploads, fmt.Errorf(fmt.Sprintf("Successfully uploaded %d bookmarks but failed to upload covers", numUploads))
	}
	if k.Settings.UploadCovers {
		for _, book := range uploadedBooks.Results {
			// We don't want to overwrite user uploaded covers or covers already present
			if !strings.Contains(book.CoverURL, "uploaded_book_covers") {
				bookDetail, err := k.FindBookOnDevice(book.SourceURL)
				logger.Log.Info(bookDetail)
				if err != nil {
					logger.Log.Error(fmt.Sprintf("Failed to retrieve %s", book.SourceURL))
					continue
				}
				coverID := kobo.ContentIDToImageID(book.SourceURL)
				coverPath := kobo.CoverTypeLibFull.GeneratePath(false, coverID)
				absCoverPath := path.Join(k.SelectedKobo.MntPath, "/", coverPath)
				coverBytes, err := ioutil.ReadFile(absCoverPath)
				if err != nil {
					logger.Log.Error(fmt.Sprintf("Failed to load cover for %s. Location %s", book.SourceURL, absCoverPath))
				}
				var base64Encoding string
				mimeType := http.DetectContentType(coverBytes)
				logger.Log.Info(mimeType)
				switch mimeType {
				case "image/jpeg":
					base64Encoding += "data:image/jpeg;base64,"
				case "image/png":
					base64Encoding += "data:image/png;base64,"
				}
				base64Encoding += base64.StdEncoding.EncodeToString(coverBytes)
				err = readwise.UploadCover(base64Encoding, book.ID, k.Settings.ReadwiseToken)
				if err != nil {
					logger.Log.Error(fmt.Sprintf("Failed to upload cover for %s", book.SourceURL))
				}
				logger.Log.Debug(fmt.Sprintf("Successfully uploaded cover for %s", book.SourceURL))
			}
		}
	}
	return numUploads, nil
}
