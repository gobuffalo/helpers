package encoders

import "html/template"

func Raw(s string) template.HTML {
	return template.HTML(s)
}
