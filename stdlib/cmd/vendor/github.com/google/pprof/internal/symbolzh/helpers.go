package symbolzh

import (
	"cmd/vendor/github.com/google/pprof/internal/symbolz"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	SymbolizeKey = "Symbolize"
)

func New() hctx.Map {
	return hctx.Map{

		SymbolizeKey: Symbolize,
	}
}

// Symbolize symbolizes profile p by parsing data returned by a symbolz
// handler. syms receives the symbolz query (hex addresses separated by &#39;+&#39;)
// and returns the symbolz output in a string. If force is false, it will only
// symbolize locations from mappings not already marked as HasFunctions. Never
// attempts symbolization of addresses from unsymbolizable system
// mappings as those may look negative - e.g. &#34;[vsyscall]&#34;.
var Symbolize = symbolz.Symbolize
