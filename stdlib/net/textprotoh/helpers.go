package textprotoh

import (
	"net/textproto"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CanonicalMIMEHeaderKeyKey = "CanonicalMIMEHeaderKey"

	DialKey = "Dial"

	NewConnKey = "NewConn"

	NewReaderKey = "NewReader"

	NewWriterKey = "NewWriter"

	TrimBytesKey = "TrimBytes"

	TrimStringKey = "TrimString"
)

func New() hctx.Map {
	return hctx.Map{

		CanonicalMIMEHeaderKeyKey: CanonicalMIMEHeaderKey,

		DialKey: Dial,

		NewConnKey: NewConn,

		NewReaderKey: NewReader,

		NewWriterKey: NewWriter,

		TrimBytesKey: TrimBytes,

		TrimStringKey: TrimString,
	}
}

// CanonicalMIMEHeaderKey returns the canonical format of the
// MIME header key s. The canonicalization converts the first
// letter and any letter following a hyphen to upper case;
// the rest are converted to lowercase. For example, the
// canonical key for &#34;accept-encoding&#34; is &#34;Accept-Encoding&#34;.
// MIME header keys are assumed to be ASCII only.
// If s contains a space or invalid header field bytes, it is
// returned without modifications.
var CanonicalMIMEHeaderKey = textproto.CanonicalMIMEHeaderKey

// Dial connects to the given address on the given network using net.Dial
// and then returns a new Conn for the connection.
var Dial = textproto.Dial

// NewConn returns a new Conn using conn for I/O.
var NewConn = textproto.NewConn

// NewReader returns a new Reader reading from r.
//
// To avoid denial of service attacks, the provided bufio.Reader
// should be reading from an io.LimitReader or similar Reader to bound
// the size of responses.
var NewReader = textproto.NewReader

// NewWriter returns a new Writer writing to w.
var NewWriter = textproto.NewWriter

// TrimBytes returns b without leading and trailing ASCII space.
var TrimBytes = textproto.TrimBytes

// TrimString returns s without leading and trailing ASCII space.
var TrimString = textproto.TrimString
