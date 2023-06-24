package backend

import (
	"strconv"

	"github.com/adrg/xdg"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// var DB *sql.DB
var DB *sqlx.DB

// Database
func OpenDB() {
	dbFilePath, err := xdg.DataFile("Save Tracker/db/saveTracker.db")
	if err != nil {
		Logger.Error().Timestamp().Str("Error", err.Error())
		return
	}
	db, err := sqlx.Open("sqlite3", dbFilePath)
	if err != nil {
		Logger.Error().Timestamp().Msg("DB Connection could not be made")
		return
	}
	_, err = db.Exec(CreateTables)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return
	}
	DB = db
	Logger.Info().Timestamp().Msg("DB has been opened.")
}

func CloseDB() error {
	Logger.Info().Timestamp().Msg("DB has been closed.")
	return DB.Close()
}

// Saves
func GetSaves() ([]Save, error) {
	var saves []Save
	err := DB.Select(&saves, AllSaves)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, err
	}
	return saves, nil
}

func GetSingleSave(ID int) (Save, error) {
	var save Save
	err := DB.QueryRowx(SingleSave, ID).StructScan(&save)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return save, err
	}
	return save, nil
}

func GetSaveImage(ID int) string {
	var save Save
	err := DB.QueryRowx(SingleSaveImage, ID).StructScan(&save)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return ""
	}
	return save.SaveImage.String
}

func UpdateSaveImage(id int, filePath string) error {
	_, err := DB.Exec(SaveImageUpdate, filePath, id)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
	}
	return err
}

func AddSave(newSave Save) (int64, error) {
	result, err := DB.NamedExec(NewSave, newSave)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return 0, nil
	}
	return result.LastInsertId()
}

func GetNumSaves() int {
	row := DB.QueryRow(NumSaves)
	var numSaves int
	err := row.Scan(&numSaves)
	if err != nil {
		Logger.Error().Timestamp().Msg("numSaves")
		Logger.Error().Timestamp().Msg(err.Error())
		return 0
	}
	return numSaves
}

// Seasons
func AddSeason(saveID int, teamID int, year string) (int64, error) {
	result, err := DB.Exec(NewSeason, teamID, saveID, year)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return 0, err
	}
	seasonID, err := result.LastInsertId()
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return 0, err
	}
	return seasonID, nil
}

func CheckIfSeasonsInSave(saveID int) bool {
	row := DB.QueryRow(NumSeasonsToSave, saveID)
	var numSeasons int
	err := row.Scan(&numSeasons)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return false
	}
	if numSeasons == 0 {
		return false
	}
	return true
}

func GetNumSeasons() int {
	row := DB.QueryRow(NumSeasons)
	var numSeasons int
	err := row.Scan(&numSeasons)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return 0
	}
	return numSeasons
}

// Teams
func GetTeams() ([]Team, error) {
	var teams []Team
	err := DB.Select(&teams, AllTeams)
	return teams, err
}

func AddTeam(team Team) (int64, error) {
	teams, err := GetTeams()
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return 0, err
	}
	teamsInDB := make(map[string]Team)
	for _, team := range teams {
		teamsInDB[team.TeamName] = team
	}
	if _, ok := teamsInDB[team.TeamName]; ok {
		if teamsInDB[team.TeamName].Country == team.Country {
			return teamsInDB[team.TeamName].TeamID.Int64, nil
		}
	}

	result, err := DB.NamedExec(NewTeam, team)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return 0, err
	}
	teamID, err := result.LastInsertId()
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return 0, err
	}
	return teamID, nil
}

// Player Season
func addPlayerSeason(playerID int, seasonID int) (int64, error) {
	result, err := DB.Exec(NewPlayerSeason, playerID, seasonID)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return 0, err
	}
	playerSeasonID, err := result.LastInsertId()
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return 0, err
	}
	return playerSeasonID, nil
}

// Players
func GetSavePlayers(saveID int) ([]PlayerInfo, error) {
	var players []PlayerInfo
	err := DB.Select(&players, SavePlayers, saveID)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, err
	}
	return players, nil
}

