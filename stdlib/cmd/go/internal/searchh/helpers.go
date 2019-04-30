package searchh

import (
	"cmd/go/internal/search"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CleanPatternsKey = "CleanPatterns"

	ImportPathsKey = "ImportPaths"

	ImportPathsQuietKey = "ImportPathsQuiet"

	InDirKey = "InDir"

	IsMetaPackageKey = "IsMetaPackage"

	IsRelativePathKey = "IsRelativePath"

	IsStandardImportPathKey = "IsStandardImportPath"

	MatchPackagesKey = "MatchPackages"

	MatchPackagesInFSKey = "MatchPackagesInFS"

	MatchPatternKey = "MatchPattern"

	SetModRootKey = "SetModRoot"

	TreeCanMatchPatternKey = "TreeCanMatchPattern"

	WarnUnmatchedKey = "WarnUnmatched"
)

func New() hctx.Map {
	return hctx.Map{

		CleanPatternsKey: CleanPatterns,

		ImportPathsKey: ImportPaths,

		ImportPathsQuietKey: ImportPathsQuiet,

		InDirKey: InDir,

		IsMetaPackageKey: IsMetaPackage,

		IsRelativePathKey: IsRelativePath,

		IsStandardImportPathKey: IsStandardImportPath,

		MatchPackagesKey: MatchPackages,

		MatchPackagesInFSKey: MatchPackagesInFS,

		MatchPatternKey: MatchPattern,

		SetModRootKey: SetModRoot,

		TreeCanMatchPatternKey: TreeCanMatchPattern,

		WarnUnmatchedKey: WarnUnmatched,
	}
}

// CleanPatterns returns the patterns to use for the given
// command line. It canonicalizes the patterns but does not
// evaluate any matches.
var CleanPatterns = search.CleanPatterns

// ImportPaths returns the matching paths to use for the given command line.
// It calls ImportPathsQuiet and then WarnUnmatched.
var ImportPaths = search.ImportPaths

// ImportPathsQuiet is like ImportPaths but does not warn about patterns with no matches.
var ImportPathsQuiet = search.ImportPathsQuiet

// InDir checks whether path is in the file tree rooted at dir.
// If so, InDir returns an equivalent path relative to dir.
// If not, InDir returns an empty string.
// InDir makes some effort to succeed even in the presence of symbolic links.
// TODO(rsc): Replace internal/test.inDir with a call to this function for Go 1.12.
var InDir = search.InDir

// IsMetaPackage checks if name is a reserved package name that expands to multiple packages.
var IsMetaPackage = search.IsMetaPackage

// IsRelativePath reports whether pattern should be interpreted as a directory
// path relative to the current directory, as opposed to a pattern matching
// import paths.
var IsRelativePath = search.IsRelativePath

// IsStandardImportPath reports whether $GOROOT/src/path should be considered
// part of the standard distribution. For historical reasons we allow people to add
// their own code to $GOROOT instead of using $GOPATH, but we assume that
// code will start with a domain name (dot in the first element).
//
// Note that this function is meant to evaluate whether a directory found in GOROOT
// should be treated as part of the standard library. It should not be used to decide
// that a directory found in GOPATH should be rejected: directories in GOPATH
// need not have dots in the first element, and they just take their chances
// with future collisions in the standard library.
var IsStandardImportPath = search.IsStandardImportPath

// MatchPackages returns all the packages that can be found
// under the $GOPATH directories and $GOROOT matching pattern.
// The pattern is either &#34;all&#34; (all packages), &#34;std&#34; (standard packages),
// &#34;cmd&#34; (standard commands), or a path including &#34;...&#34;.
var MatchPackages = search.MatchPackages

// MatchPackagesInFS is like allPackages but is passed a pattern
// beginning ./ or ../, meaning it should scan the tree rooted
// at the given directory. There are ... in the pattern too.
// (See go help packages for pattern syntax.)
var MatchPackagesInFS = search.MatchPackagesInFS

// MatchPattern(pattern)(name) reports whether
// name matches pattern. Pattern is a limited glob
// pattern in which &#39;...&#39; means &#39;any string&#39; and there
// is no other special syntax.
// Unfortunately, there are two special cases. Quoting &#34;go help packages&#34;:
//
// First, /... at the end of the pattern can match an empty string,
// so that net/... matches both net and packages in its subdirectories, like net/http.
// Second, any slash-separted pattern element containing a wildcard never
// participates in a match of the &#34;vendor&#34; element in the path of a vendored
// package, so that ./... does not match packages in subdirectories of
// ./vendor or ./mycode/vendor, but ./vendor/... and ./mycode/vendor/... do.
// Note, however, that a directory named vendor that itself contains code
// is not a vendored package: cmd/vendor would be a command named vendor,
// and the pattern cmd/... matches it.
var MatchPattern = search.MatchPattern

var SetModRoot = search.SetModRoot

// TreeCanMatchPattern(pattern)(name) reports whether
// name or children of name can possibly match pattern.
// Pattern is the same limited glob accepted by matchPattern.
var TreeCanMatchPattern = search.TreeCanMatchPattern

// WarnUnmatched warns about patterns that didn&#39;t match any packages.
var WarnUnmatched = search.WarnUnmatched
