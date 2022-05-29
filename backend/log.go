package backend

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/adrg/xdg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logFile = fmt.Sprintf("october/logs/%s.log", time.Now().Format("20060102150405"))

func ConfigureLogger() {
	logPath, err := xdg.DataFile(logFile)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	multiOutput := io.MultiWriter(os.Stdout, f)
	log.Logger = log.Output(multiOutput).Level(zerolog.DebugLevel)
}
