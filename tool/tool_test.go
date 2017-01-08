package tool_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/matryer/codeform/source"
	"github.com/matryer/codeform/tool"
	"github.com/matryer/is"
)

func TestExecute(t *testing.T) {
	is := is.New(t)

	srcCode := source.Reader("source.go", strings.NewReader(`package something

type Inter1 interface {
	Inter1Method1(a, b int) error
	Inter1Method2(c, d int) error
}

type Inter2 interface {
	Inter2Method1(a, b int) error
	Inter2Method2(c, d int) error
}`))
	srcTmpl := source.Reader("template.tpl", strings.NewReader(
		`{{ range .Packages }}{{ range .Interfaces }}{{ .Name }} {{ end }}{{ end }}`,
	))

	j := tool.Job{
		Code:     srcCode,
		Template: srcTmpl,
	}

	var buf bytes.Buffer
	err := j.Execute(&buf)
	is.NoErr(err)
	is.Equal(buf.String(), `Inter1 Inter2 `)

}
