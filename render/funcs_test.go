package render_test

import (
	"testing"

	"github.com/matryer/codeform/model"
	"github.com/matryer/codeform/render"
	"github.com/matryer/is"
)

func TestCamel(t *testing.T) {
	is := is.New(t)
	is.Equal(render.Camel("SomeService"), "someService")
	is.Equal(render.Camel("someService"), "someService")
	is.Equal(render.Camel("A"), "a")
	is.Equal(render.Camel("a"), "a")
	is.Equal(render.Camel(""), "")
}

func TestArgList(t *testing.T) {
	is := is.New(t)
	args := model.Args{
		{Name: "name", Type: model.Type{Name: "string"}},
		{Name: "age", Type: model.Type{Name: "int"}},
		{Name: "ok", Type: model.Type{Name: "bool"}},
	}
	is.Equal(render.ArgList(args), `name string, age int, ok bool`)
	args = model.Args{
		{Name: "", Anonymous: true, Type: model.Type{Name: "string"}},
		{Name: "age", Type: model.Type{Name: "int"}},
		{Name: "ok", Type: model.Type{Name: "bool"}},
	}
	is.Equal(render.ArgList(args), `arg1 string, age int, ok bool`)
	is.Equal(render.ArgList(nil), ``)
}

func TestArgListTypes(t *testing.T) {
	is := is.New(t)
	args := model.Args{
		{Name: "err", Type: model.Type{Name: "error"}},
	}
	is.Equal(render.ArgListTypes(args), `error`)
	args = model.Args{
		{Name: "name", Type: model.Type{Name: "string"}},
		{Name: "age", Type: model.Type{Name: "int"}},
		{Name: "ok", Type: model.Type{Name: "bool"}},
	}
	is.Equal(render.ArgListTypes(args), `(string, int, bool)`)
	is.Equal(render.ArgListTypes(nil), ``)
}

func TestArgListNames(t *testing.T) {
	is := is.New(t)
	args := model.Args{
		{Name: "name", Type: model.Type{Name: "string"}},
		{Name: "age", Type: model.Type{Name: "int"}},
		{Name: "ok", Type: model.Type{Name: "bool"}},
	}
	is.Equal(render.ArgListNames(args), `name, age, ok`)
	args = model.Args{
		{Name: "", Anonymous: true, Type: model.Type{Name: "string"}},
		{Name: "age", Type: model.Type{Name: "int"}},
		{Name: "ok", Type: model.Type{Name: "bool"}},
	}
	is.Equal(render.ArgListNames(args), `arg1, age, ok`)
	is.Equal(render.ArgListNames(nil), ``)
}

func TestSignature(t *testing.T) {
	is := is.New(t)
	fn := model.Func{
		Args: model.Args{
			{Name: "name", Type: model.Type{Name: "string"}},
			{Name: "age", Type: model.Type{Name: "int"}},
			{Name: "ok", Type: model.Type{Name: "bool"}},
		},
		ReturnArgs: model.Args{
			{Name: "", Type: model.Type{Name: "string"}},
			{Name: "", Type: model.Type{Name: "error"}},
		},
	}
	is.Equal(render.Signature(fn), `(name string, age int, ok bool) (string, error)`)
}
