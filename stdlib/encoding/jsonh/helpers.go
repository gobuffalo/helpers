package jsonh

import (
	"encoding/json"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CompactKey = "Compact"

	HTMLEscapeKey = "HTMLEscape"

	IndentKey = "Indent"

	MarshalKey = "Marshal"

	MarshalIndentKey = "MarshalIndent"

	NewDecoderKey = "NewDecoder"

	NewEncoderKey = "NewEncoder"

	UnmarshalKey = "Unmarshal"

	ValidKey = "Valid"
)

func New() hctx.Map {
	return hctx.Map{

		CompactKey: Compact,

		HTMLEscapeKey: HTMLEscape,

		IndentKey: Indent,

		MarshalKey: Marshal,

		MarshalIndentKey: MarshalIndent,

		NewDecoderKey: NewDecoder,

		NewEncoderKey: NewEncoder,

		UnmarshalKey: Unmarshal,

		ValidKey: Valid,
	}
}

// Compact appends to dst the JSON-encoded src with
// insignificant space characters elided.
var Compact = json.Compact

// HTMLEscape appends to dst the JSON-encoded src with &lt;, &gt;, &amp;, U+2028 and U+2029
// characters inside string literals changed to \u003c, \u003e, \u0026, \u2028, \u2029
// so that the JSON will be safe to embed inside HTML &lt;script&gt; tags.
// For historical reasons, web browsers don&#39;t honor standard HTML
// escaping within &lt;script&gt; tags, so an alternative JSON encoding must
// be used.
var HTMLEscape = json.HTMLEscape

// Indent appends to dst an indented form of the JSON-encoded src.
// Each element in a JSON object or array begins on a new,
// indented line beginning with prefix followed by one or more
// copies of indent according to the indentation nesting.
// The data appended to dst does not begin with the prefix nor
// any indentation, to make it easier to embed inside other formatted JSON data.
// Although leading space characters (space, tab, carriage return, newline)
// at the beginning of src are dropped, trailing space characters
// at the end of src are preserved and copied to dst.
// For example, if src has no trailing spaces, neither will dst;
// if src ends in a trailing newline, so will dst.
var Indent = json.Indent

