package backend

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/pgaskin/koboutils/v2/kobo"
)

type Kobo struct {
	Name       string `json:"name"`
	Storage    int    `json:"storage"`
	DisplayPPI int    `json:"display_ppi"`
	MntPath    string `json:"mnt_path"`
	DbPath     string `json:"db_path"`
}

type HighlightCounts struct {
	Total      int64 `json:"total"`
	Sideloaded int64 `json:"sideloaded"`
	Official   int64 `json:"official"`
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

func GetKoboMetadata(detectedPaths []string, logger *slog.Logger) []Kobo {
	var kobos []Kobo
	for _, path := range detectedPaths {
		_, _, deviceId, err := kobo.ParseKoboVersion(path)
		if err != nil {
			logger.Error("Failed to parse Kobo version",
				slog.String("error", err.Error()),
				slog.String("kobo_path", path),
			)
		}
		logger.Info("Found attached device",
			slog.String("device_id", deviceId),
			slog.String("kobo_path", path),
		)
		device, found := kobo.DeviceByID(deviceId)
		if found {
			kobos = append(kobos, Kobo{
				Name:       device.Name(),
				Storage:    device.StorageGB(),
				DisplayPPI: device.DisplayPPI(),
				MntPath:    path,
				DbPath:     fmt.Sprintf("%s/.kobo/KoboReader.sqlite", path),
			})
			continue
		}
		fallbackSkus := getFallbackSkus()
		deviceIdBits := strings.Split(deviceId, ",")
		deviceIdGuid := deviceIdBits[len(deviceIdBits)-1]
		if fallback, ok := fallbackSkus[deviceIdGuid]; ok {
			kobos = append(kobos, Kobo{
				Name:       fallback.Name,
				Storage:    fallback.Storage,
				DisplayPPI: fallback.DisplayPPI,
				MntPath:    path,
				DbPath:     fmt.Sprintf("%s/.kobo/KoboReader.sqlite", path),
			})
			continue
		}
		logger.Warn("Found a device that isn't officially supported but will likely still operate just fine",
			slog.String("device_id", deviceId),
		)
		// We can handle unsupported Kobos in future but at present, there are none
		kobos = append(kobos, Kobo{
			MntPath: path,
			DbPath:  fmt.Sprintf("%s/.kobo/KoboReader.sqlite", path),
		})
	}
	return kobos
}

func (k *Kobo) ListDeviceContent(includeStoreBought bool, logger *slog.Logger) ([]Content, error) {
	var content []Content
	logger.Debug("Retrieving content list from device")
	result := Conn.Where(&Content{ContentType: "6", VolumeIndex: -1})
	if !includeStoreBought {
		result = result.Where("ContentID LIKE '%file:///%'")
	}
	result = result.Order("___PercentRead desc, title asc").Find(&content)
	if result.Error != nil {
		logger.Error("Failed to retrieve content from device",
			slog.String("error", result.Error.Error()),
		)
		return nil, result.Error
	}
	logger.Debug("Successfully retrieved device content",
		slog.Int("content_count", len(content)),
	)
	return content, nil
}

func (k *Kobo) ListDeviceBookmarks(includeStoreBought bool, logger *slog.Logger) ([]Bookmark, error) {
	var bookmarks []Bookmark
	logger.Debug("Retrieving bookmarks from device")
	result := Conn
	if !includeStoreBought {
		result = result.Where("VolumeID LIKE '%file:///%'")
	}
	result = result.Order("VolumeID ASC, ChapterProgress ASC").Find(&bookmarks).Limit(1)
	if result.Error != nil {
		logger.Error("Failed to retrieve bookmarks from device",
			slog.String("error", result.Error.Error()),
		)
		return nil, result.Error
	}
	logger.Debug("Successfully retrieved device bookmarks",
		slog.Int("bookmark_count", len(bookmarks)),
	)
	return bookmarks, nil
}

func (k *Kobo) BuildContentIndex(content []Content, logger *slog.Logger) map[string]Content {
	logger.Debug("Building an index out of device content")
	contentIndex := make(map[string]Content)
	for _, item := range content {
		contentIndex[item.ContentID] = item
	}
	logger.Debug("Built content index",
		slog.Int("index_count", len(contentIndex)),
	)
	return contentIndex
}

func (k *Kobo) CountDeviceBookmarks(logger *slog.Logger) HighlightCounts {
	var totalCount int64
	var officialCount int64
	var sideloadedCount int64
	result := Conn.Model(&Bookmark{}).Count(&totalCount)
	if result.Error != nil {
		logger.Error("Failed to count bookmarks on device",
			slog.String("error", result.Error.Error()),
		)
	}
	Conn.Model(&Bookmark{}).Where("VolumeID LIKE '%file:///%'").Count(&sideloadedCount)
	Conn.Model(&Bookmark{}).Where("VolumeID NOT LIKE '%file:///%'").Count(&officialCount)
	return HighlightCounts{
		Total:      totalCount,
		Official:   officialCount,
		Sideloaded: sideloadedCount,
	}
}

func getFallbackSkus() map[string]Kobo {
	return map[string]Kobo{
		"00000000-0000-0000-0000-000000000386": {
			Name:       "Kobo Clara 2E",
			Storage:    16,
			DisplayPPI: 300,
		},
		"00000000-0000-0000-0000-000000000389": {
			Name:       "Kobo Elipsa 2E",
			Storage:    32,
			DisplayPPI: 227,
		},
	}
}
