package main

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/marcus-crane/october/backend"
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
	backend.ConfigureLogger()
	log.Debug().Msg("Logger should be initialised now")
	log.Info().Msg("Backend is about to start up")
}

func (a *App) shutdown(ctx context.Context) {
	log.Info().Msg("Shutting down. Goodbye!")
}
