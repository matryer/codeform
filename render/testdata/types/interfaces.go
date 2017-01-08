package pkgname

type Interface1 interface{}

type Person interface {
	Greet(name string) (string, error)
	ShakeHand(level int) error
	Whisper(messages ...string)
}

type Interface3 interface {
	TheMethod(arg1, arg2 string) error
}
