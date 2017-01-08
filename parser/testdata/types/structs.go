package pkgname

type Struct1 struct{}

type Struct2 struct {
	Field1 int  `json:"field_one"`
	Field2 bool `something:"else"`
	Field3 string
}

func (s *Struct2) Method1(a, b int) error {
	return nil
}

func (Struct2) Method2(a, b, c, d int) error {
	return nil
}

func (*Struct2) Method3(a int) error {
	return nil
}

type Struct3 struct {
	other1 Struct1
	other2 Struct2
}
