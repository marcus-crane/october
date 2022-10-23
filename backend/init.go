package backend

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/pgaskin/koboutils/v2/kobo"
	"github.com/rs/zerolog/log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Backend struct {
	SelectedKobo   Kobo
	ConnectedKobos map[string]Kobo
	RuntimeContext *context.Context
	Settings       *Settings
	Readwise       *Readwise
	Kobo           *Kobo
	Content        *Content
	Bookmark       *Bookmark
}

func StartBackend(ctx *context.Context) *Backend {
	settings, err := LoadSettings()
	if err != nil {
		log.Error().Msg("Failed to load settings")
	}
	return &Backend{
		SelectedKobo:   Kobo{},
		ConnectedKobos: map[string]Kobo{},
		RuntimeContext: ctx,
		Settings:       settings,
		Readwise:       &Readwise{},
		Kobo:           &Kobo{},
		Content:        &Content{},
		Bookmark:       &Bookmark{},
	}
}

func (b *Backend) GetSettings() *Settings {
	return b.Settings
}

func (b *Backend) GetContent() *Content {
	return b.Content
}

func (b *Backend) GetBookmark() *Bookmark {
	return b.Bookmark
}

func (b *Backend) DetectKobos() []Kobo {
	connectedKobos, err := kobo.Find()
	log.Info().Msg(fmt.Sprintf("Kobos found: %d", len(connectedKobos)))
	if err != nil {
		panic(err)
	}
	kobos := GetKoboMetadata(connectedKobos)
	for _, kb := range kobos {
		b.ConnectedKobos[kb.MntPath] = kb
	}
	return kobos
}

func (b *Backend) GetSelectedKobo() Kobo {
	return b.SelectedKobo
}

func (b *Backend) SelectKobo(devicePath string) error {
	if val, ok := b.ConnectedKobos[devicePath]; ok {
		b.SelectedKobo = val
	} else {
		b.SelectedKobo = Kobo{
			Name:       "Local Database",
			Storage:    0,
			DisplayPPI: 0,
			MntPath:    devicePath,
			DbPath:     devicePath,
		}
	}
	if err := OpenConnection(b.SelectedKobo.DbPath); err != nil {
		return err
	}
	return nil
}

func (b *Backend) PromptForLocalDBPath() error {
	selectedFile, err := runtime.OpenFileDialog(*b.RuntimeContext, runtime.OpenDialogOptions{
		Title: "Select local Kobo database",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "sqlite (*.sqlite;*.sqlite3)",
				Pattern:     "*.sqlite;*.sqlite3",
			},
		},
	})
	if err != nil {
		return err
	}
	// The user has cancelled the dialog so we just do nothing
	if selectedFile == "" {
		return errors.New("canceled selection")
	}
	return b.SelectKobo(selectedFile)
}

func (b *Backend) ForwardToReadwise() (int, error) {
	content, err := b.Kobo.ListDeviceContent()
	if err != nil {
		return 0, err
	}
	contentIndex := b.Kobo.BuildContentIndex(content)
	bookmarks, err := b.Kobo.ListDeviceBookmarks()
	if err != nil {
		return 0, err
	}
	payload, err := BuildPayload(bookmarks, contentIndex)
	if err != nil {
		return 0, err
	}
	numUploads, err := b.Readwise.SendBookmarks(payload, b.Settings.ReadwiseToken)
	if err != nil {
		return 0, err
	}
	uploadedBooks, err := b.Readwise.RetrieveUploadedBooks(b.Settings.ReadwiseToken)
	if err != nil {
		return numUploads, fmt.Errorf(fmt.Sprintf("Successfully uploaded %d bookmarks but failed to upload covers", numUploads))
	}
	if b.Settings.UploadCovers {
		for _, book := range uploadedBooks.Results {
			// We don't want to overwrite user uploaded covers or covers already present
			if !strings.Contains(book.CoverURL, "uploaded_book_covers") {
				coverID := kobo.ContentIDToImageID(book.SourceURL)
				coverPath := kobo.CoverTypeLibFull.GeneratePath(false, coverID)
				absCoverPath := path.Join(b.SelectedKobo.MntPath, "/", coverPath)
				coverBytes, err := os.ReadFile(absCoverPath)
				if err != nil {
					log.Error().Str("cover", book.SourceURL).Str("location", absCoverPath).Msg("Failed to load cover")
				}
				var base64Encoding string
				mimeType := http.DetectContentType(coverBytes)
				switch mimeType {
				case "image/jpeg":
					base64Encoding += "data:image/jpeg;base64,"
				case "image/png":
					base64Encoding += "data:image/png;base64,"
				}
				base64Encoding += base64.StdEncoding.EncodeToString(coverBytes)
				err = b.Readwise.UploadCover(base64Encoding, book.ID, b.Settings.ReadwiseToken)
				if err != nil {
					log.Error().Str("cover", book.SourceURL).Msg("Failed to load cover")
				}
				log.Info().Str("cover", book.SourceURL).Msg("Successfully uploaded cover")
			}
		}
	}
	return numUploads, nil
}
