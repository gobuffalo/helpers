package urlh

import (
	"net/url"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ParseKey = "Parse"

	ParseQueryKey = "ParseQuery"

	ParseRequestURIKey = "ParseRequestURI"

	PathEscapeKey = "PathEscape"

	PathUnescapeKey = "PathUnescape"

	QueryEscapeKey = "QueryEscape"

	QueryUnescapeKey = "QueryUnescape"

	UserKey = "User"

	UserPasswordKey = "UserPassword"
)

func New() hctx.Map {
	return hctx.Map{

		ParseKey: Parse,

		ParseQueryKey: ParseQuery,

		ParseRequestURIKey: ParseRequestURI,

		PathEscapeKey: PathEscape,

		PathUnescapeKey: PathUnescape,

		QueryEscapeKey: QueryEscape,

		QueryUnescapeKey: QueryUnescape,

		UserKey: User,

		UserPasswordKey: UserPassword,
	}
}

// Parse parses rawurl into a URL structure.
//
// The rawurl may be relative (a path, without a host) or absolute
// (starting with a scheme). Trying to parse a hostname and path
// without a scheme is invalid but may not necessarily return an
// error, due to parsing ambiguities.
var Parse = url.Parse

// ParseQuery parses the URL-encoded query string and returns
// a map listing the values specified for each key.
// ParseQuery always returns a non-nil map containing all the
// valid query parameters found; err describes the first decoding error
// encountered, if any.
//
// Query is expected to be a list of key=value settings separated by
// ampersands or semicolons. A setting without an equals sign is
// interpreted as a key set to an empty value.
var ParseQuery = url.ParseQuery

// ParseRequestURI parses rawurl into a URL structure. It assumes that
// rawurl was received in an HTTP request, so the rawurl is interpreted
// only as an absolute URI or an absolute path.
// The string rawurl is assumed not to have a #fragment suffix.
// (Web browsers strip #fragment before sending the URL to a web server.)
var ParseRequestURI = url.ParseRequestURI

// PathEscape escapes the string so it can be safely placed
// inside a URL path segment.
var PathEscape = url.PathEscape

// PathUnescape does the inverse transformation of PathEscape,
// converting each 3-byte encoded substring of the form &#34;%AB&#34; into the
// hex-decoded byte 0xAB. It also converts &#39;+&#39; into &#39; &#39; (space).
// It returns an error if any % is not followed by two hexadecimal
// digits.
//
// PathUnescape is identical to QueryUnescape except that it does not
// unescape &#39;+&#39; to &#39; &#39; (space).
var PathUnescape = url.PathUnescape

// QueryEscape escapes the string so it can be safely placed
// inside a URL query.
var QueryEscape = url.QueryEscape

// QueryUnescape does the inverse transformation of QueryEscape,
// converting each 3-byte encoded substring of the form &#34;%AB&#34; into the
// hex-decoded byte 0xAB. It also converts &#39;+&#39; into &#39; &#39; (space).
// It returns an error if any % is not followed by two hexadecimal
// digits.
var QueryUnescape = url.QueryUnescape

// User returns a Userinfo containing the provided username
// and no password set.
var User = url.User

// UserPassword returns a Userinfo containing the provided username
// and password.
//
// This functionality should only be used with legacy web sites.
// RFC 2396 warns that interpreting Userinfo this way
// ``is NOT RECOMMENDED, because the passing of authentication
// information in clear text (such as URI) has proven to be a
// security risk in almost every case where it has been used.&#39;&#39;
var UserPassword = url.UserPassword
