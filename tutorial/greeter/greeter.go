package greeter

type Greeter interface {
	Greet(name string) (string, error)
	Reset()
}

type Signoff interface {
	Signoff(name string) string
}
