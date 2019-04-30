package aesh

import (
	"crypto/aes"

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

// NewCipher creates and returns a new cipher.Block.
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
var NewCipher = aes.NewCipher
