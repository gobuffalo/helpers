package filepathh

import (
	"path/filepath"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AbsKey = "Abs"

	BaseKey = "Base"

	CleanKey = "Clean"

	DirKey = "Dir"

	EvalSymlinksKey = "EvalSymlinks"

	ExtKey = "Ext"

	FromSlashKey = "FromSlash"

	GlobKey = "Glob"

	HasPrefixKey = "HasPrefix"

	HasPrefixKey = "HasPrefix"

	HasPrefixKey = "HasPrefix"

	IsAbsKey = "IsAbs"

	IsAbsKey = "IsAbs"

	IsAbsKey = "IsAbs"

	JoinKey = "Join"

	MatchKey = "Match"

	RelKey = "Rel"

	SplitKey = "Split"

	SplitListKey = "SplitList"

	ToSlashKey = "ToSlash"

	VolumeNameKey = "VolumeName"

	WalkKey = "Walk"
)

func New() hctx.Map {
	return hctx.Map{

		AbsKey: Abs,

		BaseKey: Base,

		CleanKey: Clean,

		DirKey: Dir,

		EvalSymlinksKey: EvalSymlinks,

		ExtKey: Ext,

		FromSlashKey: FromSlash,

		GlobKey: Glob,

		HasPrefixKey: HasPrefix,

		HasPrefixKey: HasPrefix,

		HasPrefixKey: HasPrefix,

		IsAbsKey: IsAbs,

		IsAbsKey: IsAbs,

		IsAbsKey: IsAbs,

		JoinKey: Join,

		MatchKey: Match,

		RelKey: Rel,

		SplitKey: Split,

		SplitListKey: SplitList,

		ToSlashKey: ToSlash,

		VolumeNameKey: VolumeName,

		WalkKey: Walk,
	}
}

// Abs returns an absolute representation of path.
// If the path is not absolute it will be joined with the current
// working directory to turn it into an absolute path. The absolute
// path name for a given file is not guaranteed to be unique.
// Abs calls Clean on the result.
var Abs = filepath.Abs

// Base returns the last element of path.
// Trailing path separators are removed before extracting the last element.
// If the path is empty, Base returns &#34;.&#34;.
// If the path consists entirely of separators, Base returns a single separator.
var Base = filepath.Base

// Clean returns the shortest path name equivalent to path
// by purely lexical processing. It applies the following rules
// iteratively until no further processing can be done:
//
// 	1. Replace multiple Separator elements with a single one.
// 	2. Eliminate each . path name element (the current directory).
// 	3. Eliminate each inner .. path name element (the parent directory)
// 	   along with the non-.. element that precedes it.
// 	4. Eliminate .. elements that begin a rooted path:
// 	   that is, replace &#34;/..&#34; by &#34;/&#34; at the beginning of a path,
// 	   assuming Separator is &#39;/&#39;.
//
// The returned path ends in a slash only if it represents a root directory,
// such as &#34;/&#34; on Unix or `C:\` on Windows.
//
// Finally, any occurrences of slash are replaced by Separator.
//
// If the result of this process is an empty string, Clean
// returns the string &#34;.&#34;.
//
// See also Rob Pike, ``Lexical File Names in Plan 9 or
// Getting Dot-Dot Right,&#39;&#39;
// https://9p.io/sys/doc/lexnames.html
var Clean = filepath.Clean

// Dir returns all but the last element of path, typically the path&#39;s directory.
// After dropping the final element, Dir calls Clean on the path and trailing
// slashes are removed.
// If the path is empty, Dir returns &#34;.&#34;.
// If the path consists entirely of separators, Dir returns a single separator.
// The returned path does not end in a separator unless it is the root directory.
var Dir = filepath.Dir

// EvalSymlinks returns the path name after the evaluation of any symbolic
// links.
// If path is relative the result will be relative to the current directory,
// unless one of the components is an absolute symbolic link.
// EvalSymlinks calls Clean on the result.
var EvalSymlinks = filepath.EvalSymlinks

// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final element of path; it is empty if there is
// no dot.
var Ext = filepath.Ext

// FromSlash returns the result of replacing each slash (&#39;/&#39;) character
// in path with a separator character. Multiple slashes are replaced
// by multiple separators.
var FromSlash = filepath.FromSlash

// Glob returns the names of all files matching pattern or nil
// if there is no matching file. The syntax of patterns is the same
// as in Match. The pattern may describe hierarchical names such as
// /usr/*/bin/ed (assuming the Separator is &#39;/&#39;).
//
// Glob ignores file system errors such as I/O errors reading directories.
// The only possible returned error is ErrBadPattern, when pattern
// is malformed.
var Glob = filepath.Glob

// HasPrefix exists for historical compatibility and should not be used.
//
// Deprecated: HasPrefix does not respect path boundaries and
// does not ignore case when required.
var HasPrefix = filepath.HasPrefix

// HasPrefix exists for historical compatibility and should not be used.
//
// Deprecated: HasPrefix does not respect path boundaries and
// does not ignore case when required.
var HasPrefix = filepath.HasPrefix

// HasPrefix exists for historical compatibility and should not be used.
//
// Deprecated: HasPrefix does not respect path boundaries and
// does not ignore case when required.
var HasPrefix = filepath.HasPrefix

// IsAbs reports whether the path is absolute.
var IsAbs = filepath.IsAbs

// IsAbs reports whether the path is absolute.
var IsAbs = filepath.IsAbs

// IsAbs reports whether the path is absolute.
var IsAbs = filepath.IsAbs

// Join joins any number of path elements into a single path, adding
// a Separator if necessary. Join calls Clean on the result; in particular,
// all empty strings are ignored.
// On Windows, the result is a UNC path if and only if the first path
// element is a UNC path.
var Join = filepath.Join

// Match reports whether name matches the shell file name pattern.
// The pattern syntax is:
//
// 	pattern:
// 		{ term }
// 	term:
// 		&#39;*&#39;         matches any sequence of non-Separator characters
// 		&#39;?&#39;         matches any single non-Separator character
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
//
// On Windows, escaping is disabled. Instead, &#39;\\&#39; is treated as
// path separator.
var Match = filepath.Match

// Rel returns a relative path that is lexically equivalent to targpath when
// joined to basepath with an intervening separator. That is,
// Join(basepath, Rel(basepath, targpath)) is equivalent to targpath itself.
// On success, the returned path will always be relative to basepath,
// even if basepath and targpath share no elements.
// An error is returned if targpath can&#39;t be made relative to basepath or if
// knowing the current working directory would be necessary to compute it.
// Rel calls Clean on the result.
var Rel = filepath.Rel

// Split splits path immediately following the final Separator,
// separating it into a directory and file name component.
// If there is no Separator in path, Split returns an empty dir
// and file set to path.
// The returned values have the property that path = dir+file.
var Split = filepath.Split

// SplitList splits a list of paths joined by the OS-specific ListSeparator,
// usually found in PATH or GOPATH environment variables.
// Unlike strings.Split, SplitList returns an empty slice when passed an empty
// string.
var SplitList = filepath.SplitList

// ToSlash returns the result of replacing each separator character
// in path with a slash (&#39;/&#39;) character. Multiple separators are
// replaced by multiple slashes.
var ToSlash = filepath.ToSlash

// VolumeName returns leading volume name.
// Given &#34;C:\foo\bar&#34; it returns &#34;C:&#34; on Windows.
// Given &#34;\\host\share\foo&#34; it returns &#34;\\host\share&#34;.
// On other platforms it returns &#34;&#34;.
var VolumeName = filepath.VolumeName

// Walk walks the file tree rooted at root, calling walkFn for each file or
// directory in the tree, including root. All errors that arise visiting files
// and directories are filtered by walkFn. The files are walked in lexical
// order, which makes the output deterministic but means that for very
// large directories Walk can be inefficient.
// Walk does not follow symbolic links.
var Walk = filepath.Walk
