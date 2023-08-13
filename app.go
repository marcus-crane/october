package main

import (
	"context"

	"github.com/marcus-crane/october/backend"
	"github.com/sirupsen/logrus"
)

// App struct
type App struct {
	ctx      context.Context
	portable bool
}

// NewApp creates a new App application struct
func NewApp(portable bool) *App {
	return &App{
		portable: portable,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) domReady(ctx context.Context) {
	a.ctx = ctx
	backend.StartLogger(a.portable)
	logrus.WithContext(ctx).Info("Logger should be initialised now")
	logrus.WithContext(ctx).Info("Backend is about to start up")
}

func (a *App) shutdown(ctx context.Context) {
	logrus.WithContext(ctx).Info("Shutting down. Goodbye!")
	backend.CloseLogFile()
}
