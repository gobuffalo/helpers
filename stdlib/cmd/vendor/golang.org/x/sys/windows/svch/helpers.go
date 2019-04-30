package svch

import (
	"cmd/vendor/golang.org/x/sys/windows/svc"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	IsAnInteractiveSessionKey = "IsAnInteractiveSession"

	RunKey = "Run"

	StatusHandleKey = "StatusHandle"
)

func New() hctx.Map {
	return hctx.Map{

		IsAnInteractiveSessionKey: IsAnInteractiveSession,

		RunKey: Run,

		StatusHandleKey: StatusHandle,
	}
}

// IsAnInteractiveSession determines if calling process is running interactively.
// It queries the process token for membership in the Interactive group.
// http://stackoverflow.com/questions/2668851/how-do-i-detect-that-my-application-is-running-as-service-or-in-an-interactive-s
var IsAnInteractiveSession = svc.IsAnInteractiveSession

// Run executes service name by calling appropriate handler function.
var Run = svc.Run

// StatusHandle returns service status handle. It is safe to call this function
// from inside the Handler.Execute because then it is guaranteed to be set.
// This code will have to change once multiple services are possible per process.
var StatusHandle = svc.StatusHandle
