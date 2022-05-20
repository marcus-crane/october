package backend

import (
	"fmt"
	"os"
	"time"

	"github.com/adrg/xdg"
	"github.com/rs/zerolog/log"
)

var logFile = fmt.Sprintf("october/logs/%s.log", time.Now().Format("20060102150405"))

func ConfigureLogger() {
	logPath, err := xdg.DataFile(logFile)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.Logger = log.Output(f)
}
