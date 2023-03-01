package backend

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

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
	// sourceInfo, err := source.Stat()
	// if err != nil {
	// 	Logger.Error().Timestamp().Msg(err.Error())
	// 	return ""
	// }

	newFileType := filepath.Ext(filePath)
	newFilePath, err := xdg.DataFile("Save Tracker/images/" + fmt.Sprintf("%d", id) + "_picture" + newFileType)
	// newFilePath, err := xdg.DataFile("Save Tracker/images/" + fmt.Sprintf("%d", id) + "_" + sourceInfo.Name())
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return ""
	}

	files, err := os.ReadDir(xdg.DataHome + "\\Save Tracker\\images")
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return ""
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), fmt.Sprintf("%d", id)+"_picture") {
			fileName, err := xdg.DataFile("Save Tracker/images/" + file.Name())
			if err != nil {
				Logger.Error().Timestamp().Msg(err.Error())
				return ""
			}
			err = os.Remove(fileName)
			if err != nil {
				Logger.Error().Timestamp().Msg(err.Error())
			}
		}
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
