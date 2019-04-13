package helpers

import (
	"fmt"
	"html/template"
)

type Debug int

// Debug by verbosely printing out using 'pre' tags.
func (d Debug) Debug(v interface{}) template.HTML {
	return template.HTML(fmt.Sprintf("<pre>%s</pre>", d.Inspect(v)))
}

func (Debug) Inspect(v interface{}) string {
	return fmt.Sprintf("%+v", v)
}
