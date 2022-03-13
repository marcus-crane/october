package readwise

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcus-crane/october/pkg/device"
	"github.com/marcus-crane/october/pkg/logger"
)

func TestMain(m *testing.M) {
	logger.Init()
	code := m.Run()
	os.Exit(code)
}

func TestBuildPayload_NoBookmarks(t *testing.T) {
	var expected Response
	var contentIndex map[string]device.Content
	var bookmarks []device.Bookmark
	var actual, _ = BuildPayload(bookmarks, contentIndex)
	assert.Equal(t, expected, actual)
}

func TestBuildPayload_BookmarksPresent(t *testing.T) {
	highlights := []device.Highlight{{
		Text:          "Hello World",
		Title:         "A Book",
		Author:        "Computer",
		SourceType:    sourceType,
		Category:      sourceCategory,
		Note:          "Making a note here",
		HighlightedAt: "2006-01-02T15:04:05+00:00",
	}}
	expected := Response{Highlights: highlights}
	contentIndex := map[string]device.Content{"abc123": {ContentID: "abc123", Title: "A Book", Attribution: "Computer"}}
	bookmarks := []device.Bookmark{
		{VolumeID: "abc123", Text: "Hello World", DateCreated: "2006-01-02T15:04:05.000", Annotation: "Making a note here"},
	}
	var actual, _ = BuildPayload(bookmarks, contentIndex)
	assert.Equal(t, expected, actual)
}
