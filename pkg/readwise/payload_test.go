package readwise

import (
	"testing"

	"github.com/marcus-crane/october/pkg/device"
	"github.com/stretchr/testify/assert"
)

func TestBuildPayload(t *testing.T) {
	var expected Response
	contentIndex := []map[string]device.Content{{"abc123": {ContentID: "abc123"}}}
	var bookmarks []device.Bookmark
	var actual, _ = BuildPayload(bookmarks, contentIndex)
	assert.Equal(t, expected, actual)
}
