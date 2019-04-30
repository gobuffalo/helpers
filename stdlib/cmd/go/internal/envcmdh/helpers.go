package envcmdh

import (
	"cmd/go/internal/envcmd"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ExtraEnvVarsKey = "ExtraEnvVars"

	ExtraEnvVarsCostlyKey = "ExtraEnvVarsCostly"

	MkEnvKey = "MkEnv"
)

func New() hctx.Map {
	return hctx.Map{

		ExtraEnvVarsKey: ExtraEnvVars,

		ExtraEnvVarsCostlyKey: ExtraEnvVarsCostly,

		MkEnvKey: MkEnv,
	}
}

// ExtraEnvVars returns environment variables that should not leak into child processes.
var ExtraEnvVars = envcmd.ExtraEnvVars

// ExtraEnvVarsCostly returns environment variables that should not leak into child processes
// but are costly to evaluate.
var ExtraEnvVarsCostly = envcmd.ExtraEnvVarsCostly

var MkEnv = envcmd.MkEnv
