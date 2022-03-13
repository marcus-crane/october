package readwise

import (
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/marcus-crane/october/pkg/device"
	"github.com/marcus-crane/october/pkg/logger"
)

var (
	sourceCategory = "books"
	sourceType     = "OctoberForKobo"
)

func BuildPayload(bookmarks []device.Bookmark, contentIndex map[string]device.Content) (Response, error) {
	var payload Response
	for _, entry := range bookmarks {
		source := contentIndex[entry.VolumeID]
		t, err := time.Parse("2006-01-02T15:04:05.000", entry.DateCreated)
		if err != nil {
			logger.Log.Errorw(fmt.Sprintf("Failed to parse timestamp %s from bookmark", entry.DateCreated), "bookmark", entry)
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
			logger.Log.Infow("Found an entry with neither highlighted text nor an annotation so skipping to next item", "source", source, "bookmark", entry)
			fmt.Printf("Ignoring entry from %s", source.Title)
			continue
		}
		if source.Title == "" {
			// While Kepubs have a title in the Kobo database, the same can't be guaranteed for epubs at all.
			// In that event, we just fall back to using the filename
			sourceFile, err := url.Parse(entry.VolumeID)
			if err != nil {
				// While extremely unlikely, we should handle the case where a VolumeID doesn't have a suffix. This condition is only
				// triggered for completely busted names such as control codes given url.Parse will happen take URLs without a protocol
				// or even just arbitrary strings. That reminds me though:
				// TODO: Test exports with non-epub files
				logger.Log.Errorw("Failed to retrieve epub title. This is not a hard requirement so will send a payload with a dummy title.", "source", source, "bookmark", entry)
				source.Title = "Unknown Book"
				source.Attribution = "Unknown Author"
				goto sendhighlight
			}
			filename := path.Base(sourceFile.Path)
			logger.Log.Debugw(fmt.Sprintf("No source title. Constructing title from filename: %s", filename))
			source.Title = strings.TrimSuffix(filename, ".epub")
		}
	sendhighlight:
		highlight := device.Highlight{
			Text:          text,
			Title:         source.Title,
			Author:        source.Attribution,
			SourceType:    sourceType,
			Category:      sourceCategory,
			Note:          entry.Annotation,
			HighlightedAt: createdAt,
		}
		logger.Log.Debugw("Succesfully built highlight", "highlight", highlight)
		payload.Highlights = append(payload.Highlights, highlight)
	}
	logger.Log.Infow(fmt.Sprintf("Successfully parsed %d highlights", len(payload.Highlights)))
	return payload, nil
}
