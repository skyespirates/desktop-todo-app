package main

import (
	"database/sql"
	"embed"
	"log"
	"skyes-app/internal/repository"
	"skyes-app/internal/storage"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	_ "modernc.org/sqlite"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

var DB *sql.DB

func main() {
	// Create an instance of the app structure

	db, err := storage.OpenDB("todo-app")
	if err != nil {
		log.Fatal(err)
	}

	err = storage.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTodoRepo(db)

	app := NewApp(repo)

	// Create application with options
	err = wails.Run(&options.App{
		Title:            "Skyes App",
		Width:            720,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		// Debug:            options.Debug{OpenInspectorOnStartup: true},
		Bind: []any{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
