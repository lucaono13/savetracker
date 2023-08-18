package backend

import (
	"fmt"
	"os"
	"time"

	"github.com/adrg/xdg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	logFileHandler *os.File
	logFile        = fmt.Sprintf("Save Tracker/logs/%s.json", time.Now().Format("20060102"))
	Logger         zerolog.Logger
)

func StartLogger() {
	files, err := os.ReadDir(xdg.DataHome + "/Save Tracker/logs")
	if err != nil {
		log.Fatal().Err(err).Msg("Error finding logs directory")
	}
	if len(files) > 4 {
		for _, file := range files[:len(files)-4] {
			os.Remove(xdg.DataHome + "/Save Tracker/logs/" + file.Name())
		}
	}
	// Create the directory for the log file
	logFilePath, err := xdg.DataFile(logFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating log file destination")
	}
	// Open the log file (create if it doesn't exist)
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating log file")
	} else {

		logFileHandler = file
		Logger = zerolog.New(logFileHandler).With().Timestamp().Caller().Logger()
	}
	Logger.Info().Timestamp().Msg("Logger has been created and is starting.")
}

func EndLogging() {
	logFileHandler.Close()
}
