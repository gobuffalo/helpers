package unitcheckerh

import (
	"cmd/vendor/golang.org/x/tools/go/analysis/unitchecker"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	MainKey = "Main"

	RunKey = "Run"
)

func New() hctx.Map {
	return hctx.Map{

		MainKey: Main,

		RunKey: Run,
	}
}

// Main is the main function of a vet-like analysis tool that must be
// invoked by a build system to analyze a single package.
//
// The protocol required by &#39;go vet -vettool=...&#39; is that the tool must support:
//
//      -flags          describe flags in JSON
//      -V=full         describe executable for build caching
//      foo.cfg         perform separate modular analyze on the single
//                      unit described by a JSON config file foo.cfg.
var Main = unitchecker.Main

// Run reads the *.cfg file, runs the analysis,
// and calls os.Exit with an appropriate error code.
// It assumes flags have already been set.
var Run = unitchecker.Run
