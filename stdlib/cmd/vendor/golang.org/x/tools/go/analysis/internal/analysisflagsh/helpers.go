package analysisflagsh

import (
	"cmd/vendor/golang.org/x/tools/go/analysis/internal/analysisflags"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	HelpKey = "Help"

	ParseKey = "Parse"

	PrintPlainKey = "PrintPlain"
)

func New() hctx.Map {
	return hctx.Map{

		HelpKey: Help,

		ParseKey: Parse,

		PrintPlainKey: PrintPlain,
	}
}

// Help implements the help subcommand for a multichecker or unitchecker
// style command. The optional args specify the analyzers to describe.
// Help calls log.Fatal if no such analyzer exists.
var Help = analysisflags.Help

// Parse creates a flag for each of the analyzer&#39;s flags,
// including (in multi mode) a flag named after the analyzer,
// parses the flags, then filters and returns the list of
// analyzers enabled by flags.
var Parse = analysisflags.Parse

// PrintPlain prints a diagnostic in plain text form,
// with context specified by the -c flag.
var PrintPlain = analysisflags.PrintPlain
