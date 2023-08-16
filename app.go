package main

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
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
	Error         string                     `json:"Error"`
	String        string                     `json:"String"`
	Integer       int                        `json:"Integer"`
	Save          backend.Save               `json:"Save"`
	SaveList      []backend.Save             `json:"SaveList"`
	Matches       []backend.Match            `json:"Matches"`
	InTransfers   []backend.Transfer         `json:"InTransfers"`
	OutTransfers  []backend.Transfer         `json:"OutTransfers"`
	Currency      string                     `json:"Currency"`
	Outfielders   []backend.PlayerSquadView  `json:"Outfielders"`
	Goalies       []backend.PlayerSquadView  `json:"Goalies"`
	OutTotals     []backend.PlayerTotalsView `json:"OutTotals"`
	GKTotals      []backend.PlayerTotalsView `json:"GKTotals"`
	TopGls        []backend.TopResults       `json:"TopGls"`
	TopAsts       []backend.TopResults       `json:"TopAsts"`
	TopApps       []backend.TopResults       `json:"TopApps"`
	TopRat        []backend.TopResults       `json:"TopAvg"`
	TopTransfers  []backend.TopTransfers     `json:"TopTrfs"`
	TopLoans      []backend.TopTransfers     `json:"TopLoans"`
	AvgInFee      float32                    `json:"AvgInFee"`
	AvgOutFee     float32                    `json:"AvgOutFee"`
	Trophies      []backend.Trophy           `json:"Trophies"`
	ImageFile     string                     `json:"ImageFile"`
	ImageB64      string                     `json:"b64Image"`
	SinglePlayer  []backend.PlayerPageInfo   `json:"OnePlayer"`
	SingPlayerAvg backend.PlayerSumsAvgs     `json:"PlayerAvgSum"`
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
}

func (a *App) domReady(ctx context.Context) {
	a.ctx = ctx
	needUpdate, url, err := backend.CheckVersion()
	if err != nil {
		return
	}
	if needUpdate {
		runtime.EventsEmit(a.ctx, "updateVersion", url)
	}
}

func (a *App) shutdown(ctx context.Context) bool {
	backend.Logger.Info().Msg("Shutting down...")
	backend.CloseDB()
	backend.EndLogging()

	return false
}

func (a *App) aboutDialog(ctx context.Context) {
	a.ctx = ctx
	runtime.EventsEmit(a.ctx, "aboutDialog", backend.GetVersion(), backend.GetDBVersion())
}

func (a *App) openGithubIssues(ctx context.Context) {
	a.ctx = ctx
	runtime.EventsEmit(a.ctx, "githubIssues")
}

// Open File Dialog
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
		backend.Logger.Error().Msg(err.Error())
		return ""
	}
	return file
}

// Selecting the files exported from the game.
func (a *App) SelectFileParse(fileType string) string {
	return SelectExportedFile(a.ctx, fileType)
}

// Save Functions
func (a *App) RetrieveSaves() ErrorReturn {
	saves, err := backend.GetSaves()
	if err != nil {
		return ErrorReturn{
			Error: "Error retrieving all saves",
		}
	}
	return ErrorReturn{
		SaveList: saves,
	}
}

func (a *App) AddNewSave(saveName string, managerName string, gameVersion string, currency string) ErrorReturn {
	newSave := backend.Save{
		SaveName:    saveName,
		ManagerName: managerName,
		GameVersion: gameVersion,
		SaveID:      backend.NullInt64{NullInt64: sql.NullInt64{Valid: false}},
		SaveImage:   backend.NullString{NullString: sql.NullString{Valid: false}},
		Currency:    currency,
	}
	addedID, err := backend.AddSave(newSave)
	if err != nil {
		return ErrorReturn{
			Error: "Error adding save",
		}
	}
	return ErrorReturn{
		Integer: int(addedID),
	}
}

func (a *App) SingleSave(id int) ErrorReturn {
	saveID, err := backend.GetSingleSave(id)
	if err != nil {
		return ErrorReturn{
			Error: "Could not retrieve save",
		}
	}
	return ErrorReturn{
		Save: saveID,
	}
}

func (a *App) UploadSaveImage(id int) ErrorReturn {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Save Image",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg;*.jpeg)",
				Pattern:     "*.png;*.jpg;*.jpeg",
			},
		},
	})
	if file == "" {
		return ErrorReturn{}
	}
	if err != nil {
		backend.Logger.Error().Msg(err.Error())
		return ErrorReturn{
			Error: "Error getting save image. Check log file for more details.",
		}
	}
	// Have backend deal with the new image
	// return backend.NewImage(id, file)
	err = backend.UpdateSaveImage(id, file)
	if err != nil {
		return ErrorReturn{
			Error: " Error updating save image. Check log file for more details.",
		}
	}
	return ErrorReturn{
		ImageFile: file,
	}
}

