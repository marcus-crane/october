package main

import (
	"context"
	"fmt"
	"github.com/adrg/xdg"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

var (
	configFilename = "october/config.json"
)

// App struct
type App struct {
	ctx      context.Context
	settings *Settings
	logger   *zap.SugaredLogger

	KoboService *KoboService
}

// NewApp creates a new App application struct
func NewApp() (*App, error) {
	configPath, err := xdg.ConfigFile(configFilename)
	if err != nil {
		panic(err)
	}
	settings, err := loadSettings(configPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load settings")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialise settings")
	}
	logFile := fmt.Sprintf("october/logs/%s.log", time.Now().Format("2006-01-02"))
	logPath, err := xdg.DataFile(logFile)
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout", logPath}
	logger, err := config.Build()
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialise logger")
	}
	defer logger.Sync()
	sugaredLogger := logger.Sugar()
	app := &App{
		settings: settings,
		logger:   sugaredLogger,
	}
	app.KoboService = NewKoboService(settings, sugaredLogger)
	app.logger.Debugw("October is fully initialised")
	return app, nil
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	a.logger.Infow("Starting up October")
}

// domReady is called after the front-end dom has been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
	a.logger.Debugw("Frontend DOM is ready")
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	a.logger.Infow("Shutting down October")
}
