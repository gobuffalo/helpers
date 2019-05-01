package x509h

import (
	"crypto/x509"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CreateCertificateKey = "CreateCertificate"

	CreateCertificateRequestKey = "CreateCertificateRequest"

	DecryptPEMBlockKey = "DecryptPEMBlock"

	EncryptPEMBlockKey = "EncryptPEMBlock"

	IsEncryptedPEMBlockKey = "IsEncryptedPEMBlock"

	MarshalECPrivateKeyKey = "MarshalECPrivateKey"

	MarshalPKCS1PrivateKeyKey = "MarshalPKCS1PrivateKey"

	MarshalPKCS1PublicKeyKey = "MarshalPKCS1PublicKey"

	MarshalPKCS8PrivateKeyKey = "MarshalPKCS8PrivateKey"

	MarshalPKIXPublicKeyKey = "MarshalPKIXPublicKey"

	NewCertPoolKey = "NewCertPool"

	ParseCRLKey = "ParseCRL"

	ParseCertificateKey = "ParseCertificate"

	ParseCertificateRequestKey = "ParseCertificateRequest"

	ParseCertificatesKey = "ParseCertificates"

	ParseDERCRLKey = "ParseDERCRL"

	ParseECPrivateKeyKey = "ParseECPrivateKey"

	ParsePKCS1PrivateKeyKey = "ParsePKCS1PrivateKey"

	ParsePKCS1PublicKeyKey = "ParsePKCS1PublicKey"

	ParsePKCS8PrivateKeyKey = "ParsePKCS8PrivateKey"

	ParsePKIXPublicKeyKey = "ParsePKIXPublicKey"

	SystemCertPoolKey = "SystemCertPool"
)

func New() hctx.Map {
	return hctx.Map{

		CreateCertificateKey: CreateCertificate,

		CreateCertificateRequestKey: CreateCertificateRequest,

		DecryptPEMBlockKey: DecryptPEMBlock,

		EncryptPEMBlockKey: EncryptPEMBlock,

		IsEncryptedPEMBlockKey: IsEncryptedPEMBlock,

		MarshalECPrivateKeyKey: MarshalECPrivateKey,

		MarshalPKCS1PrivateKeyKey: MarshalPKCS1PrivateKey,

		MarshalPKCS1PublicKeyKey: MarshalPKCS1PublicKey,

		MarshalPKCS8PrivateKeyKey: MarshalPKCS8PrivateKey,

		MarshalPKIXPublicKeyKey: MarshalPKIXPublicKey,

		NewCertPoolKey: NewCertPool,

		ParseCRLKey: ParseCRL,

		ParseCertificateKey: ParseCertificate,

		ParseCertificateRequestKey: ParseCertificateRequest,

		ParseCertificatesKey: ParseCertificates,

		ParseDERCRLKey: ParseDERCRL,

		ParseECPrivateKeyKey: ParseECPrivateKey,

		ParsePKCS1PrivateKeyKey: ParsePKCS1PrivateKey,

		ParsePKCS1PublicKeyKey: ParsePKCS1PublicKey,

		ParsePKCS8PrivateKeyKey: ParsePKCS8PrivateKey,

		ParsePKIXPublicKeyKey: ParsePKIXPublicKey,

		SystemCertPoolKey: SystemCertPool,
	}
}

// CreateCertificate creates a new X.509v3 certificate based on a template.
// The following members of template are used: AuthorityKeyId,
// BasicConstraintsValid, DNSNames, ExcludedDNSDomains, ExtKeyUsage,
// IsCA, KeyUsage, MaxPathLen, MaxPathLenZero, NotAfter, NotBefore,
// PermittedDNSDomains, PermittedDNSDomainsCritical, SerialNumber,
// SignatureAlgorithm, Subject, SubjectKeyId, and UnknownExtKeyUsage.
//
// The certificate is signed by parent. If parent is equal to template then the
// certificate is self-signed. The parameter pub is the public key of the
// signee and priv is the private key of the signer.
//
// The returned slice is the certificate in DER encoding.
//
// All keys types that are implemented via crypto.Signer are supported (This
// includes *rsa.PublicKey and *ecdsa.PublicKey.)
//
// The AuthorityKeyId will be taken from the SubjectKeyId of parent, if any,
// unless the resulting certificate is self-signed. Otherwise the value from
// template will be used.
var CreateCertificate = x509.CreateCertificate

