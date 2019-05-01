package decls

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/gogen"
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
	err := filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			base := filepath.Base(p)
			for _, x := range []string{"testdata", "internal", "vendor", ".git", ".idea", "syscall", "unsafe", "gob"} {
				if base == x {
					return filepath.SkipDir
				}
			}
		}

		if filepath.Ext(p) != ".go" {
			return nil
		}

		for _, pre := range prefixes {
			if strings.HasPrefix(p, pre) {
				return nil
			}
		}
		for _, suf := range suffixes {
			if strings.HasSuffix(p, suf) {
				return nil
			}
		}

		of, err := os.Open(p)
		if err != nil {
			return err
		}
		defer of.Close()

		f := genny.NewFile(p, of)
		for _, x := range []string{"// +build ignore", "// +build race", "// +build msan"} {
			if strings.Contains(f.String(), x) {
				return nil
			}
		}

		pf, err := gogen.ParseFileMode(f, parser.ParseComments)
		if err != nil {
			return err
		}

		for _, x := range []string{"main", "encoding/gob", "syscall", "syscall/js"} {
			if pf.Ast.Name.String() == x {
				return nil
			}
		}

		pn := strings.TrimPrefix(p, root)
		pn = strings.TrimPrefix(pn, "/")
		pn = filepath.Dir(pn)
		// fmt.Println(pn)
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
		var dex []Decl
		dm := map[string]Decl{}
		for _, d := range v.Decls {
			func() {
				if d.Name == "New" {
					return
				}
				y, ok := dm[d.Name]
				if !ok {
					defer func() {
						dm[d.Name] = y
					}()
					y = Decl{
						Name: d.Name,
					}
				}
				y.Doc += d.Doc
			}()
		}

		for _, v := range dm {
			dex = append(dex, v)
		}
		v.Decls = dex

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
