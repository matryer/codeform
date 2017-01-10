package pkgname

import (
	"io"

	"github.com/matryer/codeform/parser/testdata/otherpackage"
)

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

var customType *Struct1

var externalType *otherpackage.ExternalStruct
var externalType2 otherpackage.ExternalStruct

var r io.Reader
