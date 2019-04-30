package sysh

import (
	"runtime/internal/sys"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	Bswap32Key = "Bswap32"

	Bswap32Key = "Bswap32"

	Bswap64Key = "Bswap64"

	Bswap64Key = "Bswap64"

	Ctz32Key = "Ctz32"

	Ctz32Key = "Ctz32"

	Ctz64Key = "Ctz64"

	Ctz64Key = "Ctz64"

	Ctz8Key = "Ctz8"

	Ctz8Key = "Ctz8"
)

func New() hctx.Map {
	return hctx.Map{

		Bswap32Key: Bswap32,

		Bswap32Key: Bswap32,

		Bswap64Key: Bswap64,

		Bswap64Key: Bswap64,

		Ctz32Key: Ctz32,

		Ctz32Key: Ctz32,

		Ctz64Key: Ctz64,

		Ctz64Key: Ctz64,

		Ctz8Key: Ctz8,

		Ctz8Key: Ctz8,
	}
}

var Bswap32 = sys.Bswap32

// Bswap32 returns its input with byte order reversed
// 0x01020304 -&gt; 0x04030201
var Bswap32 = sys.Bswap32

var Bswap64 = sys.Bswap64

// Bswap64 returns its input with byte order reversed
// 0x0102030405060708 -&gt; 0x0807060504030201
var Bswap64 = sys.Bswap64

var Ctz32 = sys.Ctz32

// Ctz32 counts trailing (low-order) zeroes,
// and if all are zero, then 32.
var Ctz32 = sys.Ctz32

var Ctz64 = sys.Ctz64

// Ctz64 counts trailing (low-order) zeroes,
// and if all are zero, then 64.
var Ctz64 = sys.Ctz64

var Ctz8 = sys.Ctz8

// Ctz8 returns the number of trailing zero bits in x; the result is 8 for x == 0.
var Ctz8 = sys.Ctz8
