package kobo

type Bookmark struct {
	BookmarkID               string  `db:"BookmarkID"`
	VolumeID                 string  `db:"VolumeID"`
	ContentID                string  `db:"ContentID"`
	StartContainerPath       string  `db:"StartContainerPath"`
	StartContainerChild      string  `db:"StartContainerChild"`
	StartContainerChildIndex int     `db:"StartContainerChildIndex"`
	StartOffset              int     `db:"StartOffset"`
	EndContainerPath         string  `db:"EndContainerPath"`
	EndContainerChildIndex   int     `db:"EndContainerChildIndex"`
	EndOffset                int     `db:"EndOffset"`
	Text                     string  `db:"Text"`
	Annotation               string  `db:"Annotation"`
	ExtraAnnotationData      string  `db:"ExtraAnnotationData"`
	DateCreated              string  `db:"DateCreated"`
	ChapterProgress          float64 `db:"ChapterProgress"`
	Hidden                   bool    `db:"Hidden"`
	Version                  string  `db:"Version"`
	DateModified             string  `db:"DateModified"`
	Creator                  string  `db:"Creator"`
	UUID                     string  `db:"UUID"`
	UserID                   string  `db:"UserID"`
	SyncTime                 string  `db:"SyncTime"`
	Published                bool    `db:"Published"`
	ContextString            string  `db:"ContextString"`
	Type                     string  `db:"Type"`
}

func CountBookmarks(kobo *Kobo) (int, error) {
	var count int
	if err := kobo.dbClient.Get(&count, "SELECT count(*) FROM Bookmark"); err != nil {
		return count, err
	}
	return count, nil
}

// QueryDistinctVolumes retrieves all books that contains highlights and reside (or did at one point)
// on the device itself. We exclude Kobo store books for simplicity where the path is a GUID which is
// not useful in the context of retrieving the underlying epub for scanning.
func QueryDistinctVolumes(kobo *Kobo) ([]string, error) {
	var contentWithBookmarks []string
	if err := kobo.dbClient.Select(
		&contentWithBookmarks,
		"SELECT DISTINCT VolumeID FROM Bookmark WHERE VolumeID LIKE '%file:///%' ORDER BY VolumeID;",
	); err != nil {
		return contentWithBookmarks, err
	}
	return contentWithBookmarks, nil
}
