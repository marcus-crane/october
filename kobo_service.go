package main

import (
	"context"
	"fmt"
  "log"

  "github.com/pgaskin/koboutils/v2/kobo"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type KoboService struct {
	ctx          context.Context
	SelectedKobo Kobo
	ConnectedDB  *gorm.DB
	Content      []Content
	Bookmarks    []Bookmark
}

type Kobo struct {
	Name       string `json:"name"`
	Storage    int    `json:"storage"`
	DisplayPPI int    `json:"display_ppi"`
	MntPath    string `json:"mnt_path"`
	DbPath     string `json:"db_path"`
}

type Content struct {
	ContentID              string `gorm:"column:ContentID"`
	ContentType            string `gorm:"column:ContentType"`
	MimeType               string `gorm:"column:MimeType"`
	BookID                 string
	BookTitle              string `gorm:"column:BookTitle"`
	ImageId                string
	Title                  string
	Attribution            string
	Description            string
	DateCreated            string
	ShortCoverKey          string
	AdobeLocation          string `gorm:"column:adobe_location"`
	Publisher              string
	IsEncrypted            bool
	DateLastRead           string
	FirstTimeReading       bool
	ChapterIDBookmarked    string
	ParagraphBookmarked    int
	BookmarkWordOffset     int
	NumShortcovers         int
	VolumeIndex            int `gorm:"column:VolumeIndex"`
	NumPages               int `gorm:"column:___NumPages"`
	ReadStatus             int
	SyncTime               string `gorm:"column:___SyncTime"`
	UserID                 string `gorm:"column:___UserID"`
	PublicationId          string
	FileOffset             int    `gorm:"column:___FileOffset"`
	FileSize               int    `gorm:"column:___FileSize"`
	PercentRead            string `gorm:"column:___PercentRead"`
	ExpirationStatus       int    `gorm:"column:___ExpirationStatus"`
	CurrentChapterProgress float32
}

type Bookmark struct {
	BookmarkID               string `gorm:"column:BookmarkID"`
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
	ChapterProgress          string
	Hidden                   string
	Version                  string
	DateModified             string
	Creator                  string
	UUID                     string
	UserID                   string
	SyncTime                 string
	Published                string
	ContextString            string
}

func (Content) TableName() string {
	return "Content"
}

func (Bookmark) TableName() string {
	return "Bookmark"
}

func NewKoboService(ctx context.Context) *KoboService {
	return &KoboService{
		ctx: ctx,
	}
}
func (k *KoboService) DetectKobos() []Kobo {
	var kobos []Kobo
	connectedKobos, err := kobo.Find()
	if err != nil {
		panic(err)
	}
	for _, koboPath := range connectedKobos {
    log.Print(koboPath)
		_, _, deviceId, err := kobo.ParseKoboVersion(koboPath)
		if err != nil {
			panic(err)
		}
		device, found := kobo.DeviceByID(deviceId)
		if !found {
			continue
		}
		kobos = append(kobos, Kobo{
			Name:       device.Name(),
			Storage:    device.StorageGB(),
			DisplayPPI: device.DisplayPPI(),
			MntPath:    koboPath,
			DbPath:     fmt.Sprintf("%s/.kobo/KoboReader.sqlite", koboPath),
		})
	}
  log.Print(kobos)
	return kobos
}

func (k *KoboService) SelectKobo(devicePath string) bool {
	_, _, deviceId, err := kobo.ParseKoboVersion(devicePath)
	if err != nil {
		panic(err)
	}
	device, found := kobo.DeviceByID(deviceId)
	if !found {
		panic("device detached?")
	}
	k.SelectedKobo = Kobo{
		Name:       device.Name(),
		Storage:    device.StorageGB(),
		DisplayPPI: device.DisplayPPI(),
		MntPath:    devicePath,
		DbPath:     fmt.Sprintf("%s/.kobo/KoboReader.sqlite", devicePath),
	}
  err = k.OpenDBConnection(k.SelectedKobo.DbPath)
  if err != nil {
    panic(err)
    return false
  }
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
    panic(err)
	}
	k.ConnectedDB = db
	return nil
}

func (k *KoboService) PromptForLocalDBPath() error {
	selectedFile, err := runtime.OpenFileDialog(k.ctx, runtime.OpenDialogOptions{
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

func (k *KoboService) ListDeviceContent() error {
	var content []Content
	result := k.ConnectedDB.Where(
		&Content{ContentType: "6", MimeType: "application/x-kobo-epub+zip", VolumeIndex: -1},
	).Order("___PercentRead desc, title asc").Find(&content)
	if result.Error != nil {
		return result.Error
	}
	k.Content = content
	return nil
}

func (k *KoboService) ListDeviceBookmarks() []Bookmark {
	var bookmarks []Bookmark
	result := k.ConnectedDB.Find(&bookmarks)
	if result.Error != nil {
    log.Print(result.Error)
	}
	k.Bookmarks = bookmarks
	return bookmarks
}
