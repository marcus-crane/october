package backend

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

var logFileHandle *os.File

var logFile = fmt.Sprintf("october/logs/%s.txt", time.Now().Format("20060102150405"))

func StartLogger(portable bool, level slog.Leveler) (*slog.Logger, error) {
	logPath, err := LocateDataFile(logFile, portable)
	if err != nil {
		panic("Failed to create location to store logfiles")
	}
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return &slog.Logger{}, err
	}
	handlerOpts := slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}
	handler := slog.NewTextHandler(file, &handlerOpts)
	return slog.New(handler), nil
}

func CloseLogFile() {
	logFileHandle.Close()
}

// Dummy until real handler ships in Go 1.24
type discardHandler struct {
	slog.TextHandler
}

func (d *discardHandler) Enabled(context.Context, slog.Level) bool { return false }
