package backend

import (
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/pgaskin/koboutils/v2/kobo"
)

type Kobo struct {
	Name       string `json:"name"`
	Storage    int    `json:"storage"`
	DisplayPPI int    `json:"display_ppi"`
	MntPath    string `json:"mnt_path"`
	DbPath     string `json:"db_path"`
}

type Content struct {
	ContentID               string `gorm:"column:ContentID" json:"content_id"`
	ContentType             string `gorm:"column:ContentType" json:"content_type"`
	MimeType                string `gorm:"column:MimeType" json:"mime_type"`
	BookID                  string `json:"book_id"`
	BookTitle               string `gorm:"column:BookTitle" json:"book_title"`
	ImageId                 string `json:"image_id"`
	Title                   string `gorm:"column:Title" json:"title"`
	Attribution             string `gorm:"column:Attribution" json:"attribution"`
	Description             string `gorm:"column:Description" json:"description"`
	DateCreated             string `gorm:"column:DateCreated" json:"date_created"`
	ShortCoverKey           string
	AdobeLocation           string `gorm:"column:adobe_location"`
	Publisher               string
	IsEncrypted             bool
	DateLastRead            string `json:"date_last_read"`
	FirstTimeReading        bool
	ChapterIDBookmarked     string
	ParagraphBookmarked     int
	BookmarkWordOffset      int
	NumShortcovers          int
	VolumeIndex             int `gorm:"column:VolumeIndex"`
	NumPages                int `gorm:"column:___NumPages" json:"num_pages"`
	ReadStatus              int
	SyncTime                string `gorm:"column:___SyncTime"`
	UserID                  string `gorm:"column:___UserID"`
	PublicationId           string
	FileOffset              int    `gorm:"column:___FileOffset"`
	FileSize                int    `gorm:"column:___FileSize"`
	PercentRead             string `gorm:"column:___PercentRead" json:"percent_read"`
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
	ISBN                    string
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
	BookmarkID               string `gorm:"BookmarkID" json:"bookmark_id"`
	VolumeID                 string `gorm:"column:VolumeID" json:"volume_id"`
	ContentID                string `json:"content_id"`
	StartContainerPath       string
	StartContainerChild      string
	StartContainerChildIndex string
	StartOffset              string
	EndContainerPath         string
	EndContainerChildIndex   string
	EndOffset                string
	Text                     string `gorm:"Text" json:"text"`
	Annotation               string `gorm:"Annotation" json:"annotation"`
	ExtraAnnotationData      string `json:"extra_annotation_data"`
	DateCreated              string `json:"date_created"`
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

func GetKoboMetadata(detectedPaths []string) []Kobo {
	var kobos []Kobo
	for _, path := range detectedPaths {
		_, _, deviceId, err := kobo.ParseKoboVersion(path)
		if err != nil {
			log.WithField("kobo_path", path).WithError(err).Error("Failed to parse Kobo version")
		}
		log.WithField("device_id", deviceId).Info("Found an attached device")
		device, found := kobo.DeviceByID(deviceId)
		if !found {
			log.WithField("device_id", deviceId).Warn("Found a device that isn't officially supported but will likely still operate just fine")
			// We can handle unsupported Kobos in future but at present, there are none
			kobos = append(kobos, Kobo{
				MntPath: path,
				DbPath:  fmt.Sprintf("%s/.kobo/KoboReader.sqlite", path),
			})
		} else {
			kobos = append(kobos, Kobo{
				Name:       device.Name(),
				Storage:    device.StorageGB(),
				DisplayPPI: device.DisplayPPI(),
				MntPath:    path,
				DbPath:     fmt.Sprintf("%s/.kobo/KoboReader.sqlite", path),
			})
		}
	}
	return kobos
}

func (k *Kobo) ListBooksOnDevice() ([]Content, error) {
	var content []Content
	result := Conn.Where(
		&Content{ContentType: "6", VolumeIndex: -1, MimeType: "application/x-kobo-epub+zip"},
	).Order("DateLastRead desc, title asc").Find(&content)
	if result.Error != nil {
		return nil, result.Error
	}
	return content, nil
}

func (k *Kobo) ListBookmarksByID(contentID string) ([]Bookmark, error) {
	var bookmark []Bookmark
	result := Conn.Where(
		&Bookmark{VolumeID: contentID},
	).Find(&bookmark)
	if result.Error != nil {
		log.WithError(result.Error).WithField("content_id", contentID).Error("Encountered an error while trying to list bookmarks by ID")
		return nil, result.Error
	}
	return bookmark, nil
}

func (k *Kobo) FindBookOnDevice(bookID string) (Content, error) {
	var content Content
	log.WithField("book_id", bookID).Debug("Retrieving a book that has been uploaded to Readwise previously")
	result := Conn.Where(&Content{ContentType: "6", VolumeIndex: -1, ContentID: bookID}).Find(&content)
	if result.Error != nil {
		log.WithError(result.Error).WithField("book_id", bookID).Error("Failed to retrieve content from device")
		return content, result.Error
	}
	log.WithField("title", content.Title).Debug("Successfully retrieved content from device DB")
	return content, nil
}

func (k *Kobo) ListDeviceContent() ([]Content, error) {
	var content []Content
	log.Debug("Retrieving content list from device")
	result := Conn.Where(
		&Content{ContentType: "6", VolumeIndex: -1},
	).Order("___PercentRead desc, title asc").Find(&content)
	if result.Error != nil {
		log.WithError(result.Error).Error("Failed to retrieve content from device")
		return nil, result.Error
	}
	log.WithField("content_count", len(content)).Debug("Successfully retrieved device content")
	return content, nil
}

func (k *Kobo) ListDeviceBookmarks() ([]Bookmark, error) {
	var bookmarks []Bookmark
	log.Debug("Retrieving bookmarks from device")
	result := Conn.Order("VolumeID ASC, ChapterProgress ASC").Find(&bookmarks).Limit(1)
	if result.Error != nil {
		log.WithError(result.Error).Error("Failed to retrieve bookmarks from device")
		return nil, result.Error
	}
	log.WithField("bookmark_count", len(bookmarks)).Debug("Successfully retrieved device bookmarks")
	return bookmarks, nil
}

func (k *Kobo) BuildContentIndex(content []Content) map[string]Content {
	log.Debug("Building an index out of device content")
	contentIndex := make(map[string]Content)
	for _, item := range content {
		contentIndex[item.ContentID] = item
	}
	log.WithField("index_count", len(contentIndex)).Debug("Built content index")
	return contentIndex
}

func (k *Kobo) CountDeviceBookmarks() int64 {
	var count int64
	result := Conn.Model(&Bookmark{}).Count(&count)
	if result.Error != nil {
		log.WithError(result.Error).Error("Failed to count bookmarks on device")
	}
	return count
}
