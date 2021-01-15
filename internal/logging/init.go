package logging

import (
	// External modules
	"BotsBot/internal/configuration"
	"fmt"
	"path"
	"runtime"

	log "github.com/sirupsen/logrus"
)

// InitLogger init logrus and configure log pattern
func InitLogger() {
	// Set logger formatter
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			functionName := path.Base(f.Function)
			return fmt.Sprintf("%s()", functionName), fmt.Sprintf(" %s:%d", filename, f.Line)
		},
	})
}

// ParseLogLevel - function for parse logging level
func ParseLogLevel() log.Level {
	level, err := log.ParseLevel(configuration.LoggingLevel)
	if err != nil {
		log.Fatal("Level parse error: ", err)
	}

	return level
}
