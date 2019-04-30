package measurementh

import (
	"cmd/vendor/github.com/google/pprof/internal/measurement"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	CommonValueTypeKey = "CommonValueType"

	LabelKey = "Label"

	PercentageKey = "Percentage"

	ScaleKey = "Scale"

	ScaleProfilesKey = "ScaleProfiles"

	ScaledLabelKey = "ScaledLabel"
)

func New() hctx.Map {
	return hctx.Map{

		CommonValueTypeKey: CommonValueType,

		LabelKey: Label,

		PercentageKey: Percentage,

		ScaleKey: Scale,

		ScaleProfilesKey: ScaleProfiles,

		ScaledLabelKey: ScaledLabel,
	}
}

// CommonValueType returns the finest type from a set of compatible
// types.
var CommonValueType = measurement.CommonValueType

// Label returns the label used to describe a certain measurement.
var Label = measurement.Label

// Percentage computes the percentage of total of a value, and encodes
// it as a string. At least two digits of precision are printed.
var Percentage = measurement.Percentage

// Scale a measurement from an unit to a different unit and returns
// the scaled value and the target unit. The returned target unit
// will be empty if uninteresting (could be skipped).
var Scale = measurement.Scale

// ScaleProfiles updates the units in a set of profiles to make them
// compatible. It scales the profiles to the smallest unit to preserve
// data.
var ScaleProfiles = measurement.ScaleProfiles

// ScaledLabel scales the passed-in measurement (if necessary) and
// returns the label used to describe a float measurement.
var ScaledLabel = measurement.ScaledLabel
