package docs

import (
	"fmt"
	"go/ast"
	"go/parser"
	"path"
	"path/filepath"
	"sort"
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

	helpers, err := findHelpers()
	if err != nil {
		return g, err
	}
	ctx.Set("helpers", helpers)
	g.Transformer(plushgen.Transformer(ctx))
	return g, nil
}

var prefixes = []string{"genny", "helptest", "helpers/cmd", "hctx"}
var suffixes = []string{"_test.go"}

func findHelpers() ([]helper, error) {
	var helpers []helper
	box := packr.Folder(".")
	err := box.Walk(func(p string, f packd.File) error {
		if filepath.Ext(f.Name()) != ".go" {
			return nil
		}

		for _, pre := range prefixes {
			if strings.HasPrefix(f.Name(), pre) {
				return nil
			}
		}
		for _, suf := range suffixes {
			if strings.HasSuffix(f.Name(), suf) {
				return nil
			}
		}
		pf, err := gogen.ParseFileMode(f, parser.ParseComments)
		if err != nil {
			return err
		}
		pkg := filepath.Dir(p)
		pkg = path.Join("github.com/gobuffalo/helpers", pkg)
		for _, d := range pf.Ast.Decls {
			ast.Inspect(d, func(n ast.Node) bool {
				fn, ok := n.(*ast.FuncDecl)
				if !ok {
					return false
				}
				if unicode.IsLower(rune(fn.Name.Name[0])) {
					return false
				}
				h := helper{
					Name:     fmt.Sprintf("%s#%s", pf.Ast.Name, fn.Name),
					FullName: fmt.Sprintf("%s#%s", pkg, fn.Name.String()),
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

	sort.Slice(helpers, func(a, b int) bool {
		return helpers[a].Name < helpers[b].Name
	})
	return helpers, err
}
