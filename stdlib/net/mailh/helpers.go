package mailh

import (
	"net/mail"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ParseAddressKey = "ParseAddress"

	ParseAddressListKey = "ParseAddressList"

	ParseDateKey = "ParseDate"

	ReadMessageKey = "ReadMessage"
)

func New() hctx.Map {
	return hctx.Map{

		ParseAddressKey: ParseAddress,

		ParseAddressListKey: ParseAddressList,

		ParseDateKey: ParseDate,

		ReadMessageKey: ReadMessage,
	}
}

// Parses a single RFC 5322 address, e.g. &#34;Barry Gibbs &lt;bg@example.com&gt;&#34;
var ParseAddress = mail.ParseAddress

// ParseAddressList parses the given string as a list of addresses.
var ParseAddressList = mail.ParseAddressList

// ParseDate parses an RFC 5322 date string.
var ParseDate = mail.ParseDate

// ReadMessage reads a message from r.
// The headers are parsed, and the body of the message will be available
// for reading from msg.Body.
var ReadMessage = mail.ReadMessage
