package traceh

import (
	"runtime/trace"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	IsEnabledKey = "IsEnabled"

	LogKey = "Log"

	LogfKey = "Logf"

	NewTaskKey = "NewTask"

	StartKey = "Start"

	StartRegionKey = "StartRegion"

	StopKey = "Stop"

	WithRegionKey = "WithRegion"
)

func New() hctx.Map {
	return hctx.Map{

		IsEnabledKey: IsEnabled,

		LogKey: Log,

		LogfKey: Logf,

		NewTaskKey: NewTask,

		StartKey: Start,

		StartRegionKey: StartRegion,

		StopKey: Stop,

		WithRegionKey: WithRegion,
	}
}

// IsEnabled reports whether tracing is enabled.
// The information is advisory only. The tracing status
// may have changed by the time this function returns.
var IsEnabled = trace.IsEnabled

// Log emits a one-off event with the given category and message.
// Category can be empty and the API assumes there are only a handful of
// unique categories in the system.
var Log = trace.Log

// Logf is like Log, but the value is formatted using the specified format spec.
var Logf = trace.Logf

// NewTask creates a task instance with the type taskType and returns
// it along with a Context that carries the task.
// If the input context contains a task, the new task is its subtask.
//
// The taskType is used to classify task instances. Analysis tools
// like the Go execution tracer may assume there are only a bounded
// number of unique task types in the system.
//
// The returned end function is used to mark the task&#39;s end.
// The trace tool measures task latency as the time between task creation
// and when the end function is called, and provides the latency
// distribution per task type.
// If the end function is called multiple times, only the first
// call is used in the latency measurement.
//
//   ctx, task := trace.NewTask(ctx, &#34;awesomeTask&#34;)
//   trace.WithRegion(ctx, &#34;preparation&#34;, prepWork)
//   // preparation of the task
//   go func() {  // continue processing the task in a separate goroutine.
//       defer task.End()
//       trace.WithRegion(ctx, &#34;remainingWork&#34;, remainingWork)
//   }()
var NewTask = trace.NewTask

// Start enables tracing for the current program.
// While tracing, the trace will be buffered and written to w.
// Start returns an error if tracing is already enabled.
var Start = trace.Start

// StartRegion starts a region and returns a function for marking the
// end of the region. The returned Region&#39;s End function must be called
// from the same goroutine where the region was started.
// Within each goroutine, regions must nest. That is, regions started
// after this region must be ended before this region can be ended.
// Recommended usage is
//
//     defer trace.StartRegion(ctx, &#34;myTracedRegion&#34;).End()
var StartRegion = trace.StartRegion

// Stop stops the current tracing, if any.
// Stop only returns after all the writes for the trace have completed.
var Stop = trace.Stop

// WithRegion starts a region associated with its calling goroutine, runs fn,
// and then ends the region. If the context carries a task, the region is
// associated with the task. Otherwise, the region is attached to the background
// task.
//
// The regionType is used to classify regions, so there should be only a
// handful of unique region types.
var WithRegion = trace.WithRegion
