package logger

import (
	"fmt"
	"net/url"
	"os"
	"runtime"
	"time"

	"github.com/adrg/xdg"
	"go.uber.org/zap"
)

var logFile = fmt.Sprintf("october/logs/%s.log", time.Now().Format("2006-01-02"))
var Log *zap.SugaredLogger

// newWinFileSink creates a log sink on Windows machines as zap, by default,
// doesn't support Windows paths. A workaround is to create a fake winfile
// scheme and register it with zap instead. This workaround is taken from
// the Github issue at https://github.com/uber-go/zap/issues/621
func newWinFileSink(u *url.URL) (zap.Sink, error) {
	// Remove leading slash left by url.Parse()
	return os.OpenFile(u.Path[1:], os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
}

func Init() {
	logPath, err := xdg.DataFile(logFile)
	config := zap.NewProductionConfig()
	if runtime.GOOS == "windows" {
		err := zap.RegisterSink("winfile", newWinFileSink)
		if err != nil {
			panic("failed to register winfile sink")
		}
		logPath = "winfile:///" + logPath
	}
	config.OutputPaths = []string{"stdout", logPath}
	logger, err := config.Build()
	if err != nil {
		panic("failed to initialise logger")
	}
	defer logger.Sync()
	Log = logger.Sugar()
}
