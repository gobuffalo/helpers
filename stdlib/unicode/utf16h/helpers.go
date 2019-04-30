package utf16h

import (
	"unicode/utf16"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DecodeKey = "Decode"

	DecodeRuneKey = "DecodeRune"

	EncodeKey = "Encode"

	EncodeRuneKey = "EncodeRune"

	IsSurrogateKey = "IsSurrogate"
)

func New() hctx.Map {
	return hctx.Map{

		DecodeKey: Decode,

		DecodeRuneKey: DecodeRune,

		EncodeKey: Encode,

		EncodeRuneKey: EncodeRune,

		IsSurrogateKey: IsSurrogate,
	}
}

// Decode returns the Unicode code point sequence represented
// by the UTF-16 encoding s.
var Decode = utf16.Decode

// DecodeRune returns the UTF-16 decoding of a surrogate pair.
// If the pair is not a valid UTF-16 surrogate pair, DecodeRune returns
// the Unicode replacement code point U+FFFD.
var DecodeRune = utf16.DecodeRune

// Encode returns the UTF-16 encoding of the Unicode code point sequence s.
var Encode = utf16.Encode

// EncodeRune returns the UTF-16 surrogate pair r1, r2 for the given rune.
// If the rune is not a valid Unicode code point or does not need encoding,
// EncodeRune returns U+FFFD, U+FFFD.
var EncodeRune = utf16.EncodeRune

// IsSurrogate reports whether the specified Unicode code point
// can appear in a surrogate pair.
var IsSurrogate = utf16.IsSurrogate
