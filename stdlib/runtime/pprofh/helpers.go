package pprofh

import (
	"runtime/pprof"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DoKey = "Do"

	ForLabelsKey = "ForLabels"

	LabelKey = "Label"

	LabelsKey = "Labels"

	LookupKey = "Lookup"

	NewProfileKey = "NewProfile"

	ProfilesKey = "Profiles"

	SetGoroutineLabelsKey = "SetGoroutineLabels"

	StartCPUProfileKey = "StartCPUProfile"

	StopCPUProfileKey = "StopCPUProfile"

	WithLabelsKey = "WithLabels"

	WriteHeapProfileKey = "WriteHeapProfile"
)

func New() hctx.Map {
	return hctx.Map{

		DoKey: Do,

		ForLabelsKey: ForLabels,

		LabelKey: Label,

		LabelsKey: Labels,

		LookupKey: Lookup,

		NewProfileKey: NewProfile,

		ProfilesKey: Profiles,

		SetGoroutineLabelsKey: SetGoroutineLabels,

		StartCPUProfileKey: StartCPUProfile,

		StopCPUProfileKey: StopCPUProfile,

		WithLabelsKey: WithLabels,

		WriteHeapProfileKey: WriteHeapProfile,
	}
}

// Do calls f with a copy of the parent context with the
// given labels added to the parent&#39;s label map.
// Each key/value pair in labels is inserted into the label map in the
// order provided, overriding any previous value for the same key.
// The augmented label map will be set for the duration of the call to f
// and restored once f returns.
var Do = pprof.Do

// ForLabels invokes f with each label set on the context.
// The function f should return true to continue iteration or false to stop iteration early.
var ForLabels = pprof.ForLabels

// Label returns the value of the label with the given key on ctx, and a boolean indicating
// whether that label exists.
var Label = pprof.Label

// Labels takes an even number of strings representing key-value pairs
// and makes a LabelSet containing them.
// A label overwrites a prior label with the same key.
var Labels = pprof.Labels

// Lookup returns the profile with the given name, or nil if no such profile exists.
var Lookup = pprof.Lookup

// NewProfile creates a new profile with the given name.
// If a profile with that name already exists, NewProfile panics.
// The convention is to use a &#39;import/path.&#39; prefix to create
// separate name spaces for each package.
// For compatibility with various tools that read pprof data,
// profile names should not contain spaces.
var NewProfile = pprof.NewProfile

// Profiles returns a slice of all the known profiles, sorted by name.
var Profiles = pprof.Profiles

// SetGoroutineLabels sets the current goroutine&#39;s labels to match ctx.
// This is a lower-level API than Do, which should be used instead when possible.
var SetGoroutineLabels = pprof.SetGoroutineLabels

// StartCPUProfile enables CPU profiling for the current process.
// While profiling, the profile will be buffered and written to w.
// StartCPUProfile returns an error if profiling is already enabled.
//
// On Unix-like systems, StartCPUProfile does not work by default for
// Go code built with -buildmode=c-archive or -buildmode=c-shared.
// StartCPUProfile relies on the SIGPROF signal, but that signal will
// be delivered to the main program&#39;s SIGPROF signal handler (if any)
// not to the one used by Go. To make it work, call os/signal.Notify
// for syscall.SIGPROF, but note that doing so may break any profiling
// being done by the main program.
var StartCPUProfile = pprof.StartCPUProfile

// StopCPUProfile stops the current CPU profile, if any.
// StopCPUProfile only returns after all the writes for the
// profile have completed.
var StopCPUProfile = pprof.StopCPUProfile

// WithLabels returns a new context.Context with the given labels added.
// A label overwrites a prior label with the same key.
var WithLabels = pprof.WithLabels

// WriteHeapProfile is shorthand for Lookup(&#34;heap&#34;).WriteTo(w, 0).
// It is preserved for backwards compatibility.
var WriteHeapProfile = pprof.WriteHeapProfile
