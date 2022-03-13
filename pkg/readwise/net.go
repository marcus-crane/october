package readwise

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"

	"github.com/marcus-crane/october/pkg/logger"
)

var (
	authEndpoint       = "https://readwise.io/api/v2/auth/"
	booksEndpoint      = "https://readwise.io/api/v2/books?page_size=1000&category=books&source=OctoberForKobo"
	coverEndpoint      = "https://readwise.io/api/v2/books/%s"
	highlightsEndpoint = "https://readwise.io/api/v2/highlights/"
)

func SendBookmarks(payload Response, token string) (int, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Token %s", token)).
		SetHeader("User-Agent", "october/1.0.0 <https://github.com/marcus-crane/october>").
		SetBody(payload).
		Post(highlightsEndpoint)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Failed to send request to Readwise: code %d", resp.StatusCode()))
	}
	logger.Log.Debug(resp.Body())
	if resp.StatusCode() != 200 {
		logger.Log.Errorw("Received a non-200 response from Readwise", "status", resp.StatusCode(), "response", string(resp.Body()))
		return 0, errors.New(fmt.Sprintf("Received a non-200 status code from Readwise: code %d", resp.StatusCode()))
	}
	logger.Log.Infow(fmt.Sprintf("Successfully sent %d bookmarks to Readwise", len(payload.Highlights)))
	return len(payload.Highlights), nil
}

func CheckTokenValidity(token string) error {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Token %s", token)).
		SetHeader("User-Agent", "october/1.0.0 <https://github.com/marcus-crane/october>").
		Get(authEndpoint)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 204 {
		logger.Log.Error("Received a non-204 response from Readwise", "status", resp.StatusCode(), "response", string(resp.Body()))
		return err
	}
	logger.Log.Info("Successfully authenticated with the Readwise API")
	return nil
}

func RetrieveUploadedBooks(token string) (BookListResponse, error) {
	bookList := BookListResponse{}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Token %s", token)).
		SetHeader("User-Agent", "october/1.0.0 <https://github.com/marcus-crane/october>").
		Get(booksEndpoint)
	if err != nil {
		return bookList, err
	}
	if resp.StatusCode() != 204 {
		logger.Log.Error("Received a non-204 response from Readwise", "status", resp.StatusCode(), "response", string(resp.Body()))
		return bookList, err
	}
	defer resp.RawBody().Close()
	body, _ := ioutil.ReadAll(resp.RawBody())
	err = json.Unmarshal(body, &bookList)
	if err != nil {
		logger.Log.Error("Failed to unmarshal response from Readwise", "status", resp.StatusCode())
		return bookList, err
	}
	logger.Log.Info(fmt.Sprintf("Successfully retrieved %d books from Readwise API", bookList.Count))
	return bookList, nil
}

func UploadCover(encodedCover string, bookId string, token string) error {
	coverUpdate := CoverUpdate{Cover: encodedCover}
	cover, err := json.Marshal(coverUpdate)
	if err != nil {
		return err
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Token %s", token)).
		SetHeader("User-Agent", "october/1.0.0 <https://github.com/marcus-crane/october>").
		SetBody(cover).
		Patch(fmt.Sprintf(coverEndpoint, bookId))
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		logger.Log.Errorw("Received a non-200 response from Readwise", "status", resp.StatusCode(), "response", string(resp.Body()))
		return fmt.Errorf("failed to upload cover for book with id %s", bookId)
	}
	return nil
}
