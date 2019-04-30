package analysish

import (
	"cmd/vendor/golang.org/x/tools/go/analysis"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	ValidateKey = "Validate"
)

func New() hctx.Map {
	return hctx.Map{

		ValidateKey: Validate,
	}
}

// Validate reports an error if any of the analyzers are misconfigured.
// Checks include:
// that the name is a valid identifier;
// that analyzer names are unique;
// that the Requires graph is acylic;
// that analyzer fact types are unique;
// that each fact type is a pointer.
var Validate = analysis.Validate
