package semverh

import (
	"cmd/go/internal/semver"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	BuildKey = "Build"

	CanonicalKey = "Canonical"

	CompareKey = "Compare"

	IsValidKey = "IsValid"

	MajorKey = "Major"

	MajorMinorKey = "MajorMinor"

	MaxKey = "Max"

	PrereleaseKey = "Prerelease"
)

func New() hctx.Map {
	return hctx.Map{

		BuildKey: Build,

		CanonicalKey: Canonical,

		CompareKey: Compare,

		IsValidKey: IsValid,

		MajorKey: Major,

		MajorMinorKey: MajorMinor,

		MaxKey: Max,

		PrereleaseKey: Prerelease,
	}
}

// Build returns the build suffix of the semantic version v.
// For example, Build(&#34;v2.1.0+meta&#34;) == &#34;+meta&#34;.
// If v is an invalid semantic version string, Build returns the empty string.
var Build = semver.Build

// Canonical returns the canonical formatting of the semantic version v.
// It fills in any missing .MINOR or .PATCH and discards build metadata.
// Two semantic versions compare equal only if their canonical formattings
// are identical strings.
// The canonical invalid semantic version is the empty string.
var Canonical = semver.Canonical

// Compare returns an integer comparing two versions according to
// according to semantic version precedence.
// The result will be 0 if v == w, -1 if v &lt; w, or +1 if v &gt; w.
//
// An invalid semantic version string is considered less than a valid one.
// All invalid semantic version strings compare equal to each other.
var Compare = semver.Compare

// IsValid reports whether v is a valid semantic version string.
var IsValid = semver.IsValid

// Major returns the major version prefix of the semantic version v.
// For example, Major(&#34;v2.1.0&#34;) == &#34;v2&#34;.
// If v is an invalid semantic version string, Major returns the empty string.
var Major = semver.Major

// MajorMinor returns the major.minor version prefix of the semantic version v.
// For example, MajorMinor(&#34;v2.1.0&#34;) == &#34;v2.1&#34;.
// If v is an invalid semantic version string, MajorMinor returns the empty string.
var MajorMinor = semver.MajorMinor

// Max canonicalizes its arguments and then returns the version string
// that compares greater.
var Max = semver.Max

// Prerelease returns the prerelease suffix of the semantic version v.
// For example, Prerelease(&#34;v2.1.0-pre+meta&#34;) == &#34;-pre&#34;.
// If v is an invalid semantic version string, Prerelease returns the empty string.
var Prerelease = semver.Prerelease
