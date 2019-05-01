package xmlh

import (
	"encoding/xml"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CopyTokenKey = "CopyToken"

	EscapeKey = "Escape"

	EscapeTextKey = "EscapeText"

	MarshalKey = "Marshal"

	MarshalIndentKey = "MarshalIndent"

	NewDecoderKey = "NewDecoder"

	NewEncoderKey = "NewEncoder"

	NewTokenDecoderKey = "NewTokenDecoder"

	UnmarshalKey = "Unmarshal"
)

func New() hctx.Map {
	return hctx.Map{

		CopyTokenKey: CopyToken,

		EscapeKey: Escape,

		EscapeTextKey: EscapeText,

		MarshalKey: Marshal,

		MarshalIndentKey: MarshalIndent,

		NewDecoderKey: NewDecoder,

		NewEncoderKey: NewEncoder,

		NewTokenDecoderKey: NewTokenDecoder,

		UnmarshalKey: Unmarshal,
	}
}

// CopyToken returns a copy of a Token.
var CopyToken = xml.CopyToken

// Escape is like EscapeText but omits the error return value.
// It is provided for backwards compatibility with Go 1.0.
// Code targeting Go 1.1 or later should use EscapeText.
var Escape = xml.Escape

// EscapeText writes to w the properly escaped XML equivalent
// of the plain text data s.
var EscapeText = xml.EscapeText

// Marshal returns the XML encoding of v.
//
// Marshal handles an array or slice by marshaling each of the elements.
// Marshal handles a pointer by marshaling the value it points at or, if the
// pointer is nil, by writing nothing. Marshal handles an interface value by
// marshaling the value it contains or, if the interface value is nil, by
// writing nothing. Marshal handles all other data by writing one or more XML
// elements containing the data.
//
// The name for the XML elements is taken from, in order of preference:
//     - the tag on the XMLName field, if the data is a struct
//     - the value of the XMLName field of type Name
//     - the tag of the struct field used to obtain the data
//     - the name of the struct field used to obtain the data
//     - the name of the marshaled type
//
// The XML element for a struct contains marshaled elements for each of the
// exported fields of the struct, with these exceptions:
//     - the XMLName field, described above, is omitted.
//     - a field with tag &#34;-&#34; is omitted.
//     - a field with tag &#34;name,attr&#34; becomes an attribute with
//       the given name in the XML element.
//     - a field with tag &#34;,attr&#34; becomes an attribute with the
//       field name in the XML element.
//     - a field with tag &#34;,chardata&#34; is written as character data,
//       not as an XML element.
//     - a field with tag &#34;,cdata&#34; is written as character data
//       wrapped in one or more &lt;![CDATA[ ... ]]&gt; tags, not as an XML element.
//     - a field with tag &#34;,innerxml&#34; is written verbatim, not subject
//       to the usual marshaling procedure.
//     - a field with tag &#34;,comment&#34; is written as an XML comment, not
//       subject to the usual marshaling procedure. It must not contain
//       the &#34;--&#34; string within it.
//     - a field with a tag including the &#34;omitempty&#34; option is omitted
//       if the field value is empty. The empty values are false, 0, any
//       nil pointer or interface value, and any array, slice, map, or
//       string of length zero.
//     - an anonymous struct field is handled as if the fields of its
//       value were part of the outer struct.
//
// If a field uses a tag &#34;a&gt;b&gt;c&#34;, then the element c will be nested inside
// parent elements a and b. Fields that appear next to each other that name
// the same parent will be enclosed in one XML element.
//
// If the XML name for a struct field is defined by both the field tag and the
// struct&#39;s XMLName field, the names must match.
//
// See MarshalIndent for an example.
//
// Marshal will return an error if asked to marshal a channel, function, or map.
var Marshal = xml.Marshal

// MarshalIndent works like Marshal, but each XML element begins on a new
// indented line that starts with prefix and is followed by one or more
// copies of indent according to the nesting depth.
var MarshalIndent = xml.MarshalIndent

// NewDecoder creates a new XML parser reading from r.
// If r does not implement io.ByteReader, NewDecoder will
// do its own buffering.
var NewDecoder = xml.NewDecoder

// NewEncoder returns a new encoder that writes to w.
var NewEncoder = xml.NewEncoder

// NewTokenDecoder creates a new XML parser using an underlying token stream.
var NewTokenDecoder = xml.NewTokenDecoder

