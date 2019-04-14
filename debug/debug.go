package debug

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/helpers/hctx"
)

func New() hctx.Map {
	return hctx.Map{
		"debug":   Debug,
		"inspect": Inspect,
	}
}

// Debug by verbosely printing out using 'pre' tags.
func Debug(v interface{}) template.HTML {
	return template.HTML(fmt.Sprintf("<pre>%s</pre>", Inspect(v)))
}
