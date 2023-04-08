package main

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"net/http"

	"github.com/lucaono13/savetracker/backend"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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
	backend.StartBackend()
	// backend.StartLogger()
}

func (a *App) domReady(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) bool {
	backend.Logger.Info().Msg("Shutting down...")
	backend.CloseDB()
	// fmt.Println("written")
	backend.EndLogging()

	return false
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) RetrieveSaves() []backend.Save {

	return backend.GetSaves()
}

func (a *App) AddNewSave(saveName string, managerName string, gameVersion int) int {
	newSave := backend.Save{
		SaveName:    saveName,
		ManagerName: managerName,
		GameVersion: gameVersion,
		SaveID:      backend.NullInt64{NullInt64: sql.NullInt64{Valid: false}},
		SaveImage:   backend.NullString{NullString: sql.NullString{Valid: false}},
	}
	// addedID, err := backend.AddSave(saveName, managerName, gameVersion)
	// if err != nil {
	// 	return 0
	// }
	addedID, err := backend.AddSave(newSave)
	if err != nil {
		return 0
	}
	return int(addedID)
}

func Log(msg string) {
	backend.Logger.Info().Timestamp().Msg(msg)
}

func (a *App) SingleSave(id int) backend.Save {
	return backend.GetSingleSave(id)
}

func (a *App) SingleImage(id int) string {
	return backend.GetSaveImage(id)
}

func (a *App) GetImage(path string) string {
	path = strings.TrimSpace(path)
	bytes, err := os.ReadFile(path)
	if err != nil {
		backend.Logger.Error().Timestamp().Msg(err.Error())
		return string(bytes)
	}
	var base64Encoding string
	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += base64.StdEncoding.EncodeToString(bytes)
	return base64Encoding
}

func (a *App) UploadSaveImage(id int) string {
	// Open file dialog for just images
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Save Image",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg;*.jpeg)",
				Pattern:     "*.png;*.jpg;*.jpeg",
			},
		},
	})
	if err != nil {
		backend.Logger.Error().Timestamp().Msg(err.Error())
		return ""
	}
	// Have backend deal with the new image
	return backend.NewImage(id, file)

}

func SelectExportedFile(ctx context.Context, exported string) string {
	file, err := runtime.OpenFileDialog(ctx, runtime.OpenDialogOptions{
		Title: "Select " + exported + " File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "HTML File (*.html)",
				Pattern:     "*.html",
			},
		},
	})
	if err != nil {
		backend.Logger.Error().Timestamp().Msg(err.Error())
		return ""
	}
	return file
}

// Selecting the files exported from the game.
func (a *App) SelectSquadFile() string {
	return SelectExportedFile(a.ctx, "Squad")
}

func (a *App) SelectScheduleFile() string {
	return SelectExportedFile(a.ctx, "Schedule")
}

func (a *App) SelectTransfersFile() string {
	return SelectExportedFile(a.ctx, "Transfers")
}
