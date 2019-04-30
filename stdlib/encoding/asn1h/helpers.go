package asn1h

import (
	"encoding/asn1"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	MarshalKey = "Marshal"

	MarshalWithParamsKey = "MarshalWithParams"

	UnmarshalKey = "Unmarshal"

	UnmarshalWithParamsKey = "UnmarshalWithParams"
)

func New() hctx.Map {
	return hctx.Map{

		MarshalKey: Marshal,

		MarshalWithParamsKey: MarshalWithParams,

		UnmarshalKey: Unmarshal,

		UnmarshalWithParamsKey: UnmarshalWithParams,
	}
}

// Marshal returns the ASN.1 encoding of val.
//
// In addition to the struct tags recognised by Unmarshal, the following can be
// used:
//
// 	ia5:         causes strings to be marshaled as ASN.1, IA5String values
// 	omitempty:   causes empty slices to be skipped
// 	printable:   causes strings to be marshaled as ASN.1, PrintableString values
// 	utf8:        causes strings to be marshaled as ASN.1, UTF8String values
// 	utc:         causes time.Time to be marshaled as ASN.1, UTCTime values
// 	generalized: causes time.Time to be marshaled as ASN.1, GeneralizedTime values
var Marshal = asn1.Marshal

// MarshalWithParams allows field parameters to be specified for the
// top-level element. The form of the params is the same as the field tags.
var MarshalWithParams = asn1.MarshalWithParams

// Unmarshal parses the DER-encoded ASN.1 data structure b
// and uses the reflect package to fill in an arbitrary value pointed at by val.
// Because Unmarshal uses the reflect package, the structs
// being written to must use upper case field names.
//
// An ASN.1 INTEGER can be written to an int, int32, int64,
// or *big.Int (from the math/big package).
// If the encoded value does not fit in the Go type,
// Unmarshal returns a parse error.
//
// An ASN.1 BIT STRING can be written to a BitString.
//
// An ASN.1 OCTET STRING can be written to a []byte.
//
// An ASN.1 OBJECT IDENTIFIER can be written to an
// ObjectIdentifier.
//
// An ASN.1 ENUMERATED can be written to an Enumerated.
//
// An ASN.1 UTCTIME or GENERALIZEDTIME can be written to a time.Time.
//
// An ASN.1 PrintableString, IA5String, or NumericString can be written to a string.
//
// Any of the above ASN.1 values can be written to an interface{}.
// The value stored in the interface has the corresponding Go type.
// For integers, that type is int64.
//
// An ASN.1 SEQUENCE OF x or SET OF x can be written
// to a slice if an x can be written to the slice&#39;s element type.
//
// An ASN.1 SEQUENCE or SET can be written to a struct
// if each of the elements in the sequence can be
// written to the corresponding element in the struct.
//
// The following tags on struct fields have special meaning to Unmarshal:
//
// 	application specifies that an APPLICATION tag is used
// 	private     specifies that a PRIVATE tag is used
// 	default:x   sets the default value for optional integer fields (only used if optional is also present)
// 	explicit    specifies that an additional, explicit tag wraps the implicit one
// 	optional    marks the field as ASN.1 OPTIONAL
// 	set         causes a SET, rather than a SEQUENCE type to be expected
// 	tag:x       specifies the ASN.1 tag number; implies ASN.1 CONTEXT SPECIFIC
//
// If the type of the first field of a structure is RawContent then the raw
// ASN1 contents of the struct will be stored in it.
//
// If the type name of a slice element ends with &#34;SET&#34; then it&#39;s treated as if
// the &#34;set&#34; tag was set on it. This can be used with nested slices where a
// struct tag cannot be given.
//
// Other ASN.1 types are not supported; if it encounters them,
// Unmarshal returns a parse error.
var Unmarshal = asn1.Unmarshal

// UnmarshalWithParams allows field parameters to be specified for the
// top-level element. The form of the params is the same as the field tags.
var UnmarshalWithParams = asn1.UnmarshalWithParams
