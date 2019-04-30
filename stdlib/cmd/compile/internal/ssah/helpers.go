package ssah

import (
	"cmd/compile/internal/ssa"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	BuildFuncDebugKey = "BuildFuncDebug"

	CompileKey = "Compile"

	DebugNameMatchKey = "DebugNameMatch"

	IsGlobalAddrKey = "IsGlobalAddr"

	IsNewObjectKey = "IsNewObject"

	IsReadOnlyGlobalAddrKey = "IsReadOnlyGlobalAddr"

	IsSanitizerSafeAddrKey = "IsSanitizerSafeAddr"

	IsStackAddrKey = "IsStackAddr"

	LosesStmtMarkKey = "LosesStmtMark"

	NeedsFixUpKey = "NeedsFixUp"

	NewConfigKey = "NewConfig"

	NewFuncKey = "NewFunc"

	NewHTMLWriterKey = "NewHTMLWriter"

	NewSparseTreeHelperKey = "NewSparseTreeHelper"

	NewTypesKey = "NewTypes"

	PhaseOptionKey = "PhaseOption"

	ReachableBlocksKey = "ReachableBlocks"

	StructMakeOpKey = "StructMakeOp"
)

func New() hctx.Map {
	return hctx.Map{

		BuildFuncDebugKey: BuildFuncDebug,

		CompileKey: Compile,

		DebugNameMatchKey: DebugNameMatch,

		IsGlobalAddrKey: IsGlobalAddr,

		IsNewObjectKey: IsNewObject,

		IsReadOnlyGlobalAddrKey: IsReadOnlyGlobalAddr,

		IsSanitizerSafeAddrKey: IsSanitizerSafeAddr,

		IsStackAddrKey: IsStackAddr,

		LosesStmtMarkKey: LosesStmtMark,

		NeedsFixUpKey: NeedsFixUp,

		NewConfigKey: NewConfig,

		NewFuncKey: NewFunc,

		NewHTMLWriterKey: NewHTMLWriter,

		NewSparseTreeHelperKey: NewSparseTreeHelper,

		NewTypesKey: NewTypes,

		PhaseOptionKey: PhaseOption,

		ReachableBlocksKey: ReachableBlocks,

		StructMakeOpKey: StructMakeOp,
	}
}

// BuildFuncDebug returns debug information for f.
// f must be fully processed, so that each Value is where it will be when
// machine code is emitted.
var BuildFuncDebug = ssa.BuildFuncDebug

// Compile is the main entry point for this package.
// Compile modifies f so that on return:
//   · all Values in f map to 0 or 1 assembly instructions of the target architecture
//   · the order of f.Blocks is the order to emit the Blocks
//   · the order of b.Values is the order to emit the Values in each Block
//   · f has a non-nil regAlloc field
var Compile = ssa.Compile

var DebugNameMatch = ssa.DebugNameMatch

// IsGlobalAddr reports whether v is known to be an address of a global.
var IsGlobalAddr = ssa.IsGlobalAddr

// IsNewObject reports whether v is a pointer to a freshly allocated &amp; zeroed object at memory state mem.
// TODO: Be more precise. We really want &#34;IsNilPointer&#34; for the particular field in question.
// Right now, we can only detect a new object before any writes have been done to it.
// We could ignore non-pointer writes, writes to offsets which
// are known not to overlap the write in question, etc.
var IsNewObject = ssa.IsNewObject

// IsReadOnlyGlobalAddr reports whether v is known to be an address of a read-only global.
var IsReadOnlyGlobalAddr = ssa.IsReadOnlyGlobalAddr

// IsSanitizerSafeAddr reports whether v is known to be an address
// that doesn&#39;t need instrumentation.
var IsSanitizerSafeAddr = ssa.IsSanitizerSafeAddr

// IsStackAddr reports whether v is known to be an address of a stack slot.
var IsStackAddr = ssa.IsStackAddr

// LosesStmtMark reports whether a prog with op as loses its statement mark on the way to DWARF.
// The attributes from some opcodes are lost in translation.
// TODO: this is an artifact of how funcpctab combines information for instructions at a single PC.
// Should try to fix it there.
var LosesStmtMark = ssa.LosesStmtMark

// NeedsFixUp reports whether the division needs fix-up code.
var NeedsFixUp = ssa.NeedsFixUp

// NewConfig returns a new configuration object for the given architecture.
var NewConfig = ssa.NewConfig

// NewFunc returns a new, empty function object.
// Caller must set f.Config and f.Cache before using f.
var NewFunc = ssa.NewFunc

var NewHTMLWriter = ssa.NewHTMLWriter

// NewSparseTreeHelper returns a SparseTreeHelper for use
// in the gc package, for example in phi-function placement.
var NewSparseTreeHelper = ssa.NewSparseTreeHelper

// NewTypes creates and populates a Types.
var NewTypes = ssa.NewTypes

// PhaseOption sets the specified flag in the specified ssa phase,
// returning empty string if this was successful or a string explaining
// the error if it was not.
// A version of the phase name with &#34;_&#34; replaced by &#34; &#34; is also checked for a match.
// If the phase name begins a &#39;~&#39; then the rest of the underscores-replaced-with-blanks
// version is used as a regular expression to match the phase name(s).
//
// Special cases that have turned out to be useful:
//  ssa/check/on enables checking after each phase
//  ssa/all/time enables time reporting for all phases
//
// See gc/lex.go for dissection of the option string.
// Example uses:
//
// GO_GCFLAGS=-d=ssa/generic_cse/time,ssa/generic_cse/stats,ssa/generic_cse/debug=3 ./make.bash
//
// BOOT_GO_GCFLAGS=-d=&#39;ssa/~^.*scc$/off&#39; GO_GCFLAGS=&#39;-d=ssa/~^.*scc$/off&#39; ./make.bash
var PhaseOption = ssa.PhaseOption

// ReachableBlocks returns the reachable blocks in f.
var ReachableBlocks = ssa.ReachableBlocks

// StructMakeOp returns the opcode to construct a struct with the
// given number of fields.
var StructMakeOp = ssa.StructMakeOp
