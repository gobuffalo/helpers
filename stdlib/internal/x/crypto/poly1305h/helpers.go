package poly1305h

import (
	"internal/x/crypto/poly1305"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	SumKey = "Sum"

	SumKey = "Sum"

	SumKey = "Sum"

	SumKey = "Sum"

	VerifyKey = "Verify"
)

func New() hctx.Map {
	return hctx.Map{

		SumKey: Sum,

		SumKey: Sum,

		SumKey: Sum,

		SumKey: Sum,

		VerifyKey: Verify,
	}
}

// Sum generates an authenticator for m using a one-time key and puts the
// 16-byte result into out. Authenticating two different messages with the same
// key allows an attacker to forge messages at will.
var Sum = poly1305.Sum

// Sum generates an authenticator for m using a one-time key and puts the
// 16-byte result into out. Authenticating two different messages with the same
// key allows an attacker to forge messages at will.
var Sum = poly1305.Sum

// Sum generates an authenticator for msg using a one-time key and puts the
// 16-byte result into out. Authenticating two different messages with the same
// key allows an attacker to forge messages at will.
var Sum = poly1305.Sum

// Sum generates an authenticator for m using a one-time key and puts the
// 16-byte result into out. Authenticating two different messages with the same
// key allows an attacker to forge messages at will.
var Sum = poly1305.Sum

// Verify returns true if mac is a valid authenticator for m with the given
// key.
var Verify = poly1305.Verify
