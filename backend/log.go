package backend

import (
	"fmt"
	"os"
	"time"

	"github.com/adrg/xdg"
	log "github.com/sirupsen/logrus"
)

var logFileHandle *os.File

var logFile = fmt.Sprintf("october/logs/%s.json", time.Now().Format("20060102150405"))

func StartLogger() {
	logPath, err := xdg.DataFile(logFile)
	if err != nil {
		panic("Failed to create location to store logfiles")
	}
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logFileHandle = file
		log.SetOutput(file)
	} else {
		log.WithError(err).Error(err)
		log.Error("Failed to create log file, using stdout")
	}
}

func CloseLogFile() {
	logFileHandle.Close()
}
