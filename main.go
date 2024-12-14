package main

import (
	"embed"
	"fmt"
	slog "log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/marcus-crane/october/backend"
	"github.com/marcus-crane/october/cli"
	"github.com/wailsapp/wails/v2"
	wlogger "github.com/wailsapp/wails/v2/pkg/logger"
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

var loglevel = slog.LevelDebug

// Builds with this set to true cause files to be created
// in the same directory as the running executable
var portablebuild = "false"

func main() {
	isPortable := false
	isPortable, _ = strconv.ParseBool(portablebuild)

	logger, err := backend.StartLogger(isPortable, loglevel)
	if err != nil {
		panic("Failed to set up logger")
	}

	usr_cache_dir, _ := os.UserCacheDir()
	usr_config_dir, _ := os.UserConfigDir()

	logger.Info("Initialising October",
		slog.String("version", version),
		slog.String("loglevel", loglevel.String()),
		slog.Bool("portable", isPortable),
		slog.String("user_cache_dir", usr_cache_dir),
		slog.String("user_config_dir", usr_config_dir),
	)

	if cli.IsCLIInvokedExplicitly(os.Args) {
		logger.Info("CLI command invoked",
			slog.String("args", strings.Join(os.Args, " ")),
		)
		cli.Invoke(isPortable, version, logger)
		return
	}

	// Create an instance of the app structure
	app := NewApp(isPortable, logger)

	backend := backend.StartBackend(&app.ctx, version, isPortable, logger)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "October",
		Width:  1366,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		LogLevel:           wlogger.DEBUG,
		LogLevelProduction: wlogger.DEBUG,
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
		logger.Error("Wails runtime exited",
			slog.String("error", err.Error()),
		)
	}

	// TODO: Close file
}
