package gorooth

import (
	"internal/goroot"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	IsStandardPackageKey = "IsStandardPackage"

	IsStandardPackageKey = "IsStandardPackage"
)

func New() hctx.Map {
	return hctx.Map{

		IsStandardPackageKey: IsStandardPackage,

		IsStandardPackageKey: IsStandardPackage,
	}
}

// IsStandardPackage reports whether path is a standard package,
// given goroot and compiler.
var IsStandardPackage = goroot.IsStandardPackage

// IsStandardPackage reports whether path is a standard package,
// given goroot and compiler.
var IsStandardPackage = goroot.IsStandardPackage
