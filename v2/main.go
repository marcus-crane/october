package main

import (
	"fmt"
	"log"
	"sync"

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

	connection := kobo.NewKobo("/Users/marcus/Desktop/Kobo", "/Users/marcus/Desktop/Kobo/.kobo/KoboReader.sqlite")
	if err := connection.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %+v", err)
	}
	defer connection.Disconnect()
	numBookmarks, err := kobo.CountBookmarks(&connection)
	if err != nil {
		log.Fatalf("Failed to count bookmarks")
	}
	log.Printf("Your device has %d bookmarks", numBookmarks)
	numContent, err := kobo.CountContent(&connection)
	if err != nil {
		log.Fatalf("Failed to count content")
	}
	log.Printf("Your device has %d non-unique pieces of content", numContent)
	volumes, err := kobo.QueryDistinctVolumes(&connection)
	if err != nil {
		log.Fatalf("Failed to check volumes")
	}
	log.Printf("Detected %d unique books containing highlights and notes", len(volumes))
	var wg sync.WaitGroup
	for _, volume := range volumes {
		wg.Add(1)
		go func(connection kobo.Kobo, volume string) {
			defer wg.Done()
			fmt.Printf("Saw %s and tried to ping DB with result %+v\n", volume, connection.Ping())
		}(connection, volume)
	}
	wg.Wait()
}
