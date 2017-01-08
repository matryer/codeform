package query_test

import (
	"testing"

	"github.com/matryer/codeform/query"
	"github.com/matryer/is"
)

func TestQueryNothing(t *testing.T) {
	is := is.New(t)
	code, err := query.New().Run(*testcode)
	is.NoErr(err)
	is.Equal(len(code.Packages), 3)
}

func TestQueryName(t *testing.T) {
	is := is.New(t)
	code, err := query.New().Package("package1", "package2").Name("interface1", "struct5").Run(*testcode)
	is.NoErr(err)
	is.Equal(len(code.Packages), 2)
	is.Equal(len(code.Packages[0].Interfaces), 1)
	is.Equal(code.Packages[0].Interfaces[0].Name, "interface1")
	is.Equal(len(code.Packages[1].Structs), 1)
	is.Equal(code.Packages[1].Structs[0].Name, "struct5")
}

func TestQueryPackage(t *testing.T) {
	is := is.New(t)
	code, err := query.New().Package("package1").Interface("interface1", "interface2").Run(*testcode)
	is.NoErr(err)
	is.Equal(len(code.Packages), 1) // should not find interfaces in package3
	is.Equal(len(code.Packages[0].Interfaces), 2)
}

func TestQueryInterface(t *testing.T) {
	is := is.New(t)
	code, err := query.New().Package("package1").Interface("interface1", "interface2").Run(*testcode)
	is.NoErr(err)
	is.Equal(len(code.Packages), 1)
	is.Equal(len(code.Packages[0].Interfaces), 2)
	is.Equal(code.Packages[0].Interfaces[0].Name, "interface1")
	is.Equal(code.Packages[0].Interfaces[1].Name, "interface2")
}

func TestQueryStruct(t *testing.T) {
	is := is.New(t)
	code, err := query.New().Struct("struct2", "struct3").Run(*testcode)
	is.NoErr(err)
	is.Equal(len(code.Packages), 1)
	is.Equal(len(code.Packages[0].Structs), 2)
	is.Equal(code.Packages[0].Structs[0].Name, "struct2")
	is.Equal(code.Packages[0].Structs[1].Name, "struct3")
}

func TestQueryFunc(t *testing.T) {
	is := is.New(t)
	code, err := query.New().Func("func2", "func3").Run(*testcode)
	is.NoErr(err)
	is.Equal(len(code.Packages), 1)
	is.Equal(len(code.Packages[0].Funcs), 2)
	is.Equal(code.Packages[0].Funcs[0].Name, "func2")
	is.Equal(code.Packages[0].Funcs[1].Name, "func3")
}
