package chacha20poly1305h

import (
	"internal/x/crypto/chacha20poly1305"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewKey = "New"
)

func New() hctx.Map {
	return hctx.Map{

		NewKey: New,
	}
}

// New returns a ChaCha20-Poly1305 AEAD that uses the given, 256-bit key.
var New = chacha20poly1305.New
