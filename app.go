package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"io/ioutil"
	"net/http"

	"github.com/lucaono13/savetracker/backend"
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
	addedID, err := backend.AddSave(saveName, managerName, gameVersion)
	if err != nil {
		return 0
	}
	return addedID
}

func Log(msg string) {
	backend.Logger.Info().Timestamp().Msg(msg)
}

func (a *App) SingleSave(id int) backend.Save {
	return backend.GetSingleSave(id)
}

func (a *App) GetImage(path string) string {
	path = strings.TrimSpace(path)
	backend.Logger.Debug().Timestamp().Msg(path)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		backend.Logger.Error().Msg(err.Error())
		return string(bytes)

	}

	var base64Encoding string

	mimeType := http.DetectContentType(bytes)

	backend.Logger.Debug().Msg(mimeType)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += toB64(bytes)
	return base64Encoding
}

func toB64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
