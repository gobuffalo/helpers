package httpgutsh

import (
	"internal/x/net/http/httpguts"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	HeaderValuesContainsTokenKey = "HeaderValuesContainsToken"

	IsTokenRuneKey = "IsTokenRune"

	PunycodeHostPortKey = "PunycodeHostPort"

	ValidHeaderFieldNameKey = "ValidHeaderFieldName"

	ValidHeaderFieldValueKey = "ValidHeaderFieldValue"

	ValidHostHeaderKey = "ValidHostHeader"

	ValidTrailerHeaderKey = "ValidTrailerHeader"
)

func New() hctx.Map {
	return hctx.Map{

		HeaderValuesContainsTokenKey: HeaderValuesContainsToken,

		IsTokenRuneKey: IsTokenRune,

		PunycodeHostPortKey: PunycodeHostPort,

		ValidHeaderFieldNameKey: ValidHeaderFieldName,

		ValidHeaderFieldValueKey: ValidHeaderFieldValue,

		ValidHostHeaderKey: ValidHostHeader,

		ValidTrailerHeaderKey: ValidTrailerHeader,
	}
}

// HeaderValuesContainsToken reports whether any string in values
// contains the provided token, ASCII case-insensitively.
var HeaderValuesContainsToken = httpguts.HeaderValuesContainsToken

var IsTokenRune = httpguts.IsTokenRune

// PunycodeHostPort returns the IDNA Punycode version
// of the provided &#34;host&#34; or &#34;host:port&#34; string.
var PunycodeHostPort = httpguts.PunycodeHostPort

// ValidHeaderFieldName reports whether v is a valid HTTP/1.x header name.
// HTTP/2 imposes the additional restriction that uppercase ASCII
// letters are not allowed.
//
//  RFC 7230 says:
//   header-field   = field-name &#34;:&#34; OWS field-value OWS
//   field-name     = token
//   token          = 1*tchar
//   tchar = &#34;!&#34; / &#34;#&#34; / &#34;$&#34; / &#34;%&#34; / &#34;&amp;&#34; / &#34;&#39;&#34; / &#34;*&#34; / &#34;+&#34; / &#34;-&#34; / &#34;.&#34; /
//           &#34;^&#34; / &#34;_&#34; / &#34;`&#34; / &#34;|&#34; / &#34;~&#34; / DIGIT / ALPHA
var ValidHeaderFieldName = httpguts.ValidHeaderFieldName

// ValidHeaderFieldValue reports whether v is a valid &#34;field-value&#34; according to
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html#sec4.2 :
//
//        message-header = field-name &#34;:&#34; [ field-value ]
//        field-value    = *( field-content | LWS )
//        field-content  = &lt;the OCTETs making up the field-value
//                         and consisting of either *TEXT or combinations
//                         of token, separators, and quoted-string&gt;
//
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec2.html#sec2.2 :
//
//        TEXT           = &lt;any OCTET except CTLs,
//                          but including LWS&gt;
//        LWS            = [CRLF] 1*( SP | HT )
//        CTL            = &lt;any US-ASCII control character
//                         (octets 0 - 31) and DEL (127)&gt;
//
// RFC 7230 says:
//  field-value    = *( field-content / obs-fold )
//  obj-fold       =  N/A to http2, and deprecated
//  field-content  = field-vchar [ 1*( SP / HTAB ) field-vchar ]
//  field-vchar    = VCHAR / obs-text
//  obs-text       = %x80-FF
//  VCHAR          = &#34;any visible [USASCII] character&#34;
//
// http2 further says: &#34;Similarly, HTTP/2 allows header field values
// that are not valid. While most of the values that can be encoded
// will not alter header field parsing, carriage return (CR, ASCII
// 0xd), line feed (LF, ASCII 0xa), and the zero character (NUL, ASCII
// 0x0) might be exploited by an attacker if they are translated
// verbatim. Any request or response that contains a character not
// permitted in a header field value MUST be treated as malformed
// (Section 8.1.2.6). Valid characters are defined by the
// field-content ABNF rule in Section 3.2 of [RFC7230].&#34;
//
// This function does not (yet?) properly handle the rejection of
// strings that begin or end with SP or HTAB.
var ValidHeaderFieldValue = httpguts.ValidHeaderFieldValue

// ValidHostHeader reports whether h is a valid host header.
var ValidHostHeader = httpguts.ValidHostHeader

// ValidTrailerHeader reports whether name is a valid header field name to appear
// in trailers.
// See RFC 7230, Section 4.1.2
var ValidTrailerHeader = httpguts.ValidTrailerHeader
