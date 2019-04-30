package debugh

import (
	"cmd/vendor/golang.org/x/sys/windows/svc/debug"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewKey = "New"

	RunKey = "Run"
)

func New() hctx.Map {
	return hctx.Map{

		NewKey: New,

		RunKey: Run,
	}
}

// New creates new ConsoleLog.
var New = debug.New

// Run executes service name by calling appropriate handler function.
// The process is running on console, unlike real service. Use Ctrl+C to
// send &#34;Stop&#34; command to your service.
var Run = debug.Run
