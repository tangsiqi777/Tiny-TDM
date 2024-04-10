package main

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"tinytdm/backend/service"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	connectionStorageService := service.GetOneConnectionStorageService()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "tiny-tdm",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			connectionStorageService.Startup(ctx)
			app.startup(ctx)
		},
		Bind: []interface{}{
			connectionStorageService,
			app,
		},
		Windows: &windows.Options{

			//WebviewIsTransparent:              true, //网页透明
			//WindowIsTranslucent:               true, //窗口透明
			DisableFramelessWindowDecorations: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
