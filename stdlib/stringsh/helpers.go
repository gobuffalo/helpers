package stringsh

import (
	"strings"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CompareKey = "Compare"

	ContainsKey = "Contains"

	ContainsAnyKey = "ContainsAny"

	ContainsRuneKey = "ContainsRune"

	CountKey = "Count"

	EqualFoldKey = "EqualFold"

	FieldsKey = "Fields"

	FieldsFuncKey = "FieldsFunc"

	HasPrefixKey = "HasPrefix"

	HasSuffixKey = "HasSuffix"

	IndexKey = "Index"

	IndexAnyKey = "IndexAny"

	IndexByteKey = "IndexByte"

	IndexFuncKey = "IndexFunc"

	IndexRuneKey = "IndexRune"

	JoinKey = "Join"

	LastIndexKey = "LastIndex"

	LastIndexAnyKey = "LastIndexAny"

	LastIndexByteKey = "LastIndexByte"

	LastIndexFuncKey = "LastIndexFunc"

	MapKey = "Map"

	NewReaderKey = "NewReader"

	NewReplacerKey = "NewReplacer"

	RepeatKey = "Repeat"

	ReplaceKey = "Replace"

	ReplaceAllKey = "ReplaceAll"

	SplitKey = "Split"

	SplitAfterKey = "SplitAfter"

	SplitAfterNKey = "SplitAfterN"

	SplitNKey = "SplitN"

	TitleKey = "Title"

	ToLowerKey = "ToLower"

	ToLowerSpecialKey = "ToLowerSpecial"

	ToTitleKey = "ToTitle"

	ToTitleSpecialKey = "ToTitleSpecial"

	ToUpperKey = "ToUpper"

	ToUpperSpecialKey = "ToUpperSpecial"

	TrimKey = "Trim"

	TrimFuncKey = "TrimFunc"

	TrimLeftKey = "TrimLeft"

	TrimLeftFuncKey = "TrimLeftFunc"

	TrimPrefixKey = "TrimPrefix"

	TrimRightKey = "TrimRight"

	TrimRightFuncKey = "TrimRightFunc"

	TrimSpaceKey = "TrimSpace"

	TrimSuffixKey = "TrimSuffix"
)

func New() hctx.Map {
	return hctx.Map{

		CompareKey: Compare,

		ContainsKey: Contains,

		ContainsAnyKey: ContainsAny,

		ContainsRuneKey: ContainsRune,

		CountKey: Count,

		EqualFoldKey: EqualFold,

		FieldsKey: Fields,

		FieldsFuncKey: FieldsFunc,

		HasPrefixKey: HasPrefix,

		HasSuffixKey: HasSuffix,

		IndexKey: Index,

		IndexAnyKey: IndexAny,

		IndexByteKey: IndexByte,

		IndexFuncKey: IndexFunc,

		IndexRuneKey: IndexRune,

		JoinKey: Join,

		LastIndexKey: LastIndex,

		LastIndexAnyKey: LastIndexAny,

		LastIndexByteKey: LastIndexByte,

		LastIndexFuncKey: LastIndexFunc,

		MapKey: Map,

		NewReaderKey: NewReader,

		NewReplacerKey: NewReplacer,

		RepeatKey: Repeat,

		ReplaceKey: Replace,

		ReplaceAllKey: ReplaceAll,

		SplitKey: Split,

		SplitAfterKey: SplitAfter,

		SplitAfterNKey: SplitAfterN,

		SplitNKey: SplitN,

		TitleKey: Title,

		ToLowerKey: ToLower,

		ToLowerSpecialKey: ToLowerSpecial,

		ToTitleKey: ToTitle,

		ToTitleSpecialKey: ToTitleSpecial,

		ToUpperKey: ToUpper,

		ToUpperSpecialKey: ToUpperSpecial,

		TrimKey: Trim,

		TrimFuncKey: TrimFunc,

		TrimLeftKey: TrimLeft,

		TrimLeftFuncKey: TrimLeftFunc,

		TrimPrefixKey: TrimPrefix,

		TrimRightKey: TrimRight,

		TrimRightFuncKey: TrimRightFunc,

		TrimSpaceKey: TrimSpace,

		TrimSuffixKey: TrimSuffix,
	}
}

// Compare returns an integer comparing two strings lexicographically.
// The result will be 0 if a==b, -1 if a &lt; b, and +1 if a &gt; b.
//
// Compare is included only for symmetry with package bytes.
// It is usually clearer and always faster to use the built-in
// string comparison operators ==, &lt;, &gt;, and so on.
var Compare = strings.Compare

// Contains reports whether substr is within s.
var Contains = strings.Contains

// ContainsAny reports whether any Unicode code points in chars are within s.
var ContainsAny = strings.ContainsAny

// ContainsRune reports whether the Unicode code point r is within s.
var ContainsRune = strings.ContainsRune

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
var Count = strings.Count

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under Unicode case-folding.
var EqualFold = strings.EqualFold

// Fields splits the string s around each instance of one or more consecutive white space
// characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
// empty slice if s contains only white space.
var Fields = strings.Fields

// FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c)
// and returns an array of slices of s. If all code points in s satisfy f(c) or the
// string is empty, an empty slice is returned.
// FieldsFunc makes no guarantees about the order in which it calls f(c).
// If f does not return consistent results for a given c, FieldsFunc may crash.
var FieldsFunc = strings.FieldsFunc

// HasPrefix tests whether the string s begins with prefix.
var HasPrefix = strings.HasPrefix

// HasSuffix tests whether the string s ends with suffix.
var HasSuffix = strings.HasSuffix

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
var Index = strings.Index

