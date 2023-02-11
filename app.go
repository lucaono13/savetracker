package main

import (
	"context"
	"fmt"

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

func (a *App) shutdown(ctx context.Context) {
	backend.Logger.Info().Msg("Shutting down...")
	backend.EndLogging()
	backend.CloseDB()
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
