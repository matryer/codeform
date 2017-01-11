package samepackage

import (
	"github.com/matryer/codeform/parser/testdata/otherpackage"
)

// Something refers to an external type.
func Something(val *otherpackage.ExternalStruct) {}
