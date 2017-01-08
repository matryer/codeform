package parser

import (
	"testing"

	"github.com/matryer/codeform/source"
	"github.com/matryer/is"
)

func TestConsts(t *testing.T) {
	is := is.New(t)
	code, err := New(source.MustLocal("./testdata/types")).Parse()
	is.NoErr(err) // Parse()
	is.OK(code != nil)
	is.Equal(len(code.Packages), 1)
	pkg := code.Packages[0]

	is.Equal(len(pkg.Consts), 6)
	for _, v := range pkg.Consts {
		switch v.Name {
		case "pi":
			is.Equal(v.Type.Name, "float64")
		case "constantName":
			is.Equal(v.Type.Name, "string")
		case "c1":
			is.Equal(v.Type.Name, "int")
		case "c2":
			is.Equal(v.Type.Name, "int")
		case "c3":
			is.Equal(v.Type.Name, "int")
		case "c4":
			is.Equal(v.Type.Name, "int")
		}
	}

}
