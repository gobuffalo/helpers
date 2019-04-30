package traceh

import (
	"internal/trace"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	GoroutineStatsKey = "GoroutineStats"

	MutatorUtilizationKey = "MutatorUtilization"

	NewMMUCurveKey = "NewMMUCurve"

	NewWriterKey = "NewWriter"

	ParseKey = "Parse"

	PrintKey = "Print"

	PrintEventKey = "PrintEvent"

	RelatedGoroutinesKey = "RelatedGoroutines"
)

func New() hctx.Map {
	return hctx.Map{

		GoroutineStatsKey: GoroutineStats,

		MutatorUtilizationKey: MutatorUtilization,

		NewMMUCurveKey: NewMMUCurve,

		NewWriterKey: NewWriter,

		ParseKey: Parse,

		PrintKey: Print,

		PrintEventKey: PrintEvent,

		RelatedGoroutinesKey: RelatedGoroutines,
	}
}

// GoroutineStats generates statistics for all goroutines in the trace.
var GoroutineStats = trace.GoroutineStats

// MutatorUtilization returns a set of mutator utilization functions
// for the given trace. Each function will always end with 0
// utilization. The bounds of each function are implicit in the first
// and last event; outside of these bounds each function is undefined.
//
// If the UtilPerProc flag is not given, this always returns a single
// utilization function. Otherwise, it returns one function per P.
var MutatorUtilization = trace.MutatorUtilization

// NewMMUCurve returns an MMU curve for the given mutator utilization
// function.
var NewMMUCurve = trace.NewMMUCurve

var NewWriter = trace.NewWriter

// Parse parses, post-processes and verifies the trace.
var Parse = trace.Parse

// Print dumps events to stdout. For debugging.
var Print = trace.Print

// PrintEvent dumps the event to stdout. For debugging.
var PrintEvent = trace.PrintEvent

// RelatedGoroutines finds a set of goroutines related to goroutine goid.
var RelatedGoroutines = trace.RelatedGoroutines
