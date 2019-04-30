package dwarfh

import (
	"cmd/internal/dwarf"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AppendSleb128Key = "AppendSleb128"

	AppendUleb128Key = "AppendUleb128"

	EnableLoggingKey = "EnableLogging"

	GetAbbrevKey = "GetAbbrev"

	HasChildrenKey = "HasChildren"

	PutAbstractFuncKey = "PutAbstractFunc"

	PutAttrsKey = "PutAttrs"

	PutConcreteFuncKey = "PutConcreteFunc"

	PutDefaultFuncKey = "PutDefaultFunc"

	PutInlinedFuncKey = "PutInlinedFunc"

	PutIntConstKey = "PutIntConst"

	PutRangesKey = "PutRanges"

	Sleb128putKey = "Sleb128put"

	Uleb128putKey = "Uleb128put"
)

func New() hctx.Map {
	return hctx.Map{

		AppendSleb128Key: AppendSleb128,

		AppendUleb128Key: AppendUleb128,

		EnableLoggingKey: EnableLogging,

		GetAbbrevKey: GetAbbrev,

		HasChildrenKey: HasChildren,

		PutAbstractFuncKey: PutAbstractFunc,

		PutAttrsKey: PutAttrs,

		PutConcreteFuncKey: PutConcreteFunc,

		PutDefaultFuncKey: PutDefaultFunc,

		PutInlinedFuncKey: PutInlinedFunc,

		PutIntConstKey: PutIntConst,

		PutRangesKey: PutRanges,

		Sleb128putKey: Sleb128put,

		Uleb128putKey: Uleb128put,
	}
}

// AppendSleb128 appends v to b using DWARF&#39;s signed LEB128 encoding.
var AppendSleb128 = dwarf.AppendSleb128

// AppendUleb128 appends v to b using DWARF&#39;s unsigned LEB128 encoding.
var AppendUleb128 = dwarf.AppendUleb128

var EnableLogging = dwarf.EnableLogging

// GetAbbrev returns the contents of the .debug_abbrev section.
var GetAbbrev = dwarf.GetAbbrev

// HasChildren reports whether &#39;die&#39; uses an abbrev that supports children.
var HasChildren = dwarf.HasChildren

// Emit DWARF attributes and child DIEs for an &#39;abstract&#39; subprogram.
// The abstract subprogram DIE for a function contains its
// location-independent attributes (name, type, etc). Other instances
// of the function (any inlined copy of it, or the single out-of-line
// &#39;concrete&#39; instance) will contain a pointer back to this abstract
// DIE (as a space-saving measure, so that name/type etc doesn&#39;t have
// to be repeated for each inlined copy).
var PutAbstractFunc = dwarf.PutAbstractFunc

// PutAttrs writes the attributes for a DIE to symbol &#39;s&#39;.
//
// Note that we can (and do) add arbitrary attributes to a DIE, but
// only the ones actually listed in the Abbrev will be written out.
var PutAttrs = dwarf.PutAttrs

// Emit DWARF attributes and child DIEs for a &#39;concrete&#39; subprogram,
// meaning the out-of-line copy of a function that was inlined at some
// point during the compilation of its containing package. The first
// attribute for a concrete DIE is a reference to the &#39;abstract&#39; DIE
// for the function (which holds location-independent attributes such
// as name, type), then the remainder of the attributes are specific
// to this instance (location, frame base, etc).
var PutConcreteFunc = dwarf.PutConcreteFunc

// Emit DWARF attributes and child DIEs for a subprogram. Here
// &#39;default&#39; implies that the function in question was not inlined
// when its containing package was compiled (hence there is no need to
// emit an abstract version for it to use as a base for inlined
// routine records).
var PutDefaultFunc = dwarf.PutDefaultFunc

// Emit DWARF attributes and child DIEs for an inlined subroutine. The
// first attribute of an inlined subroutine DIE is a reference back to
// its corresponding &#39;abstract&#39; DIE (containing location-independent
// attributes such as name, type, etc). Inlined subroutine DIEs can
// have other inlined subroutine DIEs as children.
var PutInlinedFunc = dwarf.PutInlinedFunc

// PutIntConst writes a DIE for an integer constant
var PutIntConst = dwarf.PutIntConst

// PutRanges writes a range table to sym. All addresses in ranges are
// relative to some base address. If base is not nil, then they&#39;re
// relative to the start of base. If base is nil, then the caller must
// arrange a base address some other way (such as a DW_AT_low_pc
// attribute).
var PutRanges = dwarf.PutRanges

// Sleb128put appends v to s using DWARF&#39;s signed LEB128 encoding.
var Sleb128put = dwarf.Sleb128put

// Uleb128put appends v to s using DWARF&#39;s unsigned LEB128 encoding.
var Uleb128put = dwarf.Uleb128put
