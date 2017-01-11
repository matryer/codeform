package parser

import (
	"testing"

	"github.com/matryer/codeform/source"
	"github.com/matryer/is"
)

func TestDifferentPackage(t *testing.T) {
	is := is.New(t)
	p := New(source.MustLocal("./testdata/samepackage"))

	p.TargetPackage = ""
	code, err := p.Parse()
	is.NoErr(err) // Parse()
	is.OK(code != nil)
	is.Equal(len(code.Packages), 1)
	is.Equal(len(code.Packages[0].Funcs), 1)
	is.Equal(len(code.Packages[0].Funcs[0].Args), 1)
	is.Equal(code.Packages[0].Funcs[0].Args[0].Type.Name, "ExternalStruct")
	is.Equal(code.Packages[0].Funcs[0].Args[0].Type.Fullname, "*otherpackage.ExternalStruct")
	is.Equal(code.TargetPackageName, "samepackage")
	is.Equal(len(code.Packages[0].Imports), 1)
	is.Equal(code.Packages[0].Imports[0].Name, "github.com/matryer/codeform/parser/testdata/otherpackage")

	p.TargetPackage = "github.com/matryer/codeform/parser/testdata/samepackage"
	code, err = p.Parse()
	is.NoErr(err) // Parse()
	is.OK(code != nil)
	is.Equal(len(code.Packages), 1)
	is.Equal(len(code.Packages[0].Funcs), 1)
	is.Equal(len(code.Packages[0].Funcs[0].Args), 1)
	is.Equal(code.Packages[0].Funcs[0].Args[0].Type.Name, "ExternalStruct")
	is.Equal(code.Packages[0].Funcs[0].Args[0].Type.Fullname, "*otherpackage.ExternalStruct")
	is.Equal(code.TargetPackageName, "samepackage")
	is.Equal(len(code.Packages[0].Imports), 1)
	is.Equal(code.Packages[0].Imports[0].Name, "github.com/matryer/codeform/parser/testdata/otherpackage")

	p.TargetPackage = "github.com/matryer/codeform/parser/testdata/otherpackage"
	code, err = p.Parse()
	is.NoErr(err) // Parse()
	is.OK(code != nil)
	is.Equal(len(code.Packages), 1)
	is.Equal(len(code.Packages[0].Funcs), 1)
	is.Equal(len(code.Packages[0].Funcs[0].Args), 1)
	is.Equal(code.Packages[0].Funcs[0].Args[0].Type.Name, "*ExternalStruct")
	is.Equal(code.Packages[0].Funcs[0].Args[0].Type.Fullname, "*ExternalStruct")
	is.Equal(code.TargetPackageName, "otherpackage")
	is.Equal(len(code.Packages[0].Imports), 0)

	p.TargetPackage = "github.com/matryer/codeform/parser/testdata/otherpackage"
	code, err = p.Parse()
	is.NoErr(err) // Parse()
	is.OK(code != nil)
	is.Equal(len(code.Packages), 1)
	is.Equal(len(code.Packages[0].Funcs), 1)
	is.Equal(len(code.Packages[0].Funcs[0].Args), 1)
	is.Equal(code.Packages[0].Funcs[0].Args[0].Type.Name, "*ExternalStruct")
	is.Equal(code.Packages[0].Funcs[0].Args[0].Type.Fullname, "*ExternalStruct")
	is.Equal(code.TargetPackageName, "otherpackage")
	is.Equal(len(code.Packages[0].Imports), 0)

}
