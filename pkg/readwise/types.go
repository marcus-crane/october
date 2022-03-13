package readwise

type Response struct {
	Highlights []Highlight `json:"highlights"`
}

type Highlight struct {
	Text          string `json:"text"`
	Title         string `json:"title,omitempty"`
	Author        string `json:"author,omitempty"`
	SourceURL     string `json:"source_url"`
	SourceType    string `json:"source_type"`
	Category      string `json:"category"`
	Note          string `json:"note,omitempty"`
	HighlightedAt string `json:"highlighted_at,omitempty"`
}

type CoverUpdate struct {
	Cover string `json:"cover"`
}

type BookListResponse struct {
	Count   int `json:"count"`
	Results struct {
		ID        int    `json:"id"`
		CoverURL  string `json:"cover_image_url"`
		SourceURL string `json:"source_url"`
	}
}
