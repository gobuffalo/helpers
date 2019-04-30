package decls

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gobuffalo/gogen"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
)

type Pkg struct {
	Full  string
	Short string
	Decls []Decl
	Doc   string
}

type Decl struct {
	Name string
	Doc  string
}

var prefixes = []string{"genny", "helptest", "helpers/cmd", "hctx"}
var suffixes = []string{"_test.go"}

func FindDecls(root string) ([]Pkg, error) {
	var helpers []Pkg
	m := map[string]Pkg{}
	box := packr.Folder(root)
	err := box.Walk(func(p string, f packd.File) error {
		if filepath.Ext(f.Name()) != ".go" {
			return nil
		}

		for _, x := range []string{"testdata", "internal", "vendor"} {
			if strings.Contains(p, x) {
				return nil
			}
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
		pn := strings.TrimPrefix(p, root)
		pn = strings.TrimPrefix(pn, "/")
		pn = filepath.Dir(pn)
		if pn == "." {
			pn = ""
		}

		pkg := Pkg{
			Full:  pn,
			Short: filepath.Base(pn),
		}
		if g, ok := m[pn]; ok {
			pkg = g
		}

		pf, err := gogen.ParseFileMode(f, parser.ParseComments)
		if err != nil {
			return err
		}
		if pf.Ast.Doc != nil {
			pkg.Doc = comment(pf.Ast.Doc.Text())
		}

		for _, d := range pf.Ast.Decls {

			switch decl := d.(type) {
			case *ast.GenDecl:
				for _, spec := range decl.Specs {
					switch spec := spec.(type) {

					case *ast.ValueSpec:
						for _, id := range spec.Names {
							x := id.Obj.Decl
							if vs, ok := x.(*ast.ValueSpec); ok {

								if len(vs.Values) == 0 {
									continue
								}
								if _, ok := vs.Values[0].(*ast.SelectorExpr); ok {
									if !id.IsExported() {
										continue
									}
									dh := Decl{
										Name: id.Name,
									}

									if vs.Doc != nil {
										dh.Doc = comment(vs.Doc.Text())
									}
									pkg.Decls = append(pkg.Decls, dh)
								}

							}
						}
					}
				}
			case *ast.FuncDecl:
				if !decl.Name.IsExported() || decl.Recv != nil || len(decl.Name.String()) == 0 {
					continue
				}
				dh := Decl{
					Name: decl.Name.String(),
				}

				if decl.Doc != nil {
					dh.Doc = comment(decl.Doc.Text())
				}
				pkg.Decls = append(pkg.Decls, dh)
			}
		}
		m[pn] = pkg

		return nil
	})

	for _, v := range m {
		sort.Slice(v.Decls, func(a, b int) bool {
			return v.Decls[a].Name < v.Decls[b].Name
		})
		helpers = append(helpers, v)
	}

	sort.Slice(helpers, func(a, b int) bool {
		return helpers[a].Full < helpers[b].Full
	})
	return helpers, err
}

func comment(s string) string {
	var lines []string
	scan := bufio.NewScanner(strings.NewReader(s))
	for scan.Scan() {
		lines = append(lines, fmt.Sprintf("// %s", scan.Text()))
	}

	return strings.Join(lines, "\n")
}
