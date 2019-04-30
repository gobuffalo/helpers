package dsah

import (
	"crypto/dsa"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	GenerateKeyKey = "GenerateKey"

	GenerateParametersKey = "GenerateParameters"

	SignKey = "Sign"

	VerifyKey = "Verify"
)

func New() hctx.Map {
	return hctx.Map{

		GenerateKeyKey: GenerateKey,

		GenerateParametersKey: GenerateParameters,

		SignKey: Sign,

		VerifyKey: Verify,
	}
}

// GenerateKey generates a public&amp;private key pair. The Parameters of the
// PrivateKey must already be valid (see GenerateParameters).
var GenerateKey = dsa.GenerateKey

// GenerateParameters puts a random, valid set of DSA parameters into params.
// This function can take many seconds, even on fast machines.
var GenerateParameters = dsa.GenerateParameters

// Sign signs an arbitrary length hash (which should be the result of hashing a
// larger message) using the private key, priv. It returns the signature as a
// pair of integers. The security of the private key depends on the entropy of
// rand.
//
// Note that FIPS 186-3 section 4.6 specifies that the hash should be truncated
// to the byte-length of the subgroup. This function does not perform that
// truncation itself.
//
// Be aware that calling Sign with an attacker-controlled PrivateKey may
// require an arbitrary amount of CPU.
var Sign = dsa.Sign

// Verify verifies the signature in r, s of hash using the public key, pub. It
// reports whether the signature is valid.
//
// Note that FIPS 186-3 section 4.6 specifies that the hash should be truncated
// to the byte-length of the subgroup. This function does not perform that
// truncation itself.
var Verify = dsa.Verify
