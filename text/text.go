package text

import "github.com/gobuffalo/helpers/hctx"

const (
	MarkdownKey = "markdown"
	TruncateKey = "truncate"
)

func New() hctx.Map {
	return hctx.Map{
		MarkdownKey: Markdown,
		TruncateKey: Truncate,
	}
}
