package docs

import (
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/helpers/internal/decls"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/plushgen"
)

type helper struct {
	Name     string
	FullName string
	Doc      string
}

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, err
	}

	if err := g.Box(packr.New("github.com/gobuffalo/helpers/genny/docs/templates", "../docs/templates")); err != nil {
		return g, err
	}

	ctx := plush.NewContext()
	ctx.Set("opts", opts)

	helpers, err := decls.FindDecls(".")
	if err != nil {
		return g, err
	}
	ctx.Set("helpers", helpers)
	g.Transformer(plushgen.Transformer(ctx))
	return g, nil
}
