package coverh

import (
	"github.com/gobuffalo/helpers/hctx"
	"golang.org/x/tools/cover"
)

const (
	ParseProfilesKey = "ParseProfiles"
)

func New() hctx.Map {
	return hctx.Map{

		ParseProfilesKey: ParseProfiles,
	}
}

// ParseProfiles parses profile data in the specified file and returns a
// Profile for each source file described therein.
var ParseProfiles = cover.ParseProfiles
