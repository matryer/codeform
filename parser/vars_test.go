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
	is.Equal(len(code.Packages), 1) // should be one package
	pkg := code.Packages[0]

	is.Equal(len(pkg.Vars), 13)
	for _, v := range pkg.Vars {
		switch v.Name {
		case "var1":
			is.Equal(v.Type.Fullname, "int")
			is.Equal(v.Type.Name, "int")
		case "var2":
			is.Equal(v.Type.Fullname, "string")
			is.Equal(v.Type.Name, "string")
		case "var3":
			is.Equal(v.Type.Fullname, "bool")
			is.Equal(v.Type.Name, "bool")
		case "number":
			is.Equal(v.Type.Fullname, "int")
			is.Equal(v.Type.Name, "int")
		case "name":
			is.Equal(v.Type.Fullname, "string")
			is.Equal(v.Type.Name, "string")
		case "preset":
			is.Equal(v.Type.Fullname, "int")
			is.Equal(v.Type.Name, "int")
		case "channel":
			is.Equal(v.Type.Fullname, "chan []byte")
			is.Equal(v.Type.Name, "chan []byte")
		case "amap":
			is.Equal(v.Type.Fullname, "map[string]int")
			is.Equal(v.Type.Name, "map[string]int")
		case "customType":
			is.Equal(v.Type.Fullname, "StructInSameFile")
			is.Equal(v.Type.Name, "StructInSameFile")
		case "customTypePointer":
			is.Equal(v.Type.Fullname, "*Struct1")
			is.Equal(v.Type.Name, "*Struct1")
		case "externalTypePointer":
			is.Equal(v.Type.Fullname, "*otherpackage.ExternalStruct")
			is.Equal(v.Type.Name, "ExternalStruct")
		case "externalType":
			is.Equal(v.Type.Fullname, "otherpackage.ExternalStruct")
			is.Equal(v.Type.Name, "ExternalStruct")
		case "r":
			is.Equal(v.Type.Fullname, "io.Reader")
			is.Equal(v.Type.Name, "Reader")
		}
	}

}
