package debug

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DebugKey   = "debug"
	InspectKey = "inspect"
)

func New() hctx.Map {
	return hctx.Map{
		DebugKey:   Debug,
		InspectKey: Inspect,
	}
}

// Debug by verbosely printing out using 'pre' tags.
func Debug(v interface{}) template.HTML {
	return template.HTML(fmt.Sprintf("<pre>%s</pre>", Inspect(v)))
}
