package readwise

import "github.com/marcus-crane/october/pkg/kobo"

type Response struct {
	Highlights []kobo.Highlight `json:"highlights"`
}
