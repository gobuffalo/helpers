package base64h

import (
	"encoding/base64"

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

// NewDecoder constructs a new base64 stream decoder.
var NewDecoder = base64.NewDecoder

// NewEncoder returns a new base64 stream encoder. Data written to
// the returned writer will be encoded using enc and then written to w.
// Base64 encodings operate in 4-byte blocks; when finished
// writing, the caller must Close the returned encoder to flush any
// partially written blocks.
var NewEncoder = base64.NewEncoder

// NewEncoding returns a new padded Encoding defined by the given alphabet,
// which must be a 64-byte string that does not contain the padding character
// or CR / LF (&#39;\r&#39;, &#39;\n&#39;).
// The resulting Encoding uses the default padding character (&#39;=&#39;),
// which may be changed or disabled via WithPadding.
var NewEncoding = base64.NewEncoding
