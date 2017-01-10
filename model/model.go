// Package model contains the types that describe Go code.
package model

// Code represents the Go code as a data model.
type Code struct {
	// Packages represents a list of packages.
	Packages []Package `json:"packages"`
}

// Package represents a single Go package.
type Package struct {
	// Name is the name of the package.
	Name string `json:"name"`
	// Imports is a list of packages that need importing.
	Imports []string
	// Funcs represents the global functions.
	Funcs []Func `json:"funcs"`
	// Interfaces represents the global interfaces in the
	// package.
	Interfaces []Interface `json:"interfaces"`
	// Vars are the global variables defined in this
	// package.
	Vars []Var `json:"vars"`
	// Consts are the global constants defined in this
	// package.
	Consts []Var `json:"consts"`
	// Structs represents the global structs defined in this
	// package.
	Structs []Struct `json:"structs"`
}

// Struct represents a struct.
type Struct struct {
	// Name is the name of the struct.
	Name string `json:"name"`
	// Methods are the functions that have this struct as
	// its receiver.
	Methods []Func `json:"methods"`
	// Fields are the fields in this struct.
	Fields []Var `json:"fields"`
}

// Interface represents an interface.
type Interface struct {
	// Name is the name of the interface.
	Name string `json:"name"`
	// Methods is the methods that make up this
	// interface.
	Methods []Func `json:"methods"`
}

// Func represents a function.
type Func struct {
	// Name is the name of the function.
	Name string `json:"name"`
	// Args represents the input arguments.
	Args Args `json:"args"`
	// ReturnArgs represents the output arguments.
	ReturnArgs Args `json:"returnArgs"`
	// Variadic indicates whether the final input argument is
	// variadic or not.
	Variadic bool `json:"variadic"`
}

// Args represents a set of argument variables.
type Args []Var

// Var represents a variable.
type Var struct {
	// Name is the name of this variable.
	Name string `json:"name"`
	// Anonymous indicates whether the variable has
	// a name or not.
	Anonymous bool `json:"anonymous"`
	// Type is the type of this variable.
	Type Type `json:"type"`
	// Tag is the field tag for this variable or an empty
	// string if it doesn't have one.
	Tag string `json:"tag"`
	// Variadic indicates whether this argument is
	// variadic or not.
	Variadic bool `json:"variadic"`
}

// Type represents a type.
type Type struct {
	// Name is the name of the type.
	Name string `json:"name"`
}
