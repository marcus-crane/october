package readwise

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"

	"github.com/marcus-crane/october/pkg/logger"
)

var (
	authEndpoint       = "https://readwise.io/api/v2/auth/"
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
