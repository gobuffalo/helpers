package traceh

import (
	"runtime/trace"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	StartKey = "Start"

	StopKey = "Stop"
)

func New() hctx.Map {
	return hctx.Map{

		StartKey: Start,

		StopKey: Stop,
	}
}

// Start enables tracing for the current program.
// While tracing, the trace will be buffered and written to w.
// Start returns an error if tracing is already enabled.
var Start = trace.Start

// Stop stops the current tracing, if any.
// Stop only returns after all the writes for the trace have completed.
var Stop = trace.Stop
