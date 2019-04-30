package rc4h

import (
	"crypto/rc4"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewCipherKey = "NewCipher"
)

func New() hctx.Map {
	return hctx.Map{

		NewCipherKey: NewCipher,
	}
}

// NewCipher creates and returns a new Cipher. The key argument should be the
// RC4 key, at least 1 byte and at most 256 bytes.
var NewCipher = rc4.NewCipher
