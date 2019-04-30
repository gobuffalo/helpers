package atomich

import (
	"runtime/internal/atomic"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	And8Key = "And8"

	And8Key = "And8"

	And8Key = "And8"

	And8Key = "And8"

	And8Key = "And8"

	And8Key = "And8"

	And8Key = "And8"

	And8Key = "And8"

	And8Key = "And8"

	CasKey = "Cas"

	CasKey = "Cas"

	Cas64Key = "Cas64"

	Cas64Key = "Cas64"

	Cas64Key = "Cas64"

	Cas64Key = "Cas64"

	Cas64Key = "Cas64"

	Cas64Key = "Cas64"

	Cas64Key = "Cas64"

	Cas64Key = "Cas64"

	Cas64Key = "Cas64"

	CasRelKey = "CasRel"

	CasRelKey = "CasRel"

	CasRelKey = "CasRel"

	CasRelKey = "CasRel"

	CasRelKey = "CasRel"

	CasRelKey = "CasRel"

	CasRelKey = "CasRel"

	CasRelKey = "CasRel"

	CasRelKey = "CasRel"

	Casp1Key = "Casp1"

	Casp1Key = "Casp1"

	CasuintptrKey = "Casuintptr"

	CasuintptrKey = "Casuintptr"

	LoadKey = "Load"

	LoadKey = "Load"

	LoadKey = "Load"

	LoadKey = "Load"

	LoadKey = "Load"

	LoadKey = "Load"

	LoadKey = "Load"

	LoadKey = "Load"

	LoadKey = "Load"

	Load64Key = "Load64"

	Load64Key = "Load64"

	Load64Key = "Load64"

	Load64Key = "Load64"

	Load64Key = "Load64"

	Load64Key = "Load64"

	Load64Key = "Load64"

	Load64Key = "Load64"

	Load64Key = "Load64"

	LoadAcqKey = "LoadAcq"

	LoadAcqKey = "LoadAcq"

	LoadAcqKey = "LoadAcq"

	LoadAcqKey = "LoadAcq"

	LoadAcqKey = "LoadAcq"

	LoadAcqKey = "LoadAcq"

	LoadAcqKey = "LoadAcq"

	LoadAcqKey = "LoadAcq"

	LoadAcqKey = "LoadAcq"

	Loadint64Key = "Loadint64"

	Loadint64Key = "Loadint64"

	LoadpKey = "Loadp"

	LoadpKey = "Loadp"

	LoadpKey = "Loadp"

	LoadpKey = "Loadp"

	LoadpKey = "Loadp"

	LoadpKey = "Loadp"

	LoadpKey = "Loadp"

	LoadpKey = "Loadp"

	LoadpKey = "Loadp"

	LoaduintKey = "Loaduint"

	LoaduintKey = "Loaduint"

	LoaduintptrKey = "Loaduintptr"

	LoaduintptrKey = "Loaduintptr"

	Or8Key = "Or8"

	Or8Key = "Or8"

	Or8Key = "Or8"

	Or8Key = "Or8"

	Or8Key = "Or8"

	Or8Key = "Or8"

	Or8Key = "Or8"

	Or8Key = "Or8"

	Or8Key = "Or8"

	StoreKey = "Store"

	StoreKey = "Store"

	StoreKey = "Store"

	StoreKey = "Store"

	StoreKey = "Store"

	StoreKey = "Store"

	StoreKey = "Store"

	StoreKey = "Store"

	StoreKey = "Store"

	Store64Key = "Store64"

	Store64Key = "Store64"

	Store64Key = "Store64"

	Store64Key = "Store64"

	Store64Key = "Store64"

	Store64Key = "Store64"

	Store64Key = "Store64"

	Store64Key = "Store64"

	Store64Key = "Store64"

	StoreRelKey = "StoreRel"

	StoreRelKey = "StoreRel"

	StoreRelKey = "StoreRel"

	StoreRelKey = "StoreRel"

	StoreRelKey = "StoreRel"

	StoreRelKey = "StoreRel"

	StoreRelKey = "StoreRel"

	StoreRelKey = "StoreRel"

	StoreRelKey = "StoreRel"

	StorepNoWBKey = "StorepNoWB"

	StorepNoWBKey = "StorepNoWB"

	StorepNoWBKey = "StorepNoWB"

	StorepNoWBKey = "StorepNoWB"

	StorepNoWBKey = "StorepNoWB"

	StorepNoWBKey = "StorepNoWB"

	StorepNoWBKey = "StorepNoWB"

	StorepNoWBKey = "StorepNoWB"

	StorepNoWBKey = "StorepNoWB"

	StoreuintptrKey = "Storeuintptr"

	StoreuintptrKey = "Storeuintptr"

	XaddKey = "Xadd"

	XaddKey = "Xadd"

	XaddKey = "Xadd"

	XaddKey = "Xadd"

	XaddKey = "Xadd"

	XaddKey = "Xadd"

	XaddKey = "Xadd"

	XaddKey = "Xadd"

	XaddKey = "Xadd"

	Xadd64Key = "Xadd64"

	Xadd64Key = "Xadd64"

	Xadd64Key = "Xadd64"

	Xadd64Key = "Xadd64"

	Xadd64Key = "Xadd64"

	Xadd64Key = "Xadd64"

	Xadd64Key = "Xadd64"

	Xadd64Key = "Xadd64"

	Xadd64Key = "Xadd64"

	Xaddint64Key = "Xaddint64"

	Xaddint64Key = "Xaddint64"

	XadduintptrKey = "Xadduintptr"

	XadduintptrKey = "Xadduintptr"

	XadduintptrKey = "Xadduintptr"

	XadduintptrKey = "Xadduintptr"

	XadduintptrKey = "Xadduintptr"

	XadduintptrKey = "Xadduintptr"

	XadduintptrKey = "Xadduintptr"

	XadduintptrKey = "Xadduintptr"

	XadduintptrKey = "Xadduintptr"

	XchgKey = "Xchg"

	XchgKey = "Xchg"

	XchgKey = "Xchg"

	XchgKey = "Xchg"

	XchgKey = "Xchg"

	XchgKey = "Xchg"

	XchgKey = "Xchg"

	XchgKey = "Xchg"

	XchgKey = "Xchg"

	Xchg64Key = "Xchg64"

	Xchg64Key = "Xchg64"

	Xchg64Key = "Xchg64"

	Xchg64Key = "Xchg64"

	Xchg64Key = "Xchg64"

	Xchg64Key = "Xchg64"

	Xchg64Key = "Xchg64"

	Xchg64Key = "Xchg64"

	Xchg64Key = "Xchg64"

	XchguintptrKey = "Xchguintptr"

	XchguintptrKey = "Xchguintptr"

	XchguintptrKey = "Xchguintptr"

	XchguintptrKey = "Xchguintptr"

	XchguintptrKey = "Xchguintptr"

	XchguintptrKey = "Xchguintptr"

	XchguintptrKey = "Xchguintptr"

	XchguintptrKey = "Xchguintptr"

	XchguintptrKey = "Xchguintptr"
)