func AddPlayersInfo(saveID int, info []PlayerInfo) error {
	tx, err := DB.Beginx()
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return err
		// return "Error starting transaction.\nCheck log file for more details."
	}
	defer tx.Rollback()
	oneNat, err := tx.PrepareNamed(AddPlayer1Nat)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return err
		// return "Error preparing statement.\nCheck log file for more details."
	}
	twoNat, err := tx.PrepareNamed(AddPlayer2Nat)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return err
		// return "Error preparing statement.\nCheck log file for more details."
	}

	currentPlayers, err := GetSavePlayers(saveID)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return err
		// return "Error getting players in save.\nCheck log file for more details."
	}
	playersMap := MakePlayersMap(currentPlayers)
	for _, player := range info {
		if _, ok := playersMap[player.UID]; !ok {
			player.SaveID = saveID
			// currentPlayers[player.UID] = int(player.PlayerID)
			// fmt.Println(player, player.SecNat)
			if player.SecNat.Valid {
				_, err := twoNat.Exec(player)
				if err != nil {
					Logger.Error().Timestamp().Msg(err.Error())
					return err
					// return "Error executing adding player.\nCheck log file for more details."
				}
			} else {
				_, err := oneNat.Exec(player)
				if err != nil {
					Logger.Error().Timestamp().Msg(err.Error())
					return err
					// return "Error executing adding player.\nCheck log file for more details."
				}
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return err
		// return "Error committing transaction.\nCheck log file for more details."
	}

	return nil
}

func AddPlayersStats(saveID int, seasonID int, stats []PlayerSeason) error {
	players, err := GetSavePlayers(saveID)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return err
		// return "Error Getting Players for save " + strconv.Itoa(saveID) + ".\nCheck log file for more details."
	}
	playersMap := MakePlayersMap(players)
	for _, player := range stats {
		if playerID, ok := playersMap[player.UID]; ok {
			playerSeasonID, err := addPlayerSeason(playerID, seasonID)
			if err != nil {
				Logger.Error().Timestamp().Msg(err.Error())
				return err
				// return "Error adding Player Season.\nCheck log file for more details."
			}
			player.PlayerSeasonID = int(playerSeasonID)
			_, err = DB.NamedExec(NewStats, player)
			if err != nil {
				Logger.Error().Timestamp().Msg(err.Error())
				return err
				// return "Error executing stats INSERT.\nCheck log file for more details."
			}
			_, err = DB.NamedExec(NewAttrs, player)
			if err != nil {
				Logger.Error().Timestamp().Msg(err.Error())
				return err
				// return "Error executing attribute INSERT.\nCheck log file for more details."
			}
		} else {
			Logger.Error().Timestamp().Msg("Player with UID: " + strconv.Itoa(player.UID) + " does not exist in DB.")
		}
	}
	return nil
}

func GetSinglePlayer(playerID int) ([]PlayerPageInfo, PlayerSumsAvgs, error) {
	var (
		player   []PlayerPageInfo
		sumsAvgs PlayerSumsAvgs
	)

	err := DB.Select(&player, SinglePlayer, playerID)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, PlayerSumsAvgs{}, err
	}
	err = DB.QueryRowx(SinglePlayerSumsAvgs, playerID).StructScan(&sumsAvgs)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, PlayerSumsAvgs{}, err
	}
	return player, sumsAvgs, nil
}

func GetSavePlayersStats(saveID int) ([]PlayerSquadView, []PlayerSquadView, error) {
	var (
		players []PlayerSquadView
		goalies []PlayerSquadView
	)
	err := DB.Select(&players, SaveOutfieldPlayers, saveID)
	if err != nil {
		Logger.Error().Timestamp().Msg("2")
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, err
	}
	err = DB.Select(&goalies, SaveGoaliePlayers, saveID)
	if err != nil {
		Logger.Error().Timestamp().Msg("1")
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, err
	}
	return players, goalies, nil
}

func GetSavePlayersTotals(saveID int) ([]PlayerTotalsView, []PlayerTotalsView, error) {
	var (
		players []PlayerTotalsView
		goalies []PlayerTotalsView
	)
	err := DB.Select(&players, TotalsSaveOutfieldPlayers, saveID)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, err
	}
	err = DB.Select(&goalies, TotalsSaveGoalies, saveID)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, err
	}
	return players, goalies, nil
}

