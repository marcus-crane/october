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
		    if (a.slice) {
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
	    text: string;
	    annotation: string;
	    extra_annotation_data: string;
	    date_created: string;
	
	    static createFrom(source: any = {}) {
	        return new Bookmark(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.bookmark_id = source["bookmark_id"];
	        this.volume_id = source["volume_id"];
	        this.content_id = source["content_id"];
	        this.text = source["text"];
	        this.annotation = source["annotation"];
	        this.extra_annotation_data = source["extra_annotation_data"];
	        this.date_created = source["date_created"];
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
	    date_last_read: string;
	    num_pages: number;
	    percent_read: string;
	
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
	        this.date_last_read = source["date_last_read"];
	        this.num_pages = source["num_pages"];
	        this.percent_read = source["percent_read"];
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
		    if (a.slice) {
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

