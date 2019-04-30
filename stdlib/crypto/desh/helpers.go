package desh

import (
	"crypto/des"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewCipherKey = "NewCipher"

	NewTripleDESCipherKey = "NewTripleDESCipher"
)

func New() hctx.Map {
	return hctx.Map{

		NewCipherKey: NewCipher,

		NewTripleDESCipherKey: NewTripleDESCipher,
	}
}

// NewCipher creates and returns a new cipher.Block.
var NewCipher = des.NewCipher

// NewTripleDESCipher creates and returns a new cipher.Block.
var NewTripleDESCipher = des.NewTripleDESCipher
