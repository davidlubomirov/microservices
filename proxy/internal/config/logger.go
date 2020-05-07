package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

var (
	SystemLogger = logrus.New()
)

func SetupLogger() *logrus.Logger {
	SystemLogger.Out = os.Stdout
	SystemLogger.SetLevel(logrus.DebugLevel)
	SystemLogger.Formatter = &logrus.JSONFormatter{}

	return SystemLogger
}
