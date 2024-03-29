package kobo

type Content struct {
	ContentID                    string  `db:"ContentID"`
	ContentType                  string  `db:"ContentType"`
	MimeType                     string  `db:"MimeType"`
	BookID                       string  `db:"BookID"`
	BookTitle                    string  `db:"BookTitle"`
	ImageId                      string  `db:"ImageId"`
	Title                        string  `db:"Title"`
	Attribution                  string  `db:"Attribution"`
	Description                  string  `db:"Description"`
	DateCreated                  string  `db:"DateCreated"`
	ShortCoverKey                string  `db:"ShortCoverKey"`
	AdobeLocation                string  `db:"adobe_location"`
	Publisher                    string  `db:"Publisher"`
	IsEncrypted                  bool    `db:"IsEncrypted"`
	DateLastRead                 string  `db:"DateLastRead"`
	FirstTimeReading             bool    `db:"FirstTimeReading"`
	ChapterIDBookmarked          string  `db:"ChapterIDBookmarked"`
	ParagraphBookmarked          int     `db:"ParagraphBookmarked"`
	BookmarkWordOffset           int     `db:"BookmarkWordOffset"`
	NumShortcovers               int     `db:"NumShortcovers"`
	VolumeIndex                  int     `db:"VolumeIndex"`
	NumPages                     int     `db:"___NumPages"`
	ReadStatus                   int     `db:"ReadStatus"`
	SyncTime                     string  `db:"___SyncTime"`
	UserID                       string  `db:"___UserID"`
	PublicationId                string  `db:"PublicationId"`
	FileOffset                   int     `db:"___FileOffset"`
	FileSize                     int     `db:"___FileSize"`
	PercentRead                  int     `db:"___PercentRead"`
	ExpirationStatus             int     `db:"___ExpirationStatus"`
	FavouritesIndex              int     `db:"FavouritesIndex"`
	Accessibility                int     `db:"Accessibility"`
	ContentURL                   string  `db:"ContentURL"`
	Language                     string  `db:"Language"`
	BookshelfTags                string  `db:"BookshelfTags"`
	IsDownloaded                 bool    `db:"IsDownloaded"`
	FeedbackType                 int     `db:"FeedbackType"`
	AverageRating                int     `db:"AverageRating"`
	Depth                        int     `db:"Depth"`
	PageProgressDirection        string  `db:"PageProgressDirection"`
	InWishlist                   bool    `db:"InWishlist"`
	ISBN                         string  `db:"ISBN"`
	WishlistedDate               string  `db:"WishlistedDate"`
	FeedbackTypeSynced           int     `db:"FeedbackTypeSynced"`
	IsSocialEnabled              bool    `db:"IsSocialEnabled"`
	EpubType                     int     `db:"EpubType"`
	Monetization                 int     `db:"Monetization"`
	ExternalId                   string  `db:"ExternalId"`
	Series                       string  `db:"Series"`
	SeriesNumber                 string  `db:"SeriesNumber"`
	Subtitle                     string  `db:"Subtitle"`
	WordCount                    int     `db:"WordCount"`
	Fallback                     string  `db:"Fallback"`
	RestOfBookEstimate           int     `db:"RestOfBookEstimate"`
	CurrentChapterEstimate       int     `db:"CurrentChapterEstimate"`
	CurrentChapterProgress       float32 `db:"CurrentChapterProgress"`
	PocketStatus                 int     `db:"PocketStatus"`
	UnsyncedPocketChanges        string  `db:"UnsyncedPocketChanges"`
	ImageUrl                     string  `db:"ImageUrl"`
	DateAdded                    string  `db:"DateAdded"`
	WorkId                       string  `db:"WorkId"`
	Properties                   string  `db:"Properties"`
	RenditionSpread              string  `db:"RenditionSpread"`
	RatingCount                  int     `db:"RatingCount"`
	ReviewsSyncDate              string  `db:"ReviewsSyncDate"`
	MediaOverlay                 string  `db:"MediaOverlay"`
	MediaOverlayType             string  `db:"MediaOverlayType"`
	RedirectPreviewUrl           bool    `db:"RedirectPreviewUrl"`
	PreviewFileSize              int     `db:"PreviewFileSize"`
	EntitlementId                string  `db:"EntitlementId"`
	CrossRevisionId              string  `db:"CrossRevisionId"`
	DownloadUrl                  string  `db:"DownloadUrl"`
	ReadStateSynced              bool    `db:"ReadStateSynced"`
	TimesStartedReading          int     `db:"TimesStartedReading"`
	TimeSpentReading             int     `db:"TimeSpentReading"`
	LastTimeStartedReading       string  `db:"LastTimeStartedReading"`
	LastTimeFinishedReading      string  `db:"LastTimeFinishedReading"`
	ApplicableSubscriptions      string  `db:"ApplicableSubscriptions"`
	ExternalIds                  string  `db:"ExternalIds"`
	PurchaseRevisionId           string  `db:"PurchaseRevisionId"`
	SeriesID                     string  `db:"SeriesID"`
	SeriesNumberFloat            float64 `db:"SeriesNumberFloat"`
	AdobeLoanExpiration          string  `db:"AdobeLoanExpiration"`
	HideFromHomePage             bool    `db:"HideFromHomePage"`
	IsInternetArchive            bool    `db:"IsInternetArchive"`
	TitleKana                    string  `db:"titleKana"`
	SubtitleKana                 string  `db:"subtitleKana"`
	SeriesKana                   string  `db:"seriesKana"`
	AttributionKana              string  `db:"attributionKana"`
	PublisherKana                string  `db:"publisherKana"`
	IsPurchaseable               bool    `db:"IsPurchaseable"`
	IsSupported                  bool    `db:"IsSupported"`
	AnnotationsSyncToken         string  `db:"AnnotationsSyncToken"`
	DateModified                 string  `db:"DateModified"`
	StorePages                   int     `db:"StorePages"`
	StoreWordCount               int     `db:"StoreWordCount"`
	StoreTimeToReadLowerEstimate int     `db:"StoreTimeToReadLowerEstimate"`
	StoreTimeToReadUpperEstimate int     `db:"StoreTimeToReadUpperEstimate"`
	Duration                     int     `db:"Duration"`
	IsAbridged                   bool    `db:"IsAbridged"`
}

func CountContent(kobo *Kobo) (count int, err error) {
	if err := kobo.dbClient.Get(
		&count,
		"SELECT count(*) FROM content WHERE ContentType = ? AND VolumeIndex = ? AND MimeType = ?",
		6, -1, "application/x-kobo-epub+zip",
	); err != nil {
		return count, err
	}
	return count, nil
}
