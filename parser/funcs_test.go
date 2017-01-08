package parser

import (
	"testing"

	"github.com/matryer/codeform/source"
	"github.com/matryer/is"
)

func TestPackages(t *testing.T) {
	is := is.New(t)
	code, err := New(source.MustLocal("./testdata/packages")).Parse()
	is.NoErr(err) // Parse()
	is.OK(code != nil)
	is.Equal(len(code.Packages), 2)
	for _, pkg := range code.Packages {
		switch pkg.Name {
		case "pkgname", "pkgname_test":
		default:
			is.Fail() // bad package name
		}
	}
}

func TestFuncs(t *testing.T) {
	is := is.New(t)
	code, err := New(source.MustLocal("./testdata/types")).Parse()
	is.NoErr(err) // Parse()
	is.OK(code != nil)
	is.Equal(len(code.Packages[0].Funcs), 7)

	is.Equal(code.Packages[0].Funcs[0].Name, "Func1")
	is.Equal(len(code.Packages[0].Funcs[0].Args), 0)

	is.Equal(code.Packages[0].Funcs[1].Name, "Func2")
	is.Equal(len(code.Packages[0].Funcs[1].Args), 1)
	is.Equal(code.Packages[0].Funcs[1].Args[0].Anonymous, false)
	is.Equal(code.Packages[0].Funcs[1].Args[0].Name, "a")
	is.Equal(code.Packages[0].Funcs[1].Args[0].Type.Name, "int")

	is.Equal(code.Packages[0].Funcs[2].Name, "Func3")
	is.Equal(len(code.Packages[0].Funcs[2].Args), 1)
	is.Equal(len(code.Packages[0].Funcs[2].ReturnArgs), 1)
	is.Equal(code.Packages[0].Funcs[2].ReturnArgs[0].Anonymous, true)
	is.Equal(code.Packages[0].Funcs[2].ReturnArgs[0].Name, "")
	is.Equal(code.Packages[0].Funcs[2].ReturnArgs[0].Type.Name, "int")

	is.Equal(code.Packages[0].Funcs[3].Name, "Func4")
	is.Equal(len(code.Packages[0].Funcs[3].Args), 2)

	is.Equal(code.Packages[0].Funcs[4].Name, "Func5")
	is.Equal(len(code.Packages[0].Funcs[4].Args), 2)
	is.Equal(len(code.Packages[0].Funcs[4].ReturnArgs), 2)
	is.Equal(code.Packages[0].Funcs[4].ReturnArgs[0].Anonymous, false)
	is.Equal(code.Packages[0].Funcs[4].ReturnArgs[0].Name, "c")
	is.Equal(code.Packages[0].Funcs[4].ReturnArgs[0].Type.Name, "int")
	is.Equal(code.Packages[0].Funcs[4].ReturnArgs[1].Anonymous, false)
	is.Equal(code.Packages[0].Funcs[4].ReturnArgs[1].Name, "d")
	is.Equal(code.Packages[0].Funcs[4].ReturnArgs[1].Type.Name, "int")

	is.Equal(code.Packages[0].Funcs[5].Name, "Func6")
	is.Equal(code.Packages[0].Funcs[5].Variadic, false)
	is.Equal(len(code.Packages[0].Funcs[5].Args), 1)
	is.Equal(code.Packages[0].Funcs[5].Args[0].Anonymous, true)
	is.Equal(code.Packages[0].Funcs[5].Args[0].Name, "")
	is.Equal(code.Packages[0].Funcs[5].Args[0].Type.Name, "int")

	is.Equal(code.Packages[0].Funcs[6].Name, "Func7")
	is.Equal(code.Packages[0].Funcs[6].Variadic, true)
	is.Equal(len(code.Packages[0].Funcs[6].Args), 1)
	is.Equal(code.Packages[0].Funcs[6].Args[0].Anonymous, false)
	is.Equal(code.Packages[0].Funcs[6].Args[0].Name, "names")
	is.Equal(code.Packages[0].Funcs[6].Args[0].Variadic, true)
	is.Equal(code.Packages[0].Funcs[6].Args[0].Type.Name, "[]string")

}
