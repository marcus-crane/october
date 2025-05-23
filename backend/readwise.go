package backend

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
	"time"
)

const (
	HIGHLIGHT_REQUEST_BATCH_MAX = 2000
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

type Readwise struct {
	logger    *slog.Logger
	UserAgent string
}

func (r *Readwise) CheckTokenValidity(token string) error {
	req, err := http.NewRequest("GET", AuthEndpoint, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", token))
	req.Header.Add("User-Agent", r.UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		return errors.New(resp.Status)
	}
	if resp.StatusCode != 204 {
		return errors.New(resp.Status)
	}
	r.logger.Info("Successfully validated token against the Readwise API")
	return nil
}

func (r *Readwise) SendBookmarks(payloads []Response, token string) (int, error) {
	// TODO: This is dumb, we count stuff that this function doesn't need to know about + we already know the size from earlier
	submittedHighlights := 0
	for _, payload := range payloads {
		data, err := json.Marshal(payload)
		if err != nil {
			return 0, fmt.Errorf("failed to marshal bookmark payload: %+v", err)
		}
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPost, HighlightsEndpoint, bytes.NewBuffer(data))
		if err != nil {
			return 0, fmt.Errorf("failed to construct bookmark request: %+v", err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("Token %s", token))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("User-Agent", r.UserAgent)
		resp, err := client.Do(req)
		if err != nil {
			return 0, fmt.Errorf("failed to send request to Readwise: code %d", resp.StatusCode)
		}
		if resp.StatusCode != 200 {
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err == nil {
				r.logger.Error("Received a non-200 response from Readwise",
					slog.Int("status_code", resp.StatusCode),
					slog.String("response", string(body)),
				)
			}
			return 0, fmt.Errorf("received a non-200 status code from Readwise: code %d", resp.StatusCode)
		}
		submittedHighlights += len(payload.Highlights)
	}
	r.logger.Info("Successfully sent bookmarks to Readwise",
		slog.Int("batch_count", len(payloads)),
	)
	return submittedHighlights, nil
}

func (r *Readwise) RetrieveUploadedBooks(token string) (BookListResponse, error) {
	bookList := BookListResponse{}
	headers := map[string][]string{
		"Authorization": {fmt.Sprintf("Token %s", token)},
		"User-Agent":    {r.UserAgent},
	}
	client := http.Client{}
	remoteURL, err := url.Parse(BooksEndpoint)
	if err != nil {
		r.logger.Error("Failed to parse URL for Readwise book upload endpoint",
			slog.String("error", err.Error()),
		)
	}
	request := http.Request{
		Method: "GET",
		URL:    remoteURL,
		Header: headers,
	}
	res, err := client.Do(&request)
	if err != nil {
		r.logger.Error("An unexpected error occurred while retrieving uploads from Readwise",
			slog.String("error", err.Error()),
			slog.Int("status_code", res.StatusCode),
		)
		return bookList, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			panic(err)
		}
	}(res.Body)
	b, err := httputil.DumpResponse(res, true)
	if err != nil {
		r.logger.Error("Encountered an error while dumping response from Readwise",
			slog.String("error", err.Error()),
		)
		return bookList, err
	}
	if res.StatusCode != 200 {
		r.logger.Error("Received a non-200 response from Readwise",
			slog.Int("status", res.StatusCode),
			slog.String("body", string(b)),
		)
		return bookList, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		r.logger.Error("Failed to parse response from Readwise",
			slog.String("error", err.Error()),
		)
		return bookList, err
	}
	err = json.Unmarshal(body, &bookList)
	if err != nil {
		r.logger.Error("Failed to unmarshal response from Readwise",
			slog.String("error", err.Error()),
			slog.Int("status", res.StatusCode),
			slog.String("body", string(b)),
		)
		return bookList, err
	}
	r.logger.Info("Successfully retrieved books from Readwise API",
		slog.Int("book_count", bookList.Count),
	)
	return bookList, nil
}

func (r *Readwise) UploadCover(encodedCover string, bookId int, token string) error {
	client := &http.Client{}
	body := map[string]interface{}{
		"cover_image": encodedCover,
	}
	data, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal cover upload payload: %+v", err)
	}
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf(CoverEndpoint, bookId), bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to construct cover upload request: %+v", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", token))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", r.UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to upload cover to Readwise: code %d", resp.StatusCode)
	}
	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			r.logger.Error("Received a non-200 response from Readwise",
				slog.Int("status", resp.StatusCode),
				slog.String("body", string(body)),
			)
		}
		return fmt.Errorf("failed to upload cover for book with id %d", bookId)
	}
	return nil
}