// Marshal returns the JSON encoding of v.
//
// Marshal traverses the value v recursively.
// If an encountered value implements the Marshaler interface
// and is not a nil pointer, Marshal calls its MarshalJSON method
// to produce JSON. If no MarshalJSON method is present but the
// value implements encoding.TextMarshaler instead, Marshal calls
// its MarshalText method and encodes the result as a JSON string.
// The nil pointer exception is not strictly necessary
// but mimics a similar, necessary exception in the behavior of
// UnmarshalJSON.
//
// Otherwise, Marshal uses the following type-dependent default encodings:
//
// Boolean values encode as JSON booleans.
//
// Floating point, integer, and Number values encode as JSON numbers.
//
// String values encode as JSON strings coerced to valid UTF-8,
// replacing invalid bytes with the Unicode replacement rune.
// The angle brackets &#34;&lt;&#34; and &#34;&gt;&#34; are escaped to &#34;\u003c&#34; and &#34;\u003e&#34;
// to keep some browsers from misinterpreting JSON output as HTML.
// Ampersand &#34;&amp;&#34; is also escaped to &#34;\u0026&#34; for the same reason.
// This escaping can be disabled using an Encoder that had SetEscapeHTML(false)
// called on it.
//
// Array and slice values encode as JSON arrays, except that
// []byte encodes as a base64-encoded string, and a nil slice
// encodes as the null JSON value.
//
// Struct values encode as JSON objects.
// Each exported struct field becomes a member of the object, using the
// field name as the object key, unless the field is omitted for one of the
// reasons given below.
//
// The encoding of each struct field can be customized by the format string
// stored under the &#34;json&#34; key in the struct field&#39;s tag.
// The format string gives the name of the field, possibly followed by a
// comma-separated list of options. The name may be empty in order to
// specify options without overriding the default field name.
//
// The &#34;omitempty&#34; option specifies that the field should be omitted
// from the encoding if the field has an empty value, defined as
// false, 0, a nil pointer, a nil interface value, and any empty array,
// slice, map, or string.
//
// As a special case, if the field tag is &#34;-&#34;, the field is always omitted.
// Note that a field with name &#34;-&#34; can still be generated using the tag &#34;-,&#34;.
//
// Examples of struct field tags and their meanings:
//
//   // Field appears in JSON as key &#34;myName&#34;.
//   Field int `json:&#34;myName&#34;`
//
//   // Field appears in JSON as key &#34;myName&#34; and
//   // the field is omitted from the object if its value is empty,
//   // as defined above.
//   Field int `json:&#34;myName,omitempty&#34;`
//
//   // Field appears in JSON as key &#34;Field&#34; (the default), but
//   // the field is skipped if empty.
//   // Note the leading comma.
//   Field int `json:&#34;,omitempty&#34;`
//
//   // Field is ignored by this package.
//   Field int `json:&#34;-&#34;`
//
//   // Field appears in JSON as key &#34;-&#34;.
//   Field int `json:&#34;-,&#34;`
//
// The &#34;string&#34; option signals that a field is stored as JSON inside a
// JSON-encoded string. It applies only to fields of string, floating point,
// integer, or boolean types. This extra level of encoding is sometimes used
// when communicating with JavaScript programs:
//
//    Int64String int64 `json:&#34;,string&#34;`
//
// The key name will be used if it&#39;s a non-empty string consisting of
// only Unicode letters, digits, and ASCII punctuation except quotation
// marks, backslash, and comma.
//
// Anonymous struct fields are usually marshaled as if their inner exported fields
// were fields in the outer struct, subject to the usual Go visibility rules amended
// as described in the next paragraph.
// An anonymous struct field with a name given in its JSON tag is treated as
// having that name, rather than being anonymous.
// An anonymous struct field of interface type is treated the same as having
// that type as its name, rather than being anonymous.
//
// The Go visibility rules for struct fields are amended for JSON when
// deciding which field to marshal or unmarshal. If there are
// multiple fields at the same level, and that level is the least
// nested (and would therefore be the nesting level selected by the
// usual Go rules), the following extra rules apply:
//
// 1) Of those fields, if any are JSON-tagged, only tagged fields are considered,
// even if there are multiple untagged fields that would otherwise conflict.
//
// 2) If there is exactly one field (tagged or not according to the first rule), that is selected.
//
// 3) Otherwise there are multiple fields, and all are ignored; no error occurs.
//
// Handling of anonymous struct fields is new in Go 1.1.
// Prior to Go 1.1, anonymous struct fields were ignored. To force ignoring of
// an anonymous struct field in both current and earlier versions, give the field
// a JSON tag of &#34;-&#34;.
//
// Map values encode as JSON objects. The map&#39;s key type must either be a
// string, an integer type, or implement encoding.TextMarshaler. The map keys
// are sorted and used as JSON object keys by applying the following rules,
// subject to the UTF-8 coercion described for string values above:
//   - string keys are used directly
//   - encoding.TextMarshalers are marshaled
//   - integer keys are converted to strings
//
// Pointer values encode as the value pointed to.
// A nil pointer encodes as the null JSON value.
//
// Interface values encode as the value contained in the interface.
// A nil interface value encodes as the null JSON value.
//
// Channel, complex, and function values cannot be encoded in JSON.
// Attempting to encode such a value causes Marshal to return
// an UnsupportedTypeError.
//
// JSON cannot represent cyclic data structures and Marshal does not
// handle them. Passing cyclic structures to Marshal will result in
// an infinite recursion.
var Marshal = json.Marshal

