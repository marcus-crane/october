package backend

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var logFileHandle *os.File

var logFile = fmt.Sprintf("october/logs/%s.json", time.Now().Format("20060102150405"))

func StartLogger(portable bool) {
	logPath, err := LocateDataFile(logFile, portable)
	if err != nil {
		panic("Failed to create location to store logfiles")
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logFileHandle = file
		logrus.SetOutput(file)
	} else {
		logrus.WithError(err).Error(err)
		logrus.Error("Failed to create log file, using stdout")
	}
}

func CloseLogFile() {
	logFileHandle.Close()
}
