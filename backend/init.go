package backend

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"github.com/pgaskin/koboutils/v2/kobo"

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
	logger         *slog.Logger
	version        string
	portable       bool
}

func StartBackend(ctx *context.Context, version string, portable bool, logger *slog.Logger) (*Backend, error) {
	settings, err := LoadSettings(portable, logger)
	logger.Info("Successfully parsed settings file",
		slog.String("path", settings.path),
		slog.Bool("upload_store_highlights", settings.UploadStoreHighlights),
		slog.Bool("upload_covers", settings.UploadCovers),
	)
	if err != nil {
		logger.Error("Failed to load settings",
			slog.String("error", err.Error()),
		)
		return &Backend{}, err
	}
	return &Backend{
		SelectedKobo:   Kobo{},
		ConnectedKobos: map[string]Kobo{},
		RuntimeContext: ctx,
		Settings:       settings,
		Readwise: &Readwise{
			UserAgent: fmt.Sprintf(UserAgentFmt, version),
		},
		Kobo:     &Kobo{},
		Content:  &Content{},
		Bookmark: &Bookmark{},
		logger:   logger,
		version:  version,
		portable: portable,
	}, nil
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
	onboardingComplete := false
	if b.Settings.ReadwiseToken != "" {
		onboardingComplete = true
	}
	return fmt.Sprintf("<details><summary>System Details</summary><ul><li>Version: %s</li><li>Platform: %s</li><li>Architecture: %s</li><li>Onboarding Complete: %t</li></details>", b.version, runtime.GOOS, runtime.GOARCH, onboardingComplete)
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
	b.logger.Info("Opening logs in system file explorer",
		slog.String("command", explorerCommand),
		slog.String("os", runtime.GOOS),
	)
	logLocation, err := LocateDataFile("october/logs", b.portable)
	if err != nil {
		b.logger.Error("Failed to determine XDG data location for opening log location in explorer",
			slog.String("error", err.Error()),
		)
	}
	b.logger.Debug("Executing command to open system file explorer",
		slog.String("command", explorerCommand),
		slog.String("os", runtime.GOOS),
	)
	// We will always get an error because the file explorer doesn't exit so it is unable to
	// return a 0 successful exit code until y'know, the user exits the window
	_ = exec.Command(explorerCommand, logLocation).Run()
}

func (b *Backend) DetectKobos() []Kobo {
	connectedKobos, err := kobo.Find()
	if err != nil {
		b.logger.Error("Failed to detect any connected Kobos")
		panic(err)
	}
	kobos := GetKoboMetadata(connectedKobos)
	b.logger.Info("Found one or more kobos",
		"count", len(kobos),
	)
	for _, kb := range kobos {
		b.logger.Info("Found connected device",
			slog.String("mount_path", kb.MntPath),
			slog.String("database_path", kb.DbPath),
			slog.String("name", kb.Name),
			slog.Int("display_ppi", kb.DisplayPPI),
			slog.Int("storage", kb.Storage),
		)
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
		b.logger.Info("No device found at path. Selecting local database",
			slog.String("device_path", devicePath),
		)
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
	highlightBreakdown := b.Kobo.CountDeviceBookmarks()
	slog.Info("Got highlight counts from device",
		slog.Int("highlight_count_sideload", int(highlightBreakdown.Sideloaded)),
		slog.Int("highlight_count_official", int(highlightBreakdown.Official)),
		slog.Int("highlight_count_total", int(highlightBreakdown.Total)),
	)
	if highlightBreakdown.Total == 0 {
		slog.Error("Tried to submit highlights when there are none on device.")
		return 0, fmt.Errorf("Your device doesn't seem to have any highlights so there is nothing left to sync.")
	}
	includeStoreBought := b.Settings.UploadStoreHighlights
	if !includeStoreBought && highlightBreakdown.Sideloaded == 0 {
		slog.Error("Tried to submit highlights with no sideloaded highlights + store-bought syncing disabled. Result is that no highlights would be fetched.")
		return 0, fmt.Errorf("You have disabled store-bought syncing but you don't have any sideloaded highlights either. This combination means there are no highlights left to be synced.")
	}
	content, err := b.Kobo.ListDeviceContent(includeStoreBought)
	if err != nil {
		slog.Error("Received an error trying to list content from device",
			slog.String("error", err.Error()),
		)
		return 0, err
	}
	contentIndex := b.Kobo.BuildContentIndex(content)
	bookmarks, err := b.Kobo.ListDeviceBookmarks(includeStoreBought)
	if err != nil {
		slog.Error("Received an error trying to list bookmarks from device",
			slog.String("error", err.Error()),
		)
		return 0, err
	}
	payload, err := BuildPayload(bookmarks, contentIndex)
	if err != nil {
		slog.Error("Received an error trying to build Readwise payload",
			slog.String("error", err.Error()),
		)
		return 0, err
	}
	numUploads, err := b.Readwise.SendBookmarks(payload, b.Settings.ReadwiseToken)
	if err != nil {
		slog.Error("Received an error trying to send bookmarks to Readwise",
			slog.String("error", err.Error()),
		)
		return 0, err
	}
	slog.Info("Successfully uploaded bookmarks to Readwise",
		slog.Int("payload_count", numUploads),
	)
	if b.Settings.UploadCovers {
		uploadedBooks, err := b.Readwise.RetrieveUploadedBooks(b.Settings.ReadwiseToken)
		if err != nil {
			slog.Error("Failed to retrieve uploaded titles from Readwise",
				slog.String("error", err.Error()),
			)
			return numUploads, fmt.Errorf("Successfully uploaded %d bookmarks", numUploads)
		}
		slog.Info("Retrieved uploaded books from Readwise for cover insertion",
			slog.Int("book_count", uploadedBooks.Count),
		)
		for _, book := range uploadedBooks.Results {
			slog.Info("Checking cover status for book",
				slog.Int("book_id", book.ID),
				slog.String("book_title", book.Title),
			)
			// We don't want to overwrite user uploaded covers or covers already present
			if !strings.Contains(book.CoverURL, "uploaded_book_covers") {
				coverID := kobo.ContentIDToImageID(book.SourceURL)
				coverPath := kobo.CoverTypeLibFull.GeneratePath(false, coverID)
				absCoverPath := path.Join(b.SelectedKobo.MntPath, "/", coverPath)
				coverBytes, err := os.ReadFile(absCoverPath)
				if err != nil {
					slog.Warn("Failed to load cover from disc. Skipping to next book.",
						slog.String("error", err.Error()),
						slog.String("cover_path", absCoverPath),
						slog.String("cover_id", book.SourceURL),
					)
					continue
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
					slog.Error("Failed to upload cover to Readwise. Skipping to next book.",
						slog.String("error", err.Error()),
						slog.String("cover_url", book.SourceURL),
					)
					continue
				}
				slog.Debug("Successfully uploaded cover to Readwise",
					slog.String("cover_url", book.SourceURL),
				)
			}
			slog.Info("Cover already exists for book. Skipping as we don't know if this was us prior or a user upload.")
		}
	}
	return numUploads, nil
}
