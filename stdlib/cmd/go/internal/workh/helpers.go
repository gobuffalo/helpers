package workh

import (
	"cmd/go/internal/work"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AddBuildFlagsKey = "AddBuildFlags"

	BuildInitKey = "BuildInit"

	BuildInstallFuncKey = "BuildInstallFunc"

	FindExecCmdKey = "FindExecCmd"

	InstallPackagesKey = "InstallPackages"
)

func New() hctx.Map {
	return hctx.Map{

		AddBuildFlagsKey: AddBuildFlags,

		BuildInitKey: BuildInit,

		BuildInstallFuncKey: BuildInstallFunc,

		FindExecCmdKey: FindExecCmd,

		InstallPackagesKey: InstallPackages,
	}
}

// addBuildFlags adds the flags common to the build, clean, get,
// install, list, run, and test commands.
var AddBuildFlags = work.AddBuildFlags

var BuildInit = work.BuildInit

// BuildInstallFunc is the action for installing a single package or executable.
var BuildInstallFunc = work.BuildInstallFunc

// FindExecCmd derives the value of ExecCmd to use.
// It returns that value and leaves ExecCmd set for direct use.
var FindExecCmd = work.FindExecCmd

var InstallPackages = work.InstallPackages
