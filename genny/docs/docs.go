package docs

import (
	"fmt"
	"go/ast"
	"go/parser"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/gogen"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/plushgen"
)

type helper struct {
	Name string
	Doc  string
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

	helpers, err := findHelpers()
	if err != nil {
		return g, err
	}
	ctx.Set("helpers", helpers)
	g.Transformer(plushgen.Transformer(ctx))
	return g, nil
}

func findHelpers() ([]helper, error) {
	var helpers []helper
	box := packr.Folder(".")
	err := box.Walk(func(path string, f packd.File) error {
		if filepath.Ext(f.Name()) != ".go" {
			return nil
		}
		if strings.HasSuffix(f.Name(), "_test.go") {
			return nil
		}
		pf, err := gogen.ParseFileMode(f, parser.ParseComments)
		if err != nil {
			return err
		}
		pkg := pf.Ast.Name
		for _, d := range pf.Ast.Decls {
			ast.Inspect(d, func(n ast.Node) bool {
				fn, ok := n.(*ast.FuncDecl)
				if !ok {
					return false
				}
				if unicode.IsLower(rune(fn.Name.Name[0])) {
					return false
				}
				fname := fmt.Sprintf("%s#%s", pkg, fn.Name)
				h := helper{
					Name: fname,
				}
				if fn.Doc != nil {
					h.Doc = fn.Doc.Text()
				}
				helpers = append(helpers, h)
				return true
			})
		}
		return nil
	})
	return helpers, err
}
