package backend

import (
	"database/sql"

	"github.com/adrg/xdg"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Save struct {
	SaveID      NullInt64  `json:"id" db:"saveID"`
	ManagerName string     `json:"managerName" db:"managerName"`
	GameVersion int        `json:"gameVersion" db:"gameVersion"`
	SaveName    string     `json:"saveName" db:"saveName"`
	SaveImage   NullString `json:"saveImage" db:"saveImage"`
}

// var DB *sql.DB
var DB *sqlx.DB

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

func GetSaves() []Save {
	var saves []Save
	err := DB.Select(&saves, AllSaves)
	switch {
	case err == sql.ErrNoRows:
		Logger.Info().Timestamp().Msg("No Saves Found")
	case err != nil:
		Logger.Error().Timestamp().Msg(err.Error())
	}
	return saves
}

// func AddSave(saveName string, managerName string, gameVersion int) (int, error) {
func AddSave(newSave Save) (int64, error) {
	result, err := DB.NamedExec(NewSave, newSave)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return 0, nil
	}
	return result.LastInsertId()
}

func GetSingleSave(ID int) Save {
	var save Save
	err := DB.QueryRowx(SingleSave, ID).StructScan(&save)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
	}
	return save
}

func GetSaveImage(ID int) string {
	var save Save
	err := DB.QueryRowx(SingleSaveImage, ID).StructScan(&save)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
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
