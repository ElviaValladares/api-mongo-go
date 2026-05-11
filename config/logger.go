package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Logger = log.New()

func InitLogger() {

	Logger.SetFormatter(&log.JSONFormatter{})

	Logger.SetOutput(os.Stdout)

	Logger.SetLevel(log.InfoLevel)
}