package logfx

import (
	"log"
	"os"

	"go.uber.org/fx"
)

// Module provides the *log.Logger
var Module = fx.Options(
	fx.Provide(NewLogger),
)

// NewLogger instantiates a logger which can be used throughout
func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "[fxdemo] ", 0)
	logger.Print("Instantiating a logger")
	return logger

}
