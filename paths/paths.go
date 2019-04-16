package paths

import "github.com/gobuffalo/helpers/hctx"

const (
	PathForKey = "pathFor"
)

func New() hctx.Map {
	return hctx.Map{
		PathForKey: PathFor,
	}
}