func BuildPayload(bookmarks []Bookmark, contentIndex map[string]Content, logger *slog.Logger) ([]Response, error) {
	var payloads []Response
	var currentBatch Response
	for count, entry := range bookmarks {
		// If max payload size is reached, start building another batch which will be sent separately
		if count > 0 && (count%HIGHLIGHT_REQUEST_BATCH_MAX == 0) {
			fmt.Println(count / HIGHLIGHT_REQUEST_BATCH_MAX)
			payloads = append(payloads, currentBatch)
			currentBatch = Response{}
		}
		source := contentIndex[entry.VolumeID]
		logger.Debug("Parsing highlight",
			slog.String("title", source.Title),
		)
		var createdAt string
		if entry.DateCreated == "" {
			logger.Warn("No date created for bookmark. Defaulting to date last modified.",
				slog.String("title", source.Title),
				slog.String("volume_id", entry.VolumeID),
			)
			if entry.DateModified == "" {
				logger.Warn("No date modified for bookmark. Default to current date.",
					slog.String("title", source.Title),
					slog.String("volume_id", entry.VolumeID),
				)
				createdAt = time.Now().Format("2006-01-02T15:04:05-07:00")
			} else {
				t, err := time.Parse("2006-01-02T15:04:05Z", entry.DateModified)
				if err != nil {
					logger.Error("Failed to parse a valid timestamp from date modified field",
						slog.String("error", err.Error()),
						slog.String("title", source.Title),
						slog.String("volume_id", entry.VolumeID),
						slog.String("date_modified", entry.DateModified),
					)
					return []Response{}, err
				}
				createdAt = t.Format("2006-01-02T15:04:05-07:00")
			}
		} else {
			t, err := time.Parse("2006-01-02T15:04:05.000", entry.DateCreated)
			if err != nil {
				logger.Error("Failed to parse a valid timestamp from date created field",
					slog.String("error", err.Error()),
					slog.String("title", source.Title),
					slog.String("volume_id", entry.VolumeID),
					slog.String("date_modified", entry.DateModified),
				)
				return []Response{}, err
			}
			createdAt = t.Format("2006-01-02T15:04:05-07:00")
		}
		text := NormaliseText(entry.Text)
		if entry.Annotation != "" && text == "" {
			// I feel like this state probably shouldn't be possible but we'll handle it anyway
			// since it's useful to surface annotations, regardless of highlights. We put a
			// glaring placeholder here because the text field is required by the Readwise API.
			text = "Placeholder for attached annotation"
		}
		if entry.Annotation == "" && text == "" {
			// This state should be impossible but stranger things have happened so worth a sanity check
			logger.Warn("Found an entry with neither highlighted text nor an annotation so skipping entry",
				slog.String("title", source.Title),
				slog.String("volume_id", entry.VolumeID),
			)
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
				logger.Warn("Failed to retrieve epub title. This is not a hard requirement so sending with a dummy title.",
					slog.String("error", err.Error()),
					slog.String("title", source.Title),
					slog.String("volume_id", entry.VolumeID),
				)
				goto sendhighlight
			}
			filename := path.Base(sourceFile.Path)
			logger.Debug("No source title. Constructing title from filename",
				slog.String("filename", filename),
			)
			source.Title = strings.TrimSuffix(filename, ".epub")
		}
	sendhighlight:
		highlightChunks := splitHighlight(text, MaxHighlightLen)
		for _, chunk := range highlightChunks {
			highlight := Highlight{
				Text:          chunk,
				Title:         source.Title,
				Author:        source.Attribution,
				SourceURL:     entry.VolumeID,
				SourceType:    SourceType,
				Category:      SourceCategory,
				Note:          entry.Annotation,
				HighlightedAt: createdAt,
			}
			currentBatch.Highlights = append(currentBatch.Highlights, highlight)
		}
		logger.Debug("Successfully compiled highlights for book",
			slog.String("title", source.Title),
			slog.String("volume_id", entry.VolumeID),
			slog.Int("chunks", len(highlightChunks)),
		)
	}
	payloads = append(payloads, currentBatch)
	logger.Info("Succcessfully parsed highlights",
		slog.Int("highlight_count", len(currentBatch.Highlights)),
		slog.Int("batch_count", len(payloads)),
	)
	return payloads, nil
}

func NormaliseText(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\n", " ")
	return s
}
