package backend

import (
	"fmt"
	"io"
	"os"

	"github.com/adrg/xdg"
)

func NewImage(id int, filePath string) string {
	// Open source image
	source, err := os.Open(filePath)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
	}
	defer source.Close()

	// Get image info, specifically to get the base name
	sourceInfo, err := source.Stat()
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return ""
	}

	newFilePath, err := xdg.DataFile("Save Tracker/images/" + fmt.Sprintf("%d", id) + "_" + sourceInfo.Name())
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return ""
	}

	// Create destination image in correct location
	destination, err := os.Create(newFilePath)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return ""
	}
	defer destination.Close()

	// Copy image
	_, err = io.Copy(destination, source)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return ""
	}

	// Add to database
	err = UpdateSaveImage(id, newFilePath)
	if err != nil {
		return ""
	}

	return newFilePath
}