func GetAllPlayersStats() ([]PlayerSquadView, []PlayerSquadView, error) {
	var (
		players []PlayerSquadView
		goalies []PlayerSquadView
	)
	err := DB.Select(&players, AllSaveOutfieldPlayers)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, err
	}
	err = DB.Select(&goalies, AllSaveGoaliePlayers)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, err
	}
	return players, goalies, nil
}

func GetAllPlayersTotals() ([]PlayerTotalsView, []PlayerTotalsView, error) {
	var (
		players []PlayerTotalsView
		goalies []PlayerTotalsView
	)
	err := DB.Select(&players, AllTotalsSaveOutfieldPlayers)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, err
	}
	err = DB.Select(&goalies, AllTotalsSaveGoalies)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, err
	}
	return players, goalies, nil
}

// Transfers
func AddTransfers(ins []Transfer, outs []Transfer, seasonID int) error {
	tx, err := DB.Beginx()
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return err
		// return "Error starting transaction.\nCheck log file for more details."
	}
	defer tx.Rollback()

	transferStmt, err := tx.PrepareNamed(NewTransfer)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return err
	}

	for _, transfer := range ins {
		transfer.SeasonID = seasonID
		// fmt.Println(transfer.PotentialFee.Int64)
		_, err := transferStmt.Exec(transfer)
		if err != nil {
			Logger.Error().Timestamp().Msg(err.Error())
			return err
		}
	}
	for _, transfer := range outs {
		transfer.SeasonID = seasonID
		_, err := transferStmt.Exec(transfer)
		if err != nil {
			Logger.Error().Timestamp().Msg(err.Error())
			return err
		}
	}
	return tx.Commit()
}

func GetSaveTransfers(saveID int, transferIn int) ([]Transfer, string, error) {
	var transfers []Transfer
	err := DB.Select(&transfers, SaveTransfers, saveID, transferIn)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, "", err
	}
	var save Save
	err = DB.QueryRowx(GetCurrency, saveID).StructScan(&save)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, "", err
	}
	// return save.SaveImage.String
	return transfers, save.Currency, nil
}

func GetAllTransfers(transferIn int) ([]Transfer, error) {
	var transfers []Transfer
	err := DB.Select(&transfers, AllSaveTransfers, transferIn)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, err
	}
	// var save Save
	// err = DB.QueryRowx(GetCurrency, saveID).StructScan(&save)
	// if err != nil {
	// 	Logger.Error().Timestamp().Msg(err.Error())
	// 	return nil, "", err
	// }
	// return save.SaveImage.String
	return transfers, nil
}

func GetTransfersStats(saveID int) ([]TopTransfers, float32, float32, error) {
	var (
		topTransfers  []TopTransfers
		avgOut, avgIn float32
	)
	err := DB.Select(&topTransfers, MostTransfers, saveID)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, 0.0, 0.0, err
	}
	row := DB.QueryRow(AvgTransfersOut, saveID)
	err = row.Scan(&avgOut)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, 0.0, 0.0, err
	}
	row = DB.QueryRow(AvgTransfersIn, saveID)
	err = row.Scan(&avgIn)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, 0.0, 0.0, err
	}
	return topTransfers, avgIn, avgOut, nil
}

func GetAllTransersStats() ([]TopTransfers, float32, float32, error) {
	var (
		topTransfers  []TopTransfers
		avgOut, avgIn float32
	)
	err := DB.Select(&topTransfers, AllMostTransfers)
	if err != nil {
		Logger.Error().Timestamp().Msg("all transfers")
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, 0.0, 0.0, err
	}
	row := DB.QueryRow(AllAvgTransfersOut)
	err = row.Scan(&avgOut)
	if err != nil {
		Logger.Error().Timestamp().Msg("avgout")
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, 0.0, 0.0, err
	}
	row = DB.QueryRow(AllAvgTransfersIn)
	err = row.Scan(&avgIn)
	if err != nil {
		Logger.Error().Timestamp().Msg("avgin")
		Logger.Error().Timestamp().Msg("err " + err.Error())
		return nil, 0.0, 0.0, err
	}
	return topTransfers, avgIn, avgOut, nil
}

