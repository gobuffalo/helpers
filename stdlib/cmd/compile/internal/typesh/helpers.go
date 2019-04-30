package typesh

import (
	"cmd/compile/internal/types"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CleanroomDoKey = "CleanroomDo"

	FakeRecvTypeKey = "FakeRecvType"

	HaspointersKey = "Haspointers"

	Haspointers1Key = "Haspointers1"

	IdenticalKey = "Identical"

	IdenticalIgnoreTagsKey = "IdenticalIgnoreTags"

	ImportedPkgListKey = "ImportedPkgList"

	InternStringKey = "InternString"

	IsDclstackValidKey = "IsDclstackValid"

	IsExportedKey = "IsExported"

	MarkdclKey = "Markdcl"

	NewKey = "New"

	NewArrayKey = "NewArray"

	NewChanKey = "NewChan"

	NewChanArgsKey = "NewChanArgs"

	NewDDDArrayKey = "NewDDDArray"

	NewDDDFieldKey = "NewDDDField"

	NewFieldKey = "NewField"

	NewFuncArgsKey = "NewFuncArgs"

	NewMapKey = "NewMap"

	NewPkgKey = "NewPkg"

	NewPtrKey = "NewPtr"

	NewSliceKey = "NewSlice"

	NewTupleKey = "NewTuple"

	PopdclKey = "Popdcl"

	PushdclKey = "Pushdcl"

	SubstAnyKey = "SubstAny"
)

func New() hctx.Map {
	return hctx.Map{

		CleanroomDoKey: CleanroomDo,

		FakeRecvTypeKey: FakeRecvType,

		HaspointersKey: Haspointers,

		Haspointers1Key: Haspointers1,

		IdenticalKey: Identical,

		IdenticalIgnoreTagsKey: IdenticalIgnoreTags,

		ImportedPkgListKey: ImportedPkgList,

		InternStringKey: InternString,

		IsDclstackValidKey: IsDclstackValid,

		IsExportedKey: IsExported,

		MarkdclKey: Markdcl,

		NewKey: New,

		NewArrayKey: NewArray,

		NewChanKey: NewChan,

		NewChanArgsKey: NewChanArgs,

		NewDDDArrayKey: NewDDDArray,

		NewDDDFieldKey: NewDDDField,

		NewFieldKey: NewField,

		NewFuncArgsKey: NewFuncArgs,

		NewMapKey: NewMap,

		NewPkgKey: NewPkg,

		NewPtrKey: NewPtr,

		NewSliceKey: NewSlice,

		NewTupleKey: NewTuple,

		PopdclKey: Popdcl,

		PushdclKey: Pushdcl,

		SubstAnyKey: SubstAny,
	}
}

// CleanroomDo invokes f in an environment with no preexisting packages.
// For testing of import/export only.
var CleanroomDo = types.CleanroomDo

// FakeRecvType returns the singleton type used for interface method receivers.
var FakeRecvType = types.FakeRecvType

var Haspointers = types.Haspointers

var Haspointers1 = types.Haspointers1

// Identical reports whether t1 and t2 are identical types, following
// the spec rules. Receiver parameter types are ignored.
var Identical = types.Identical

// IdenticalIgnoreTags is like Identical, but it ignores struct tags
// for struct identity.
var IdenticalIgnoreTags = types.IdenticalIgnoreTags

// ImportedPkgList returns the list of directly imported packages.
// The list is sorted by package path.
var ImportedPkgList = types.ImportedPkgList

var InternString = types.InternString

var IsDclstackValid = types.IsDclstackValid

// IsExported reports whether name is an exported Go symbol (that is,
// whether it begins with an upper-case letter).
var IsExported = types.IsExported

// Markdcl records the start of a new block scope for declarations.
var Markdcl = types.Markdcl

// New returns a new Type of the specified kind.
var New = types.New

// NewArray returns a new fixed-length array Type.
var NewArray = types.NewArray

// NewChan returns a new chan Type with direction dir.
var NewChan = types.NewChan

// NewChanArgs returns a new TCHANARGS type for channel type c.
var NewChanArgs = types.NewChanArgs

// NewDDDArray returns a new [...]T array Type.
var NewDDDArray = types.NewDDDArray

// NewDDDField returns a new TDDDFIELD type for slice type s.
var NewDDDField = types.NewDDDField

var NewField = types.NewField

// NewFuncArgs returns a new TFUNCARGS type for func type f.
var NewFuncArgs = types.NewFuncArgs

// NewMap returns a new map Type with key type k and element (aka value) type v.
var NewMap = types.NewMap

// NewPkg returns a new Pkg for the given package path and name.
// Unless name is the empty string, if the package exists already,
// the existing package name and the provided name must match.
var NewPkg = types.NewPkg

// NewPtr returns the pointer type pointing to t.
var NewPtr = types.NewPtr

// NewSlice returns the slice Type with element type elem.
var NewSlice = types.NewSlice

var NewTuple = types.NewTuple

// Popdcl pops the innermost block scope and restores all symbol declarations
// to their previous state.
var Popdcl = types.Popdcl

// Pushdcl pushes the current declaration for symbol s (if any) so that
// it can be shadowed by a new declaration within a nested block scope.
var Pushdcl = types.Pushdcl

// SubstAny walks t, replacing instances of &#34;any&#34; with successive
// elements removed from types.  It returns the substituted type.
var SubstAny = types.SubstAny
