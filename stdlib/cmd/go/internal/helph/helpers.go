package helph

import (
	"cmd/go/internal/help"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	HelpKey = "Help"

	PrintUsageKey = "PrintUsage"
)

func New() hctx.Map {
	return hctx.Map{

		HelpKey: Help,

		PrintUsageKey: PrintUsage,
	}
}

// Help implements the &#39;help&#39; command.
var Help = help.Help

var PrintUsage = help.PrintUsage
