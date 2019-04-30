package atomich

import (
	"sync/atomic"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AddInt32Key = "AddInt32"

	AddInt64Key = "AddInt64"

	AddUint32Key = "AddUint32"

	AddUint64Key = "AddUint64"

	AddUintptrKey = "AddUintptr"

	CompareAndSwapInt32Key = "CompareAndSwapInt32"

	CompareAndSwapInt64Key = "CompareAndSwapInt64"

	CompareAndSwapPointerKey = "CompareAndSwapPointer"

	CompareAndSwapUint32Key = "CompareAndSwapUint32"

	CompareAndSwapUint64Key = "CompareAndSwapUint64"

	CompareAndSwapUintptrKey = "CompareAndSwapUintptr"

	LoadInt32Key = "LoadInt32"

	LoadInt64Key = "LoadInt64"

	LoadPointerKey = "LoadPointer"

	LoadUint32Key = "LoadUint32"

	LoadUint64Key = "LoadUint64"

	LoadUintptrKey = "LoadUintptr"

	StoreInt32Key = "StoreInt32"

	StoreInt64Key = "StoreInt64"

	StorePointerKey = "StorePointer"

	StoreUint32Key = "StoreUint32"

	StoreUint64Key = "StoreUint64"

	StoreUintptrKey = "StoreUintptr"

	SwapInt32Key = "SwapInt32"

	SwapInt64Key = "SwapInt64"

	SwapPointerKey = "SwapPointer"

	SwapUint32Key = "SwapUint32"

	SwapUint64Key = "SwapUint64"

	SwapUintptrKey = "SwapUintptr"
)

func New() hctx.Map {
	return hctx.Map{

		AddInt32Key: AddInt32,

		AddInt64Key: AddInt64,

		AddUint32Key: AddUint32,

		AddUint64Key: AddUint64,

		AddUintptrKey: AddUintptr,

		CompareAndSwapInt32Key: CompareAndSwapInt32,

		CompareAndSwapInt64Key: CompareAndSwapInt64,

		CompareAndSwapPointerKey: CompareAndSwapPointer,

		CompareAndSwapUint32Key: CompareAndSwapUint32,

		CompareAndSwapUint64Key: CompareAndSwapUint64,

		CompareAndSwapUintptrKey: CompareAndSwapUintptr,

		LoadInt32Key: LoadInt32,

		LoadInt64Key: LoadInt64,

		LoadPointerKey: LoadPointer,

		LoadUint32Key: LoadUint32,

		LoadUint64Key: LoadUint64,

		LoadUintptrKey: LoadUintptr,

		StoreInt32Key: StoreInt32,

		StoreInt64Key: StoreInt64,

		StorePointerKey: StorePointer,

		StoreUint32Key: StoreUint32,

		StoreUint64Key: StoreUint64,

		StoreUintptrKey: StoreUintptr,

		SwapInt32Key: SwapInt32,

		SwapInt64Key: SwapInt64,

		SwapPointerKey: SwapPointer,

		SwapUint32Key: SwapUint32,

		SwapUint64Key: SwapUint64,

		SwapUintptrKey: SwapUintptr,
	}
}

// AddInt32 atomically adds delta to *addr and returns the new value.
var AddInt32 = atomic.AddInt32

// AddInt64 atomically adds delta to *addr and returns the new value.
var AddInt64 = atomic.AddInt64

// AddUint32 atomically adds delta to *addr and returns the new value.
// To subtract a signed positive constant value c from x, do AddUint32(&amp;x, ^uint32(c-1)).
// In particular, to decrement x, do AddUint32(&amp;x, ^uint32(0)).
var AddUint32 = atomic.AddUint32

// AddUint64 atomically adds delta to *addr and returns the new value.
// To subtract a signed positive constant value c from x, do AddUint64(&amp;x, ^uint64(c-1)).
// In particular, to decrement x, do AddUint64(&amp;x, ^uint64(0)).
var AddUint64 = atomic.AddUint64

// AddUintptr atomically adds delta to *addr and returns the new value.
var AddUintptr = atomic.AddUintptr

// CompareAndSwapInt32 executes the compare-and-swap operation for an int32 value.
var CompareAndSwapInt32 = atomic.CompareAndSwapInt32

// CompareAndSwapInt64 executes the compare-and-swap operation for an int64 value.
var CompareAndSwapInt64 = atomic.CompareAndSwapInt64

// CompareAndSwapPointer executes the compare-and-swap operation for a unsafe.Pointer value.
var CompareAndSwapPointer = atomic.CompareAndSwapPointer

// CompareAndSwapUint32 executes the compare-and-swap operation for a uint32 value.
var CompareAndSwapUint32 = atomic.CompareAndSwapUint32

// CompareAndSwapUint64 executes the compare-and-swap operation for a uint64 value.
var CompareAndSwapUint64 = atomic.CompareAndSwapUint64

// CompareAndSwapUintptr executes the compare-and-swap operation for a uintptr value.
var CompareAndSwapUintptr = atomic.CompareAndSwapUintptr

// LoadInt32 atomically loads *addr.
var LoadInt32 = atomic.LoadInt32

// LoadInt64 atomically loads *addr.
var LoadInt64 = atomic.LoadInt64

// LoadPointer atomically loads *addr.
var LoadPointer = atomic.LoadPointer

// LoadUint32 atomically loads *addr.
var LoadUint32 = atomic.LoadUint32

// LoadUint64 atomically loads *addr.
var LoadUint64 = atomic.LoadUint64

// LoadUintptr atomically loads *addr.
var LoadUintptr = atomic.LoadUintptr

// StoreInt32 atomically stores val into *addr.
var StoreInt32 = atomic.StoreInt32

// StoreInt64 atomically stores val into *addr.
var StoreInt64 = atomic.StoreInt64

// StorePointer atomically stores val into *addr.
var StorePointer = atomic.StorePointer

// StoreUint32 atomically stores val into *addr.
var StoreUint32 = atomic.StoreUint32

// StoreUint64 atomically stores val into *addr.
var StoreUint64 = atomic.StoreUint64

// StoreUintptr atomically stores val into *addr.
var StoreUintptr = atomic.StoreUintptr

// SwapInt32 atomically stores new into *addr and returns the previous *addr value.
var SwapInt32 = atomic.SwapInt32

// SwapInt64 atomically stores new into *addr and returns the previous *addr value.
var SwapInt64 = atomic.SwapInt64

// SwapPointer atomically stores new into *addr and returns the previous *addr value.
var SwapPointer = atomic.SwapPointer

// SwapUint32 atomically stores new into *addr and returns the previous *addr value.
var SwapUint32 = atomic.SwapUint32

// SwapUint64 atomically stores new into *addr and returns the previous *addr value.
var SwapUint64 = atomic.SwapUint64

// SwapUintptr atomically stores new into *addr and returns the previous *addr value.
var SwapUintptr = atomic.SwapUintptr
