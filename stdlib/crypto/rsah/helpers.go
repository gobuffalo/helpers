package rsah

import (
	"crypto/rsa"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DecryptOAEPKey = "DecryptOAEP"

	DecryptPKCS1v15Key = "DecryptPKCS1v15"

	DecryptPKCS1v15SessionKeyKey = "DecryptPKCS1v15SessionKey"

	EncryptOAEPKey = "EncryptOAEP"

	EncryptPKCS1v15Key = "EncryptPKCS1v15"

	GenerateKeyKey = "GenerateKey"

	GenerateMultiPrimeKeyKey = "GenerateMultiPrimeKey"

	SignPKCS1v15Key = "SignPKCS1v15"

	SignPSSKey = "SignPSS"

	VerifyPKCS1v15Key = "VerifyPKCS1v15"

	VerifyPSSKey = "VerifyPSS"
)

func New() hctx.Map {
	return hctx.Map{

		DecryptOAEPKey: DecryptOAEP,

		DecryptPKCS1v15Key: DecryptPKCS1v15,

		DecryptPKCS1v15SessionKeyKey: DecryptPKCS1v15SessionKey,

		EncryptOAEPKey: EncryptOAEP,

		EncryptPKCS1v15Key: EncryptPKCS1v15,

		GenerateKeyKey: GenerateKey,

		GenerateMultiPrimeKeyKey: GenerateMultiPrimeKey,

		SignPKCS1v15Key: SignPKCS1v15,

		SignPSSKey: SignPSS,

		VerifyPKCS1v15Key: VerifyPKCS1v15,

		VerifyPSSKey: VerifyPSS,
	}
}

// OAEP is parameterised by a hash function that is used as a random oracle.
// Encryption and decryption of a given message must use the same hash function
// and sha256.New() is a reasonable choice.
//
// The random parameter, if not nil, is used to blind the private-key operation
// and avoid timing side-channel attacks. Blinding is purely internal to this
// function – the random data need not match that used when encrypting.
//
// The label parameter must match the value given when encrypting. See
// EncryptOAEP for details.
var DecryptOAEP = rsa.DecryptOAEP

// DecryptPKCS1v15 decrypts a plaintext using RSA and the padding scheme from PKCS#1 v1.5.
// If rand != nil, it uses RSA blinding to avoid timing side-channel attacks.
//
// Note that whether this function returns an error or not discloses secret
// information. If an attacker can cause this function to run repeatedly and
// learn whether each instance returned an error then they can decrypt and
// forge signatures as if they had the private key. See
// DecryptPKCS1v15SessionKey for a way of solving this problem.
var DecryptPKCS1v15 = rsa.DecryptPKCS1v15

// DecryptPKCS1v15SessionKey decrypts a session key using RSA and the padding scheme from PKCS#1 v1.5.
// If rand != nil, it uses RSA blinding to avoid timing side-channel attacks.
// It returns an error if the ciphertext is the wrong length or if the
// ciphertext is greater than the public modulus. Otherwise, no error is
// returned. If the padding is valid, the resulting plaintext message is copied
// into key. Otherwise, key is unchanged. These alternatives occur in constant
// time. It is intended that the user of this function generate a random
// session key beforehand and continue the protocol with the resulting value.
// This will remove any possibility that an attacker can learn any information
// about the plaintext.
// See ``Chosen Ciphertext Attacks Against Protocols Based on the RSA
// Encryption Standard PKCS #1&#39;&#39;, Daniel Bleichenbacher, Advances in Cryptology
// (Crypto &#39;98).
//
// Note that if the session key is too small then it may be possible for an
// attacker to brute-force it. If they can do that then they can learn whether
// a random value was used (because it&#39;ll be different for the same ciphertext)
// and thus whether the padding was correct. This defeats the point of this
// function. Using at least a 16-byte key will protect against this attack.
var DecryptPKCS1v15SessionKey = rsa.DecryptPKCS1v15SessionKey

// EncryptOAEP encrypts the given message with RSA-OAEP.
//
// OAEP is parameterised by a hash function that is used as a random oracle.
// Encryption and decryption of a given message must use the same hash function
// and sha256.New() is a reasonable choice.
//
// The random parameter is used as a source of entropy to ensure that
// encrypting the same message twice doesn&#39;t result in the same ciphertext.
//
// The label parameter may contain arbitrary data that will not be encrypted,
// but which gives important context to the message. For example, if a given
// public key is used to decrypt two types of messages then distinct label
// values could be used to ensure that a ciphertext for one purpose cannot be
// used for another by an attacker. If not required it can be empty.
//
// The message must be no longer than the length of the public modulus minus
// twice the hash length, minus a further 2.
var EncryptOAEP = rsa.EncryptOAEP

// EncryptPKCS1v15 encrypts the given message with RSA and the padding
// scheme from PKCS#1 v1.5.  The message must be no longer than the
// length of the public modulus minus 11 bytes.
//
// The rand parameter is used as a source of entropy to ensure that
// encrypting the same message twice doesn&#39;t result in the same
// ciphertext.
//
// WARNING: use of this function to encrypt plaintexts other than
// session keys is dangerous. Use RSA OAEP in new protocols.
var EncryptPKCS1v15 = rsa.EncryptPKCS1v15

// GenerateKey generates an RSA keypair of the given bit size using the
// random source random (for example, crypto/rand.Reader).
var GenerateKey = rsa.GenerateKey

// GenerateMultiPrimeKey generates a multi-prime RSA keypair of the given bit
// size and the given random source, as suggested in [1]. Although the public
// keys are compatible (actually, indistinguishable) from the 2-prime case,
// the private keys are not. Thus it may not be possible to export multi-prime
// private keys in certain formats or to subsequently import them into other
// code.
//
// Table 1 in [2] suggests maximum numbers of primes for a given size.
//
// [1] US patent 4405829 (1972, expired)
// [2] http://www.cacr.math.uwaterloo.ca/techreports/2006/cacr2006-16.pdf
var GenerateMultiPrimeKey = rsa.GenerateMultiPrimeKey

// SignPKCS1v15 calculates the signature of hashed using
// RSASSA-PKCS1-V1_5-SIGN from RSA PKCS#1 v1.5.  Note that hashed must
// be the result of hashing the input message using the given hash
// function. If hash is zero, hashed is signed directly. This isn&#39;t
// advisable except for interoperability.
//
// If rand is not nil then RSA blinding will be used to avoid timing
// side-channel attacks.
//
// This function is deterministic. Thus, if the set of possible
// messages is small, an attacker may be able to build a map from
// messages to signatures and identify the signed messages. As ever,
// signatures provide authenticity, not confidentiality.
var SignPKCS1v15 = rsa.SignPKCS1v15

// SignPSS calculates the signature of hashed using RSASSA-PSS [1].
// Note that hashed must be the result of hashing the input message using the
// given hash function. The opts argument may be nil, in which case sensible
// defaults are used.
var SignPSS = rsa.SignPSS

// VerifyPKCS1v15 verifies an RSA PKCS#1 v1.5 signature.
// hashed is the result of hashing the input message using the given hash
// function and sig is the signature. A valid signature is indicated by
// returning a nil error. If hash is zero then hashed is used directly. This
// isn&#39;t advisable except for interoperability.
var VerifyPKCS1v15 = rsa.VerifyPKCS1v15

// VerifyPSS verifies a PSS signature.
// hashed is the result of hashing the input message using the given hash
// function and sig is the signature. A valid signature is indicated by
// returning a nil error. The opts argument may be nil, in which case sensible
// defaults are used.
var VerifyPSS = rsa.VerifyPSS
