package main

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/marcus-crane/october/backend"
	"github.com/marcus-crane/october/pkg/logger"
)

var (
	configFilename = "october/config.json"
)

// App struct
type App struct {
	ctx      context.Context
	Settings backend.Settings

	KoboService *KoboService
}

// NewApp creates a new App application struct
func NewApp() (*App, error) {
	loadedSettings, err := backend.LoadSettings()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load settings")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialise settings")
	}
	app := &App{
		Settings: loadedSettings,
	}
	app.KoboService = NewKoboService(loadedSettings)
	logger.Log.Debugw("October is fully initialised")
	return app, nil
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	backend.ConfigureLogger()
	log.Debug().Msg("Logger should be initialised now")
	log.Info().Msg("Backend is about to start up")

	a.KoboService.SetContext(ctx)
}

// domReady is called after the front-end dom has been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
	log.Debug().Msg("Frontend DOM is ready")
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	log.Info().Msg("Shutting down October")
}
