package srcimporterh

import (
	"go/internal/srcimporter"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewKey = "New"
)

func New() hctx.Map {
	return hctx.Map{

		NewKey: New,
	}
}

// NewImporter returns a new Importer for the given context, file set, and map
// of packages. The context is used to resolve import paths to package paths,
// and identifying the files belonging to the package. If the context provides
// non-nil file system functions, they are used instead of the regular package
// os functions. The file set is used to track position information of package
// files; and imported packages are added to the packages map.
var New = srcimporter.New
