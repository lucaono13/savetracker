package backend

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Save struct {
	gorm.Model
	ManagerName string
	GameVersion int
}

var DB *gorm.DB

func OpenDB(filePath string) error {
	db, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	if err != nil {
		Logger.Error().Timestamp().Msg("DB Connection could not be made.")
	}
	DB = db
	return nil
}

func GetSaves() []Save {
	DB.AutoMigrate(&Save{})
	DB.Create(&Save{
		ManagerName: "Gianluca",
		GameVersion: 22,
	})
	var saves []Save
	DB.Find(&saves)
	return saves
}
