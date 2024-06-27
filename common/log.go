package common

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func InitLog() {
	// Log as JSON instead of the default ASCII formatter.
	Logger.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	Logger.SetOutput(os.Stdout)

	Logger.SetLevel(logrus.InfoLevel)
}
