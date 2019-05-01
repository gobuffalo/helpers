package syslogh

import (
	"log/syslog"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DialKey = "Dial"

	NewLoggerKey = "NewLogger"
)

func New() hctx.Map {
	return hctx.Map{

		DialKey: Dial,

		NewLoggerKey: NewLogger,
	}
}

// Dial establishes a connection to a log daemon by connecting to
// address raddr on the specified network. Each write to the returned
// writer sends a log message with the facility and severity
// (from priority) and tag. If tag is empty, the os.Args[0] is used.
// If network is empty, Dial will connect to the local syslog server.
// Otherwise, see the documentation for net.Dial for valid values
// of network and raddr.
var Dial = syslog.Dial

// NewLogger creates a log.Logger whose output is written to the
// system log service with the specified priority, a combination of
// the syslog facility and severity. The logFlag argument is the flag
// set passed through to log.New to create the Logger.
var NewLogger = syslog.NewLogger
