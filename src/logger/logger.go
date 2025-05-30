package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func Init() {
	logEnv := os.Getenv("ENVIRONMENT")
	var raw *zap.Logger
	var err error

	if strings.ToLower(logEnv) == "production" {
		raw, err = zap.NewProduction()
	} else {
		raw, err = zap.NewDevelopment()
	}

	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	Log = raw.Sugar()
}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}
