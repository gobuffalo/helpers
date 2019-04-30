package elliptich

import (
	"crypto/elliptic"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	GenerateKeyKey = "GenerateKey"

	MarshalKey = "Marshal"

	P224Key = "P224"

	P256Key = "P256"

	P384Key = "P384"

	P521Key = "P521"

	UnmarshalKey = "Unmarshal"
)

func New() hctx.Map {
	return hctx.Map{

		GenerateKeyKey: GenerateKey,

		MarshalKey: Marshal,

		P224Key: P224,

		P256Key: P256,

		P384Key: P384,

		P521Key: P521,

		UnmarshalKey: Unmarshal,
	}
}

// GenerateKey returns a public/private key pair. The private key is
// generated using the given reader, which must return random data.
var GenerateKey = elliptic.GenerateKey

// Marshal converts a point into the uncompressed form specified in section 4.3.6 of ANSI X9.62.
var Marshal = elliptic.Marshal

// P224 returns a Curve which implements P-224 (see FIPS 186-3, section D.2.2).
//
// The cryptographic operations are implemented using constant-time algorithms.
var P224 = elliptic.P224

// P256 returns a Curve which implements P-256 (see FIPS 186-3, section D.2.3)
//
// The cryptographic operations are implemented using constant-time algorithms.
var P256 = elliptic.P256

// P384 returns a Curve which implements P-384 (see FIPS 186-3, section D.2.4)
//
// The cryptographic operations do not use constant-time algorithms.
var P384 = elliptic.P384

// P521 returns a Curve which implements P-521 (see FIPS 186-3, section D.2.5)
//
// The cryptographic operations do not use constant-time algorithms.
var P521 = elliptic.P521

// Unmarshal converts a point, serialized by Marshal, into an x, y pair.
// It is an error if the point is not in uncompressed form or is not on the curve.
// On error, x = nil.
var Unmarshal = elliptic.Unmarshal
