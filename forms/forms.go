package forms

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags"
)

const (
	FormKey    = "form"
	FormForKey = "formFor"
)

func New() hctx.Map {
	return hctx.Map{
		FormKey:    Form,
		FormForKey: FormFor,
		"form_for": FormFor, // Deprecated
	}
}

type helperable interface {
	SetAuthenticityToken(string)
	Append(...tags.Body)
	htmler
}

func helper(opts tags.Options, help hctx.HelperContext, fn func(opts tags.Options) helperable) (template.HTML, error) {
	hn := "f"
	if n, ok := opts["var"]; ok {
		hn = n.(string)
		delete(opts, "var")
	}
	if opts["errors"] == nil && help.Value("errors") != nil {
		opts["errors"] = help.Value("errors")
	}
	form := fn(opts)
	if help.Value("authenticity_token") != nil && opts["method"] != "GET" {
		form.SetAuthenticityToken(fmt.Sprint(help.Value("authenticity_token")))
	}
	help.Set(hn, form)
	s, err := help.Block()
	if err != nil {
		return "", err
	}
	form.Append(s)
	return form.HTML(), nil
}

// htmler generates HTML source
type htmler interface {
	HTML() template.HTML
}
