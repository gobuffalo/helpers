package bytesh

import (
	"bytes"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CompareKey = "Compare"

	ContainsKey = "Contains"

	ContainsAnyKey = "ContainsAny"

	ContainsRuneKey = "ContainsRune"

	CountKey = "Count"

	EqualKey = "Equal"

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

	NewBufferKey = "NewBuffer"

	NewBufferStringKey = "NewBufferString"

	NewReaderKey = "NewReader"

	RepeatKey = "Repeat"

	ReplaceKey = "Replace"

	RunesKey = "Runes"

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

		EqualKey: Equal,

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

		NewBufferKey: NewBuffer,

		NewBufferStringKey: NewBufferString,

		NewReaderKey: NewReader,

		RepeatKey: Repeat,

		ReplaceKey: Replace,

		RunesKey: Runes,

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

// Compare returns an integer comparing two byte slices lexicographically.
// The result will be 0 if a==b, -1 if a &lt; b, and +1 if a &gt; b.
// A nil argument is equivalent to an empty slice.
var Compare = bytes.Compare

// Contains reports whether subslice is within b.
var Contains = bytes.Contains

// ContainsAny reports whether any of the UTF-8-encoded code points in chars are within b.
var ContainsAny = bytes.ContainsAny

// ContainsRune reports whether the rune is contained in the UTF-8-encoded byte slice b.
var ContainsRune = bytes.ContainsRune

// Count counts the number of non-overlapping instances of sep in s.
// If sep is an empty slice, Count returns 1 + the number of UTF-8-encoded code points in s.
var Count = bytes.Count

// Equal returns a boolean reporting whether a and b
// are the same length and contain the same bytes.
// A nil argument is equivalent to an empty slice.
var Equal = bytes.Equal

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under Unicode case-folding.
var EqualFold = bytes.EqualFold

// Fields interprets s as a sequence of UTF-8-encoded code points.
// It splits the slice s around each instance of one or more consecutive white space
// characters, as defined by unicode.IsSpace, returning a slice of subslices of s or an
// empty slice if s contains only white space.
var Fields = bytes.Fields

// FieldsFunc interprets s as a sequence of UTF-8-encoded code points.
// It splits the slice s at each run of code points c satisfying f(c) and
// returns a slice of subslices of s. If all code points in s satisfy f(c), or
// len(s) == 0, an empty slice is returned.
// FieldsFunc makes no guarantees about the order in which it calls f(c).
// If f does not return consistent results for a given c, FieldsFunc may crash.
var FieldsFunc = bytes.FieldsFunc

// HasPrefix tests whether the byte slice s begins with prefix.
var HasPrefix = bytes.HasPrefix

// HasSuffix tests whether the byte slice s ends with suffix.
var HasSuffix = bytes.HasSuffix

// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.
var Index = bytes.Index

// IndexAny interprets s as a sequence of UTF-8-encoded Unicode code points.
// It returns the byte index of the first occurrence in s of any of the Unicode
// code points in chars. It returns -1 if chars is empty or if there is no code
// point in common.
var IndexAny = bytes.IndexAny

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
var IndexByte = bytes.IndexByte

// IndexFunc interprets s as a sequence of UTF-8-encoded code points.
// It returns the byte index in s of the first Unicode
// code point satisfying f(c), or -1 if none do.
var IndexFunc = bytes.IndexFunc

// IndexRune interprets s as a sequence of UTF-8-encoded code points.
// It returns the byte index of the first occurrence in s of the given rune.
// It returns -1 if rune is not present in s.
// If r is utf8.RuneError, it returns the first instance of any
// invalid UTF-8 byte sequence.
var IndexRune = bytes.IndexRune

// Join concatenates the elements of s to create a new byte slice. The separator
// sep is placed between elements in the resulting slice.
var Join = bytes.Join

// LastIndex returns the index of the last instance of sep in s, or -1 if sep is not present in s.
var LastIndex = bytes.LastIndex

// LastIndexAny interprets s as a sequence of UTF-8-encoded Unicode code
// points. It returns the byte index of the last occurrence in s of any of
// the Unicode code points in chars. It returns -1 if chars is empty or if
// there is no code point in common.
var LastIndexAny = bytes.LastIndexAny

// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
var LastIndexByte = bytes.LastIndexByte

// LastIndexFunc interprets s as a sequence of UTF-8-encoded code points.
// It returns the byte index in s of the last Unicode
// code point satisfying f(c), or -1 if none do.
var LastIndexFunc = bytes.LastIndexFunc

// Map returns a copy of the byte slice s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the byte slice with no replacement. The characters in s and the
// output are interpreted as UTF-8-encoded code points.
var Map = bytes.Map

// NewBuffer creates and initializes a new Buffer using buf as its
// initial contents. The new Buffer takes ownership of buf, and the
// caller should not use buf after this call. NewBuffer is intended to
// prepare a Buffer to read existing data. It can also be used to size
// the internal buffer for writing. To do that, buf should have the
// desired capacity but a length of zero.
//
// In most cases, new(Buffer) (or just declaring a Buffer variable) is
// sufficient to initialize a Buffer.
var NewBuffer = bytes.NewBuffer

// NewBufferString creates and initializes a new Buffer using string s as its
// initial contents. It is intended to prepare a buffer to read an existing
// string.
//
// In most cases, new(Buffer) (or just declaring a Buffer variable) is
// sufficient to initialize a Buffer.
var NewBufferString = bytes.NewBufferString

// NewReader returns a new Reader reading from b.
var NewReader = bytes.NewReader

// Repeat returns a new byte slice consisting of count copies of b.
//
// It panics if count is negative or if
// the result of (len(b) * count) overflows.
var Repeat = bytes.Repeat

// Replace returns a copy of the slice s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the slice
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune slice.
// If n &lt; 0, there is no limit on the number of replacements.
var Replace = bytes.Replace

// Runes interprets s as a sequence of UTF-8-encoded code points.
// It returns a slice of runes (Unicode code points) equivalent to s.
var Runes = bytes.Runes

// Split slices s into all subslices separated by sep and returns a slice of
// the subslices between those separators.
// If sep is empty, Split splits after each UTF-8 sequence.
// It is equivalent to SplitN with a count of -1.
var Split = bytes.Split

// SplitAfter slices s into all subslices after each instance of sep and
// returns a slice of those subslices.
// If sep is empty, SplitAfter splits after each UTF-8 sequence.
// It is equivalent to SplitAfterN with a count of -1.
var SplitAfter = bytes.SplitAfter

// SplitAfterN slices s into subslices after each instance of sep and
// returns a slice of those subslices.
// If sep is empty, SplitAfterN splits after each UTF-8 sequence.
// The count determines the number of subslices to return:
//   n &gt; 0: at most n subslices; the last subslice will be the unsplit remainder.
//   n == 0: the result is nil (zero subslices)
//   n &lt; 0: all subslices
var SplitAfterN = bytes.SplitAfterN

// SplitN slices s into subslices separated by sep and returns a slice of
// the subslices between those separators.
// If sep is empty, SplitN splits after each UTF-8 sequence.
// The count determines the number of subslices to return:
//   n &gt; 0: at most n subslices; the last subslice will be the unsplit remainder.
//   n == 0: the result is nil (zero subslices)
//   n &lt; 0: all subslices
var SplitN = bytes.SplitN

// Title treats s as UTF-8-encoded bytes and returns a copy with all Unicode letters that begin
// words mapped to their title case.
//
// BUG(rsc): The rule Title uses for word boundaries does not handle Unicode punctuation properly.
var Title = bytes.Title

// ToLower treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their lower case.
var ToLower = bytes.ToLower

// ToLowerSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their
// lower case, giving priority to the special casing rules.
var ToLowerSpecial = bytes.ToLowerSpecial

// ToTitle treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their title case.
var ToTitle = bytes.ToTitle

// ToTitleSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their
// title case, giving priority to the special casing rules.
var ToTitleSpecial = bytes.ToTitleSpecial

// ToUpper treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters within it mapped to their upper case.
var ToUpper = bytes.ToUpper

// ToUpperSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their
// upper case, giving priority to the special casing rules.
var ToUpperSpecial = bytes.ToUpperSpecial

// Trim returns a subslice of s by slicing off all leading and
// trailing UTF-8-encoded code points contained in cutset.
var Trim = bytes.Trim

// TrimFunc returns a subslice of s by slicing off all leading and trailing
// UTF-8-encoded code points c that satisfy f(c).
var TrimFunc = bytes.TrimFunc

// TrimLeft returns a subslice of s by slicing off all leading
// UTF-8-encoded code points contained in cutset.
var TrimLeft = bytes.TrimLeft

// TrimLeftFunc treats s as UTF-8-encoded bytes and returns a subslice of s by slicing off
// all leading UTF-8-encoded code points c that satisfy f(c).
var TrimLeftFunc = bytes.TrimLeftFunc

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn&#39;t start with prefix, s is returned unchanged.
var TrimPrefix = bytes.TrimPrefix

// TrimRight returns a subslice of s by slicing off all trailing
// UTF-8-encoded code points that are contained in cutset.
var TrimRight = bytes.TrimRight

// TrimRightFunc returns a subslice of s by slicing off all trailing
// UTF-8-encoded code points c that satisfy f(c).
var TrimRightFunc = bytes.TrimRightFunc

// TrimSpace returns a subslice of s by slicing off all leading and
// trailing white space, as defined by Unicode.
var TrimSpace = bytes.TrimSpace

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn&#39;t end with suffix, s is returned unchanged.
var TrimSuffix = bytes.TrimSuffix
