package pkgname

type Interface1 interface{}

type Interface2 interface {
	Method1()
	Method2()
}

type Interface3 interface {
	TheMethod(arg1, arg2 string) error
}
