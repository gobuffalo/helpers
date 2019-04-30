package smtph

import (
	"net/smtp"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CRAMMD5AuthKey = "CRAMMD5Auth"

	DialKey = "Dial"

	NewClientKey = "NewClient"

	PlainAuthKey = "PlainAuth"

	SendMailKey = "SendMail"
)

func New() hctx.Map {
	return hctx.Map{

		CRAMMD5AuthKey: CRAMMD5Auth,

		DialKey: Dial,

		NewClientKey: NewClient,

		PlainAuthKey: PlainAuth,

		SendMailKey: SendMail,
	}
}

// CRAMMD5Auth returns an Auth that implements the CRAM-MD5 authentication
// mechanism as defined in RFC 2195.
// The returned Auth uses the given username and secret to authenticate
// to the server using the challenge-response mechanism.
var CRAMMD5Auth = smtp.CRAMMD5Auth

// Dial returns a new Client connected to an SMTP server at addr.
// The addr must include a port, as in &#34;mail.example.com:smtp&#34;.
var Dial = smtp.Dial

// NewClient returns a new Client using an existing connection and host as a
// server name to be used when authenticating.
var NewClient = smtp.NewClient

// PlainAuth returns an Auth that implements the PLAIN authentication
// mechanism as defined in RFC 4616. The returned Auth uses the given
// username and password to authenticate to host and act as identity.
// Usually identity should be the empty string, to act as username.
//
// PlainAuth will only send the credentials if the connection is using TLS
// or is connected to localhost. Otherwise authentication will fail with an
// error, without sending the credentials.
var PlainAuth = smtp.PlainAuth

// SendMail connects to the server at addr, switches to TLS if
// possible, authenticates with the optional mechanism a if possible,
// and then sends an email from address from, to addresses to, with
// message msg.
// The addr must include a port, as in &#34;mail.example.com:smtp&#34;.
//
// The addresses in the to parameter are the SMTP RCPT addresses.
//
// The msg parameter should be an RFC 822-style email with headers
// first, a blank line, and then the message body. The lines of msg
// should be CRLF terminated. The msg headers should usually include
// fields such as &#34;From&#34;, &#34;To&#34;, &#34;Subject&#34;, and &#34;Cc&#34;.  Sending &#34;Bcc&#34;
// messages is accomplished by including an email address in the to
// parameter but not including it in the msg headers.
//
// The SendMail function and the net/smtp package are low-level
// mechanisms and provide no support for DKIM signing, MIME
// attachments (see the mime/multipart package), or other mail
// functionality. Higher-level packages exist outside of the standard
// library.
var SendMail = smtp.SendMail