// IndexAny returns the index of the first instance of any Unicode code point
// from chars in s, or -1 if no Unicode code point from chars is present in s.
var IndexAny = strings.IndexAny

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
var IndexByte = strings.IndexByte

// IndexFunc returns the index into s of the first Unicode
// code point satisfying f(c), or -1 if none do.
var IndexFunc = strings.IndexFunc

// IndexRune returns the index of the first instance of the Unicode code point
// r, or -1 if rune is not present in s.
// If r is utf8.RuneError, it returns the first instance of any
// invalid UTF-8 byte sequence.
var IndexRune = strings.IndexRune

// Join concatenates the elements of a to create a single string. The separator string
// sep is placed between elements in the resulting string.
var Join = strings.Join

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
var LastIndex = strings.LastIndex

// LastIndexAny returns the index of the last instance of any Unicode code
// point from chars in s, or -1 if no Unicode code point from chars is
// present in s.
var LastIndexAny = strings.LastIndexAny

// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
var LastIndexByte = strings.LastIndexByte

// LastIndexFunc returns the index into s of the last
// Unicode code point satisfying f(c), or -1 if none do.
var LastIndexFunc = strings.LastIndexFunc

// Map returns a copy of the string s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the string with no replacement.
var Map = strings.Map

// NewReader returns a new Reader reading from s.
// It is similar to bytes.NewBufferString but more efficient and read-only.
var NewReader = strings.NewReader

// NewReplacer returns a new Replacer from a list of old, new string
// pairs. Replacements are performed in the order they appear in the
// target string, without overlapping matches.
var NewReplacer = strings.NewReplacer

// Repeat returns a new string consisting of count copies of the string s.
//
// It panics if count is negative or if
// the result of (len(s) * count) overflows.
var Repeat = strings.Repeat

// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
// If n &lt; 0, there is no limit on the number of replacements.
var Replace = strings.Replace

// ReplaceAll returns a copy of the string s with all
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
var ReplaceAll = strings.ReplaceAll

// Split slices s into all substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// If s does not contain sep and sep is not empty, Split returns a
// slice of length 1 whose only element is s.
//
// If sep is empty, Split splits after each UTF-8 sequence. If both s
// and sep are empty, Split returns an empty slice.
//
// It is equivalent to SplitN with a count of -1.
var Split = strings.Split

// SplitAfter slices s into all substrings after each instance of sep and
// returns a slice of those substrings.
//
// If s does not contain sep and sep is not empty, SplitAfter returns
// a slice of length 1 whose only element is s.
//
// If sep is empty, SplitAfter splits after each UTF-8 sequence. If
// both s and sep are empty, SplitAfter returns an empty slice.
//
// It is equivalent to SplitAfterN with a count of -1.
var SplitAfter = strings.SplitAfter

// SplitAfterN slices s into substrings after each instance of sep and
// returns a slice of those substrings.
//
// The count determines the number of substrings to return:
//   n &gt; 0: at most n substrings; the last substring will be the unsplit remainder.
//   n == 0: the result is nil (zero substrings)
//   n &lt; 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for SplitAfter.
var SplitAfterN = strings.SplitAfterN

// SplitN slices s into substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// The count determines the number of substrings to return:
//   n &gt; 0: at most n substrings; the last substring will be the unsplit remainder.
//   n == 0: the result is nil (zero substrings)
//   n &lt; 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for Split.
var SplitN = strings.SplitN

// Title returns a copy of the string s with all Unicode letters that begin words
// mapped to their title case.
//
// BUG(rsc): The rule Title uses for word boundaries does not handle Unicode punctuation properly.
var Title = strings.Title

// ToLower returns a copy of the string s with all Unicode letters mapped to their lower case.
var ToLower = strings.ToLower

// ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their
// lower case using the case mapping specified by c.
var ToLowerSpecial = strings.ToLowerSpecial

// ToTitle returns a copy of the string s with all Unicode letters mapped to their title case.
var ToTitle = strings.ToTitle

// ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their
// title case, giving priority to the special casing rules.
var ToTitleSpecial = strings.ToTitleSpecial

// ToUpper returns a copy of the string s with all Unicode letters mapped to their upper case.
var ToUpper = strings.ToUpper

// ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their
// upper case using the case mapping specified by c.
var ToUpperSpecial = strings.ToUpperSpecial

// Trim returns a slice of the string s with all leading and
// trailing Unicode code points contained in cutset removed.
var Trim = strings.Trim

// TrimFunc returns a slice of the string s with all leading
// and trailing Unicode code points c satisfying f(c) removed.
var TrimFunc = strings.TrimFunc

// TrimLeft returns a slice of the string s with all leading
// Unicode code points contained in cutset removed.
//
// To remove a prefix, use TrimPrefix instead.
var TrimLeft = strings.TrimLeft

// TrimLeftFunc returns a slice of the string s with all leading
// Unicode code points c satisfying f(c) removed.
var TrimLeftFunc = strings.TrimLeftFunc

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn&#39;t start with prefix, s is returned unchanged.
var TrimPrefix = strings.TrimPrefix

// TrimRight returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
//
// To remove a suffix, use TrimSuffix instead.
var TrimRight = strings.TrimRight

// TrimRightFunc returns a slice of the string s with all trailing
// Unicode code points c satisfying f(c) removed.
var TrimRightFunc = strings.TrimRightFunc

// TrimSpace returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
var TrimSpace = strings.TrimSpace

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn&#39;t end with suffix, s is returned unchanged.
var TrimSuffix = strings.TrimSuffix