func (a *App) SingleImage(id int) string {
	return backend.GetSaveImage(id)
}

func (a *App) GetNumSaves() int {
	return backend.GetNumSaves()
}

func (a *App) DeleteSave(saveID int) ErrorReturn {
	err := backend.DeleteSave(saveID)
	if err != nil {
		return ErrorReturn{
			Error: "Error deleting save. Check logs for more details",
		}
	}
	runtime.EventsEmit(a.ctx, "saveDeleted")
	return ErrorReturn{}
}

func (a *App) GetSaveStory(saveID int) backend.Story {
	return backend.GetSaveStory(saveID)
}

func (a *App) UpdateSaveStory(updatedStory backend.Story) string {
	err := backend.UpdateSaveStory(updatedStory)
	if err != nil {
		return "Error updating story. Check log file for more details."
	}
	return ""
}

// Seasons
func (a *App) AddNewSeason(saveID int, season NewSeason, historicalSeason bool) ErrorReturn {
	// First parses files THEN adds to DB
	// Order to add a new season
	// Add team -> add Season -> add Players -> add Player Stats/Attributes ->
	// Add Transfers/Results -> Add Trophies
	var (
		inTransfers  []backend.Transfer
		outTransfers []backend.Transfer
		info         []backend.PlayerInfo
		stats        []backend.PlayerSeason
		err          error
	)
	team := backend.Team{
		TeamName: season.TeamName,
		Country:  season.Country,
	}
	if season.ShortName != "" {
		team.ShortName.String = season.ShortName
		team.ShortName.Valid = true
	}
	if historicalSeason {
		info, stats, err = backend.ParseStats(season.SquadFile, season.Season)
		if err != nil {
			return ErrorReturn{
				Error: "Error parsing squad file. Check log file for more details.",
			}
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

	if historicalSeason {
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
				TrophyID:        backend.NullInt64{},
				CompetitionName: trophy,
				TrophyImage:     backend.NullString{},
			})
		}
		err := backend.AddSeasonTrophies(trophies, int(seasonID))
		if err != nil {
			return ErrorReturn{
				Error: "Error adding trophies to DB. Check log file for more details.",
			}
		}
	}

	return ErrorReturn{}
}

func (a *App) GetNumSeasonsInSave(saveID int) bool {
	return backend.CheckIfSeasonsInSave(saveID)
}

func (a *App) GetNumSeasons() int {
	return backend.GetNumSeasons()
}

func (a *App) GetSavePSeasons(saveID int) int {
	return backend.GetSavePlayerSeasons(saveID)
}

func (a *App) GetAllPSeasons() int {
	return backend.GetAllPlayerSeasons()
}

// Players
func (a *App) GetSavePlayers(saveID int) ErrorReturn {
	outfielders, goalies, err := backend.GetSavePlayersStats(saveID)
	fmt.Println(err)
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

func (a *App) GetAllPlayersSeason() ErrorReturn {
	outfielders, goalies, err := backend.GetAllPlayersStats()
	if err != nil {
		return ErrorReturn{
			Error: "Error getting all players' seasons. Check log file for more details.",
		}
	}
	return ErrorReturn{
		Outfielders: outfielders,
		Goalies:     goalies,
	}
}

func (a *App) GetAllPlayersTotals() ErrorReturn {
	outfielders, goalies, err := backend.GetAllPlayersTotals()
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

func (a *App) GetSinglePlayer(playerID int) ErrorReturn {
	player, sumsAvgs, err := backend.GetSinglePlayer(playerID)
	if err != nil {
		return ErrorReturn{
			Error: "Error getting player's info. Check log file for more details.",
		}
	}
	return ErrorReturn{
		SinglePlayer:  player,
		SingPlayerAvg: sumsAvgs,
	}
}

// Results
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

func (a *App) GetAllResults() ErrorReturn {
	matches, err := backend.GetAllResults()
	if err != nil {
		return ErrorReturn{
			Error: "Error getting all results. Check log file for more details.",
		}
	}
	return ErrorReturn{
		Matches: matches,
	}
}

// Transfers
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

func (a *App) GetAllTransfers() ErrorReturn {
	inTransfers, err := backend.GetAllTransfers(1)
	if err != nil {
		return ErrorReturn{
			Error: "Error getting all transfers. Check log file for more details.",
		}
	}
	outTransfers, err := backend.GetAllTransfers(0)
	if err != nil {
		return ErrorReturn{
			Error: "Error getting all transfers. Check log file for more details.",
		}
	}
	return ErrorReturn{
		InTransfers:  inTransfers,
		OutTransfers: outTransfers,
	}
}

// Trophies
func (a *App) GetTrophies(saveID int) ErrorReturn {
	trophies, err := backend.GetSaveTrophies(saveID)
	if err != nil {
		return ErrorReturn{
			Error: "Error getting save trophies. Check log file for more details.",
		}
	}
	return ErrorReturn{
		Trophies: trophies,
	}
}

func (a *App) GetAllTrophies() ErrorReturn {
	trophies, err := backend.GetTrophies()
	if err != nil {
		return ErrorReturn{
			Error: "Error getting all trophies. Check log file for more details.",
		}
	}
	return ErrorReturn{
		Trophies: trophies,
	}
}

func (a *App) SelectNewTrophyImage(id int) ErrorReturn {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Trophy Image",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg;*.jpeg)",
				Pattern:     "*.png;*.jpg;*.jpeg",
			},
		},
	})
	if err != nil {
		backend.Logger.Error().Msg(err.Error())
		return ErrorReturn{
			Error: "Error getting trophy image. Check log file for more details.",
		}
	}
	err = backend.UpdateTrophyImage(id, file)
	if err != nil {
		backend.Logger.Error().Msg(err.Error())
		return ErrorReturn{
			Error: "Error updating trophy image. Check log file for more details.",
		}
	}
	return ErrorReturn{
		ImageFile: file,
	}
}

