package reporth

import (
	"cmd/vendor/github.com/google/pprof/internal/report"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	AddSourceTemplatesKey = "AddSourceTemplates"

	GenerateKey = "Generate"

	GetDOTKey = "GetDOT"

	NewKey = "New"

	NewDefaultKey = "NewDefault"

	PrintAssemblyKey = "PrintAssembly"

	PrintWebListKey = "PrintWebList"

	ProfileLabelsKey = "ProfileLabels"

	TextItemsKey = "TextItems"
)

func New() hctx.Map {
	return hctx.Map{

		AddSourceTemplatesKey: AddSourceTemplates,

		GenerateKey: Generate,

		GetDOTKey: GetDOT,

		NewKey: New,

		NewDefaultKey: NewDefault,

		PrintAssemblyKey: PrintAssembly,

		PrintWebListKey: PrintWebList,

		ProfileLabelsKey: ProfileLabels,

		TextItemsKey: TextItems,
	}
}

// AddSourceTemplates adds templates used by PrintWebList to t.
var AddSourceTemplates = report.AddSourceTemplates

// Generate generates a report as directed by the Report.
var Generate = report.Generate

// GetDOT returns a graph suitable for dot processing along with some
// configuration information.
var GetDOT = report.GetDOT

// New builds a new report indexing the sample values interpreting the
// samples with the provided function.
var New = report.New

// NewDefault builds a new report indexing the last sample value
// available.
var NewDefault = report.NewDefault

// PrintAssembly prints annotated disassembly of rpt to w.
var PrintAssembly = report.PrintAssembly

// PrintWebList prints annotated source listing of rpt to w.
var PrintWebList = report.PrintWebList

// ProfileLabels returns printable labels for a profile.
var ProfileLabels = report.ProfileLabels

// TextItems returns a list of text items from the report and a list
// of labels that describe the report.
var TextItems = report.TextItems
