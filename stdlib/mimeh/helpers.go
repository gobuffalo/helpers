package mimeh

import (
	"mime"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AddExtensionTypeKey = "AddExtensionType"

	ExtensionsByTypeKey = "ExtensionsByType"

	FormatMediaTypeKey = "FormatMediaType"

	ParseMediaTypeKey = "ParseMediaType"

	TypeByExtensionKey = "TypeByExtension"
)

func New() hctx.Map {
	return hctx.Map{

		AddExtensionTypeKey: AddExtensionType,

		ExtensionsByTypeKey: ExtensionsByType,

		FormatMediaTypeKey: FormatMediaType,

		ParseMediaTypeKey: ParseMediaType,

		TypeByExtensionKey: TypeByExtension,
	}
}

// AddExtensionType sets the MIME type associated with
// the extension ext to typ. The extension should begin with
// a leading dot, as in &#34;.html&#34;.
var AddExtensionType = mime.AddExtensionType

// ExtensionsByType returns the extensions known to be associated with the MIME
// type typ. The returned extensions will each begin with a leading dot, as in
// &#34;.html&#34;. When typ has no associated extensions, ExtensionsByType returns an
// nil slice.
var ExtensionsByType = mime.ExtensionsByType

// FormatMediaType serializes mediatype t and the parameters
// param as a media type conforming to RFC 2045 and RFC 2616.
// The type and parameter names are written in lower-case.
// When any of the arguments result in a standard violation then
// FormatMediaType returns the empty string.
var FormatMediaType = mime.FormatMediaType

// ParseMediaType parses a media type value and any optional
// parameters, per RFC 1521.  Media types are the values in
// Content-Type and Content-Disposition headers (RFC 2183).
// On success, ParseMediaType returns the media type converted
// to lowercase and trimmed of white space and a non-nil map.
// If there is an error parsing the optional parameter,
// the media type will be returned along with the error
// ErrInvalidMediaParameter.
// The returned map, params, maps from the lowercase
// attribute to the attribute value with its case preserved.
var ParseMediaType = mime.ParseMediaType

// TypeByExtension returns the MIME type associated with the file extension ext.
// The extension ext should begin with a leading dot, as in &#34;.html&#34;.
// When ext has no associated type, TypeByExtension returns &#34;&#34;.
//
// Extensions are looked up first case-sensitively, then case-insensitively.
//
// The built-in table is small but on unix it is augmented by the local
// system&#39;s mime.types file(s) if available under one or more of these
// names:
//
//   /etc/mime.types
//   /etc/apache2/mime.types
//   /etc/apache/mime.types
//
// On Windows, MIME types are extracted from the registry.
//
// Text types have the charset parameter set to &#34;utf-8&#34; by default.
var TypeByExtension = mime.TypeByExtension
