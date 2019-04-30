package cipherh

import (
	"crypto/cipher"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewCBCDecrypterKey = "NewCBCDecrypter"

	NewCBCEncrypterKey = "NewCBCEncrypter"

	NewCFBDecrypterKey = "NewCFBDecrypter"

	NewCFBEncrypterKey = "NewCFBEncrypter"

	NewCTRKey = "NewCTR"

	NewGCMKey = "NewGCM"

	NewGCMWithNonceSizeKey = "NewGCMWithNonceSize"

	NewGCMWithTagSizeKey = "NewGCMWithTagSize"

	NewOFBKey = "NewOFB"
)

func New() hctx.Map {
	return hctx.Map{

		NewCBCDecrypterKey: NewCBCDecrypter,

		NewCBCEncrypterKey: NewCBCEncrypter,

		NewCFBDecrypterKey: NewCFBDecrypter,

		NewCFBEncrypterKey: NewCFBEncrypter,

		NewCTRKey: NewCTR,

		NewGCMKey: NewGCM,

		NewGCMWithNonceSizeKey: NewGCMWithNonceSize,

		NewGCMWithTagSizeKey: NewGCMWithTagSize,

		NewOFBKey: NewOFB,
	}
}

// NewCBCDecrypter returns a BlockMode which decrypts in cipher block chaining
// mode, using the given Block. The length of iv must be the same as the
// Block&#39;s block size and must match the iv used to encrypt the data.
var NewCBCDecrypter = cipher.NewCBCDecrypter

// NewCBCEncrypter returns a BlockMode which encrypts in cipher block chaining
// mode, using the given Block. The length of iv must be the same as the
// Block&#39;s block size.
var NewCBCEncrypter = cipher.NewCBCEncrypter

// NewCFBDecrypter returns a Stream which decrypts with cipher feedback mode,
// using the given Block. The iv must be the same length as the Block&#39;s block
// size.
var NewCFBDecrypter = cipher.NewCFBDecrypter

// NewCFBEncrypter returns a Stream which encrypts with cipher feedback mode,
// using the given Block. The iv must be the same length as the Block&#39;s block
// size.
var NewCFBEncrypter = cipher.NewCFBEncrypter

// NewCTR returns a Stream which encrypts/decrypts using the given Block in
// counter mode. The length of iv must be the same as the Block&#39;s block size.
var NewCTR = cipher.NewCTR

// NewGCM returns the given 128-bit, block cipher wrapped in Galois Counter Mode
// with the standard nonce length.
//
// In general, the GHASH operation performed by this implementation of GCM is not constant-time.
// An exception is when the underlying Block was created by aes.NewCipher
// on systems with hardware support for AES. See the crypto/aes package documentation for details.
var NewGCM = cipher.NewGCM

// NewGCMWithNonceSize returns the given 128-bit, block cipher wrapped in Galois
// Counter Mode, which accepts nonces of the given length.
//
// Only use this function if you require compatibility with an existing
// cryptosystem that uses non-standard nonce lengths. All other users should use
// NewGCM, which is faster and more resistant to misuse.
var NewGCMWithNonceSize = cipher.NewGCMWithNonceSize

// NewGCMWithTagSize returns the given 128-bit, block cipher wrapped in Galois
// Counter Mode, which generates tags with the given length.
//
// Tag sizes between 12 and 16 bytes are allowed.
//
// Only use this function if you require compatibility with an existing
// cryptosystem that uses non-standard tag lengths. All other users should use
// NewGCM, which is more resistant to misuse.
var NewGCMWithTagSize = cipher.NewGCMWithTagSize

// NewOFB returns a Stream that encrypts or decrypts using the block cipher b
// in output feedback mode. The initialization vector iv&#39;s length must be equal
// to b&#39;s block size.
var NewOFB = cipher.NewOFB
