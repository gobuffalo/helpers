package testenvh

import (
	"internal/testenv"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	BuilderKey = "Builder"

	CleanCmdEnvKey = "CleanCmdEnv"

	GoToolKey = "GoTool"

	GoToolPathKey = "GoToolPath"

	HasCGOKey = "HasCGO"

	HasExecKey = "HasExec"

	HasExternalNetworkKey = "HasExternalNetwork"

	HasGoBuildKey = "HasGoBuild"

	HasGoRunKey = "HasGoRun"

	HasLinkKey = "HasLink"

	HasSrcKey = "HasSrc"

	HasSymlinkKey = "HasSymlink"

	MustHaveCGOKey = "MustHaveCGO"

	MustHaveExecKey = "MustHaveExec"

	MustHaveExternalNetworkKey = "MustHaveExternalNetwork"

	MustHaveGoBuildKey = "MustHaveGoBuild"

	MustHaveGoRunKey = "MustHaveGoRun"

	MustHaveLinkKey = "MustHaveLink"

	MustHaveSymlinkKey = "MustHaveSymlink"

	SkipFlakyKey = "SkipFlaky"

	SkipFlakyNetKey = "SkipFlakyNet"
)

func New() hctx.Map {
	return hctx.Map{

		BuilderKey: Builder,

		CleanCmdEnvKey: CleanCmdEnv,

		GoToolKey: GoTool,

		GoToolPathKey: GoToolPath,

		HasCGOKey: HasCGO,

		HasExecKey: HasExec,

		HasExternalNetworkKey: HasExternalNetwork,

		HasGoBuildKey: HasGoBuild,

		HasGoRunKey: HasGoRun,

		HasLinkKey: HasLink,

		HasSrcKey: HasSrc,

		HasSymlinkKey: HasSymlink,

		MustHaveCGOKey: MustHaveCGO,

		MustHaveExecKey: MustHaveExec,

		MustHaveExternalNetworkKey: MustHaveExternalNetwork,

		MustHaveGoBuildKey: MustHaveGoBuild,

		MustHaveGoRunKey: MustHaveGoRun,

		MustHaveLinkKey: MustHaveLink,

		MustHaveSymlinkKey: MustHaveSymlink,

		SkipFlakyKey: SkipFlaky,

		SkipFlakyNetKey: SkipFlakyNet,
	}
}

// Builder reports the name of the builder running this test
// (for example, &#34;linux-amd64&#34; or &#34;windows-386-gce&#34;).
// If the test is not running on the build infrastructure,
// Builder returns the empty string.
var Builder = testenv.Builder

// CleanCmdEnv will fill cmd.Env with the environment, excluding certain
// variables that could modify the behavior of the Go tools such as
// GODEBUG and GOTRACEBACK.
var CleanCmdEnv = testenv.CleanCmdEnv

// GoTool reports the path to the Go tool.
var GoTool = testenv.GoTool

// GoToolPath reports the path to the Go tool.
// It is a convenience wrapper around GoTool.
// If the tool is unavailable GoToolPath calls t.Skip.
// If the tool should be available and isn&#39;t, GoToolPath calls t.Fatal.
var GoToolPath = testenv.GoToolPath

// HasCGO reports whether the current system can use cgo.
var HasCGO = testenv.HasCGO

// HasExec reports whether the current system can start new processes
// using os.StartProcess or (more commonly) exec.Command.
var HasExec = testenv.HasExec

// HasExternalNetwork reports whether the current system can use
// external (non-localhost) networks.
var HasExternalNetwork = testenv.HasExternalNetwork

// HasGoBuild reports whether the current system can build programs with ``go build&#39;&#39;
// and then run them with os.StartProcess or exec.Command.
var HasGoBuild = testenv.HasGoBuild

// HasGoRun reports whether the current system can run programs with ``go run.&#39;&#39;
var HasGoRun = testenv.HasGoRun

// HasLink reports whether the current system can use os.Link.
var HasLink = testenv.HasLink

// HasSrc reports whether the entire source tree is available under GOROOT.
var HasSrc = testenv.HasSrc

// HasSymlink reports whether the current system can use os.Symlink.
var HasSymlink = testenv.HasSymlink

// MustHaveCGO calls t.Skip if cgo is not available.
var MustHaveCGO = testenv.MustHaveCGO

// MustHaveExec checks that the current system can start new processes
// using os.StartProcess or (more commonly) exec.Command.
// If not, MustHaveExec calls t.Skip with an explanation.
var MustHaveExec = testenv.MustHaveExec

// MustHaveExternalNetwork checks that the current system can use
// external (non-localhost) networks.
// If not, MustHaveExternalNetwork calls t.Skip with an explanation.
var MustHaveExternalNetwork = testenv.MustHaveExternalNetwork

// MustHaveGoBuild checks that the current system can build programs with ``go build&#39;&#39;
// and then run them with os.StartProcess or exec.Command.
// If not, MustHaveGoBuild calls t.Skip with an explanation.
var MustHaveGoBuild = testenv.MustHaveGoBuild

// MustHaveGoRun checks that the current system can run programs with ``go run.&#39;&#39;
// If not, MustHaveGoRun calls t.Skip with an explanation.
var MustHaveGoRun = testenv.MustHaveGoRun

// MustHaveLink reports whether the current system can use os.Link.
// If not, MustHaveLink calls t.Skip with an explanation.
var MustHaveLink = testenv.MustHaveLink

// MustHaveSymlink reports whether the current system can use os.Symlink.
// If not, MustHaveSymlink calls t.Skip with an explanation.
var MustHaveSymlink = testenv.MustHaveSymlink

var SkipFlaky = testenv.SkipFlaky

var SkipFlakyNet = testenv.SkipFlakyNet
