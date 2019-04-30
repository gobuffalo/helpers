package pprofh

import (
	"net/http/pprof"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CmdlineKey = "Cmdline"

	HandlerKey = "Handler"

	IndexKey = "Index"

	ProfileKey = "Profile"

	SymbolKey = "Symbol"

	TraceKey = "Trace"
)

func New() hctx.Map {
	return hctx.Map{

		CmdlineKey: Cmdline,

		HandlerKey: Handler,

		IndexKey: Index,

		ProfileKey: Profile,

		SymbolKey: Symbol,

		TraceKey: Trace,
	}
}

// Cmdline responds with the running program&#39;s
// command line, with arguments separated by NUL bytes.
// The package initialization registers it as /debug/pprof/cmdline.
var Cmdline = pprof.Cmdline

// Handler returns an HTTP handler that serves the named profile.
var Handler = pprof.Handler

// Index responds with the pprof-formatted profile named by the request.
// For example, &#34;/debug/pprof/heap&#34; serves the &#34;heap&#34; profile.
// Index responds to a request for &#34;/debug/pprof/&#34; with an HTML page
// listing the available profiles.
var Index = pprof.Index

// Profile responds with the pprof-formatted cpu profile.
// Profiling lasts for duration specified in seconds GET parameter, or for 30 seconds if not specified.
// The package initialization registers it as /debug/pprof/profile.
var Profile = pprof.Profile

// Symbol looks up the program counters listed in the request,
// responding with a table mapping program counters to function names.
// The package initialization registers it as /debug/pprof/symbol.
var Symbol = pprof.Symbol

// Trace responds with the execution trace in binary form.
// Tracing lasts for duration specified in seconds GET parameter, or for 1 second if not specified.
// The package initialization registers it as /debug/pprof/trace.
var Trace = pprof.Trace
