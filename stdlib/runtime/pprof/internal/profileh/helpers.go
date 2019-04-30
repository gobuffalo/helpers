package profileh

import (
	"runtime/pprof/internal/profile"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ParseKey = "Parse"

	ParseTracebacksKey = "ParseTracebacks"
)

func New() hctx.Map {
	return hctx.Map{

		ParseKey: Parse,

		ParseTracebacksKey: ParseTracebacks,
	}
}

// Parse parses a profile and checks for its validity. The input
// may be a gzip-compressed encoded protobuf or one of many legacy
// profile formats which may be unsupported in the future.
var Parse = profile.Parse

// ParseTracebacks parses a set of tracebacks and returns a newly
// populated profile. It will accept any text file and generate a
// Profile out of it with any hex addresses it can identify, including
// a process map if it can recognize one. Each sample will include a
// tag &#34;source&#34; with the addresses recognized in string format.
var ParseTracebacks = profile.ParseTracebacks
