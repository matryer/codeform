// Package defaultsource provides a staple source from which templates
// may be written.
// Accessible via a source string of "default".
// This file should contain a mixture of all possible templatable
// features.
package defaultsource

// Interface1 is an interface.
type Interface1 interface {
	Method1()
}

// Interface2 is an interface.
type Interface2 interface {
	Method2A(int)
	Method2B(a, b int)
}

// Interface3 is an interface.
type Interface3 interface {
	Method3A() error
	Method3B(int) error
	Method3C(a, b int) (bool, error)
}

// Struct1 is a struct.
type Struct1 struct {
	Field1 int
}

// Struct2 is a struct.
type Struct2 struct {
	Field1 int
	Field2 bool
	Field3 string
}

// Struct3 is a struct.
type Struct3 struct {
	Field1 string
}

// Func1 is a function.
func Func1() {}

// Func2 is a function.
func Func2(a, b int) int {
	return a
}

// Func3 is a function.
func Func3(a, b int) (bool, error) {
	return true, nil
}
