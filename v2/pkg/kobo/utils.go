package kobo

import (
	"fmt"
	"strings"
)

// getRelativeKoboPath takes the path of a file stored in a Kobo device and returns it stripped of its pathing prefixes.
// VolumeID and ContentID is where it's particularly used to retrieve books relative to the root of the filesystem
// which may either be literally on a Kobo or reading a folder from disk (such as if you're operating on files from a local backup)
func getRelativeKoboPath(path string) string {
	path = strings.ReplaceAll(path, "file://", "")
	path = strings.ReplaceAll(path, "/mnt/onboard/", "")
	return path
}

// trimContentFileName removes trailing parts from content paths, returning just
// the name of the content
//
// For example: "Vend/Technology at Vend - Vend.epub#(2)OEBPS/_projects_work.xhtml" becomes "(2)OEBPS/_projects_work.xhtml"
//
// Generally speaking, the following delimiters seem to be used for the following extensions:
//
// .kepub.epub => !! | .epub => #
func trimContentFileName(path string) string {
	// TODO: .fxl.kepub.epub is a format as well, supposedly for supporting comics?
	// https://github.com/kobolabs/epub-spec#image-based-fxl-reader
	// https://github.com/kobolabs/epub-spec#sideloading-for-testing-purposes
	// https://github.com/kobolabs/epub-spec#fixed-layout-fxl-support
	if strings.Contains(path, ".kepub.epub") {
		return strings.Split(path, ".kepub.epub!!")[1]
	}
	if strings.Contains(path, ".epub") {
		return strings.Split(path, ".epub#")[1]
	}
	// We'll just return whatever unknown format this is and surface an error
	// to the user down the line
	return path
}

// formatUsualDbPath just takes a mount path and returns the usual location
// of where you would find the underlying Kobo database
func formatUsualDbPath(mountPath string) string {
	return fmt.Sprintf("%s/.kobo/KoboReader.sqlite", mountPath)
}
