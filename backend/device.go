package backend

import (
	"fmt"

	"github.com/pgaskin/koboutils/v2/kobo"
	"github.com/rs/zerolog/log"
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
			log.Error().Msg(fmt.Sprintf("Failed to parse version for Kobo at %s", path))
		}
		log.Debug().Msg(fmt.Sprintf("Found device with ID %s", deviceId))
		device, found := kobo.DeviceByID(deviceId)
		if !found {
			log.Warn().Msg("Found a device that isn't officially supported but may still work anyway")
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
		log.Info().Msg(result.Error.Error())
		return nil, result.Error
	}
	return bookmark, nil
}

func (k *Kobo) FindBookOnDevice(bookID string) (Content, error) {
	var content Content
	log.Debug().Msg("Retrieving books that have been uploaded to Readwise previously")
	result := Conn.Where(&Content{ContentType: "6", VolumeIndex: -1, ContentID: bookID}).Find(&content)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to retrieve content from device")
		return content, result.Error
	}
	log.Debug().Str("title", content.Title).Msg("Successfully retrieve content from device DB")
	return content, nil
}

func (k *Kobo) ListDeviceContent() ([]Content, error) {
	var content []Content
	log.Debug().Msg("Retrieving content from device")
	result := Conn.Where(
		&Content{ContentType: "6", VolumeIndex: -1},
	).Order("___PercentRead desc, title asc").Find(&content)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to retrieve content from device")
		return nil, result.Error
	}
	log.Debug().Int("content_count", len(content)).Msg("Successfully retrieved device content")
	return content, nil
}

func (k *Kobo) ListDeviceBookmarks() ([]Bookmark, error) {
	var bookmarks []Bookmark
	log.Debug().Msg("Retrieving bookmarks from device")
	result := Conn.Order("VolumeID ASC, ChapterProgress ASC").Find(&bookmarks).Limit(1)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to retrieve bookmarks from device")
		return nil, result.Error
	}
	log.Debug().Int("bookmark_count", len(bookmarks)).Msg("Successfully retrieved device bookmarks")
	return bookmarks, nil
}

func (k *Kobo) BuildContentIndex(content []Content) map[string]Content {
	log.Debug().Msg("Building an index out of device content")
	contentIndex := make(map[string]Content)
	for _, item := range content {
		contentIndex[item.ContentID] = item
	}
	log.Debug().Int("index_count", len(contentIndex)).Msg("Built content index")
	return contentIndex
}

func (k *Kobo) CountDeviceBookmarks() int64 {
	var count int64
	result := Conn.Model(&Bookmark{}).Count(&count)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to count bookmarks on device")
	}
	return count
}
