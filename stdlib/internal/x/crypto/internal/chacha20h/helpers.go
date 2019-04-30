package chacha20h

import (
	"internal/x/crypto/internal/chacha20"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewKey = "New"

	XORKeyStreamKey = "XORKeyStream"
)

func New() hctx.Map {
	return hctx.Map{

		NewKey: New,

		XORKeyStreamKey: XORKeyStream,
	}
}

// New creates a new ChaCha20 stream cipher with the given key and nonce.
// The initial counter value is set to 0.
var New = chacha20.New

// XORKeyStream crypts bytes from in to out using the given key and counters.
// In and out must overlap entirely or not at all. Counter contains the raw
// ChaCha20 counter bytes (i.e. block counter followed by nonce).
var XORKeyStream = chacha20.XORKeyStream
