package driverh

import (
	"cmd/vendor/github.com/google/pprof/driver"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	PProfKey = "PProf"
)

func New() hctx.Map {
	return hctx.Map{

		PProfKey: PProf,
	}
}

// PProf acquires a profile, and symbolizes it using a profile
// manager. Then it generates a report formatted according to the
// options selected through the flags package.
var PProf = driver.PProf
