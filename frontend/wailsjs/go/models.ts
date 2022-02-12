/* Do not change, this code is generated from Golang structs */

export {};

export class Content {


    static createFrom(source: any = {}) {
        return new Content(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}
export class Highlight {
    text: string;
    title?: string;
    author?: string;
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

export class Bookmark {


    static createFrom(source: any = {}) {
        return new Bookmark(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}
