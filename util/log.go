package util

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var (
	Logger zerolog.Logger
)

func init() {
	Logger = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