func New() hctx.Map {
	return hctx.Map{

		And8Key: And8,

		And8Key: And8,

		And8Key: And8,

		And8Key: And8,

		And8Key: And8,

		And8Key: And8,

		And8Key: And8,

		And8Key: And8,

		And8Key: And8,

		CasKey: Cas,

		CasKey: Cas,

		Cas64Key: Cas64,

		Cas64Key: Cas64,

		Cas64Key: Cas64,

		Cas64Key: Cas64,

		Cas64Key: Cas64,

		Cas64Key: Cas64,

		Cas64Key: Cas64,

		Cas64Key: Cas64,

		Cas64Key: Cas64,

		CasRelKey: CasRel,

		CasRelKey: CasRel,

		CasRelKey: CasRel,

		CasRelKey: CasRel,

		CasRelKey: CasRel,

		CasRelKey: CasRel,

		CasRelKey: CasRel,

		CasRelKey: CasRel,

		CasRelKey: CasRel,

		Casp1Key: Casp1,

		Casp1Key: Casp1,

		CasuintptrKey: Casuintptr,

		CasuintptrKey: Casuintptr,

		LoadKey: Load,

		LoadKey: Load,

		LoadKey: Load,

		LoadKey: Load,

		LoadKey: Load,

		LoadKey: Load,

		LoadKey: Load,

		LoadKey: Load,

		LoadKey: Load,

		Load64Key: Load64,

		Load64Key: Load64,

		Load64Key: Load64,

		Load64Key: Load64,

		Load64Key: Load64,

		Load64Key: Load64,

		Load64Key: Load64,

		Load64Key: Load64,

		Load64Key: Load64,

		LoadAcqKey: LoadAcq,

		LoadAcqKey: LoadAcq,

		LoadAcqKey: LoadAcq,

		LoadAcqKey: LoadAcq,

		LoadAcqKey: LoadAcq,

		LoadAcqKey: LoadAcq,

		LoadAcqKey: LoadAcq,

		LoadAcqKey: LoadAcq,

		LoadAcqKey: LoadAcq,

		Loadint64Key: Loadint64,

		Loadint64Key: Loadint64,

		LoadpKey: Loadp,

		LoadpKey: Loadp,

		LoadpKey: Loadp,

		LoadpKey: Loadp,

		LoadpKey: Loadp,

		LoadpKey: Loadp,

		LoadpKey: Loadp,

		LoadpKey: Loadp,

		LoadpKey: Loadp,

		LoaduintKey: Loaduint,

		LoaduintKey: Loaduint,

		LoaduintptrKey: Loaduintptr,

		LoaduintptrKey: Loaduintptr,

		Or8Key: Or8,

		Or8Key: Or8,

		Or8Key: Or8,

		Or8Key: Or8,

		Or8Key: Or8,

		Or8Key: Or8,

		Or8Key: Or8,

		Or8Key: Or8,

		Or8Key: Or8,

		StoreKey: Store,

		StoreKey: Store,

		StoreKey: Store,

		StoreKey: Store,

		StoreKey: Store,

		StoreKey: Store,

		StoreKey: Store,

		StoreKey: Store,

		StoreKey: Store,

		Store64Key: Store64,

		Store64Key: Store64,

		Store64Key: Store64,

		Store64Key: Store64,

		Store64Key: Store64,

		Store64Key: Store64,

		Store64Key: Store64,

		Store64Key: Store64,

		Store64Key: Store64,

		StoreRelKey: StoreRel,

		StoreRelKey: StoreRel,

		StoreRelKey: StoreRel,

		StoreRelKey: StoreRel,

		StoreRelKey: StoreRel,

		StoreRelKey: StoreRel,

		StoreRelKey: StoreRel,

		StoreRelKey: StoreRel,

		StoreRelKey: StoreRel,

		StorepNoWBKey: StorepNoWB,

		StorepNoWBKey: StorepNoWB,

		StorepNoWBKey: StorepNoWB,

		StorepNoWBKey: StorepNoWB,

		StorepNoWBKey: StorepNoWB,

		StorepNoWBKey: StorepNoWB,

		StorepNoWBKey: StorepNoWB,

		StorepNoWBKey: StorepNoWB,

		StorepNoWBKey: StorepNoWB,

		StoreuintptrKey: Storeuintptr,

		StoreuintptrKey: Storeuintptr,

		XaddKey: Xadd,

		XaddKey: Xadd,

		XaddKey: Xadd,

		XaddKey: Xadd,

		XaddKey: Xadd,

		XaddKey: Xadd,

		XaddKey: Xadd,

		XaddKey: Xadd,

		XaddKey: Xadd,

		Xadd64Key: Xadd64,

		Xadd64Key: Xadd64,

		Xadd64Key: Xadd64,

		Xadd64Key: Xadd64,

		Xadd64Key: Xadd64,

		Xadd64Key: Xadd64,

		Xadd64Key: Xadd64,

		Xadd64Key: Xadd64,

		Xadd64Key: Xadd64,

		Xaddint64Key: Xaddint64,

		Xaddint64Key: Xaddint64,

		XadduintptrKey: Xadduintptr,

		XadduintptrKey: Xadduintptr,

		XadduintptrKey: Xadduintptr,

		XadduintptrKey: Xadduintptr,

		XadduintptrKey: Xadduintptr,

		XadduintptrKey: Xadduintptr,

		XadduintptrKey: Xadduintptr,

		XadduintptrKey: Xadduintptr,

		XadduintptrKey: Xadduintptr,

		XchgKey: Xchg,

		XchgKey: Xchg,

		XchgKey: Xchg,

		XchgKey: Xchg,

		XchgKey: Xchg,

		XchgKey: Xchg,

		XchgKey: Xchg,

		XchgKey: Xchg,

		XchgKey: Xchg,

		Xchg64Key: Xchg64,

		Xchg64Key: Xchg64,

		Xchg64Key: Xchg64,

		Xchg64Key: Xchg64,

		Xchg64Key: Xchg64,

		Xchg64Key: Xchg64,

		Xchg64Key: Xchg64,

		Xchg64Key: Xchg64,

		Xchg64Key: Xchg64,

		XchguintptrKey: Xchguintptr,

		XchguintptrKey: Xchguintptr,

		XchguintptrKey: Xchguintptr,

		XchguintptrKey: Xchguintptr,

		XchguintptrKey: Xchguintptr,

		XchguintptrKey: Xchguintptr,

		XchguintptrKey: Xchguintptr,

		XchguintptrKey: Xchguintptr,

		XchguintptrKey: Xchguintptr,
	}
}

