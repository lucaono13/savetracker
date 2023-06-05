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

type ErrorReturn struct {
	Error        string                     `json:"Error"`
	String       string                     `json:"String"`
	Integer      int                        `json:"Integer"`
	Save         backend.Save               `json:"Save"`
	SaveList     []backend.Save             `json:"SaveList"`
	Matches      []backend.Match            `json:"Matches"`
	InTransfers  []backend.Transfer         `json:"InTransfers"`
	OutTransfers []backend.Transfer         `json:"OutTransfers"`
	Currency     string                     `json:"Currency"`
	Outfielders  []backend.PlayerSquadView  `json:"Outfielders"`
	Goalies      []backend.PlayerSquadView  `json:"Goalies"`
	OutTotals    []backend.PlayerTotalsView `json:"OutTotals"`
	GKTotals     []backend.PlayerTotalsView `json:"GKTotals"`
}

type NewSeason struct {
	TeamName      string   `json:"teamName"`
	ShortName     string   `json:"shortName"`
	Season        string   `json:"season"`
	Country       string   `json:"country"`
	TrophiesWon   []string `json:"trophiesWon"`
	SquadFile     string   `json:"squadFile"`
	ScheduleFile  string   `json:"scheduleFile"`
	TransfersFile string   `json:"transfersFile"`
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

func (a *App) RetrieveSaves() ErrorReturn {
	saves := backend.GetSaves()
	if saves == nil {
		return ErrorReturn{
			Error: "Error retrieving all saves",
		}
	}
	return ErrorReturn{
		SaveList: saves,
	}
	// return backend.GetSaves()
}

func (a *App) AddNewSave(saveName string, managerName string, gameVersion int, currency string) ErrorReturn {
	newSave := backend.Save{
		SaveName:    saveName,
		ManagerName: managerName,
		GameVersion: gameVersion,
		SaveID:      backend.NullInt64{NullInt64: sql.NullInt64{Valid: false}},
		SaveImage:   backend.NullString{NullString: sql.NullString{Valid: false}},
		Currency:    currency,
	}
	// addedID, err := backend.AddSave(saveName, managerName, gameVersion)
	// if err != nil {
	// 	return 0
	// }
	// toReturn := StringErrorReturn{}
	addedID, err := backend.AddSave(newSave)
	if err != nil {
		return ErrorReturn{
			Error: "Error adding save",
		}
		// return 0, "Error adding save."
	}
	return ErrorReturn{
		Integer: int(addedID),
	}
	// return int(addedID), "something is here"
}

func Log(msg string) {
	backend.Logger.Info().Timestamp().Msg(msg)
}

func (a *App) SingleSave(id int) ErrorReturn {
	saveID, err := backend.GetSingleSave(id)
	if err != nil {
		return ErrorReturn{
			Error: "Could not retrieve save",
		}
		// return backend.Save{}
	}
	return ErrorReturn{
		Save: saveID,
	}
}

func (a *App) SingleImage(id int) string {
	return backend.GetSaveImage(id)
}

func (a *App) GetSaveResults(saveID int) ErrorReturn {
	matches, err := backend.GetSaveResults(saveID)
	if err != nil {
		return ErrorReturn{
			Error: "Error getting results for save. Check log file for more details.",
		}
	}
	return ErrorReturn{
		Matches: matches,
	}
}

func (a *App) GetSaveTransfers(saveID int) ErrorReturn {
	inTransfers, currency, err := backend.GetSaveTransfers(saveID, 1)
	if err != nil {
		return ErrorReturn{
			Error: "Error getting transfers for save. Check log file for more details.",
		}
	}
	outTransfers, _, err := backend.GetSaveTransfers(saveID, 0)
	if err != nil {
		return ErrorReturn{
			Error: "Error getting transfers for save. Check log file for more details.",
		}
	}
	return ErrorReturn{
		InTransfers:  inTransfers,
		OutTransfers: outTransfers,
		Currency:     currency,
	}
}

func (a *App) GetSavePlayers(saveID int) ErrorReturn {
	outfielders, goalies, err := backend.GetSavePlayersStats(saveID)
	if err != nil {
		return ErrorReturn{
			Error: "Error getting players for save. Check log file for more details.",
		}
	}
	return ErrorReturn{
		Outfielders: outfielders,
		Goalies:     goalies,
	}
}

func (a *App) GetSavePlayersTotals(saveID int) ErrorReturn {
	outfielders, goalies, err := backend.GetSavePlayersTotals(saveID)
	if err != nil {
		return ErrorReturn{
			Error: "Error getting players' totals for save. Check log file for more details.",
		}
	}
	return ErrorReturn{
		OutTotals: outfielders,
		GKTotals:  goalies,
	}
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

func (a *App) SelectFileParse(fileType string) string {
	return SelectExportedFile(a.ctx, fileType)
}

func (a *App) AddNewSeason(saveID int, season NewSeason) ErrorReturn {
	// First parses files THEN adds to DB
	// Order to add a new season
	// Add team -> add Season -> add Players -> add Player Stats/Attributes ->
	// Add Transfers/Results -> Add Trophies
	var (
		inTransfers  []backend.Transfer
		outTransfers []backend.Transfer
	)
	team := backend.Team{
		TeamName: season.TeamName,
		Country:  season.Country,
	}
	if season.ShortName != "" {
		team.ShortName.String = season.ShortName
		team.ShortName.Valid = true
	}

	info, stats, err := backend.ParseStats(season.SquadFile, season.Season)
	if err != nil {
		return ErrorReturn{
			Error: "Error parsing squad file. Check log file for more details.",
		}
	}

	if season.ShortName == "" {
		inTransfers, outTransfers, err = backend.ParseTransfers(season.TransfersFile, season.TeamName)
	} else {
		inTransfers, outTransfers, err = backend.ParseTransfers(season.TransfersFile, season.ShortName)
	}
	if err != nil {
		return ErrorReturn{
			Error: "Error parsing transfer file. Check log file for more details.",
		}
	}

	results, err := backend.ParseSchedule(season.ScheduleFile)
	if err != nil {
		return ErrorReturn{
			Error: "Error parsing schedule file. Check log file for more details.",
		}
	}

	teamID, err := backend.AddTeam(team)
	if err != nil {
		return ErrorReturn{
			Error: "Error adding team. Check log file for more details.",
		}

	}
	seasonID, err := backend.AddSeason(saveID, int(teamID), season.Season)
	if err != nil {
		return ErrorReturn{
			Error: "Error adding season. Check log file for more details.",
		}
	}

	err = backend.AddPlayersInfo(saveID, info)
	if err != nil {
		return ErrorReturn{
			Error: "Error adding player info to DB. Check log file for more details.",
		}
	}
	err = backend.AddPlayersStats(saveID, int(seasonID), stats)
	if err != nil {
		return ErrorReturn{
			Error: "Error adding player stats/attributes to DB. Check log file for more details.",
		}
	}

	err = backend.AddTransfers(inTransfers, outTransfers, int(seasonID))
	if err != nil {
		return ErrorReturn{
			Error: "Error adding transfers to DB. Check log file for more details.",
		}
	}

	err = backend.AddSchedule(results, int(seasonID))
	if err != nil {
		return ErrorReturn{
			Error: "Error adding schedule to DB. Check log file for more details.",
		}
	}
	if len(season.TrophiesWon) > 0 {
		trophies := []backend.Trophy{}
		for _, trophy := range season.TrophiesWon {
			trophies = append(trophies, backend.Trophy{
				SeasonID:        int(seasonID),
				TrophyID:        backend.NullInt64{},
				CompetitionName: trophy,
				TrophyImage:     "",
			})
		}
		err := backend.AddTrophies(trophies)
		if err != nil {
			return ErrorReturn{
				Error: "Error adding trophies to DB. Check log file for more details.",
			}
		}
	}

	return ErrorReturn{}
	// return
}
