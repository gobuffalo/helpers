package sysh

import (
	"cmd/internal/sys"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	MSanSupportedKey = "MSanSupported"

	RaceDetectorSupportedKey = "RaceDetectorSupported"
)

func New() hctx.Map {
	return hctx.Map{

		MSanSupportedKey: MSanSupported,

		RaceDetectorSupportedKey: RaceDetectorSupported,
	}
}

// MSanSupported reports whether goos/goarch supports the memory
// sanitizer option. There is a copy of this function in cmd/dist/test.go.
var MSanSupported = sys.MSanSupported

// RaceDetectorSupported reports whether goos/goarch supports the race
// detector. There is a copy of this function in cmd/dist/test.go.
var RaceDetectorSupported = sys.RaceDetectorSupported
