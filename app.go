package main

import (
	"context"
	slog "log/slog"

	"github.com/marcus-crane/october/backend"
)

// App struct
type App struct {
	ctx      context.Context
	logger   *slog.Logger
	portable bool
}

// NewApp creates a new App application struct
func NewApp(portable bool, logger *slog.Logger) *App {
	logger.Debug("Initialising app struct")
	return &App{
		logger:   logger,
		portable: portable,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.logger.Debug("Calling app startup method")
	a.ctx = ctx
}

func (a *App) domReady(ctx context.Context) {
	a.logger.Debug("Calling app domReady method")
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
	a.logger.Debug("Calling app shutdown method")
	backend.CloseLogFile()
}
