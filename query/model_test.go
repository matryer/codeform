package query_test

import "github.com/matryer/codeform/model"

var interface1 = model.Interface{Name: "interface1"}
var interface2 = model.Interface{Name: "interface2"}
var interface3 = model.Interface{Name: "interface3"}
var struct1 = model.Struct{Name: "struct1"}
var struct2 = model.Struct{Name: "struct2"}
var struct3 = model.Struct{Name: "struct3"}
var struct4 = model.Struct{Name: "struct4"}
var struct5 = model.Struct{Name: "struct5"}
var struct6 = model.Struct{Name: "struct6"}
var func1 = model.Func{Name: "func1"}
var func2 = model.Func{Name: "func2"}
var func3 = model.Func{Name: "func3"}

var package1 = model.Package{
	Name: "package1",
	Interfaces: []model.Interface{
		interface1, interface2, interface3,
	},
	Structs: []model.Struct{
		struct1, struct2, struct3,
	},
	Funcs: []model.Func{
		func1, func2, func3,
	},
}
var package2 = model.Package{
	Name: "package2",
	Structs: []model.Struct{
		struct4, struct5, struct6,
	},
}
var package3 = model.Package{
	Name: "package3",
	Interfaces: []model.Interface{
		interface1, interface2, interface3,
	},
}

var testcode = &model.Code{
	Packages: []model.Package{
		package1, package2, package3,
	},
}
