export namespace backend {
	
	export class BookListEntry {
	    id: number;
	    title: string;
	    cover_image_url: string;
	    source_url: string;
	
	    static createFrom(source: any = {}) {
	        return new BookListEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.cover_image_url = source["cover_image_url"];
	        this.source_url = source["source_url"];
	    }
	}
	export class BookListResponse {
	    count: number;
	    results: BookListEntry[];
	
	    static createFrom(source: any = {}) {
	        return new BookListResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.count = source["count"];
	        this.results = this.convertValues(source["results"], BookListEntry);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Bookmark {
	    bookmark_id: string;
	    volume_id: string;
	    content_id: string;
	    StartContainerPath: string;
	    StartContainerChild: string;
	    StartContainerChildIndex: string;
	    StartOffset: string;
	    EndContainerPath: string;
	    EndContainerChildIndex: string;
	    EndOffset: string;
	    text: string;
	    annotation: string;
	    extra_annotation_data: string;
	    date_created: string;
	    ChapterProgress: number;
	    Hidden: string;
	    Version: string;
	    DateModified: string;
	    Creator: string;
	    UUID: string;
	    UserID: string;
	    SyncTime: string;
	    Published: string;
	    ContextString: string;
	    Type: string;
	
	    static createFrom(source: any = {}) {
	        return new Bookmark(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.bookmark_id = source["bookmark_id"];
	        this.volume_id = source["volume_id"];
	        this.content_id = source["content_id"];
	        this.StartContainerPath = source["StartContainerPath"];
	        this.StartContainerChild = source["StartContainerChild"];
	        this.StartContainerChildIndex = source["StartContainerChildIndex"];
	        this.StartOffset = source["StartOffset"];
	        this.EndContainerPath = source["EndContainerPath"];
	        this.EndContainerChildIndex = source["EndContainerChildIndex"];
	        this.EndOffset = source["EndOffset"];
	        this.text = source["text"];
	        this.annotation = source["annotation"];
	        this.extra_annotation_data = source["extra_annotation_data"];
	        this.date_created = source["date_created"];
	        this.ChapterProgress = source["ChapterProgress"];
	        this.Hidden = source["Hidden"];
	        this.Version = source["Version"];
	        this.DateModified = source["DateModified"];
	        this.Creator = source["Creator"];
	        this.UUID = source["UUID"];
	        this.UserID = source["UserID"];
	        this.SyncTime = source["SyncTime"];
	        this.Published = source["Published"];
	        this.ContextString = source["ContextString"];
	        this.Type = source["Type"];
	    }
	}
	export class Content {
	    content_id: string;
	    content_type: string;
	    mime_type: string;
	    book_id: string;
	    book_title: string;
	    image_id: string;
	    title: string;
	    attribution: string;
	    description: string;
	    date_created: string;
	    ShortCoverKey: string;
	    AdobeLocation: string;
	    Publisher: string;
	    IsEncrypted: boolean;
	    date_last_read: string;
	    FirstTimeReading: boolean;
	    ChapterIDBookmarked: string;
	    ParagraphBookmarked: number;
	    BookmarkWordOffset: number;
	    NumShortcovers: number;
	    VolumeIndex: number;
	    num_pages: number;
	    ReadStatus: number;
	    SyncTime: string;
	    UserID: string;
	    PublicationId: string;
	    FileOffset: number;
	    FileSize: number;
	    percent_read: string;
	    ExpirationStatus: number;
	    FavouritesIndex: number;
	    Accessibility: number;
	    ContentURL: string;
	    Language: string;
	    BookshelfTags: string;
	    IsDownloaded: boolean;
	    FeedbackType: number;
	    AverageRating: number;
	    Depth: number;
	    PageProgressDirection: string;
	    InWishlist: string;
	    ISBN: string;
	    WishlistedDate: string;
	    FeedbackTypeSynced: boolean;
	    IsSocialEnabled: boolean;
	    EpubType: string;
	    Monetization: string;
	    ExternalId: string;
	    Series: string;
	    SeriesNumber: string;
	    Subtitle: string;
	    WordCount: string;
	    Fallback: string;
	    RestOfBookEstimate: string;
	    CurrentChapterEstimate: string;
	    CurrentChapterProgress: number;
	    PocketStatus: string;
	    UnsyncedPocketChanges: string;
	    ImageUrl: string;
	    DateAdded: string;
	    WorkId: string;
	    Properties: string;
	    RenditionSpread: string;
	    RatingCount: string;
	    ReviewsSyncDate: string;
	    MediaOverlay: string;
	    RedirectPreviewUrl: boolean;
	    PreviewFileSize: number;
	    EntitlementId: string;
	    CrossRevisionId: string;
	    DownloadUrl: boolean;
	    ReadStateSynced: boolean;
	    TimesStartedReading: number;
	    TimeSpentReading: number;
	    LastTimeStartedReading: string;
	    LastTimeFinishedReading: string;
	    ApplicableSubscriptions: string;
	    ExternalIds: string;
	    PurchaseRevisionId: string;
	    SeriesID: string;
	    SeriesNumberFloat: number;
	    AdobeLoanExpiration: string;
	    HideFromHomePage: boolean;
	    IsInternetArchive: boolean;
	    TitleKana: string;
	    SubtitleKana: string;
	    SeriesKana: string;
	    AttributionKana: string;
	    PublisherKana: string;
	    IsPurchaseable: boolean;
	    IsSupported: boolean;
	    AnnotationsSyncToken: string;
	    DateModified: string;
	
	    static createFrom(source: any = {}) {
	        return new Content(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.content_id = source["content_id"];
	        this.content_type = source["content_type"];
	        this.mime_type = source["mime_type"];
	        this.book_id = source["book_id"];
	        this.book_title = source["book_title"];
	        this.image_id = source["image_id"];
	        this.title = source["title"];
	        this.attribution = source["attribution"];
	        this.description = source["description"];
	        this.date_created = source["date_created"];
	        this.ShortCoverKey = source["ShortCoverKey"];
	        this.AdobeLocation = source["AdobeLocation"];
	        this.Publisher = source["Publisher"];
	        this.IsEncrypted = source["IsEncrypted"];
	        this.date_last_read = source["date_last_read"];
	        this.FirstTimeReading = source["FirstTimeReading"];
	        this.ChapterIDBookmarked = source["ChapterIDBookmarked"];
	        this.ParagraphBookmarked = source["ParagraphBookmarked"];
	        this.BookmarkWordOffset = source["BookmarkWordOffset"];
	        this.NumShortcovers = source["NumShortcovers"];
	        this.VolumeIndex = source["VolumeIndex"];
	        this.num_pages = source["num_pages"];
	        this.ReadStatus = source["ReadStatus"];
	        this.SyncTime = source["SyncTime"];
	        this.UserID = source["UserID"];
	        this.PublicationId = source["PublicationId"];
	        this.FileOffset = source["FileOffset"];
	        this.FileSize = source["FileSize"];
	        this.percent_read = source["percent_read"];
	        this.ExpirationStatus = source["ExpirationStatus"];
	        this.FavouritesIndex = source["FavouritesIndex"];
	        this.Accessibility = source["Accessibility"];
	        this.ContentURL = source["ContentURL"];
	        this.Language = source["Language"];
	        this.BookshelfTags = source["BookshelfTags"];
	        this.IsDownloaded = source["IsDownloaded"];
	        this.FeedbackType = source["FeedbackType"];
	        this.AverageRating = source["AverageRating"];
	        this.Depth = source["Depth"];
	        this.PageProgressDirection = source["PageProgressDirection"];
	        this.InWishlist = source["InWishlist"];
	        this.ISBN = source["ISBN"];
	        this.WishlistedDate = source["WishlistedDate"];
	        this.FeedbackTypeSynced = source["FeedbackTypeSynced"];
	        this.IsSocialEnabled = source["IsSocialEnabled"];
	        this.EpubType = source["EpubType"];
	        this.Monetization = source["Monetization"];
	        this.ExternalId = source["ExternalId"];
	        this.Series = source["Series"];
	        this.SeriesNumber = source["SeriesNumber"];
	        this.Subtitle = source["Subtitle"];
	        this.WordCount = source["WordCount"];
	        this.Fallback = source["Fallback"];
	        this.RestOfBookEstimate = source["RestOfBookEstimate"];
	        this.CurrentChapterEstimate = source["CurrentChapterEstimate"];
	        this.CurrentChapterProgress = source["CurrentChapterProgress"];
	        this.PocketStatus = source["PocketStatus"];
	        this.UnsyncedPocketChanges = source["UnsyncedPocketChanges"];
	        this.ImageUrl = source["ImageUrl"];
	        this.DateAdded = source["DateAdded"];
	        this.WorkId = source["WorkId"];
	        this.Properties = source["Properties"];
	        this.RenditionSpread = source["RenditionSpread"];
	        this.RatingCount = source["RatingCount"];
	        this.ReviewsSyncDate = source["ReviewsSyncDate"];
	        this.MediaOverlay = source["MediaOverlay"];
	        this.RedirectPreviewUrl = source["RedirectPreviewUrl"];
	        this.PreviewFileSize = source["PreviewFileSize"];
	        this.EntitlementId = source["EntitlementId"];
	        this.CrossRevisionId = source["CrossRevisionId"];
	        this.DownloadUrl = source["DownloadUrl"];
	        this.ReadStateSynced = source["ReadStateSynced"];
	        this.TimesStartedReading = source["TimesStartedReading"];
	        this.TimeSpentReading = source["TimeSpentReading"];
	        this.LastTimeStartedReading = source["LastTimeStartedReading"];
	        this.LastTimeFinishedReading = source["LastTimeFinishedReading"];
	        this.ApplicableSubscriptions = source["ApplicableSubscriptions"];
	        this.ExternalIds = source["ExternalIds"];
	        this.PurchaseRevisionId = source["PurchaseRevisionId"];
	        this.SeriesID = source["SeriesID"];
	        this.SeriesNumberFloat = source["SeriesNumberFloat"];
	        this.AdobeLoanExpiration = source["AdobeLoanExpiration"];
	        this.HideFromHomePage = source["HideFromHomePage"];
	        this.IsInternetArchive = source["IsInternetArchive"];
	        this.TitleKana = source["TitleKana"];
	        this.SubtitleKana = source["SubtitleKana"];
	        this.SeriesKana = source["SeriesKana"];
	        this.AttributionKana = source["AttributionKana"];
	        this.PublisherKana = source["PublisherKana"];
	        this.IsPurchaseable = source["IsPurchaseable"];
	        this.IsSupported = source["IsSupported"];
	        this.AnnotationsSyncToken = source["AnnotationsSyncToken"];
	        this.DateModified = source["DateModified"];
	    }
	}
	export class Highlight {
	    text: string;
	    title?: string;
	    author?: string;
	    source_url: string;
	    source_type: string;
	    category: string;
	    note?: string;
	    highlighted_at?: string;
	
	    static createFrom(source: any = {}) {
	        return new Highlight(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.text = source["text"];
	        this.title = source["title"];
	        this.author = source["author"];
	        this.source_url = source["source_url"];
	        this.source_type = source["source_type"];
	        this.category = source["category"];
	        this.note = source["note"];
	        this.highlighted_at = source["highlighted_at"];
	    }
	}
	export class HighlightCounts {
	    total: number;
	    sideloaded: number;
	    official: number;
	
	    static createFrom(source: any = {}) {
	        return new HighlightCounts(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.total = source["total"];
	        this.sideloaded = source["sideloaded"];
	        this.official = source["official"];
	    }
	}
	export class Kobo {
	    name: string;
	    storage: number;
	    display_ppi: number;
	    mnt_path: string;
	    db_path: string;
	
	    static createFrom(source: any = {}) {
	        return new Kobo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.storage = source["storage"];
	        this.display_ppi = source["display_ppi"];
	        this.mnt_path = source["mnt_path"];
	        this.db_path = source["db_path"];
	    }
	}
	export class Response {
	    highlights: Highlight[];
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.highlights = this.convertValues(source["highlights"], Highlight);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Settings {
	    readwise_token: string;
	    upload_covers: boolean;
	    upload_store_highlights: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Settings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.readwise_token = source["readwise_token"];
	        this.upload_covers = source["upload_covers"];
	        this.upload_store_highlights = source["upload_store_highlights"];
	    }
	}

}

