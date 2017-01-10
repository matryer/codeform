package render_test

import (
	"bufio"
	"bytes"
	"html/template"
	"testing"

	"strings"

	"github.com/matryer/codeform/parser"
	"github.com/matryer/codeform/render"
	"github.com/matryer/codeform/source"
	"github.com/matryer/is"
)

func TestRender(t *testing.T) {
	is := is.New(t)
	code, err := parser.New(source.MustLocal("./testdata/types")).Parse()
	is.NoErr(err)
	tmpl, err := template.New("Name").Funcs(render.TemplateFuncs).Parse(testTemplate)
	is.NoErr(err)
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, code)
	is.NoErr(err)
	// check that each line of expected appears in the output
	actual := buf.String()
	s := bufio.NewScanner(strings.NewReader(expected))
	for s.Scan() {
		is.OK(strings.Contains(actual, s.Text()))
	}
}

var testTemplate = `{{ range .Packages }}package {{ .Name }}
{{- range .Interfaces }}
{{- $interface := . }}

type {{.Name}}Mock struct {
	{{- range .Methods }}
	{{ .Name }}Func func({{ .Args | ArgList }}) {{ .ReturnArgs | ArgListTypes }}
	{{- end }}
}

{{- range .Methods }}
func (m *{{$interface.Name}}Mock) {{.Name}}({{ .Args | ArgList }}) {{ .ReturnArgs | ArgListTypes }} {
	{{- if .ReturnArgs }}
	return m.{{.Name}}Func({{ .Args | ArgListNames }})
	{{- else }}
	m.{{.Name}}Func({{ .Args | ArgListNames }})
	{{- end }}
}
{{- end }}
{{- end }}
{{- end -}}`

var expected = `package pkgname

type Interface1Mock struct {
}

type Interface3Mock struct {
	TheMethodFunc func(arg1 string, arg2 string) error
}
func (m *Interface3Mock) TheMethod(arg1 string, arg2 string) error {
	return m.TheMethodFunc(arg1, arg2)
}

type PersonMock struct {
	GreetFunc func(name string) (string, error)
	ShakeHandFunc func(level int) error
	WhisperFunc func(messages ...string)
}
func (m *PersonMock) Greet(name string) (string, error) {
	return m.GreetFunc(name)
}
func (m *PersonMock) ShakeHand(level int) error {
	return m.ShakeHandFunc(level)
}
func (m *PersonMock) Whisper(messages ...string)  {
	m.WhisperFunc(messages...)
}`