// Unmarshal parses the XML-encoded data and stores the result in
// the value pointed to by v, which must be an arbitrary struct,
// slice, or string. Well-formed data that does not fit into v is
// discarded.
//
// Because Unmarshal uses the reflect package, it can only assign
// to exported (upper case) fields. Unmarshal uses a case-sensitive
// comparison to match XML element names to tag values and struct
// field names.
//
// Unmarshal maps an XML element to a struct using the following rules.
// In the rules, the tag of a field refers to the value associated with the
// key &#39;xml&#39; in the struct field&#39;s tag (see the example above).
//
//   * If the struct has a field of type []byte or string with tag
//      &#34;,innerxml&#34;, Unmarshal accumulates the raw XML nested inside the
//      element in that field. The rest of the rules still apply.
//
//   * If the struct has a field named XMLName of type Name,
//      Unmarshal records the element name in that field.
//
//   * If the XMLName field has an associated tag of the form
//      &#34;name&#34; or &#34;namespace-URL name&#34;, the XML element must have
//      the given name (and, optionally, name space) or else Unmarshal
//      returns an error.
//
//   * If the XML element has an attribute whose name matches a
//      struct field name with an associated tag containing &#34;,attr&#34; or
//      the explicit name in a struct field tag of the form &#34;name,attr&#34;,
//      Unmarshal records the attribute value in that field.
//
//   * If the XML element has an attribute not handled by the previous
//      rule and the struct has a field with an associated tag containing
//      &#34;,any,attr&#34;, Unmarshal records the attribute value in the first
//      such field.
//
//   * If the XML element contains character data, that data is
//      accumulated in the first struct field that has tag &#34;,chardata&#34;.
//      The struct field may have type []byte or string.
//      If there is no such field, the character data is discarded.
//
//   * If the XML element contains comments, they are accumulated in
//      the first struct field that has tag &#34;,comment&#34;.  The struct
//      field may have type []byte or string. If there is no such
//      field, the comments are discarded.
//
//   * If the XML element contains a sub-element whose name matches
//      the prefix of a tag formatted as &#34;a&#34; or &#34;a&gt;b&gt;c&#34;, unmarshal
//      will descend into the XML structure looking for elements with the
//      given names, and will map the innermost elements to that struct
//      field. A tag starting with &#34;&gt;&#34; is equivalent to one starting
//      with the field name followed by &#34;&gt;&#34;.
//
//   * If the XML element contains a sub-element whose name matches
//      a struct field&#39;s XMLName tag and the struct field has no
//      explicit name tag as per the previous rule, unmarshal maps
//      the sub-element to that struct field.
//
//   * If the XML element contains a sub-element whose name matches a
//      field without any mode flags (&#34;,attr&#34;, &#34;,chardata&#34;, etc), Unmarshal
//      maps the sub-element to that struct field.
//
//   * If the XML element contains a sub-element that hasn&#39;t matched any
//      of the above rules and the struct has a field with tag &#34;,any&#34;,
//      unmarshal maps the sub-element to that struct field.
//
//   * An anonymous struct field is handled as if the fields of its
//      value were part of the outer struct.
//
//   * A struct field with tag &#34;-&#34; is never unmarshaled into.
//
// Unmarshal maps an XML element to a string or []byte by saving the
// concatenation of that element&#39;s character data in the string or
// []byte. The saved []byte is never nil.
//
// Unmarshal maps an attribute value to a string or []byte by saving
// the value in the string or slice.
//
// Unmarshal maps an attribute value to an Attr by saving the attribute,
// including its name, in the Attr.
//
// Unmarshal maps an XML element or attribute value to a slice by
// extending the length of the slice and mapping the element or attribute
// to the newly created value.
//
// Unmarshal maps an XML element or attribute value to a bool by
// setting it to the boolean value represented by the string. Whitespace
// is trimmed and ignored.
//
// Unmarshal maps an XML element or attribute value to an integer or
// floating-point field by setting the field to the result of
// interpreting the string value in decimal. There is no check for
// overflow. Whitespace is trimmed and ignored.
//
// Unmarshal maps an XML element to a Name by recording the element
// name.
//
// Unmarshal maps an XML element to a pointer by setting the pointer
// to a freshly allocated value and then mapping the element to that value.
//
// A missing element or empty attribute value will be unmarshaled as a zero value.
// If the field is a slice, a zero value will be appended to the field. Otherwise, the
// field will be set to its zero value.
var Unmarshal = xml.Unmarshal
