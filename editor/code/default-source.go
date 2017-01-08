// Package defaultsource provides a staple source from which templates
// may be written.
// Accessible via a source string of "default".
// This file should contain a mixture of all possible templatable
// features.
package defaultsource

type Interface1 interface {
	Method1()
}

type Interface2 interface {
	Method2A(int)
	Method2B(a, b int)
}

type Interface3 interface {
	Method3A() error
	Method3B(int) error
	Method3C(a, b int) (bool, error)
}

type Struct1 struct {
	Field1 int
}

type Struct2 struct {
	Field1 int
	Field2 bool
	Field3 string
}

type Struct3 struct {
	Field1 string
}

func Func1() {}

func Func2(a, b int) int {
	return a
}

func Func3(a, b int) (bool, error) {
	return true, nil
}
