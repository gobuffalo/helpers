package profileh

import (
	"cmd/vendor/github.com/google/pprof/profile"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	MergeKey = "Merge"

	ParseKey = "Parse"

	ParseDataKey = "ParseData"

	ParseProcMapsKey = "ParseProcMaps"

	ParseUncompressedKey = "ParseUncompressed"
)

func New() hctx.Map {
	return hctx.Map{

		MergeKey: Merge,

		ParseKey: Parse,

		ParseDataKey: ParseData,

		ParseProcMapsKey: ParseProcMaps,

		ParseUncompressedKey: ParseUncompressed,
	}
}

// Merge merges all the profiles in profs into a single Profile.
// Returns a new profile independent of the input profiles. The merged
// profile is compacted to eliminate unused samples, locations,
// functions and mappings. Profiles must have identical profile sample
// and period types or the merge will fail. profile.Period of the
// resulting profile will be the maximum of all profiles, and
// profile.TimeNanos will be the earliest nonzero one.
var Merge = profile.Merge

// Parse parses a profile and checks for its validity. The input
// may be a gzip-compressed encoded protobuf or one of many legacy
// profile formats which may be unsupported in the future.
var Parse = profile.Parse

// ParseData parses a profile from a buffer and checks for its
// validity.
var ParseData = profile.ParseData

// ParseProcMaps parses a memory map in the format of /proc/self/maps.
// ParseMemoryMap should be called after setting on a profile to
// associate locations to the corresponding mapping based on their
// address.
var ParseProcMaps = profile.ParseProcMaps

// ParseUncompressed parses an uncompressed protobuf into a profile.
var ParseUncompressed = profile.ParseUncompressed
