package main

import (
	"embed"
	"fmt"
	"os"
	"strconv"

	"github.com/marcus-crane/october/backend"
	"github.com/marcus-crane/october/cli"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

var version = "DEV"

// Builds with this set to true cause files to be created
// in the same directory as the running executable
var portablebuild = "false"

func main() {
	isPortable := false
	isPortable, _ = strconv.ParseBool(portablebuild)

	skipCli := version == "DEV" || len(os.Args) == 2 && os.Args[1] == "launch"

	if cli.IsInvokedFromTerminal() && !skipCli {
		cli.Invoke(isPortable, version)
		return
	}

	// Create an instance of the app structure
	app := NewApp(isPortable)

	backend := backend.StartBackend(&app.ctx, version, isPortable)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "October",
		Width:  1366,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.DEBUG,
		OnStartup:          app.startup,
		OnDomReady:         app.domReady,
		OnShutdown:         app.shutdown,
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
		Linux: &linux.Options{
			Icon:                icon,
			WindowIsTranslucent: false,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
