package main

import (
	"embed"
	"fmt"

	"github.com/marcus-crane/october/backend"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

var version = "DEV"

func main() {
	// Create an instance of the app structure
	app := NewApp()

	fmt.Println(&app.ctx)

	backend := backend.StartBackend(&app.ctx)

	// Create application with options
	err := wails.Run(&options.App{
		Title:      "october",
		Width:      1366,
		Height:     768,
		Assets:     assets,
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
			backend,
		},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarDefault(),
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "October",
				Message: fmt.Sprintf("%s\nA small Wails application for retrieving Kobo highlights", version),
				Icon:    icon,
			},
		},
	})

	if err != nil {
		println("Error:", err)
	}
}
