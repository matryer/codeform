package parser

import (
	"testing"

	"github.com/matryer/codeform/source"
	"github.com/matryer/is"
)

func TestStructs(t *testing.T) {
	is := is.New(t)
	code, err := New(source.MustLocal("./testdata/types")).Parse()
	is.NoErr(err) // Parse()
	is.OK(code != nil)
	is.Equal(len(code.Packages), 1) // should be one package
	pkg := code.Packages[0]

	is.Equal(len(pkg.Structs), 4)
	is.Equal(pkg.Structs[0].Name, "Struct1")
	is.Equal(pkg.Structs[1].Name, "Struct2")

	// fields
	is.Equal(len(pkg.Structs[0].Fields), 0)
	is.Equal(len(pkg.Structs[1].Fields), 3)
	is.Equal(pkg.Structs[1].Fields[0].Name, "Field1")
	is.Equal(pkg.Structs[1].Fields[0].Type.Name, "int")
	is.Equal(pkg.Structs[1].Fields[0].Tag, `json:"field_one"`)
	is.Equal(pkg.Structs[1].Fields[1].Name, "Field2")
	is.Equal(pkg.Structs[1].Fields[1].Type.Name, "bool")
	is.Equal(pkg.Structs[1].Fields[1].Tag, `something:"else"`)
	is.Equal(pkg.Structs[1].Fields[2].Name, "Field3")
	is.Equal(pkg.Structs[1].Fields[2].Type.Name, "string")
	is.Equal(pkg.Structs[1].Fields[2].Tag, ``)

	// methods
	is.Equal(len(pkg.Structs[1].Methods), 3)
	for _, method := range pkg.Structs[1].Methods {
		switch method.Name {
		case "Method1":
			is.Equal(len(method.Args), 2)
		case "Method2":
			is.Equal(len(method.Args), 4)
		case "Method3":
			is.Equal(len(method.Args), 1)
		}
	}

}
