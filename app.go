package main

import (
	"context"

	"github.com/adrg/xdg"
	"github.com/marcus-crane/october/pkg/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/marcus-crane/october/pkg/settings"
)

var (
	configFilename = "october/config.json"
)

// App struct
type App struct {
	ctx      context.Context
	settings *settings.Settings
	logger   *zap.SugaredLogger

	KoboService *KoboService
}

// NewApp creates a new App application struct
func NewApp() (*App, error) {
	configPath, err := xdg.ConfigFile(configFilename)
	if err != nil {
		panic(err)
	}
	loadedSettings, err := settings.LoadSettings(configPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load settings")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialise settings")
	}
	logger.Init()
	app := &App{
		settings: loadedSettings,
	}
	app.KoboService = NewKoboService(loadedSettings)
	logger.Log.Debugw("October is fully initialised")
	return app, nil
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	a.KoboService.SetContext(ctx)
	logger.Log.Infow("Starting up October")
}

// domReady is called after the front-end dom has been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
	logger.Log.Debugw("Frontend DOM is ready")
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	logger.Log.Infow("Shutting down October")
}
