package pkgname

// Func1 is a function.
func Func1() {}

// Func2 is a function.
func Func2(a int) {}

// Func3 is a function.
func Func3(a int) int {
	return a
}

// Func4 is a function.
func Func4(a, b int) int {
	return a
}

// Func5 is a function.
func Func5(a, b int) (c int, d int) {
	return a, b
}

// Func6 is a function.
func Func6(int) error {
	return nil
}

// Func7 is a function with a viradic argument.
func Func7(names ...string) {}
