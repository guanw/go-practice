package logger

import (
	"go.uber.org/fx"
	"log"
	"os"
)

var Module = fx.Provide(NewLogger)

func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "", 0)
	logger.Print("Executing NewLogger")
	return logger
}
