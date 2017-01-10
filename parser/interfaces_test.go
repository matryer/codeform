package parser

import (
	"testing"

	"github.com/matryer/codeform/source"
	"github.com/matryer/is"
)

func TestInterfaces(t *testing.T) {
	is := is.New(t)
	code, err := New(source.MustLocal("./testdata/types")).Parse()
	is.NoErr(err) // Parse()
	is.OK(code != nil)
	is.Equal(len(code.Packages), 1) // should be one package
	pkg := code.Packages[0]

	is.Equal(len(pkg.Interfaces), 3)
	is.Equal(pkg.Interfaces[0].Name, "Interface1")
	is.Equal(len(pkg.Interfaces[0].Methods), 0)

	is.Equal(pkg.Interfaces[1].Name, "Interface2")
	is.Equal(len(pkg.Interfaces[1].Methods), 2)
	is.Equal(pkg.Interfaces[1].Methods[0].Name, "Method1")
	is.Equal(pkg.Interfaces[1].Methods[1].Name, "Method2")

	is.Equal(pkg.Interfaces[2].Name, "Interface3")
	is.Equal(len(pkg.Interfaces[2].Methods), 1)
	is.Equal(pkg.Interfaces[2].Methods[0].Name, "TheMethod")

}
