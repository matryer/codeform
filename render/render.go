// Package render provides utilities for rendering code using
// templates.
//
// Template functions
//
// To use the utility functions in the Go template packages, call Funcs
// passing in render.TemplateFuncs.
//
//     template.New("templatename").Funcs(render.TemplateFuncs)
package render

import (
	"fmt"
	"strings"

	"github.com/matryer/codeform/model"
)

// TemplateFuncs represents codeform utilities for templates.
var TemplateFuncs = map[string]interface{}{
	"ArgList":      ArgList,
	"ArgListNames": ArgListNames,
	"ArgListTypes": ArgListTypes,
	"Signature":    Signature,
}

// ArgList turns model.Args into a Go argument list.
//     name string, age int, ok bool
// Anonymous variables are given a name.
func ArgList(args model.Args) string {
	list := make([]string, len(args))
	for i, arg := range args {
		name := arg.Name
		if len(name) == 0 {
			name = fmt.Sprintf("arg%d", i+1)
		}
		typ := arg.Type.Name
		if arg.Variadic {
			typ = "..." + typ[2:]
		}
		list[i] = name + " " + typ
	}
	return strings.Join(list, ", ")
}

// ArgListNames turns model.Args into a comma separated list of names.
//     name, age, ok
// Anonymous variables are given a name.
func ArgListNames(args model.Args) string {
	list := make([]string, len(args))
	for i, arg := range args {
		name := arg.Name
		if len(name) == 0 {
			name = fmt.Sprintf("arg%d", i+1)
		}
		if arg.Variadic {
			name += "..."
		}
		list[i] = name
	}
	return strings.Join(list, ", ")
}

// ArgListTypes turns model.Args into a comma separated list of types.
//     (string, int, bool)
// Parentheses will be added for multiple arguments.
func ArgListTypes(args model.Args) string {
	list := make([]string, len(args))
	for i, arg := range args {
		typ := arg.Type.Name
		if arg.Variadic {
			typ = "..." + typ[2:]
		}
		list[i] = typ
	}
	if len(list) < 2 {
		return strings.Join(list, ", ")
	}
	return "(" + strings.Join(list, ", ") + ")"
}

// Signature gets the function arguments and return arguments
// as a string.
//     (name string, age int) (string, error)
func Signature(fn model.Func) string {
	sig := "(" + ArgList(fn.Args) + ")"
	if len(fn.ReturnArgs) > 0 {
		sig += " " + ArgListTypes(fn.ReturnArgs)
	}
	return sig
}