// go:noescape
var And8 = atomic.And8

// go:noescape
var And8 = atomic.And8

// go:noescape
var And8 = atomic.And8

// go:noescape
var And8 = atomic.And8

// go:noescape
var And8 = atomic.And8

// go:nosplit
// go:noinline
var And8 = atomic.And8

// go:nosplit
var And8 = atomic.And8

// go:noescape
var And8 = atomic.And8

// go:noescape
var And8 = atomic.And8

// go:noescape
var Cas = atomic.Cas

// go:nosplit
// go:noinline
var Cas = atomic.Cas

// go:noescape
var Cas64 = atomic.Cas64

// go:noescape
var Cas64 = atomic.Cas64

// go:nosplit
var Cas64 = atomic.Cas64

// go:noescape
var Cas64 = atomic.Cas64

// go:noescape
var Cas64 = atomic.Cas64

// go:noescape
var Cas64 = atomic.Cas64

// go:noescape
var Cas64 = atomic.Cas64

// go:noescape
var Cas64 = atomic.Cas64

// go:nosplit
// go:noinline
var Cas64 = atomic.Cas64

// go:noescape
var CasRel = atomic.CasRel

// go:noescape
var CasRel = atomic.CasRel

// go:noescape
var CasRel = atomic.CasRel

// go:noescape
var CasRel = atomic.CasRel

