package strh

import (
	"cmd/go/internal/str"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ContainsKey = "Contains"

	FoldDupKey = "FoldDup"

	HasFilePathPrefixKey = "HasFilePathPrefix"

	HasPathPrefixKey = "HasPathPrefix"

	SplitQuotedFieldsKey = "SplitQuotedFields"

	StringListKey = "StringList"

	ToFoldKey = "ToFold"
)

func New() hctx.Map {
	return hctx.Map{

		ContainsKey: Contains,

		FoldDupKey: FoldDup,

		HasFilePathPrefixKey: HasFilePathPrefix,

		HasPathPrefixKey: HasPathPrefix,

		SplitQuotedFieldsKey: SplitQuotedFields,

		StringListKey: StringList,

		ToFoldKey: ToFold,
	}
}

// Contains reports whether x contains s.
var Contains = str.Contains

// FoldDup reports a pair of strings from the list that are
// equal according to strings.EqualFold.
// It returns &#34;&#34;, &#34;&#34; if there are no such strings.
var FoldDup = str.FoldDup

// HasFilePathPrefix reports whether the filesystem path s
// begins with the elements in prefix.
var HasFilePathPrefix = str.HasFilePathPrefix

// HasPath reports whether the slash-separated path s
// begins with the elements in prefix.
var HasPathPrefix = str.HasPathPrefix

// SplitQuotedFields splits s into a list of fields,
// allowing single or double quotes around elements.
// There is no unescaping or other processing within
// quoted fields.
var SplitQuotedFields = str.SplitQuotedFields

// StringList flattens its arguments into a single []string.
// Each argument in args must have type string or []string.
var StringList = str.StringList

// ToFold returns a string with the property that
// 	strings.EqualFold(s, t) iff ToFold(s) == ToFold(t)
// This lets us test a large set of strings for fold-equivalent
// duplicates without making a quadratic number of calls
// to EqualFold. Note that strings.ToUpper and strings.ToLower
// do not have the desired property in some corner cases.
var ToFold = str.ToFold
