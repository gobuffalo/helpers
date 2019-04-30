package logh

import (
	"log"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	FatalKey = "Fatal"

	FatalfKey = "Fatalf"

	FatallnKey = "Fatalln"

	FlagsKey = "Flags"

	NewKey = "New"

	OutputKey = "Output"

	PanicKey = "Panic"

	PanicfKey = "Panicf"

	PaniclnKey = "Panicln"

	PrefixKey = "Prefix"

	PrintKey = "Print"

	PrintfKey = "Printf"

	PrintlnKey = "Println"

	SetFlagsKey = "SetFlags"

	SetOutputKey = "SetOutput"

	SetPrefixKey = "SetPrefix"
)

func New() hctx.Map {
	return hctx.Map{

		FatalKey: Fatal,

		FatalfKey: Fatalf,

		FatallnKey: Fatalln,

		FlagsKey: Flags,

		NewKey: New,

		OutputKey: Output,

		PanicKey: Panic,

		PanicfKey: Panicf,

		PaniclnKey: Panicln,

		PrefixKey: Prefix,

		PrintKey: Print,

		PrintfKey: Printf,

		PrintlnKey: Println,

		SetFlagsKey: SetFlags,

		SetOutputKey: SetOutput,

		SetPrefixKey: SetPrefix,
	}
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
var Fatal = log.Fatal

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
var Fatalf = log.Fatalf

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
var Fatalln = log.Fatalln

// Flags returns the output flags for the standard logger.
var Flags = log.Flags

// New creates a new Logger. The out variable sets the
// destination to which log data will be written.
// The prefix appears at the beginning of each generated log line.
// The flag argument defines the logging properties.
var New = log.New

// Output writes the output for a logging event. The string s contains
// the text to print after the prefix specified by the flags of the
// Logger. A newline is appended if the last character of s is not
// already a newline. Calldepth is the count of the number of
// frames to skip when computing the file name and line number
// if Llongfile or Lshortfile is set; a value of 1 will print the details
// for the caller of Output.
var Output = log.Output

// Panic is equivalent to Print() followed by a call to panic().
var Panic = log.Panic

// Panicf is equivalent to Printf() followed by a call to panic().
var Panicf = log.Panicf

// Panicln is equivalent to Println() followed by a call to panic().
var Panicln = log.Panicln

// Prefix returns the output prefix for the standard logger.
var Prefix = log.Prefix

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
var Print = log.Print

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
var Printf = log.Printf

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
var Println = log.Println

// SetFlags sets the output flags for the standard logger.
var SetFlags = log.SetFlags

// SetOutput sets the output destination for the standard logger.
var SetOutput = log.SetOutput

// SetPrefix sets the output prefix for the standard logger.
var SetPrefix = log.SetPrefix
