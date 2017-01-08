package parser_test

import (
	"strings"
	"testing"

	"github.com/matryer/codeform/parser"
	"github.com/matryer/codeform/source"
	"github.com/matryer/is"
)

func TestFile(t *testing.T) {
	is := is.New(t)
	code, err := parser.New(source.MustLocal("./testdata/files/file.go")).Parse()
	is.NoErr(err) // Parse
	is.Equal(len(code.Packages), 1)
	is.Equal(code.Packages[0].Name, "files")
	is.Equal(len(code.Packages[0].Structs), 1)
	is.Equal(len(code.Packages[0].Vars), 1)
}

func TestByteSource(t *testing.T) {
	is := is.New(t)
	src := source.Reader("./testdata/files/no-such-file.go", strings.NewReader(`package files

type Struct struct {
	Field string
}

var egg int = 1
`))
	code, err := parser.New(src).Parse()
	is.NoErr(err) // Parse
	is.Equal(len(code.Packages), 1)
	is.Equal(code.Packages[0].Name, "files")
	is.Equal(len(code.Packages[0].Structs), 1)
	is.Equal(len(code.Packages[0].Vars), 1)
}
