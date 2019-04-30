package driverh

import (
	"cmd/vendor/github.com/google/pprof/internal/driver"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	AddCommandKey = "AddCommand"

	PProfKey = "PProf"

	SetVariableDefaultKey = "SetVariableDefault"
)

func New() hctx.Map {
	return hctx.Map{

		AddCommandKey: AddCommand,

		PProfKey: PProf,

		SetVariableDefaultKey: SetVariableDefault,
	}
}

// AddCommand adds an additional command to the set of commands
// accepted by pprof. This enables extensions to add new commands for
// specialized visualization formats. If the command specified already
// exists, it is overwritten.
var AddCommand = driver.AddCommand

// PProf acquires a profile, and symbolizes it using a profile
// manager. Then it generates a report formatted according to the
// options selected through the flags package.
var PProf = driver.PProf

// SetVariableDefault sets the default value for a pprof
// variable. This enables extensions to set their own defaults.
var SetVariableDefault = driver.SetVariableDefault
