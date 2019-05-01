package pathh

import (
	"path"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	BaseKey = "Base"

	CleanKey = "Clean"

	DirKey = "Dir"

	ExtKey = "Ext"

	IsAbsKey = "IsAbs"

	JoinKey = "Join"

	MatchKey = "Match"

	SplitKey = "Split"
)

func New() hctx.Map {
	return hctx.Map{

		BaseKey: Base,

		CleanKey: Clean,

		DirKey: Dir,

		ExtKey: Ext,

		IsAbsKey: IsAbs,

		JoinKey: Join,

		MatchKey: Match,

		SplitKey: Split,
	}
}

// Base returns the last element of path.
// Trailing slashes are removed before extracting the last element.
// If the path is empty, Base returns &#34;.&#34;.
// If the path consists entirely of slashes, Base returns &#34;/&#34;.
var Base = path.Base

// Clean returns the shortest path name equivalent to path
// by purely lexical processing. It applies the following rules
// iteratively until no further processing can be done:
//
// 	1. Replace multiple slashes with a single slash.
// 	2. Eliminate each . path name element (the current directory).
// 	3. Eliminate each inner .. path name element (the parent directory)
// 	   along with the non-.. element that precedes it.
// 	4. Eliminate .. elements that begin a rooted path:
// 	   that is, replace &#34;/..&#34; by &#34;/&#34; at the beginning of a path.
//
// The returned path ends in a slash only if it is the root &#34;/&#34;.
//
// If the result of this process is an empty string, Clean
// returns the string &#34;.&#34;.
//
// See also Rob Pike, ``Lexical File Names in Plan 9 or
// Getting Dot-Dot Right,&#39;&#39;
// https://9p.io/sys/doc/lexnames.html
var Clean = path.Clean

// Dir returns all but the last element of path, typically the path&#39;s directory.
// After dropping the final element using Split, the path is Cleaned and trailing
// slashes are removed.
// If the path is empty, Dir returns &#34;.&#34;.
// If the path consists entirely of slashes followed by non-slash bytes, Dir
// returns a single slash. In any other case, the returned path does not end in a
// slash.
var Dir = path.Dir

// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final slash-separated element of path;
// it is empty if there is no dot.
var Ext = path.Ext

// IsAbs reports whether the path is absolute.
var IsAbs = path.IsAbs

// Join joins any number of path elements into a single path, adding a
// separating slash if necessary. The result is Cleaned; in particular,
// all empty strings are ignored.
var Join = path.Join

// Match reports whether name matches the shell file name pattern.
// The pattern syntax is:
//
// 	pattern:
// 		{ term }
// 	term:
// 		&#39;*&#39;         matches any sequence of non-/ characters
// 		&#39;?&#39;         matches any single non-/ character
// 		&#39;[&#39; [ &#39;^&#39; ] { character-range } &#39;]&#39;
// 		            character class (must be non-empty)
// 		c           matches character c (c != &#39;*&#39;, &#39;?&#39;, &#39;\\&#39;, &#39;[&#39;)
// 		&#39;\\&#39; c      matches character c
//
// 	character-range:
// 		c           matches character c (c != &#39;\\&#39;, &#39;-&#39;, &#39;]&#39;)
// 		&#39;\\&#39; c      matches character c
// 		lo &#39;-&#39; hi   matches character c for lo &lt;= c &lt;= hi
//
// Match requires pattern to match all of name, not just a substring.
// The only possible returned error is ErrBadPattern, when pattern
// is malformed.
var Match = path.Match

// Split splits path immediately following the final slash,
// separating it into a directory and file name component.
// If there is no slash in path, Split returns an empty dir and
// file set to path.
// The returned values have the property that path = dir+file.
var Split = path.Split
