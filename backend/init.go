package backend

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"github.com/adrg/xdg"

	"github.com/pgaskin/koboutils/v2/kobo"
	log "github.com/sirupsen/logrus"

	updater "github.com/marcus-crane/october/internal/updater"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
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
	version        string
}

func StartBackend(ctx *context.Context, version string) *Backend {
	settings, err := LoadSettings()
	if err != nil {
		log.WithContext(*ctx).WithError(err).Error("Failed to load settings")
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
		version:        version,
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

func (b *Backend) GetPlainSystemDetails() string {
	return fmt.Sprintf("%s (%s %s)", b.version, runtime.GOOS, runtime.GOARCH)
}

func (b *Backend) FormatSystemDetails() string {
	return fmt.Sprintf("<details><summary>System Details</summary><ul><li>Version: %s</li><li>Platform: %s</li><li>Architecture: %s</li></details>", b.version, runtime.GOOS, runtime.GOARCH)
}

func (b *Backend) NavigateExplorerToLogLocation() {
	var explorerCommand string
	if runtime.GOOS == "windows" {
		explorerCommand = "explorer.exe"
	}
	if runtime.GOOS == "darwin" {
		explorerCommand = "open"
	}
	if runtime.GOOS == "linux" {
		explorerCommand = "xdg-open"
	}
	logLocation, err := xdg.DataFile("october/logs")
	if err != nil {
		log.WithError(err).Error("Failed to determine XDG data location for opening log location in explorer")
	}
	// We will always get an error because the file explorer doesn't exit so it is unable to
	// return a 0 successful exit code until y'know, the user exits the window
	_ = exec.Command(explorerCommand, logLocation).Run()
}

func (b *Backend) DetectKobos() []Kobo {
	connectedKobos, err := kobo.Find()
	log.WithField("kobos_found", len(connectedKobos)).Info("Detected one or more Kobos")
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

func (b *Backend) CheckForUpdate() (bool, string) {
	return updater.CheckForNewerVersion(b.version)
}

func (b *Backend) PerformUpdate() (bool, error) {
	if runtime.GOOS == "darwin" {
		return updater.PerformUpdateDarwin(b.version)
	} else {
		return updater.PerformUpdate(b.version)
	}
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
	selectedFile, err := wailsRuntime.OpenFileDialog(*b.RuntimeContext, wailsRuntime.OpenDialogOptions{
		Title: "Select local Kobo database",
		Filters: []wailsRuntime.FileFilter{
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
					log.WithError(err).WithFields(log.Fields{"cover": book.SourceURL, "location": absCoverPath}).Warn("Failed to load cover. Carrying on")
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
					log.WithError(err).WithField("cover", book.SourceURL).Error("Failed to upload cover to Readwise")
				}
				log.WithField("cover", book.SourceURL).Debug("Successfully uploaded cover to Readwise")
			}
		}
	}
	return numUploads, nil
}
