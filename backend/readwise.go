package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
	"time"

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

func (r *Readwise) CheckTokenValidity(token string) error {
	req, err := http.NewRequest("GET", AuthEndpoint, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", token))
	req.Header.Add("User-Agent", UserAgent)
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

func (r *Readwise) SendBookmarks(payload Response, token string) (int, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Token %s", token)).
		SetHeader("User-Agent", UserAgent).
		SetBody(payload).
		Post(HighlightsEndpoint)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Failed to send request to Readwise: code %d", resp.StatusCode()))
	}
	if resp.StatusCode() != 200 {
		log.Error().
			Int("status", resp.StatusCode()).
			Str("response", string(resp.Body())).
			Msg("Received a non-200 response from Readwise")
		return 0, errors.New(fmt.Sprintf("Received a non-200 status code from Readwise: code %d", resp.StatusCode()))
	}
	log.Info().Int("highlight_count", len(payload.Highlights)).Msg("Successfully sent bookmarks to Readwise")
	return len(payload.Highlights), nil
}

func (r *Readwise) RetrieveUploadedBooks(token string) (BookListResponse, error) {
	bookList := BookListResponse{}
	headers := map[string][]string{
		"Authorization": {fmt.Sprintf("Token %s", token)},
		"User-Agent":    {UserAgent},
	}
	client := http.Client{}
	remoteURL, err := url.Parse(BooksEndpoint)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse books URL")
	}
	request := http.Request{
		Method: "GET",
		URL:    remoteURL,
		Header: headers,
	}
	res, err := client.Do(&request)
	if err != nil {
		log.Error().Err(err)
		return bookList, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			panic(err)
		}
	}(res.Body)
	b, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Error().Err(err)
		return bookList, err
	}
	if res.StatusCode != 200 {
		log.Error().
			Int("status", res.StatusCode).
			Msg("Received a non-200 response from Readwise")
		return bookList, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse response")
		panic(err)
	}
	err = json.Unmarshal(body, &bookList)
	if err != nil {
		log.Error().
			Err(err).
			Int("status", res.StatusCode).
			Str("body", string(b)).
			Msg("Failed to unmarshal response from Readwise")
		return bookList, err
	}
	log.Info().Int("book_count", bookList.Count).Msg("Successfully retrieved books from Readwise API")
	return bookList, nil
}

func (r *Readwise) UploadCover(encodedCover string, bookId int, token string) error {
	body := map[string]interface{}{
		"cover_image": encodedCover,
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Token %s", token)).
		SetHeader("Content-Type", "application/json").
		SetHeader("User-Agent", UserAgent).
		SetBody(body).
		Patch(fmt.Sprintf(CoverEndpoint, bookId))
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		log.Error().
			Int("status", resp.StatusCode()).
			Str("response", string(resp.Body())).
			Msg("Received a non-200 response from Readwise")
		return fmt.Errorf("failed to upload cover for book with id %d", bookId)
	}
	return nil
}

func BuildPayload(bookmarks []Bookmark, contentIndex map[string]Content) (Response, error) {
	var payload Response
	for _, entry := range bookmarks {
		source := contentIndex[entry.VolumeID]
		log.Info().Interface("source", source).Msg("Parsing entry")
		t, err := time.Parse("2006-01-02T15:04:05.000", entry.DateCreated)
		if err != nil {
			log.Error().Err(err).Interface("bookmark", entry).Msg("Failed tp parse timestamp from bookmark")
			return Response{}, err
		}
		createdAt := t.Format("2006-01-02T15:04:05-07:00")
		text := NormaliseText(entry.Text)
		if entry.Annotation != "" && text == "" {
			// I feel like this state probably shouldn't be possible but we'll handle it anyway
			// since it's useful to surface annotations, regardless of highlights. We put a
			// glaring placeholder here because the text field is required by the Readwise API.
			text = "Placeholder for attached annotation"
		}
		if entry.Annotation == "" && text == "" {
			// This state should be impossible but stranger things have happened so worth a sanity check
			log.Info().
				Interface("source", source).
				Interface("bookmark", entry).
				Msg("Found an entry with neither highlighted text nor an annotation so skipping entry")
			continue
		}
		if source.Title == "" {
			// While Kepubs have a title in the Kobo database, the same can't be guaranteed for epubs at all.
			// In that event, we just fall back to using the filename
			sourceFile, err := url.Parse(entry.VolumeID)
			if err != nil {
				// While extremely unlikely, we should handle the case where a VolumeID doesn't have a suffix. This condition is only
				// triggered for completely busted names such as control codes given url.Parse will happen take URLs without a protocol
				// or even just arbitrary strings. Given we don't set a title here, we will use the Readwise fallback which is to add
				// these highlights to a book called "Quotes" and let the user figure out their metadata situation. That reminds me though:
				// TODO: Test exports with non-epub files
				log.Error().
					Err(err).
					Interface("source", source).
					Interface("bookmark", entry).
					Msg("Failed to retrieve epub title. This is not a hard requirement so sending with a dummy title.")
				goto sendhighlight
			}
			filename := path.Base(sourceFile.Path)
			log.Debug().Str("filename", filename).Msg("No source title. Constructing title from filename")
			source.Title = strings.TrimSuffix(filename, ".epub")
		}
	sendhighlight:
		highlight := Highlight{
			Text:          text,
			Title:         source.Title,
			Author:        source.Attribution,
			SourceURL:     entry.VolumeID,
			SourceType:    SourceType,
			Category:      SourceCategory,
			Note:          entry.Annotation,
			HighlightedAt: createdAt,
		}
		log.Debug().Interface("highlight", highlight).Msg("Successfully built highlights")
		payload.Highlights = append(payload.Highlights, highlight)
	}
	log.Info().Int("highlight_count", len(payload.Highlights)).Msg("Successfully parsed highlights")
	return payload, nil
}

func NormaliseText(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\n", " ")
	return s
}
