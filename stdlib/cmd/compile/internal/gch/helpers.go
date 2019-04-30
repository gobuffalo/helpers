package gch

import (
	"cmd/compile/internal/gc"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AddAuxKey = "AddAux"

	AddAux2Key = "AddAux2"

	AddrAutoKey = "AddrAuto"

	AddrconstKey = "Addrconst"

	AutoVarKey = "AutoVar"

	AuxOffsetKey = "AuxOffset"

	CheckLoweredGetClosurePtrKey = "CheckLoweredGetClosurePtr"

	CheckLoweredPhiKey = "CheckLoweredPhi"

	DumpKey = "Dump"

	ExitKey = "Exit"

	FatalfKey = "Fatalf"

	IncomparableFieldKey = "IncomparableField"

	IsAliasKey = "IsAlias"

	IsComparableKey = "IsComparable"

	IsRegularMemoryKey = "IsRegularMemory"

	IsconstKey = "Isconst"

	MainKey = "Main"

	PatchKey = "Patch"

	RndKey = "Rnd"

	WarnKey = "Warn"

	WarnlKey = "Warnl"
)

func New() hctx.Map {
	return hctx.Map{

		AddAuxKey: AddAux,

		AddAux2Key: AddAux2,

		AddrAutoKey: AddrAuto,

		AddrconstKey: Addrconst,

		AutoVarKey: AutoVar,

		AuxOffsetKey: AuxOffset,

		CheckLoweredGetClosurePtrKey: CheckLoweredGetClosurePtr,

		CheckLoweredPhiKey: CheckLoweredPhi,

		DumpKey: Dump,

		ExitKey: Exit,

		FatalfKey: Fatalf,

		IncomparableFieldKey: IncomparableField,

		IsAliasKey: IsAlias,

		IsComparableKey: IsComparable,

		IsRegularMemoryKey: IsRegularMemory,

		IsconstKey: Isconst,

		MainKey: Main,

		PatchKey: Patch,

		RndKey: Rnd,

		WarnKey: Warn,

		WarnlKey: Warnl,
	}
}

// AddAux adds the offset in the aux fields (AuxInt and Aux) of v to a.
var AddAux = gc.AddAux

var AddAux2 = gc.AddAux2

var AddrAuto = gc.AddrAuto

var Addrconst = gc.Addrconst

// AutoVar returns a *Node and int64 representing the auto variable and offset within it
// where v should be spilled.
var AutoVar = gc.AutoVar

var AuxOffset = gc.AuxOffset

// CheckLoweredGetClosurePtr checks that v is the first instruction in the function&#39;s entry block.
// The output of LoweredGetClosurePtr is generally hardwired to the correct register.
// That register contains the closure pointer on closure entry.
var CheckLoweredGetClosurePtr = gc.CheckLoweredGetClosurePtr

// CheckLoweredPhi checks that regalloc and stackalloc correctly handled phi values.
// Called during ssaGenValue.
var CheckLoweredPhi = gc.CheckLoweredPhi

var Dump = gc.Dump

var Exit = gc.Exit

var Fatalf = gc.Fatalf

// IncomparableField returns an incomparable Field of struct Type t, if any.
var IncomparableField = gc.IncomparableField

var IsAlias = gc.IsAlias

// IsComparable reports whether t is a comparable type.
var IsComparable = gc.IsComparable

// IsRegularMemory reports whether t can be compared/hashed as regular memory.
var IsRegularMemory = gc.IsRegularMemory

var Isconst = gc.Isconst

// Main parses flags and Go source files specified in the command-line
// arguments, type-checks the parsed Go package, compiles functions to machine
// code, and finally writes the compiled package definition to disk.
var Main = gc.Main

var Patch = gc.Patch

var Rnd = gc.Rnd

var Warn = gc.Warn

var Warnl = gc.Warnl
