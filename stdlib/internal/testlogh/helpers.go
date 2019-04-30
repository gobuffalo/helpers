package testlogh

import (
	"internal/testlog"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	GetenvKey = "Getenv"

	LoggerKey = "Logger"

	OpenKey = "Open"

	SetLoggerKey = "SetLogger"

	StatKey = "Stat"
)

func New() hctx.Map {
	return hctx.Map{

		GetenvKey: Getenv,

		LoggerKey: Logger,

		OpenKey: Open,

		SetLoggerKey: SetLogger,

		StatKey: Stat,
	}
}

// Getenv calls Logger().Getenv, if a logger has been set.
var Getenv = testlog.Getenv

// Logger returns the current test logger implementation.
// It returns nil if there is no logger.
var Logger = testlog.Logger

// Open calls Logger().Open, if a logger has been set.
var Open = testlog.Open

// SetLogger sets the test logger implementation for the current process.
// It must be called only once, at process startup.
var SetLogger = testlog.SetLogger

// Stat calls Logger().Stat, if a logger has been set.
var Stat = testlog.Stat
