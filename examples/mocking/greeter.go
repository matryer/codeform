package mocking

//go:generate codeform -src . -out ./greeter_mock.go -templatesrc ./templates/mock.tpl -interface Greeter,Signoff

type Greeter interface {
	Greet(name string) (string, error)
	Reset()
}

type Signoff interface {
	Signoff(name string) string
}
