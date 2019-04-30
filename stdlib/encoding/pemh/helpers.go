package pemh

import (
	"encoding/pem"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DecodeKey = "Decode"

	EncodeKey = "Encode"

	EncodeToMemoryKey = "EncodeToMemory"
)

func New() hctx.Map {
	return hctx.Map{

		DecodeKey: Decode,

		EncodeKey: Encode,

		EncodeToMemoryKey: EncodeToMemory,
	}
}

// Decode will find the next PEM formatted block (certificate, private key
// etc) in the input. It returns that block and the remainder of the input. If
// no PEM data is found, p is nil and the whole of the input is returned in
// rest.
var Decode = pem.Decode

// Encode writes the PEM encoding of b to out.
var Encode = pem.Encode

// EncodeToMemory returns the PEM encoding of b.
//
// If b has invalid headers and cannot be encoded,
// EncodeToMemory returns nil. If it is important to
// report details about this error case, use Encode instead.
var EncodeToMemory = pem.EncodeToMemory