// Schedule
func AddSchedule(results []Match, seasonID int) error {
	tx, err := DB.Beginx()
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareNamed(NewResult)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return err
	}

	for _, result := range results {
		result.SeasonID = seasonID
		_, err := stmt.Exec(result)
		if err != nil {
			Logger.Error().Timestamp().Msg(err.Error())
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return err
	}
	return nil
}

func GetSaveResults(saveID int) ([]Match, error) {
	var matches []Match
	err := DB.Select(&matches, SaveResults, saveID)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, err
	}
	return matches, nil
}

func GetAllResults() ([]Match, error) {
	var matches []Match
	err := DB.Select(&matches, AllSaveResults)
	if err != nil {
		Logger.Error().Timestamp().Msg("results")
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, err
	}
	return matches, nil
}

// Trophies
func AddSeasonTrophies(trophies []Trophy, seasonID int) error {
	allTrophies, err := GetTrophies()
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return err
	}
	trophiesMap := MakeTrophiesMap(allTrophies)
	for _, trophy := range trophies {
		if _, ok := trophiesMap[trophy.CompetitionName]; !ok {
			row, err := DB.Exec(NewTrophy, trophy.CompetitionName)
			if err != nil {
				Logger.Error().Timestamp().Msg(err.Error())
				return err
			}
			newTrophyID, err := row.LastInsertId()
			if err != nil {
				Logger.Error().Timestamp().Msg(err.Error())
				return err
			}
			_, err = DB.Exec(NewTrophyWon, seasonID, newTrophyID)
			if err != nil {
				Logger.Error().Timestamp().Msg(err.Error())
				return err
			}
		} else {
			_, err := DB.Exec(NewTrophyWon, seasonID, trophiesMap[trophy.CompetitionName])
			if err != nil {
				Logger.Error().Timestamp().Msg(err.Error())
				return err
			}
		}
	}
	return nil
}

func GetSaveTrophies(saveID int) ([]Trophy, error) {
	var trophies []Trophy
	err := DB.Select(&trophies, SaveTrophies, saveID)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, err
	}
	return trophies, nil
}

func GetTrophies() ([]Trophy, error) {
	var trophies []Trophy
	err := DB.Select(&trophies, AllTrophies)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, err
	}
	return trophies, nil
}

func UpdateTrophyImage(id int, filePath string) error {
	_, err := DB.Exec(TrophyImageUpdate, filePath, id)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
	}
	return err
}

// Save Home Page
func GetTopPlayersX(saveID int) ([]TopResults, []TopResults, []TopResults, []TopResults, error) {
	var goals, assists, apps, ratings []TopResults
	err := DB.Select(&goals, Top5Gls, saveID)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, nil, nil, err
	}
	err = DB.Select(&assists, Top5Asts, saveID)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, nil, nil, err
	}
	err = DB.Select(&apps, Top5Apps, saveID)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, nil, nil, err
	}
	err = DB.Select(&ratings, TopRat, saveID)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, nil, nil, err
	}
	return goals, assists, apps, ratings, nil
}

// App Default Home Screen
func GetAllTops() ([]TopResults, []TopResults, []TopResults, []TopResults, error) {
	var goals, assists, apps, ratings []TopResults
	err := DB.Select(&goals, AllTop5Gls)
	if err != nil {
		Logger.Error().Timestamp().Msg("gls")
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, nil, nil, err
	}
	err = DB.Select(&assists, AllTop5Asts)
	if err != nil {
		Logger.Error().Timestamp().Msg("asts")
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, nil, nil, err
	}
	err = DB.Select(&apps, AllTop5Apps)
	if err != nil {
		Logger.Error().Timestamp().Msg("apps")
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, nil, nil, err
	}
	err = DB.Select(&ratings, AllTopRat)
	if err != nil {
		Logger.Error().Timestamp().Msg("rats")
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, nil, nil, err
	}
	return goals, assists, apps, ratings, nil
}
