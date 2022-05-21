package backend

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

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
	Count   int             `json:"count"`
	Results []BookListEntry `json:"results"`
}

type BookListEntry struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	CoverURL  string `json:"cover_image_url"`
	SourceURL string `json:"source_url"`
}

type Readwise struct{}

func NewReadwise() *Readwise {
	return &Readwise{}
}

func (r *Readwise) CheckTokenValidity(token string) error {
	req, err := http.NewRequest("GET", authEndpoint, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", token))
	req.Header.Add("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		return errors.New(resp.Status)
	}
	if resp.StatusCode != 204 {
		return errors.New(resp.Status)
	}
	log.Info().Msg("Successfully validated token against the Readwise API")
	return nil
}
