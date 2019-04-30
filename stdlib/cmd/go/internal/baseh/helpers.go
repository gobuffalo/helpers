package baseh

import (
	"cmd/go/internal/base"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AddBuildFlagsNXKey = "AddBuildFlagsNX"

	AddKnownFlagKey = "AddKnownFlag"

	AtExitKey = "AtExit"

	EnvForDirKey = "EnvForDir"

	ErrorfKey = "Errorf"

	ExitKey = "Exit"

	ExitIfErrorsKey = "ExitIfErrors"

	ExpandScannerKey = "ExpandScanner"

	FatalfKey = "Fatalf"

	GOFLAGSKey = "GOFLAGS"

	InitGOFLAGSKey = "InitGOFLAGS"

	IsTestFileKey = "IsTestFile"

	MergeEnvListsKey = "MergeEnvLists"

	RelPathsKey = "RelPaths"

	RunKey = "Run"

	RunStdinKey = "RunStdin"

	SetExitStatusKey = "SetExitStatus"

	SetFromGOFLAGSKey = "SetFromGOFLAGS"

	ShortPathKey = "ShortPath"

	StartSigHandlersKey = "StartSigHandlers"

	ToolKey = "Tool"
)

func New() hctx.Map {
	return hctx.Map{

		AddBuildFlagsNXKey: AddBuildFlagsNX,

		AddKnownFlagKey: AddKnownFlag,

		AtExitKey: AtExit,

		EnvForDirKey: EnvForDir,

		ErrorfKey: Errorf,

		ExitKey: Exit,

		ExitIfErrorsKey: ExitIfErrors,

		ExpandScannerKey: ExpandScanner,

		FatalfKey: Fatalf,

		GOFLAGSKey: GOFLAGS,

		InitGOFLAGSKey: InitGOFLAGS,

		IsTestFileKey: IsTestFile,

		MergeEnvListsKey: MergeEnvLists,

		RelPathsKey: RelPaths,

		RunKey: Run,

		RunStdinKey: RunStdin,

		SetExitStatusKey: SetExitStatus,

		SetFromGOFLAGSKey: SetFromGOFLAGS,

		ShortPathKey: ShortPath,

		StartSigHandlersKey: StartSigHandlers,

		ToolKey: Tool,
	}
}

// AddBuildFlagsNX adds the -n and -x build flags to the flag set.
var AddBuildFlagsNX = base.AddBuildFlagsNX

// AddKnownFlag adds name to the list of known flags for use in $GOFLAGS.
var AddKnownFlag = base.AddKnownFlag

var AtExit = base.AtExit

// EnvForDir returns a copy of the environment
// suitable for running in the given directory.
// The environment is the current process&#39;s environment
// but with an updated $PWD, so that an os.Getwd in the
// child will be faster.
var EnvForDir = base.EnvForDir

var Errorf = base.Errorf

var Exit = base.Exit

var ExitIfErrors = base.ExitIfErrors

// ExpandScanner expands a scanner.List error into all the errors in the list.
// The default Error method only shows the first error
// and does not shorten paths.
var ExpandScanner = base.ExpandScanner

var Fatalf = base.Fatalf

// GOFLAGS returns the flags from $GOFLAGS.
// The list can be assumed to contain one string per flag,
// with each string either beginning with -name or --name.
var GOFLAGS = base.GOFLAGS

// InitGOFLAGS initializes the goflags list from $GOFLAGS.
// If goflags is already initialized, it does nothing.
var InitGOFLAGS = base.InitGOFLAGS

// IsTestFile reports whether the source file is a set of tests and should therefore
// be excluded from coverage analysis.
var IsTestFile = base.IsTestFile

// MergeEnvLists merges the two environment lists such that
// variables with the same name in &#34;in&#34; replace those in &#34;out&#34;.
// This always returns a newly allocated slice.
var MergeEnvLists = base.MergeEnvLists

// RelPaths returns a copy of paths with absolute paths
// made relative to the current directory if they would be shorter.
var RelPaths = base.RelPaths

// Run runs the command, with stdout and stderr
// connected to the go command&#39;s own stdout and stderr.
// If the command fails, Run reports the error using Errorf.
var Run = base.Run

// RunStdin is like run but connects Stdin.
var RunStdin = base.RunStdin

var SetExitStatus = base.SetExitStatus

// SetFromGOFLAGS sets the flags in the given flag set using settings in $GOFLAGS.
var SetFromGOFLAGS = base.SetFromGOFLAGS

// ShortPath returns an absolute or relative name for path, whatever is shorter.
var ShortPath = base.ShortPath

// StartSigHandlers starts the signal handlers.
var StartSigHandlers = base.StartSigHandlers

// Tool returns the path to the named tool (for example, &#34;vet&#34;).
// If the tool cannot be found, Tool exits the process.
var Tool = base.Tool
