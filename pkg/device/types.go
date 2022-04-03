package device

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
