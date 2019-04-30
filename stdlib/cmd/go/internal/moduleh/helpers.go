package moduleh

import (
	"cmd/go/internal/module"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CanonicalVersionKey = "CanonicalVersion"

	CheckKey = "Check"

	CheckFilePathKey = "CheckFilePath"

	CheckImportPathKey = "CheckImportPath"

	CheckPathKey = "CheckPath"

	DecodePathKey = "DecodePath"

	DecodeVersionKey = "DecodeVersion"

	EncodePathKey = "EncodePath"

	EncodeVersionKey = "EncodeVersion"

	MatchPathMajorKey = "MatchPathMajor"

	SortKey = "Sort"

	SplitPathVersionKey = "SplitPathVersion"
)

func New() hctx.Map {
	return hctx.Map{

		CanonicalVersionKey: CanonicalVersion,

		CheckKey: Check,

		CheckFilePathKey: CheckFilePath,

		CheckImportPathKey: CheckImportPath,

		CheckPathKey: CheckPath,

		DecodePathKey: DecodePath,

		DecodeVersionKey: DecodeVersion,

		EncodePathKey: EncodePath,

		EncodeVersionKey: EncodeVersion,

		MatchPathMajorKey: MatchPathMajor,

		SortKey: Sort,

		SplitPathVersionKey: SplitPathVersion,
	}
}

// CanonicalVersion returns the canonical form of the version string v.
// It is the same as semver.Canonical(v) except that it preserves the special build suffix &#34;+incompatible&#34;.
var CanonicalVersion = module.CanonicalVersion

// Check checks that a given module path, version pair is valid.
// In addition to the path being a valid module path
// and the version being a valid semantic version,
// the two must correspond.
// For example, the path &#34;yaml/v2&#34; only corresponds to
// semantic versions beginning with &#34;v2.&#34;.
var Check = module.Check

// CheckFilePath checks whether a slash-separated file path is valid.
var CheckFilePath = module.CheckFilePath

// CheckImportPath checks that an import path is valid.
var CheckImportPath = module.CheckImportPath

// CheckPath checks that a module path is valid.
var CheckPath = module.CheckPath

// DecodePath returns the module path of the given safe encoding.
// It fails if the encoding is invalid or encodes an invalid path.
var DecodePath = module.DecodePath

// DecodeVersion returns the version string for the given safe encoding.
// It fails if the encoding is invalid or encodes an invalid version.
// Versions are allowed to be in non-semver form but must be valid file names
// and not contain exclamation marks.
var DecodeVersion = module.DecodeVersion

// EncodePath returns the safe encoding of the given module path.
// It fails if the module path is invalid.
var EncodePath = module.EncodePath

// EncodeVersion returns the safe encoding of the given module version.
// Versions are allowed to be in non-semver form but must be valid file names
// and not contain exclamation marks.
var EncodeVersion = module.EncodeVersion

// MatchPathMajor reports whether the semantic version v
// matches the path major version pathMajor.
var MatchPathMajor = module.MatchPathMajor

// Sort sorts the list by Path, breaking ties by comparing Versions.
var Sort = module.Sort

// SplitPathVersion returns prefix and major version such that prefix+pathMajor == path
// and version is either empty or &#34;/vN&#34; for N &gt;= 2.
// As a special case, gopkg.in paths are recognized directly;
// they require &#34;.vN&#34; instead of &#34;/vN&#34;, and for all N, not just N &gt;= 2.
var SplitPathVersion = module.SplitPathVersion
