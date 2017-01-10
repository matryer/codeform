package parser

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

// CREDIT
// Most of this was stolen from https://github.com/ernesto-jimenez/gogen/blob/master/importer/importer.go - Copyright (c) 2015 Ernesto Jiménez
// which has an MIT license

var gopathlistOnce sync.Once
var gopathlistCache []string

func gopathlist() []string {
	gopathlistOnce.Do(func() {
		gopathlistCache = filepath.SplitList(os.Getenv("GOPATH"))
	})
	return gopathlistCache
}

// smartImporter is an importer.Importer that looks in the usual places
// for packages.
type smartImporter struct {
	base     types.Importer
	packages map[string]*types.Package
}

func newSmartImporter(base types.Importer) *smartImporter {
	return &smartImporter{
		base:     base,
		packages: make(map[string]*types.Package),
	}
}

func (i *smartImporter) Import(path string) (*types.Package, error) {
	var err error
	if path == "" || path[0] == '.' {
		path, err = filepath.Abs(filepath.Clean(path))
		if err != nil {
			return nil, err
		}
		path = stripGopath(path)
	}
	if pkg, ok := i.packages[path]; ok {
		return pkg, nil
	}
	pkg, err := i.doimport(path)
	if err != nil {
		return nil, err
	}
	i.packages[path] = pkg
	return pkg, nil
}

func (i *smartImporter) doimport(p string) (*types.Package, error) {
	dir, err := lookupImport(p)
	if err != nil {
		return i.base.Import(p)
	}
	dirFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		return i.base.Import(p)
	}
	fset := token.NewFileSet()
	var files []*ast.File
	for _, fileInfo := range dirFiles {
		if fileInfo.IsDir() {
			continue
		}
		n := fileInfo.Name()
		if path.Ext(fileInfo.Name()) != ".go" {
			continue
		}
		// if i.skipTestFiles && strings.Contains(fileInfo.Name(), "_test.go") {
		// 	continue
		// }
		file := path.Join(dir, n)
		src, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		f, err := parser.ParseFile(fset, file, src, 0)
		if err != nil {
			return nil, err
		}
		files = append(files, f)
	}
	conf := types.Config{
		Importer: i,
	}
	pkg, err := conf.Check(p, fset, files, nil)
	if err != nil {
		return i.base.Import(p)
	}
	return pkg, nil
}

func lookupImport(p string) (string, error) {
	for _, gopath := range gopathlist() {
		absPath, err := filepath.Abs(path.Join(gopath, "src", p))
		if err != nil {
			return "", err
		}
		if dir, err := os.Stat(absPath); err == nil && dir.IsDir() {
			return absPath, nil
		}
	}
	return "", errors.New("not in GOPATH: " + p)
}

func stripGopath(p string) string {
	for _, gopath := range gopathlist() {
		p = strings.Replace(p, path.Join(gopath, "src")+"/", "", 1)
	}
	return p
}
