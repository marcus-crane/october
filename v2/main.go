package main

import (
	"fmt"
	"log"

	"github.com/marcus-crane/october/v2/pkg/epub"
	"github.com/marcus-crane/october/v2/pkg/kobo"
)

func main() {
	root, err := epub.LoadEpub("kobodbgen/epubs/source/james-hogg_the-private-memoirs-and-confessions-of-a-justified-sinner.epub")
	if err != nil {
		log.Fatalf("Failed to load epub: %+v", err)
	}
	// Spines are the correct order where manifests may be out of order
	// May need to check the manifest for covers
	for i, entry := range root.Spine.Itemrefs {
		idx := i + 1 // We increment by 1 as we don't trust external systems to not drop zero padding
		fmt.Printf("%d - %s\n", idx, entry.HREF)
	}

	connection := kobo.NewDirectConnection("/Users/marcus/Desktop/Kobo", "/Users/marcus/Desktop/Kobo/.kobo/KoboReader.sqlite")
	if err := connection.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %+v", err)
	}
	numBookmarks, err := kobo.BookmarksCount(&connection)
	if err != nil {
		log.Fatalf("Failed to count bookmarks")
	}
	log.Printf("Your device has %d bookmarks", numBookmarks)
	numContent, err := kobo.ContentCount(&connection)
	if err != nil {
		log.Fatalf("Failed to count content")
	}
	log.Printf("Your device has %d pieces of content", numContent)
}
