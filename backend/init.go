package backend

import (
	"context"
	"errors"
	"fmt"

	"github.com/pgaskin/koboutils/v2/kobo"
	"github.com/rs/zerolog/log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Backend struct {
	SelectedKobo   Kobo
	ConnectedKobos map[string]Kobo
	RuntimeContext *context.Context
	Settings       *Settings
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
	}
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
