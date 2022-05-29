package readwise

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/marcus-crane/october/backend"
)

func SendBookmarks(payload Response, token string) (int, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Token %s", token)).
		SetHeader("User-Agent", backend.UserAgent).
		SetBody(payload).
		Post(backend.HighlightsEndpoint)
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

func CheckTokenValidity(token string) error {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Token %s", token)).
		SetHeader("User-Agent", backend.UserAgent).
		Get(backend.AuthEndpoint)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 204 {
		log.Error().
			Int("status", resp.StatusCode()).
			Str("response", string(resp.Body())).
			Msg("Received a non-204 response from Readwise")
		return err
	}
	log.Info().Msg("Successfully authenticated with the Readwise API")
	return nil
}

func RetrieveUploadedBooks(token string) (BookListResponse, error) {
	bookList := BookListResponse{}
	headers := map[string][]string{
		"Authorization": {fmt.Sprintf("Token %s", token)},
		"User-Agent":    {backend.UserAgent},
	}
	client := http.Client{}
	remoteURL, err := url.Parse(backend.BooksEndpoint)
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

func UploadCover(encodedCover string, bookId int, token string) error {
	body := map[string]interface{}{
		"cover_image": encodedCover,
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Token %s", token)).
		SetHeader("Content-Type", "application/json").
		SetHeader("User-Agent", backend.UserAgent).
		SetBody(body).
		Patch(fmt.Sprintf(backend.CoverEndpoint, bookId))
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