// go:noescape
var CasRel = atomic.CasRel

// go:nosplit
// go:noinline
var CasRel = atomic.CasRel

// go:noescape
var CasRel = atomic.CasRel

// go:noescape
var CasRel = atomic.CasRel

// go:noescape
var CasRel = atomic.CasRel

// go:nosplit
// go:noinline
var Casp1 = atomic.Casp1

// NO go:noescape annotation; see atomic_pointer.go.
var Casp1 = atomic.Casp1

// go:noescape
var Casuintptr = atomic.Casuintptr

// go:nosplit
// go:noinline
var Casuintptr = atomic.Casuintptr

// go:nosplit
// go:noinline
var Load = atomic.Load

// go:noescape
var Load = atomic.Load

// go:nosplit
// go:noinline
var Load = atomic.Load

// go:nosplit
// go:noinline
var Load = atomic.Load

// go:noescape
var Load = atomic.Load

// go:noescape
var Load = atomic.Load

// go:noescape
var Load = atomic.Load

// go:nosplit
// go:noinline
var Load = atomic.Load

// go:noescape
var Load = atomic.Load

// go:nosplit
// go:noinline
var Load64 = atomic.Load64

// go:noescape
var Load64 = atomic.Load64

// go:nosplit
var Load64 = atomic.Load64

// go:noescape
var Load64 = atomic.Load64

// go:nosplit
// go:noinline
var Load64 = atomic.Load64

// go:noescape
var Load64 = atomic.Load64

// go:noescape
var Load64 = atomic.Load64

// go:noescape
var Load64 = atomic.Load64

// go:nosplit
// go:noinline
var Load64 = atomic.Load64

// go:noescape
var LoadAcq = atomic.LoadAcq

// go:nosplit
// go:noinline
var LoadAcq = atomic.LoadAcq

// go:noescape
var LoadAcq = atomic.LoadAcq

