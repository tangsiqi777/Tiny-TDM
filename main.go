package main

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"runtime"
	"tinytdm/backend/service"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the sqlService structure
	sqlService := service.NewRestSqlService()
	restSqlService := service.NewRestSqlService()
	connectionStorageService := service.GetOneConnectionStorageService()

	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		//runtime.Quit(sqlService.ctx)
	})

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "tiny-tdm",
		Width:     1024,
		Height:    768,
		Frameless: runtime.GOOS != "darwin",
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			connectionStorageService.Startup(ctx)
			sqlService.Startup(ctx)
		},
		Bind: []interface{}{
			connectionStorageService,
			sqlService,
			restSqlService,
		},
		Menu: AppMenu,
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
