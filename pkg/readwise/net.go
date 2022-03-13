package readwise

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/marcus-crane/october/pkg/device"
	"github.com/pkg/errors"

	"github.com/marcus-crane/october/pkg/logger"
)

var (
	highlightsEndpoint = "https://readwise.io/api/v2/highlights/"
)

func SendBookmarksToReadwise(bookmarks []device.Highlight, token string) (int, error) {
	payload := Response{
		Highlights: bookmarks,
	}
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
	logger.Log.Infow(fmt.Sprintf("Successfully sent %d bookmarks to Readwise", len(bookmarks)))
	return len(bookmarks), nil
}
