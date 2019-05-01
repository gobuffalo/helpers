package flagh

import (
	"flag"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ArgKey = "Arg"

	ArgsKey = "Args"

	BoolKey = "Bool"

	BoolVarKey = "BoolVar"

	DurationKey = "Duration"

	DurationVarKey = "DurationVar"

	Float64Key = "Float64"

	Float64VarKey = "Float64Var"

	IntKey = "Int"

	Int64Key = "Int64"

	Int64VarKey = "Int64Var"

	IntVarKey = "IntVar"

	LookupKey = "Lookup"

	NArgKey = "NArg"

	NFlagKey = "NFlag"

	NewFlagSetKey = "NewFlagSet"

	ParseKey = "Parse"

	ParsedKey = "Parsed"

	PrintDefaultsKey = "PrintDefaults"

	SetKey = "Set"

	StringKey = "String"

	StringVarKey = "StringVar"

	UintKey = "Uint"

	Uint64Key = "Uint64"

	Uint64VarKey = "Uint64Var"

	UintVarKey = "UintVar"

	UnquoteUsageKey = "UnquoteUsage"

	VarKey = "Var"

	VisitKey = "Visit"

	VisitAllKey = "VisitAll"
)

func New() hctx.Map {
	return hctx.Map{

		ArgKey: Arg,

		ArgsKey: Args,

		BoolKey: Bool,

		BoolVarKey: BoolVar,

		DurationKey: Duration,

		DurationVarKey: DurationVar,

		Float64Key: Float64,

		Float64VarKey: Float64Var,

		IntKey: Int,

		Int64Key: Int64,

		Int64VarKey: Int64Var,

		IntVarKey: IntVar,

		LookupKey: Lookup,

		NArgKey: NArg,

		NFlagKey: NFlag,

		NewFlagSetKey: NewFlagSet,

		ParseKey: Parse,

		ParsedKey: Parsed,

		PrintDefaultsKey: PrintDefaults,

		SetKey: Set,

		StringKey: String,

		StringVarKey: StringVar,

		UintKey: Uint,

		Uint64Key: Uint64,

		Uint64VarKey: Uint64Var,

		UintVarKey: UintVar,

		UnquoteUsageKey: UnquoteUsage,

		VarKey: Var,

		VisitKey: Visit,

		VisitAllKey: VisitAll,
	}
}

// Arg returns the i&#39;th command-line argument. Arg(0) is the first remaining argument
// after flags have been processed. Arg returns an empty string if the
// requested element does not exist.
var Arg = flag.Arg

// Args returns the non-flag command-line arguments.
var Args = flag.Args

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
var Bool = flag.Bool

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
var BoolVar = flag.BoolVar

// Duration defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
var Duration = flag.Duration

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
var DurationVar = flag.DurationVar

// Float64 defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
var Float64 = flag.Float64

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
var Float64Var = flag.Float64Var

// Int defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
var Int = flag.Int

// Int64 defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
var Int64 = flag.Int64

// Int64Var defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
var Int64Var = flag.Int64Var

// IntVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
var IntVar = flag.IntVar

// Lookup returns the Flag structure of the named command-line flag,
// returning nil if none exists.
var Lookup = flag.Lookup

// NArg is the number of arguments remaining after flags have been processed.
var NArg = flag.NArg

// NFlag returns the number of command-line flags that have been set.
var NFlag = flag.NFlag

// NewFlagSet returns a new, empty flag set with the specified name and
// error handling property.
var NewFlagSet = flag.NewFlagSet

// Parse parses the command-line flags from os.Args[1:]. Must be called
// after all flags are defined and before flags are accessed by the program.
var Parse = flag.Parse

// Parsed reports whether the command-line flags have been parsed.
var Parsed = flag.Parsed

// PrintDefaults prints, to standard error unless configured otherwise,
// a usage message showing the default settings of all defined
// command-line flags.
// For an integer valued flag x, the default output has the form
// 	-x int
// 		usage-message-for-x (default 7)
// The usage message will appear on a separate line for anything but
// a bool flag with a one-byte name. For bool flags, the type is
// omitted and if the flag name is one byte the usage message appears
// on the same line. The parenthetical default is omitted if the
// default is the zero value for the type. The listed type, here int,
// can be changed by placing a back-quoted name in the flag&#39;s usage
// string; the first such item in the message is taken to be a parameter
// name to show in the message and the back quotes are stripped from
// the message when displayed. For instance, given
// 	flag.String(&#34;I&#34;, &#34;&#34;, &#34;search `directory` for include files&#34;)
// the output will be
// 	-I directory
// 		search directory for include files.
var PrintDefaults = flag.PrintDefaults

// Set sets the value of the named command-line flag.
var Set = flag.Set

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
var String = flag.String

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
var StringVar = flag.StringVar

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint variable that stores the value of the flag.
var Uint = flag.Uint

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
var Uint64 = flag.Uint64

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
var Uint64Var = flag.Uint64Var

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
var UintVar = flag.UintVar

// UnquoteUsage extracts a back-quoted name from the usage
// string for a flag and returns it and the un-quoted usage.
// Given &#34;a `name` to show&#34; it returns (&#34;name&#34;, &#34;a name to show&#34;).
// If there are no back quotes, the name is an educated guess of the
// type of the flag&#39;s value, or the empty string if the flag is boolean.
var UnquoteUsage = flag.UnquoteUsage

// Var defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
var Var = flag.Var

// Visit visits the command-line flags in lexicographical order, calling fn
// for each. It visits only those flags that have been set.
var Visit = flag.Visit

// VisitAll visits the command-line flags in lexicographical order, calling
// fn for each. It visits all flags, even those not set.
var VisitAll = flag.VisitAll
