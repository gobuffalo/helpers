package text

import "github.com/gobuffalo/helpers/hctx"

func Truncate(s string, opts hctx.Map) string {
	if opts["size"] == nil {
		opts["size"] = 50
	}
	if opts["trail"] == nil {
		opts["trail"] = "..."
	}
	size := opts["size"].(int)
	if len(s) <= size {
		return s
	}
	trail := opts["trail"].(string)
	if len(trail) >= size {
		return trail
	}
	return s[:size-len(trail)] + trail
}
