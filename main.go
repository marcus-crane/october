package main

import (
	"embed"
	"fmt"

	"github.com/marcus-crane/october/backend"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

var version = "DEV"

func main() {
	// Create an instance of the app structure
	app := NewApp()

	backend := backend.StartBackend(&app.ctx)

	// Create application with options
	err := wails.Run(&options.App{
		Title:      "October",
		Width:      1024,
		Height:     768,
		Assets:     assets,
		LogLevel:   logger.DEBUG,
		OnStartup:  app.startup,
		OnDomReady: app.domReady,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
			backend,
			backend.Bookmark,
			backend.Content,
			backend.Kobo,
			backend.Readwise,
			backend.Settings,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
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