// MarshalIndent is like Marshal but applies Indent to format the output.
// Each JSON element in the output will begin on a new line beginning with prefix
// followed by one or more copies of indent according to the indentation nesting.
var MarshalIndent = json.MarshalIndent

// NewDecoder returns a new decoder that reads from r.
//
// The decoder introduces its own buffering and may
// read data from r beyond the JSON values requested.
var NewDecoder = json.NewDecoder

// NewEncoder returns a new encoder that writes to w.
var NewEncoder = json.NewEncoder

// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v. If v is nil or not a pointer,
// Unmarshal returns an InvalidUnmarshalError.
//
// Unmarshal uses the inverse of the encodings that
// Marshal uses, allocating maps, slices, and pointers as necessary,
// with the following additional rules:
//
// To unmarshal JSON into a pointer, Unmarshal first handles the case of
// the JSON being the JSON literal null. In that case, Unmarshal sets
// the pointer to nil. Otherwise, Unmarshal unmarshals the JSON into
// the value pointed at by the pointer. If the pointer is nil, Unmarshal
// allocates a new value for it to point to.
//
// To unmarshal JSON into a value implementing the Unmarshaler interface,
// Unmarshal calls that value&#39;s UnmarshalJSON method, including
// when the input is a JSON null.
// Otherwise, if the value implements encoding.TextUnmarshaler
// and the input is a JSON quoted string, Unmarshal calls that value&#39;s
// UnmarshalText method with the unquoted form of the string.
//
// To unmarshal JSON into a struct, Unmarshal matches incoming object
// keys to the keys used by Marshal (either the struct field name or its tag),
// preferring an exact match but also accepting a case-insensitive match. By
// default, object keys which don&#39;t have a corresponding struct field are
// ignored (see Decoder.DisallowUnknownFields for an alternative).
//
// To unmarshal JSON into an interface value,
// Unmarshal stores one of these in the interface value:
//
// 	bool, for JSON booleans
// 	float64, for JSON numbers
// 	string, for JSON strings
// 	[]interface{}, for JSON arrays
// 	map[string]interface{}, for JSON objects
// 	nil for JSON null
//
// To unmarshal a JSON array into a slice, Unmarshal resets the slice length
// to zero and then appends each element to the slice.
// As a special case, to unmarshal an empty JSON array into a slice,
// Unmarshal replaces the slice with a new empty slice.
//
// To unmarshal a JSON array into a Go array, Unmarshal decodes
// JSON array elements into corresponding Go array elements.
// If the Go array is smaller than the JSON array,
// the additional JSON array elements are discarded.
// If the JSON array is smaller than the Go array,
// the additional Go array elements are set to zero values.
//
// To unmarshal a JSON object into a map, Unmarshal first establishes a map to
// use. If the map is nil, Unmarshal allocates a new map. Otherwise Unmarshal
// reuses the existing map, keeping existing entries. Unmarshal then stores
// key-value pairs from the JSON object into the map. The map&#39;s key type must
// either be a string, an integer, or implement encoding.TextUnmarshaler.
//
// If a JSON value is not appropriate for a given target type,
// or if a JSON number overflows the target type, Unmarshal
// skips that field and completes the unmarshaling as best it can.
// If no more serious errors are encountered, Unmarshal returns
// an UnmarshalTypeError describing the earliest such error. In any
// case, it&#39;s not guaranteed that all the remaining fields following
// the problematic one will be unmarshaled into the target object.
//
// The JSON null value unmarshals into an interface, map, pointer, or slice
// by setting that Go value to nil. Because null is often used in JSON to mean
// ``not present,&#39;&#39; unmarshaling a JSON null into any other Go type has no effect
// on the value and produces no error.
//
// When unmarshaling quoted strings, invalid UTF-8 or
// invalid UTF-16 surrogate pairs are not treated as an error.
// Instead, they are replaced by the Unicode replacement
// character U+FFFD.
var Unmarshal = json.Unmarshal

// Valid reports whether data is a valid JSON encoding.
var Valid = json.Valid
