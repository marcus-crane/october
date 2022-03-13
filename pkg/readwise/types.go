package readwise

import "github.com/marcus-crane/october/pkg/device"

type Response struct {
	Highlights []device.Highlight `json:"highlights"`
}
