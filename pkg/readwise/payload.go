package readwise

import (
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/marcus-crane/october/pkg/device"
)

var (
	sourceCategory = "books"
	sourceType     = "OctoberForKobo"
)

func BuildPayload(bookmarks []device.Bookmark, contentIndex map[string]device.Content) (Response, error) {
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
			SourceType:    sourceType,
			Category:      sourceCategory,
			Note:          entry.Annotation,
			HighlightedAt: createdAt,
		}
		log.Debug().Interface("highlight", highlight).Msg("Successfully built highlights")
		payload.Highlights = append(payload.Highlights, highlight)
	}
	log.Info().Int("highlight_count", len(payload.Highlights)).Msg("Successfully parsed highlights")
	return payload, nil
}
