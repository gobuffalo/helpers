package raceh

import (
	"internal/race"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AcquireKey = "Acquire"

	AcquireKey = "Acquire"

	DisableKey = "Disable"

	DisableKey = "Disable"

	EnableKey = "Enable"

	EnableKey = "Enable"

	ErrorsKey = "Errors"

	ErrorsKey = "Errors"

	ReadKey = "Read"

	ReadKey = "Read"

	ReadRangeKey = "ReadRange"

	ReadRangeKey = "ReadRange"

	ReleaseKey = "Release"

	ReleaseKey = "Release"

	ReleaseMergeKey = "ReleaseMerge"

	ReleaseMergeKey = "ReleaseMerge"

	WriteKey = "Write"

	WriteKey = "Write"

	WriteRangeKey = "WriteRange"

	WriteRangeKey = "WriteRange"
)

func New() hctx.Map {
	return hctx.Map{

		AcquireKey: Acquire,

		AcquireKey: Acquire,

		DisableKey: Disable,

		DisableKey: Disable,

		EnableKey: Enable,

		EnableKey: Enable,

		ErrorsKey: Errors,

		ErrorsKey: Errors,

		ReadKey: Read,

		ReadKey: Read,

		ReadRangeKey: ReadRange,

		ReadRangeKey: ReadRange,

		ReleaseKey: Release,

		ReleaseKey: Release,

		ReleaseMergeKey: ReleaseMerge,

		ReleaseMergeKey: ReleaseMerge,

		WriteKey: Write,

		WriteKey: Write,

		WriteRangeKey: WriteRange,

		WriteRangeKey: WriteRange,
	}
}

var Acquire = race.Acquire

var Acquire = race.Acquire

var Disable = race.Disable

var Disable = race.Disable

var Enable = race.Enable

var Enable = race.Enable

var Errors = race.Errors

var Errors = race.Errors

var Read = race.Read

var Read = race.Read

var ReadRange = race.ReadRange

var ReadRange = race.ReadRange

var Release = race.Release

var Release = race.Release

var ReleaseMerge = race.ReleaseMerge

var ReleaseMerge = race.ReleaseMerge

var Write = race.Write

var Write = race.Write

var WriteRange = race.WriteRange

var WriteRange = race.WriteRange