// CreateCertificateRequest creates a new certificate request based on a
// template. The following members of template are used: Attributes, DNSNames,
// EmailAddresses, ExtraExtensions, IPAddresses, URIs, SignatureAlgorithm, and
// Subject. The private key is the private key of the signer.
//
// The returned slice is the certificate request in DER encoding.
//
// All keys types that are implemented via crypto.Signer are supported (This
// includes *rsa.PublicKey and *ecdsa.PublicKey.)
var CreateCertificateRequest = x509.CreateCertificateRequest

// DecryptPEMBlock takes a password encrypted PEM block and the password used to
// encrypt it and returns a slice of decrypted DER encoded bytes. It inspects
// the DEK-Info header to determine the algorithm used for decryption. If no
// DEK-Info header is present, an error is returned. If an incorrect password
// is detected an IncorrectPasswordError is returned. Because of deficiencies
// in the encrypted-PEM format, it&#39;s not always possible to detect an incorrect
// password. In these cases no error will be returned but the decrypted DER
// bytes will be random noise.
var DecryptPEMBlock = x509.DecryptPEMBlock

// EncryptPEMBlock returns a PEM block of the specified type holding the
// given DER-encoded data encrypted with the specified algorithm and
// password.
var EncryptPEMBlock = x509.EncryptPEMBlock

// IsEncryptedPEMBlock returns if the PEM block is password encrypted.
var IsEncryptedPEMBlock = x509.IsEncryptedPEMBlock

// MarshalECPrivateKey marshals an EC private key into ASN.1, DER format.
var MarshalECPrivateKey = x509.MarshalECPrivateKey

// MarshalPKCS1PrivateKey converts a private key to ASN.1 DER encoded form.
var MarshalPKCS1PrivateKey = x509.MarshalPKCS1PrivateKey

// MarshalPKCS1PublicKey converts an RSA public key to PKCS#1, ASN.1 DER form.
var MarshalPKCS1PublicKey = x509.MarshalPKCS1PublicKey

// MarshalPKCS8PrivateKey converts a private key to PKCS#8 encoded form.
// The following key types are supported: *rsa.PrivateKey, *ecdsa.PublicKey.
// Unsupported key types result in an error.
//
// See RFC 5208.
var MarshalPKCS8PrivateKey = x509.MarshalPKCS8PrivateKey

// MarshalPKIXPublicKey serialises a public key to DER-encoded PKIX format.
var MarshalPKIXPublicKey = x509.MarshalPKIXPublicKey

// NewCertPool returns a new, empty CertPool.
var NewCertPool = x509.NewCertPool

// ParseCRL parses a CRL from the given bytes. It&#39;s often the case that PEM
// encoded CRLs will appear where they should be DER encoded, so this function
// will transparently handle PEM encoding as long as there isn&#39;t any leading
// garbage.
var ParseCRL = x509.ParseCRL

// ParseCertificate parses a single certificate from the given ASN.1 DER data.
var ParseCertificate = x509.ParseCertificate

// ParseCertificateRequest parses a single certificate request from the
// given ASN.1 DER data.
var ParseCertificateRequest = x509.ParseCertificateRequest

// ParseCertificates parses one or more certificates from the given ASN.1 DER
// data. The certificates must be concatenated with no intermediate padding.
var ParseCertificates = x509.ParseCertificates

// ParseDERCRL parses a DER encoded CRL from the given bytes.
var ParseDERCRL = x509.ParseDERCRL

// ParseECPrivateKey parses an ASN.1 Elliptic Curve Private Key Structure.
var ParseECPrivateKey = x509.ParseECPrivateKey

// ParsePKCS1PrivateKey returns an RSA private key from its ASN.1 PKCS#1 DER encoded form.
var ParsePKCS1PrivateKey = x509.ParsePKCS1PrivateKey

// ParsePKCS1PublicKey parses a PKCS#1 public key in ASN.1 DER form.
var ParsePKCS1PublicKey = x509.ParsePKCS1PublicKey

// ParsePKCS8PrivateKey parses an unencrypted, PKCS#8 private key.
// See RFC 5208.
var ParsePKCS8PrivateKey = x509.ParsePKCS8PrivateKey

// ParsePKIXPublicKey parses a DER encoded public key. These values are
// typically found in PEM blocks with &#34;BEGIN PUBLIC KEY&#34;.
//
// Supported key types include RSA, DSA, and ECDSA. Unknown key
// types result in an error.
//
// On success, pub will be of type *rsa.PublicKey, *dsa.PublicKey,
// or *ecdsa.PublicKey.
var ParsePKIXPublicKey = x509.ParsePKIXPublicKey

// SystemCertPool returns a copy of the system cert pool.
//
// Any mutations to the returned pool are not written to disk and do
// not affect any other pool.
var SystemCertPool = x509.SystemCertPool
