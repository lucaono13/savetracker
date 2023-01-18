package backend

import (
	"database/sql"

	"github.com/adrg/xdg"
	_ "github.com/mattn/go-sqlite3"
)

type Save struct {
	SaveID      uint   `json:"id"`
	ManagerName string `json:"managerName"`
	GameVersion int    `json:"gameVersion"`
	SaveName    string `json:"saveName"`
}

var DB *sql.DB

func OpenDB() {
	dbFilePath, err := xdg.DataFile("Save Tracker/db/saveTracker.db")
	if err != nil {
		Logger.Error().Timestamp().Str("Error", err.Error())
		return
		// return err
	}

	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		Logger.Error().Timestamp().Msg("DB Connection could not be made")
		return
	}
	_, err = db.Exec(CreateTables)
	if err != nil {
		Logger.Error().Timestamp().Str("Error", err.Error())
		return
	}

	DB = db
	Logger.Info().Timestamp().Msg("DB has been opened.")
	// return nil
}

func CloseDB() error {
	Logger.Info().Timestamp().Msg("DB has been closed.")
	return DB.Close()
}

func GetSaves() []Save {
	q := `
	select * from saves;
	`

	rows, err := DB.Query(q)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
	}
	defer rows.Close()
	var saves []Save
	for rows.Next() {
		var (
			save Save
		)
		err := rows.Scan(&save.SaveID, &save.ManagerName, &save.GameVersion, &save.SaveName)
		if err != nil {
			Logger.Error().Timestamp().Msg(err.Error())
		}
		saves = append(saves, save)

	}
	return saves
}

func AddSave(saveName string, managerName string, gameVersion int) {
	result, err := DB.Exec(AddSaveQ, saveName, managerName, gameVersion)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return
	}
	rows, err := result.RowsAffected()
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
	}
	if rows != 1 {
		Logger.Error().Timestamp().Msg("Did not create 1 save.")
	} else {
		Logger.Info().Timestamp().Msg("Created new save.")
	}
}

func GetSingleSave(id int) Save {
	rows, err := DB.Query(SingleSave, id)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
	}
	// Logger.Info().Timestamp().Msg(rows)
	defer rows.Close()
	var saveExport Save
	for rows.Next() {
		err := rows.Scan(&saveExport.SaveID, &saveExport.ManagerName, &saveExport.GameVersion, &saveExport.SaveName)
		if err != nil {
			Logger.Error().Timestamp().Msg(err.Error())
		}
	}
	return saveExport
}
