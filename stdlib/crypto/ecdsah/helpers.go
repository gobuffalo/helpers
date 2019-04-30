package ecdsah

import (
	"crypto/ecdsa"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	GenerateKeyKey = "GenerateKey"

	SignKey = "Sign"

	VerifyKey = "Verify"
)

func New() hctx.Map {
	return hctx.Map{

		GenerateKeyKey: GenerateKey,

		SignKey: Sign,

		VerifyKey: Verify,
	}
}

// GenerateKey generates a public and private key pair.
var GenerateKey = ecdsa.GenerateKey

// Sign signs a hash (which should be the result of hashing a larger message)
// using the private key, priv. If the hash is longer than the bit-length of the
// private key&#39;s curve order, the hash will be truncated to that length.  It
// returns the signature as a pair of integers. The security of the private key
// depends on the entropy of rand.
var Sign = ecdsa.Sign

// Verify verifies the signature in r, s of hash using the public key, pub. Its
// return value records whether the signature is valid.
var Verify = ecdsa.Verify
