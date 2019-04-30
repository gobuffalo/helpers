package unicodeh

import (
	"unicode"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	InKey = "In"

	IsKey = "Is"

	IsControlKey = "IsControl"

	IsDigitKey = "IsDigit"

	IsGraphicKey = "IsGraphic"

	IsLetterKey = "IsLetter"

	IsLowerKey = "IsLower"

	IsMarkKey = "IsMark"

	IsNumberKey = "IsNumber"

	IsOneOfKey = "IsOneOf"

	IsPrintKey = "IsPrint"

	IsPunctKey = "IsPunct"

	IsSpaceKey = "IsSpace"

	IsSymbolKey = "IsSymbol"

	IsTitleKey = "IsTitle"

	IsUpperKey = "IsUpper"

	SimpleFoldKey = "SimpleFold"

	ToKey = "To"

	ToLowerKey = "ToLower"

	ToTitleKey = "ToTitle"

	ToUpperKey = "ToUpper"
)

func New() hctx.Map {
	return hctx.Map{

		InKey: In,

		IsKey: Is,

		IsControlKey: IsControl,

		IsDigitKey: IsDigit,

		IsGraphicKey: IsGraphic,

		IsLetterKey: IsLetter,

		IsLowerKey: IsLower,

		IsMarkKey: IsMark,

		IsNumberKey: IsNumber,

		IsOneOfKey: IsOneOf,

		IsPrintKey: IsPrint,

		IsPunctKey: IsPunct,

		IsSpaceKey: IsSpace,

		IsSymbolKey: IsSymbol,

		IsTitleKey: IsTitle,

		IsUpperKey: IsUpper,

		SimpleFoldKey: SimpleFold,

		ToKey: To,

		ToLowerKey: ToLower,

		ToTitleKey: ToTitle,

		ToUpperKey: ToUpper,
	}
}

// In reports whether the rune is a member of one of the ranges.
var In = unicode.In

// Is reports whether the rune is in the specified table of ranges.
var Is = unicode.Is

// IsControl reports whether the rune is a control character.
// The C (Other) Unicode category includes more code points
// such as surrogates; use Is(C, r) to test for them.
var IsControl = unicode.IsControl

// IsDigit reports whether the rune is a decimal digit.
var IsDigit = unicode.IsDigit

// IsGraphic reports whether the rune is defined as a Graphic by Unicode.
// Such characters include letters, marks, numbers, punctuation, symbols, and
// spaces, from categories L, M, N, P, S, Zs.
var IsGraphic = unicode.IsGraphic

// IsLetter reports whether the rune is a letter (category L).
var IsLetter = unicode.IsLetter

// IsLower reports whether the rune is a lower case letter.
var IsLower = unicode.IsLower

// IsMark reports whether the rune is a mark character (category M).
var IsMark = unicode.IsMark

// IsNumber reports whether the rune is a number (category N).
var IsNumber = unicode.IsNumber

// IsOneOf reports whether the rune is a member of one of the ranges.
// The function &#34;In&#34; provides a nicer signature and should be used in preference to IsOneOf.
var IsOneOf = unicode.IsOneOf

// IsPrint reports whether the rune is defined as printable by Go. Such
// characters include letters, marks, numbers, punctuation, symbols, and the
// ASCII space character, from categories L, M, N, P, S and the ASCII space
// character. This categorization is the same as IsGraphic except that the
// only spacing character is ASCII space, U+0020.
var IsPrint = unicode.IsPrint

// IsPunct reports whether the rune is a Unicode punctuation character
// (category P).
var IsPunct = unicode.IsPunct

// IsSpace reports whether the rune is a space character as defined
// by Unicode&#39;s White Space property; in the Latin-1 space
// this is
// 	&#39;\t&#39;, &#39;\n&#39;, &#39;\v&#39;, &#39;\f&#39;, &#39;\r&#39;, &#39; &#39;, U+0085 (NEL), U+00A0 (NBSP).
// Other definitions of spacing characters are set by category
// Z and property Pattern_White_Space.
var IsSpace = unicode.IsSpace

// IsSymbol reports whether the rune is a symbolic character.
var IsSymbol = unicode.IsSymbol

// IsTitle reports whether the rune is a title case letter.
var IsTitle = unicode.IsTitle

// IsUpper reports whether the rune is an upper case letter.
var IsUpper = unicode.IsUpper

// SimpleFold iterates over Unicode code points equivalent under
// the Unicode-defined simple case folding. Among the code points
// equivalent to rune (including rune itself), SimpleFold returns the
// smallest rune &gt; r if one exists, or else the smallest rune &gt;= 0.
// If r is not a valid Unicode code point, SimpleFold(r) returns r.
//
// For example:
// 	SimpleFold(&#39;A&#39;) = &#39;a&#39;
// 	SimpleFold(&#39;a&#39;) = &#39;A&#39;
//
// 	SimpleFold(&#39;K&#39;) = &#39;k&#39;
// 	SimpleFold(&#39;k&#39;) = &#39;\u212A&#39; (Kelvin symbol, K)
// 	SimpleFold(&#39;\u212A&#39;) = &#39;K&#39;
//
// 	SimpleFold(&#39;1&#39;) = &#39;1&#39;
//
// 	SimpleFold(-2) = -2
var SimpleFold = unicode.SimpleFold

// To maps the rune to the specified case: UpperCase, LowerCase, or TitleCase.
var To = unicode.To

// ToLower maps the rune to lower case.
var ToLower = unicode.ToLower

// ToTitle maps the rune to title case.
var ToTitle = unicode.ToTitle

// ToUpper maps the rune to upper case.
var ToUpper = unicode.ToUpper
