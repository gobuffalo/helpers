package debugh

import (
	"runtime/debug"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	FreeOSMemoryKey = "FreeOSMemory"

	PrintStackKey = "PrintStack"

	ReadBuildInfoKey = "ReadBuildInfo"

	ReadGCStatsKey = "ReadGCStats"

	SetGCPercentKey = "SetGCPercent"

	SetMaxStackKey = "SetMaxStack"

	SetMaxThreadsKey = "SetMaxThreads"

	SetPanicOnFaultKey = "SetPanicOnFault"

	SetTracebackKey = "SetTraceback"

	StackKey = "Stack"

	WriteHeapDumpKey = "WriteHeapDump"
)

func New() hctx.Map {
	return hctx.Map{

		FreeOSMemoryKey: FreeOSMemory,

		PrintStackKey: PrintStack,

		ReadBuildInfoKey: ReadBuildInfo,

		ReadGCStatsKey: ReadGCStats,

		SetGCPercentKey: SetGCPercent,

		SetMaxStackKey: SetMaxStack,

		SetMaxThreadsKey: SetMaxThreads,

		SetPanicOnFaultKey: SetPanicOnFault,

		SetTracebackKey: SetTraceback,

		StackKey: Stack,

		WriteHeapDumpKey: WriteHeapDump,
	}
}

// FreeOSMemory forces a garbage collection followed by an
// attempt to return as much memory to the operating system
// as possible. (Even if this is not called, the runtime gradually
// returns memory to the operating system in a background task.)
var FreeOSMemory = debug.FreeOSMemory

// PrintStack prints to standard error the stack trace returned by runtime.Stack.
var PrintStack = debug.PrintStack

// ReadBuildInfo returns the build information embedded
// in the running binary. The information is available only
// in binaries built with module support.
var ReadBuildInfo = debug.ReadBuildInfo

// ReadGCStats reads statistics about garbage collection into stats.
// The number of entries in the pause history is system-dependent;
// stats.Pause slice will be reused if large enough, reallocated otherwise.
// ReadGCStats may use the full capacity of the stats.Pause slice.
// If stats.PauseQuantiles is non-empty, ReadGCStats fills it with quantiles
// summarizing the distribution of pause time. For example, if
// len(stats.PauseQuantiles) is 5, it will be filled with the minimum,
// 25%, 50%, 75%, and maximum pause times.
var ReadGCStats = debug.ReadGCStats

// SetGCPercent sets the garbage collection target percentage:
// a collection is triggered when the ratio of freshly allocated data
// to live data remaining after the previous collection reaches this percentage.
// SetGCPercent returns the previous setting.
// The initial setting is the value of the GOGC environment variable
// at startup, or 100 if the variable is not set.
// A negative percentage disables garbage collection.
var SetGCPercent = debug.SetGCPercent

// SetMaxStack sets the maximum amount of memory that
// can be used by a single goroutine stack.
// If any goroutine exceeds this limit while growing its stack,
// the program crashes.
// SetMaxStack returns the previous setting.
// The initial setting is 1 GB on 64-bit systems, 250 MB on 32-bit systems.
//
// SetMaxStack is useful mainly for limiting the damage done by
// goroutines that enter an infinite recursion. It only limits future
// stack growth.
var SetMaxStack = debug.SetMaxStack

// SetMaxThreads sets the maximum number of operating system
// threads that the Go program can use. If it attempts to use more than
// this many, the program crashes.
// SetMaxThreads returns the previous setting.
// The initial setting is 10,000 threads.
//
// The limit controls the number of operating system threads, not the number
// of goroutines. A Go program creates a new thread only when a goroutine
// is ready to run but all the existing threads are blocked in system calls, cgo calls,
// or are locked to other goroutines due to use of runtime.LockOSThread.
//
// SetMaxThreads is useful mainly for limiting the damage done by
// programs that create an unbounded number of threads. The idea is
// to take down the program before it takes down the operating system.
var SetMaxThreads = debug.SetMaxThreads

// SetPanicOnFault controls the runtime&#39;s behavior when a program faults
// at an unexpected (non-nil) address. Such faults are typically caused by
// bugs such as runtime memory corruption, so the default response is to crash
// the program. Programs working with memory-mapped files or unsafe
// manipulation of memory may cause faults at non-nil addresses in less
// dramatic situations; SetPanicOnFault allows such programs to request
// that the runtime trigger only a panic, not a crash.
// SetPanicOnFault applies only to the current goroutine.
// It returns the previous setting.
var SetPanicOnFault = debug.SetPanicOnFault

// SetTraceback sets the amount of detail printed by the runtime in
// the traceback it prints before exiting due to an unrecovered panic
// or an internal runtime error.
// The level argument takes the same values as the GOTRACEBACK
// environment variable. For example, SetTraceback(&#34;all&#34;) ensure
// that the program prints all goroutines when it crashes.
// See the package runtime documentation for details.
// If SetTraceback is called with a level lower than that of the
// environment variable, the call is ignored.
var SetTraceback = debug.SetTraceback

// Stack returns a formatted stack trace of the goroutine that calls it.
// It calls runtime.Stack with a large enough buffer to capture the entire trace.
var Stack = debug.Stack

// WriteHeapDump writes a description of the heap and the objects in
// it to the given file descriptor.
//
// WriteHeapDump suspends the execution of all goroutines until the heap
// dump is completely written.  Thus, the file descriptor must not be
// connected to a pipe or socket whose other end is in the same Go
// process; instead, use a temporary file or network socket.
//
// The heap dump format is defined at https://golang.org/s/go15heapdump.
var WriteHeapDump = debug.WriteHeapDump
