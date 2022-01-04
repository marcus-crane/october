package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-resty/resty/v2"
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
	APIKey       string
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

func NewKoboService(apiKey string) *KoboService {
	return &KoboService{
		APIKey: apiKey,
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
	var ctx context.Context // noop, this doesn't actually work
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
	result := k.ConnectedDB.Where(
		&Content{ContentType: "6", MimeType: "application/x-kobo-epub+zip", VolumeIndex: -1},
	).Order("___PercentRead desc, title asc").Find(&content)
	fmt.Println(content)
	if result.Error != nil {
		return nil, result.Error
	}
	return content, nil
}

func (k *KoboService) ListDeviceBookmarks() ([]Bookmark, error) {
	var bookmarks []Bookmark
	result := k.ConnectedDB.Order("VolumeID ASC, ChapterProgress ASC").Find(&bookmarks).Limit(1)
	if result.Error != nil {
		return nil, result.Error
	}
	return bookmarks, nil
}

func (k *KoboService) BuildContentIndex(content []Content) map[string]Content {
	contentIndex := make(map[string]Content)
	for _, item := range content {
		contentIndex[item.ContentID] = item
	}
	return contentIndex
}

func (k *KoboService) CountDeviceBookmarks() int64 {
	var count int64
	result := k.ConnectedDB.Model(&Bookmark{}).Count(&count)
	if result.Error != nil {
		log.Print(result.Error)
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
	spew.Dump(bookmarks)
	spew.Dump(contentIndex)
	if err != nil {
		return nil, err
	}
	var highlights []Highlight
	for _, entry := range bookmarks {
		source := contentIndex[entry.VolumeID]
		t, err := time.Parse("2006-01-02T15:04:05.000", entry.DateCreated)
		if err != nil {
			return nil, err
		}
		createdAt := t.Format("2006-01-02T15:04:05-07:00")
		text := k.NormaliseText(entry.Text)
		if entry.Annotation != "" && text == "" {
			text = "Placeholder for attached annotation"
		}
		if entry.Annotation == "" && text == "" {
			fmt.Printf("Ignoring entry from %s", source.Title)
			continue
		}
		if source.Title == "" {
			fmt.Println("Found no source for ", entry.VolumeID)
			continue
		}
		highlight := Highlight{
			Text:          text,
			Title:         source.Title,
			Author:        source.Attribution,
			SourceType:    "Kobo",
			Category:      "books",
			Note:          entry.Annotation,
			HighlightedAt: createdAt,
		}
		highlights = append(highlights, highlight)
	}
	return highlights, nil
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
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Token %s", k.APIKey)).
		SetHeader("User-Agent", "october/1.0.0 <https://github.com/marcus-crane/october>").
		SetBody(payload).
		Post(highlightsEndpoint)
	fmt.Println(string(resp.Body()))
	if resp.StatusCode() != 200 {
		return 0, errors.New(fmt.Sprintf("Received a non-200 status code from Readwise: code %d", resp.StatusCode()))
	}
	return len(bookmarks), nil
}
