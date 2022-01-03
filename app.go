package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

var (
	configFilename = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "octowise", "octowise.config.json")
)

// App struct
type App struct {
	ctx      context.Context
	settings *settings

	KoboService *KoboService
}

// NewApp creates a new App application struct
func NewApp() (*App, error) {
	settings, err := loadSettings(configFilename)
	if err != nil {
		return nil, errors.Wrap(err, "loadSettings")
	}
	app := &App{
		settings: settings,
	}
	app.KoboService = NewKoboService(app.ctx)
	return app, nil
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