// go:nosplit
// go:noinline
var LoadAcq = atomic.LoadAcq

// go:noescape
var LoadAcq = atomic.LoadAcq

// go:noescape
var LoadAcq = atomic.LoadAcq

// go:nosplit
// go:noinline
var LoadAcq = atomic.LoadAcq

// go:noescape
var LoadAcq = atomic.LoadAcq

// go:nosplit
// go:noinline
var LoadAcq = atomic.LoadAcq

// go:nosplit
// go:noinline
var Loadint64 = atomic.Loadint64

// go:noescape
var Loadint64 = atomic.Loadint64

// go:nosplit
// go:noinline
var Loadp = atomic.Loadp

// go:noescape
var Loadp = atomic.Loadp

// go:noescape
var Loadp = atomic.Loadp

// go:noescape
var Loadp = atomic.Loadp

// go:noescape
var Loadp = atomic.Loadp

// go:nosplit
// go:noinline
var Loadp = atomic.Loadp

// go:nosplit
// go:noinline
var Loadp = atomic.Loadp

// go:noescape
var Loadp = atomic.Loadp

// go:nosplit
// go:noinline
var Loadp = atomic.Loadp

// go:nosplit
// go:noinline
var Loaduint = atomic.Loaduint

// go:noescape
var Loaduint = atomic.Loaduint

// go:noescape
var Loaduintptr = atomic.Loaduintptr

// go:nosplit
// go:noinline
var Loaduintptr = atomic.Loaduintptr

// go:nosplit
var Or8 = atomic.Or8

// go:noescape
var Or8 = atomic.Or8

// go:noescape
var Or8 = atomic.Or8

// go:nosplit
// go:noinline
var Or8 = atomic.Or8

// go:noescape
var Or8 = atomic.Or8

// go:noescape
var Or8 = atomic.Or8

// go:noescape
var Or8 = atomic.Or8

// go:noescape
var Or8 = atomic.Or8

// go:noescape
var Or8 = atomic.Or8

// go:noescape
var Store = atomic.Store

// go:noinline
// go:nosplit
var Store = atomic.Store

// go:noescape
var Store = atomic.Store

// go:noescape
var Store = atomic.Store

// go:noescape
var Store = atomic.Store

// go:noescape
var Store = atomic.Store

// go:noescape
var Store = atomic.Store

// go:noescape
var Store = atomic.Store

// go:nosplit
// go:noinline
var Store = atomic.Store

// go:nosplit
// go:noinline
var Store64 = atomic.Store64

// go:nosplit
var Store64 = atomic.Store64

// go:noescape
var Store64 = atomic.Store64

// go:noescape
var Store64 = atomic.Store64

// go:noescape
var Store64 = atomic.Store64

// go:noescape
var Store64 = atomic.Store64

// go:noescape
var Store64 = atomic.Store64

// go:noinline
// go:nosplit
var Store64 = atomic.Store64

// go:noescape
var Store64 = atomic.Store64

// go:noescape
var StoreRel = atomic.StoreRel

// go:noescape
var StoreRel = atomic.StoreRel

// go:noescape
var StoreRel = atomic.StoreRel

// go:noescape
var StoreRel = atomic.StoreRel

// go:noescape
var StoreRel = atomic.StoreRel

// go:noescape
var StoreRel = atomic.StoreRel

// go:nosplit
// go:noinline
var StoreRel = atomic.StoreRel

// go:noescape
var StoreRel = atomic.StoreRel

// go:noinline
// go:nosplit
var StoreRel = atomic.StoreRel

// NO go:noescape annotation; see atomic_pointer.go.
var StorepNoWB = atomic.StorepNoWB

// NO go:noescape annotation; see atomic_pointer.go.
var StorepNoWB = atomic.StorepNoWB

// Not noescape -- it installs a pointer to addr.
var StorepNoWB = atomic.StorepNoWB

// NO go:noescape annotation; see atomic_pointer.go.
var StorepNoWB = atomic.StorepNoWB

// go:noinline
// go:nosplit
var StorepNoWB = atomic.StorepNoWB

// NO go:noescape annotation; see atomic_pointer.go.
var StorepNoWB = atomic.StorepNoWB

// NO go:noescape annotation; see atomic_pointer.go.
// go:noinline
// go:nosplit
var StorepNoWB = atomic.StorepNoWB

