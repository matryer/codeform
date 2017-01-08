package parser

import (
	"testing"

	"github.com/matryer/codeform/source"
	"github.com/matryer/is"
)

func TestVars(t *testing.T) {
	is := is.New(t)
	code, err := New(source.MustLocal("./testdata/types")).Parse()
	is.NoErr(err) // Parse()
	is.OK(code != nil)
	is.Equal(len(code.Packages), 1)
	pkg := code.Packages[0]

	is.Equal(len(pkg.Vars), 8)
	for _, v := range pkg.Vars {
		switch v.Name {
		case "var1":
			is.Equal(v.Type.Name, "int")
		case "var2":
			is.Equal(v.Type.Name, "string")
		case "var3":
			is.Equal(v.Type.Name, "bool")
		case "number":
			is.Equal(v.Type.Name, "int")
		case "name":
			is.Equal(v.Type.Name, "string")
		case "preset":
			is.Equal(v.Type.Name, "int")
		case "channel":
			is.Equal(v.Type.Name, "chan []byte")
		case "amap":
			is.Equal(v.Type.Name, "map[string]int")
		}
	}

}
