package ascii85h

import (
	"encoding/ascii85"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DecodeKey = "Decode"

	EncodeKey = "Encode"

	MaxEncodedLenKey = "MaxEncodedLen"

	NewDecoderKey = "NewDecoder"

	NewEncoderKey = "NewEncoder"
)

func New() hctx.Map {
	return hctx.Map{

		DecodeKey: Decode,

		EncodeKey: Encode,

		MaxEncodedLenKey: MaxEncodedLen,

		NewDecoderKey: NewDecoder,

		NewEncoderKey: NewEncoder,
	}
}

// Decode decodes src into dst, returning both the number
// of bytes written to dst and the number consumed from src.
// If src contains invalid ascii85 data, Decode will return the
// number of bytes successfully written and a CorruptInputError.
// Decode ignores space and control characters in src.
// Often, ascii85-encoded data is wrapped in &lt;~ and ~&gt; symbols.
// Decode expects these to have been stripped by the caller.
//
// If flush is true, Decode assumes that src represents the
// end of the input stream and processes it completely rather
// than wait for the completion of another 32-bit block.
//
// NewDecoder wraps an io.Reader interface around Decode.
var Decode = ascii85.Decode

// Encode encodes src into at most MaxEncodedLen(len(src))
// bytes of dst, returning the actual number of bytes written.
//
// The encoding handles 4-byte chunks, using a special encoding
// for the last fragment, so Encode is not appropriate for use on
// individual blocks of a large data stream. Use NewEncoder() instead.
//
// Often, ascii85-encoded data is wrapped in &lt;~ and ~&gt; symbols.
// Encode does not add these.
var Encode = ascii85.Encode

// MaxEncodedLen returns the maximum length of an encoding of n source bytes.
var MaxEncodedLen = ascii85.MaxEncodedLen

// NewDecoder constructs a new ascii85 stream decoder.
var NewDecoder = ascii85.NewDecoder

// NewEncoder returns a new ascii85 stream encoder. Data written to
// the returned writer will be encoded and then written to w.
// Ascii85 encodings operate in 32-bit blocks; when finished
// writing, the caller must Close the returned encoder to flush any
// trailing partial block.
var NewEncoder = ascii85.NewEncoder
