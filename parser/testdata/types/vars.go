package pkgname

import (
	"io"

	"github.com/matryer/codeform/parser/testdata/otherpackage"
)

type StructInSameFile struct {
	Field int
}

var number int

var name string

var (
	var1 int
	var2 string
	var3 bool
)

var preset int = 123

var channel chan []byte

var amap map[string]int

var customTypePointer *Struct1

var customType StructInSameFile

var externalTypePointer *otherpackage.ExternalStruct

var externalType otherpackage.ExternalStruct

var r io.Reader
