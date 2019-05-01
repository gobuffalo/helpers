package stdlib

import (
	"os"
	"path/filepath"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/helpers/internal/decls"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/plushgen"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, err
	}

	g.RunFn(func(r *genny.Runner) error {
		return r.Delete("stdlib")
	})

	box := packr.New("github.com/gobuffalo/helpers/genny/stdlib/templates", "../stdlib/templates")

	helpers, err := decls.FindDecls(filepath.Join(os.Getenv("GOROOT"), "src"))

	if err != nil {
		return g, err
	}
	for _, pk := range helpers {
		s, err := box.FindString("helpers.go.plush")
		if err != nil {
			return g, err
		}
		f := genny.NewFileS(filepath.Join("stdlib", pk.Full+"h", "helpers.go.plush"), s)

		ctx := plush.NewContext()
		ctx.Set("opts", opts)
		ctx.Set("pkg", pk)
		t := plushgen.Transformer(ctx)
		f, err = t.Transform(f)
		if err != nil {
			return g, err
		}
		g.File(f)
	}
	return g, nil
}