// NO go:noescape annotation; see atomic_pointer.go.
var StorepNoWB = atomic.StorepNoWB

// StorepNoWB performs *ptr = val atomically and without a write
// barrier.
//
// NO go:noescape annotation; see atomic_pointer.go.
var StorepNoWB = atomic.StorepNoWB

// go:nosplit
// go:noinline
var Storeuintptr = atomic.Storeuintptr

// go:noescape
var Storeuintptr = atomic.Storeuintptr

// go:noescape
var Xadd = atomic.Xadd

// go:noescape
var Xadd = atomic.Xadd

// go:nosplit
// go:noinline
var Xadd = atomic.Xadd

// go:noescape
var Xadd = atomic.Xadd

// go:noescape
var Xadd = atomic.Xadd

// Atomic add and return new value.
// go:nosplit
var Xadd = atomic.Xadd

// go:noescape
var Xadd = atomic.Xadd

// go:noescape
var Xadd = atomic.Xadd

// go:noescape
var Xadd = atomic.Xadd

// go:noescape
var Xadd64 = atomic.Xadd64

// go:noescape
var Xadd64 = atomic.Xadd64

// go:nosplit
// go:noinline
var Xadd64 = atomic.Xadd64

// go:noescape
var Xadd64 = atomic.Xadd64

// go:noescape
var Xadd64 = atomic.Xadd64

// go:nosplit
var Xadd64 = atomic.Xadd64

// go:noescape
var Xadd64 = atomic.Xadd64

// go:noescape
var Xadd64 = atomic.Xadd64

// go:noescape
var Xadd64 = atomic.Xadd64

// go:nosplit
// go:noinline
var Xaddint64 = atomic.Xaddint64

// go:noescape
var Xaddint64 = atomic.Xaddint64

// go:nosplit
// go:noinline
var Xadduintptr = atomic.Xadduintptr

// go:noescape
var Xadduintptr = atomic.Xadduintptr

// go:noescape
var Xadduintptr = atomic.Xadduintptr

// go:noescape
var Xadduintptr = atomic.Xadduintptr

// go:noescape
var Xadduintptr = atomic.Xadduintptr

// go:noescape
var Xadduintptr = atomic.Xadduintptr

// go:noescape
var Xadduintptr = atomic.Xadduintptr

// go:noescape
var Xadduintptr = atomic.Xadduintptr

// go:noescape
var Xadduintptr = atomic.Xadduintptr

// go:noescape
var Xchg = atomic.Xchg

// go:noescape
var Xchg = atomic.Xchg

// go:noescape
var Xchg = atomic.Xchg

// go:noescape
var Xchg = atomic.Xchg

// go:noescape
var Xchg = atomic.Xchg

// go:noescape
var Xchg = atomic.Xchg

// go:nosplit
var Xchg = atomic.Xchg

// go:noescape
var Xchg = atomic.Xchg

// go:nosplit
// go:noinline
var Xchg = atomic.Xchg

// go:noescape
var Xchg64 = atomic.Xchg64

// go:noescape
var Xchg64 = atomic.Xchg64

// go:noescape
var Xchg64 = atomic.Xchg64

// go:noescape
var Xchg64 = atomic.Xchg64

// go:nosplit
// go:noinline
var Xchg64 = atomic.Xchg64

// go:nosplit
var Xchg64 = atomic.Xchg64

// go:noescape
var Xchg64 = atomic.Xchg64

// go:noescape
var Xchg64 = atomic.Xchg64

// go:noescape
var Xchg64 = atomic.Xchg64

// go:noescape
var Xchguintptr = atomic.Xchguintptr

// go:noescape
var Xchguintptr = atomic.Xchguintptr

// go:noescape
var Xchguintptr = atomic.Xchguintptr

// go:nosplit
var Xchguintptr = atomic.Xchguintptr

// go:noescape
var Xchguintptr = atomic.Xchguintptr

// go:nosplit
// go:noinline
var Xchguintptr = atomic.Xchguintptr

// go:noescape
var Xchguintptr = atomic.Xchguintptr

// go:noescape
var Xchguintptr = atomic.Xchguintptr

// go:noescape
var Xchguintptr = atomic.Xchguintptr
