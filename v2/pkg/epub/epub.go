package epub

import (
	"fmt"
	"os"

	"github.com/taylorskalyo/goreader/epub"
)

// NOTE TO SELF: Cover images can be retrieved like so:
// 1) Find meta name="cover" in OPF metadata and check the content
// 2) Find item entry with id set to same value as metadata content
// See https://github.com/kobolabs/epub-spec#cover-images
// Also covers may technically be SVGs so check the types and
// explicitly mention what is supported

// See https://github.com/kobolabs/epub-spec#for-epub3 for a list of nav fallbacks
// if OPF doesn't have the desired information

// LoadEpub takes a relative path to an epub file (aka VolumeID) and returns the content of the epub
// Further parsing is required to actually read files for example but it gives a full book spine
// among other useful metadata that may or may not be present in the database, particular with
// epubs (not kepubs) that have been loaded directly onto the device without using a tool like Calibre
func LoadEpub(path string) (*epub.Rootfile, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	rc, err := epub.OpenReader(path)
	if err != nil {
		return nil, err
	}
	if len(rc.Rootfiles) == 0 {
		return nil, fmt.Errorf("no root files found in epub")
	}
	return rc.Rootfiles[0], nil
}