// Save Home Page
func (a *App) GetSaveHomeRankings(saveID int, getPlayerStats bool) ErrorReturn {
	var (
		topGoals       []backend.TopResults
		topAssists     []backend.TopResults
		topAppearances []backend.TopResults
		topRatings     []backend.TopResults
		err            error
	)
	topRankings := a.GetSaveResults(saveID)
	if topRankings.Error != "" {
		return topRankings
	}
	if getPlayerStats {
		topGoals, topAssists, topAppearances, topRatings, err = backend.GetTopPlayersX(saveID)
		if err != nil {
			topRankings.Error = "Error retrieving one of top goals, assists, starts, and ratings. Check log file for more details."
			return topRankings
		}
	}
	topTransfers, topLoans, avgIn, avgOut, err := backend.GetTransfersStats(saveID)
	if err != nil {
		topRankings.Error = "Error retrieving one of the transfer stats. Check log file for more details."
		return topRankings
	}
	trophies, err := backend.GetSaveTrophies(saveID)
	if err != nil {
		topRankings.Error = "Error retrieving save trophies. Check log file for more details."
		return topRankings
	}
	topRankings.TopTransfers = topTransfers
	topRankings.TopLoans = topLoans
	topRankings.AvgInFee = avgIn
	topRankings.AvgOutFee = avgOut
	topRankings.TopGls = topGoals
	topRankings.TopAsts = topAssists
	topRankings.TopApps = topAppearances
	topRankings.TopRat = topRatings
	topRankings.Trophies = trophies
	return topRankings
}

// App Default Home Screen
func (a *App) GetAllRankings(getPlayerStats bool) ErrorReturn {
	var (
		topGoals       []backend.TopResults
		topAssists     []backend.TopResults
		topAppearances []backend.TopResults
		topRatings     []backend.TopResults
		err            error
	)
	topRankings := a.GetAllResults()
	if topRankings.Error != "" {
		return topRankings
	}
	if getPlayerStats {
		topGoals, topAssists, topAppearances, topRatings, err = backend.GetAllTops()
		if err != nil {
			topRankings.Error = "Error retrieving one of top goals, assists, starts, and ratings. Check log file for more details."
			return topRankings
		}
	}

	topTransfers, topLoans, avgIn, avgOut, err := backend.GetAllTransersStats()
	if err != nil {
		topRankings.Error = "Error retrieving one of the transfer stats. Check log file for more details."
		return topRankings
	}
	trophies, err := backend.GetTrophies()
	if err != nil {
		topRankings.Error = "Error retrieving all trophies. Check log file for more details."
		return topRankings
	}
	topRankings.TopTransfers = topTransfers
	topRankings.TopLoans = topLoans
	topRankings.AvgInFee = avgIn
	topRankings.AvgOutFee = avgOut
	topRankings.TopGls = topGoals
	topRankings.TopAsts = topAssists
	topRankings.TopApps = topAppearances
	topRankings.TopRat = topRatings
	topRankings.Trophies = trophies
	return topRankings
}

// Utility
func (a *App) GetImage(path string) ErrorReturn {
	path = strings.TrimSpace(path)
	if filepath.Ext(path) != ".png" && filepath.Ext(path) != ".jpg" && filepath.Ext(path) != ".jpeg" {
		return ErrorReturn{
			Error: "Photo is not correct file type",
		}
	}
	bytes, err := os.ReadFile(path)
	if err != nil {
		backend.Logger.Error().Msg(err.Error())
		return ErrorReturn{
			Error: "Photo file not found",
		}
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
	return ErrorReturn{
		ImageB64: base64Encoding,
	}
}

func (a *App) Log(msg string) {
	backend.Logger.Info().Msg(msg)
}
