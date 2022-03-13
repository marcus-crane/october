package readwise

import "github.com/marcus-crane/october/pkg/device"

func BuildPayload(bookmarks []device.Bookmark, contentIndex []map[string]device.Content) (Response, error) {
	var payload Response
	return payload, nil
}

//
//func (k *KoboService) BuildReadwisePayload() ([]device.Highlight, error) {
//  content, err := k.ListDeviceContent()
//  if err != nil {
//    return nil, err
//  }
//  contentIndex := k.BuildContentIndex(content)
//  bookmarks, err := k.ListDeviceBookmarks()
//  if err != nil {
//    return nil, err
//  }
//  var highlights []device.Highlight
//  logger.Log.Infow(fmt.Sprintf("Starting to build Readwise payload out of %d bookmarks", len(bookmarks)))
//  for _, entry := range bookmarks {
//    source := contentIndex[entry.VolumeID]
//    t, err := time.Parse("2006-01-02T15:04:05.000", entry.DateCreated)
//    if err != nil {
//      logger.Log.Errorw(fmt.Sprintf("Failed to parse timestamp %s from bookmark", entry.DateCreated), "bookmark", entry)
//      return nil, err
//    }
//    createdAt := t.Format("2006-01-02T15:04:05-07:00")
//    text := k.NormaliseText(entry.Text)
//    if entry.Annotation != "" && text == "" {
//      text = "Placeholder for attached annotation"
//    }
//    if entry.Annotation == "" && text == "" {
//      logger.Log.Infow("Found an entry with no annotation of text so skipping to next item", "source", source, "bookmark", entry)
//      fmt.Printf("Ignoring entry from %s", source.Title)
//      continue
//    }
//    if source.Title == "" {
//      sourceFile, err := url.Parse(entry.VolumeID)
//      if err != nil {
//        logger.Log.Errorw("No title. Fallback of using filename failed. Not required so will send with no title.", "source", source, "bookmark", entry)
//        continue
//      }
//      filename := path.Base(sourceFile.Path)
//      logger.Log.Debugw(fmt.Sprintf("No source title. Constructing title from filename: %s", filename))
//      source.Title = strings.TrimSuffix(filename, ".epub")
//    }
//    highlight := device.Highlight{
//      Text:          text,
//      Title:         source.Title,
//      Author:        source.Attribution,
//      SourceType:    "OctoberForKobo",
//      Category:      "books",
//      Note:          entry.Annotation,
//      HighlightedAt: createdAt,
//    }
//    logger.Log.Debugw("Succesfully built highlight", "highlight", highlight)
//    highlights = append(highlights, highlight)
//  }
//  logger.Log.Infow(fmt.Sprintf("Successfully parsed %d highlights", len(highlights)))
//  return highlights, nil
//}
