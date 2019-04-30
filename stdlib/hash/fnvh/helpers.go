package fnvh

import (
	"hash/fnv"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	New128Key = "New128"

	New128aKey = "New128a"

	New32Key = "New32"

	New32aKey = "New32a"

	New64Key = "New64"

	New64aKey = "New64a"
)

func New() hctx.Map {
	return hctx.Map{

		New128Key: New128,

		New128aKey: New128a,

		New32Key: New32,

		New32aKey: New32a,

		New64Key: New64,

		New64aKey: New64a,
	}
}

// New128 returns a new 128-bit FNV-1 hash.Hash.
// Its Sum method will lay the value out in big-endian byte order.
var New128 = fnv.New128

// New128a returns a new 128-bit FNV-1a hash.Hash.
// Its Sum method will lay the value out in big-endian byte order.
var New128a = fnv.New128a

// New32 returns a new 32-bit FNV-1 hash.Hash.
// Its Sum method will lay the value out in big-endian byte order.
var New32 = fnv.New32

// New32a returns a new 32-bit FNV-1a hash.Hash.
// Its Sum method will lay the value out in big-endian byte order.
var New32a = fnv.New32a

// New64 returns a new 64-bit FNV-1 hash.Hash.
// Its Sum method will lay the value out in big-endian byte order.
var New64 = fnv.New64

// New64a returns a new 64-bit FNV-1a hash.Hash.
// Its Sum method will lay the value out in big-endian byte order.
var New64a = fnv.New64a
