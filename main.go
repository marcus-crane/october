package main

import (
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options/mac"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

var version = "DEV"

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "october",
		Width:     1366,
		Height:    768,
		Assets:    assets,
		LogLevel:  logger.DEBUG,
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
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
