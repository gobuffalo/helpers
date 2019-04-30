package tlsh

import (
	"crypto/tls"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ClientKey = "Client"

	DialKey = "Dial"

	DialWithDialerKey = "DialWithDialer"

	ListenKey = "Listen"

	LoadX509KeyPairKey = "LoadX509KeyPair"

	NewLRUClientSessionCacheKey = "NewLRUClientSessionCache"

	NewListenerKey = "NewListener"

	ServerKey = "Server"

	X509KeyPairKey = "X509KeyPair"
)

func New() hctx.Map {
	return hctx.Map{

		ClientKey: Client,

		DialKey: Dial,

		DialWithDialerKey: DialWithDialer,

		ListenKey: Listen,

		LoadX509KeyPairKey: LoadX509KeyPair,

		NewLRUClientSessionCacheKey: NewLRUClientSessionCache,

		NewListenerKey: NewListener,

		ServerKey: Server,

		X509KeyPairKey: X509KeyPair,
	}
}

// Client returns a new TLS client side connection
// using conn as the underlying transport.
// The config cannot be nil: users must set either ServerName or
// InsecureSkipVerify in the config.
var Client = tls.Client

// Dial connects to the given network address using net.Dial
// and then initiates a TLS handshake, returning the resulting
// TLS connection.
// Dial interprets a nil configuration as equivalent to
// the zero configuration; see the documentation of Config
// for the defaults.
var Dial = tls.Dial

// DialWithDialer connects to the given network address using dialer.Dial and
// then initiates a TLS handshake, returning the resulting TLS connection. Any
// timeout or deadline given in the dialer apply to connection and TLS
// handshake as a whole.
//
// DialWithDialer interprets a nil configuration as equivalent to the zero
// configuration; see the documentation of Config for the defaults.
var DialWithDialer = tls.DialWithDialer

// Listen creates a TLS listener accepting connections on the
// given network address using net.Listen.
// The configuration config must be non-nil and must include
// at least one certificate or else set GetCertificate.
var Listen = tls.Listen

// LoadX509KeyPair reads and parses a public/private key pair from a pair
// of files. The files must contain PEM encoded data. The certificate file
// may contain intermediate certificates following the leaf certificate to
// form a certificate chain. On successful return, Certificate.Leaf will
// be nil because the parsed form of the certificate is not retained.
var LoadX509KeyPair = tls.LoadX509KeyPair

// NewLRUClientSessionCache returns a ClientSessionCache with the given
// capacity that uses an LRU strategy. If capacity is &lt; 1, a default capacity
// is used instead.
var NewLRUClientSessionCache = tls.NewLRUClientSessionCache

// NewListener creates a Listener which accepts connections from an inner
// Listener and wraps each connection with Server.
// The configuration config must be non-nil and must include
// at least one certificate or else set GetCertificate.
var NewListener = tls.NewListener

// Server returns a new TLS server side connection
// using conn as the underlying transport.
// The configuration config must be non-nil and must include
// at least one certificate or else set GetCertificate.
var Server = tls.Server

// X509KeyPair parses a public/private key pair from a pair of
// PEM encoded data. On successful return, Certificate.Leaf will be nil because
// the parsed form of the certificate is not retained.
var X509KeyPair = tls.X509KeyPair
