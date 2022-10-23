package main

import (
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

var version = "DEV"

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create app menus
	// appMenu := menu.NewMenu()

	// // TODO: Set only on Darwin
	// appMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut

	// helpMenu := appMenu.AddSubmenu("Help")
	// helpMenu.AddText("Report Issue", nil, func(_ *menu.CallbackData) {
	// 	runtime.BrowserOpenURL(app.ctx, "https://github.com/marcus-crane/october/issues/new")
	// })
	// helpMenu.AddText("Documentation", nil, func(_ *menu.CallbackData) {
	// 	runtime.BrowserOpenURL(app.ctx, "https://october.utf9k.net/prerequisites/")
	// })

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "October",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
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
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
