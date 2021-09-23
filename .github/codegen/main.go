package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/doc"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"path"
)

var pkg = flag.String("pkg", "", "pkg")

func main() {
	flag.Parse()

	types, err := getPackageTypes(*pkg)
	if err != nil {
		panic(err)
	}
	fmt.Println(types)
}

func getPackageTypes(pkg string) ([]string, error) {
	pd, err := build.Import(pkg, ".", 0)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	files := make(map[string]*ast.File)
	fileList := make([]*ast.File, len(pd.GoFiles))
	for i, fname := range pd.GoFiles {
		src, err := ioutil.ReadFile(path.Join(pd.Dir, fname))
		if err != nil {
			panic(err)
		}
		f, err := parser.ParseFile(fset, fname, src, parser.AllErrors)
		if err != nil {
			panic(err)
		}
		files[fname] = f
		fileList[i] = f
	}

	cfg := types.Config{
		Importer: importer.For("source", nil),
	}
	info := types.Info{
		Defs: make(map[*ast.Ident]types.Object),
	}
	tp, err := cfg.Check(pkg, fset, fileList, &info)
	if err != nil {
		panic(err)
	}

	_ = tp.Scope()

	ap, _ := ast.NewPackage(fset, files, nil, nil)
	docs := doc.New(ap, pkg, 0)

	res := []string{}
	for _, v := range docs.Types {
		res = append(res, v.Name)
	}

	return res, nil
}
