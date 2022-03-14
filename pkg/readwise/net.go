package readwise

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"

	"github.com/marcus-crane/october/pkg/logger"
)

var (
	authEndpoint       = "https://readwise.io/api/v2/auth/"
	booksEndpoint      = "https://readwise.io/api/v2/books?page_size=1000&category=books&source=OctoberForKobo"
	coverEndpoint      = "https://readwise.io/api/v2/books/%d"
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
	headers := map[string][]string{
		"Authorization": {fmt.Sprintf("Token %s", token)},
		"User-Agent":    {"october/1.0.0 <https://github.com/marcus-crane/october>"},
	}
	client := http.Client{}
	remoteURL, err := url.Parse(booksEndpoint)
	if err != nil {
		logger.Log.Error("Failed to parse books URL", "error", err)
	}
	request := http.Request{
		Method: "GET",
		URL:    remoteURL,
		Header: headers,
	}
	res, err := client.Do(&request)
	defer res.Body.Close()
	if err != nil {
		logger.Log.Error(err)
		return bookList, err
	}
	b, err := httputil.DumpResponse(res, true)
	if err != nil {
		logger.Log.Error(err)
		return bookList, err
	}
	if res.StatusCode != 200 {
		logger.Log.Error("Received a non-200 response from Readwise", "status", res.StatusCode)
		return bookList, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Log.Error("Failed to parse response")
		panic(err)
	}
	err = json.Unmarshal(body, &bookList)
	if err != nil {
		logger.Log.Error("Failed to unmarshal response from Readwise", "status", res.StatusCode, "body", string(b))
		return bookList, err
	}
	logger.Log.Info(fmt.Sprintf("Successfully retrieved %d books from Readwise API", bookList.Count))
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
		SetHeader("User-Agent", "october/1.0.0 <https://github.com/marcus-crane/october>").
		SetBody(body).
		Patch(fmt.Sprintf(coverEndpoint, bookId))
	logger.Log.Info(string(resp.Body()))
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		logger.Log.Errorw("Received a non-200 response from Readwise", "status", resp.StatusCode(), "response", string(resp.Body()))
		return fmt.Errorf("failed to upload cover for book with id %s", bookId)
	}
	return nil
}
