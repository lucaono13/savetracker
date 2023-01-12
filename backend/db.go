package backend

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Save struct {
	SaveID      uint   `json:"id"`
	ManagerName string `json:"managerName"`
	GameVersion int    `json:"gameVersion"`
}

var DB *sql.DB

func OpenDB(filePath string) error {

	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		Logger.Error().Timestamp().Msg("DB Connection could not be made")
	}
	DB = db

	// db, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	// if err != nil {
	// 	Logg	er.Error().Timestamp().Msg("DB Connection could not be made.")
	// }
	// DB = db
	Logger.Info().Timestamp().Msg("DB has been opened.")
	return nil
}

func CloseDB() error {
	Logger.Info().Timestamp().Msg("DB has been closed.")
	return DB.Close()
}

func GetSaves() []Save {
	// DB.AutoMigrate(&Save{})
	q := `
	select * from saves;
	`

	rows, err := DB.Query(q)
	if err != nil {
		Logger.Error().Timestamp().Str("Error", err.Error())
	}
	defer rows.Close()
	var saves []Save
	for rows.Next() {
		var (
			save Save
		)
		err := rows.Scan(&save.SaveID, &save.ManagerName, &save.GameVersion)
		if err != nil {
			Logger.Error().Timestamp().Str("Error", err.Error())
		}
		saves = append(saves, save)

	}

	// DB.Find(&saves)
	return saves
}

func DeleteAllSaves() []Save {
	// DB.Exec("DELETE * FROM saves")
	// DB.Delete(&Save{})
	var saves []Save
	// DB.Find(&saves)
	return saves
}

func CreateDummy() {
	q := `
	insert into saves (managerName, gameVersion) values ("Gianluca", 22)
	`
	result, err := DB.Exec(q)
	if err != nil {
		Logger.Error().Timestamp().Str("Error", err.Error())
	}
	rows, err := result.RowsAffected()
	if err != nil {
		Logger.Error().Timestamp().Str("Error", err.Error())
	}
	if rows != 1 {
		Logger.Error().Timestamp().Msgf("Should have affected 1 row, affected %d rows", rows)
	} else {
		Logger.Info().Timestamp().Msg("Added Dummy entry")
	}
}

func AddSave(managerName string, gameVersion int) {
	result, err := DB.Exec(AddSaveQ, managerName, gameVersion)
	if err != nil {
		Logger.Error().Timestamp().Str("Error", err.Error())
	}
	rows, err := result.RowsAffected()
	if err != nil {
		Logger.Error().Timestamp().Str("Error", err.Error())
	}
	if rows != 1 {
		Logger.Error().Timestamp().Msgf("Should have affected 1 row, affected %d rows", rows)
	} else {
		Logger.Info().Timestamp().Msg("Created new save.")
	}
}
