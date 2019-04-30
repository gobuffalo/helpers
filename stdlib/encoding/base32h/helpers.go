package base32h

import (
	"encoding/base32"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewDecoderKey = "NewDecoder"

	NewEncoderKey = "NewEncoder"

	NewEncodingKey = "NewEncoding"
)

func New() hctx.Map {
	return hctx.Map{

		NewDecoderKey: NewDecoder,

		NewEncoderKey: NewEncoder,

		NewEncodingKey: NewEncoding,
	}
}

// NewDecoder constructs a new base32 stream decoder.
var NewDecoder = base32.NewDecoder

// NewEncoder returns a new base32 stream encoder. Data written to
// the returned writer will be encoded using enc and then written to w.
// Base32 encodings operate in 5-byte blocks; when finished
// writing, the caller must Close the returned encoder to flush any
// partially written blocks.
var NewEncoder = base32.NewEncoder

// NewEncoding returns a new Encoding defined by the given alphabet,
// which must be a 32-byte string.
var NewEncoding = base32.NewEncoding
