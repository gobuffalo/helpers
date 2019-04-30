package fmth

import (
	"fmt"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ErrorfKey = "Errorf"

	FprintKey = "Fprint"

	FprintfKey = "Fprintf"

	FprintlnKey = "Fprintln"

	FscanKey = "Fscan"

	FscanfKey = "Fscanf"

	FscanlnKey = "Fscanln"

	PrintKey = "Print"

	PrintfKey = "Printf"

	PrintlnKey = "Println"

	ScanKey = "Scan"

	ScanfKey = "Scanf"

	ScanlnKey = "Scanln"

	SprintKey = "Sprint"

	SprintfKey = "Sprintf"

	SprintlnKey = "Sprintln"

	SscanKey = "Sscan"

	SscanfKey = "Sscanf"

	SscanlnKey = "Sscanln"
)

func New() hctx.Map {
	return hctx.Map{

		ErrorfKey: Errorf,

		FprintKey: Fprint,

		FprintfKey: Fprintf,

		FprintlnKey: Fprintln,

		FscanKey: Fscan,

		FscanfKey: Fscanf,

		FscanlnKey: Fscanln,

		PrintKey: Print,

		PrintfKey: Printf,

		PrintlnKey: Println,

		ScanKey: Scan,

		ScanfKey: Scanf,

		ScanlnKey: Scanln,

		SprintKey: Sprint,

		SprintfKey: Sprintf,

		SprintlnKey: Sprintln,

		SscanKey: Sscan,

		SscanfKey: Sscanf,

		SscanlnKey: Sscanln,
	}
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
var Errorf = fmt.Errorf

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
var Fprint = fmt.Fprint

// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
var Fprintf = fmt.Fprintf

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
var Fprintln = fmt.Fprintln

// Fscan scans text read from r, storing successive space-separated
// values into successive arguments. Newlines count as space. It
// returns the number of items successfully scanned. If that is less
// than the number of arguments, err will report why.
var Fscan = fmt.Fscan

// Fscanf scans text read from r, storing successive space-separated
// values into successive arguments as determined by the format. It
// returns the number of items successfully parsed.
// Newlines in the input must match newlines in the format.
var Fscanf = fmt.Fscanf

// Fscanln is similar to Fscan, but stops scanning at a newline and
// after the final item there must be a newline or EOF.
var Fscanln = fmt.Fscanln

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
var Print = fmt.Print

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
var Printf = fmt.Printf

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
var Println = fmt.Println

// Scan scans text read from standard input, storing successive
// space-separated values into successive arguments. Newlines count
// as space. It returns the number of items successfully scanned.
// If that is less than the number of arguments, err will report why.
var Scan = fmt.Scan

// Scanf scans text read from standard input, storing successive
// space-separated values into successive arguments as determined by
// the format. It returns the number of items successfully scanned.
// If that is less than the number of arguments, err will report why.
// Newlines in the input must match newlines in the format.
// The one exception: the verb %c always scans the next rune in the
// input, even if it is a space (or tab etc.) or newline.
var Scanf = fmt.Scanf

// Scanln is similar to Scan, but stops scanning at a newline and
// after the final item there must be a newline or EOF.
var Scanln = fmt.Scanln

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
var Sprint = fmt.Sprint

// Sprintf formats according to a format specifier and returns the resulting string.
var Sprintf = fmt.Sprintf

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
var Sprintln = fmt.Sprintln

// Sscan scans the argument string, storing successive space-separated
// values into successive arguments. Newlines count as space. It
// returns the number of items successfully scanned. If that is less
// than the number of arguments, err will report why.
var Sscan = fmt.Sscan

// Sscanf scans the argument string, storing successive space-separated
// values into successive arguments as determined by the format. It
// returns the number of items successfully parsed.
// Newlines in the input must match newlines in the format.
var Sscanf = fmt.Sscanf

// Sscanln is similar to Sscan, but stops scanning at a newline and
// after the final item there must be a newline or EOF.
var Sscanln = fmt.Sscanln
