package main

import (
	"context"
	"github.com/marcus-crane/october/backend"
	log "github.com/sirupsen/logrus"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) domReady(ctx context.Context) {
	a.ctx = ctx
	backend.StartLogger()
	log.WithContext(ctx).Info("Logger should be initialised now")
	log.WithContext(ctx).Info("Backend is about to start up")
}

func (a *App) shutdown(ctx context.Context) {
	log.WithContext(ctx).Info("Shutting down. Goodbye!")
	backend.CloseLogFile()
}
