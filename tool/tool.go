// Package tool does the work of rendering output from templates.
package tool

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/matryer/codeform/parser"
	"github.com/matryer/codeform/query"
	"github.com/matryer/codeform/render"
	"github.com/matryer/codeform/source"
	"github.com/pkg/errors"
)

// Job is a single job.
type Job struct {
	Code       *source.Source
	Template   *source.Source
	Names      []string
	Packages   []string
	Interfaces []string
	Structs    []string
	Funcs      []string
	Log        func(...interface{})
}

// Execute runs the job and writes the output to the specified io.Writer.
func (j *Job) Execute(w io.Writer) error {
	code, err := parser.New(j.Code).Parse()
	if err != nil {
		j.logf("failed to parse source: %s", err)
		return errors.Wrap(err, "parser")
	}
	q := query.New()
	q.Name(j.Names...)
	if len(j.Names) > 0 {
		j.logf("including: %s", strings.Join(j.Names, ", "))
	}
	q.Package(j.Packages...)
	if len(j.Packages) > 0 {
		j.logf("including packages: %s", strings.Join(j.Packages, ", "))
	}
	q.Interface(j.Interfaces...)
	if len(j.Interfaces) > 0 {
		j.logf("including interfaces: %s", strings.Join(j.Interfaces, ", "))
	}
	q.Struct(j.Structs...)
	if len(j.Structs) > 0 {
		j.logf("including structs: %s", strings.Join(j.Structs, ", "))
	}
	q.Func(j.Funcs...)
	if len(j.Funcs) > 0 {
		j.logf("including funcs: %s", strings.Join(j.Funcs, ", "))
	}
	code, err = q.Run(*code)
	if err != nil {
		j.logf("failed to apply query: %s", err)
		return errors.Wrap(err, "query")
	}
	tmplB, err := ioutil.ReadAll(j.Template)
	if err != nil {
		j.logf("failed to read template: %s", err)
		return errors.Wrap(err, "reading template")
	}
	templateName := filepath.Base(j.Template.Path)
	j.logf("executing template %s", templateName)
	tmpl, err := template.New(templateName).Funcs(render.TemplateFuncs).Parse(string(tmplB))
	if err != nil {
		j.logf("failed to parse template: %s", err)
		return errors.Wrap(err, "parsing template")
	}
	return tmpl.Execute(w, code)
}

func (j *Job) log(args ...interface{}) {
	if j.Log != nil {
		j.Log(args...)
	}
}
func (j *Job) logf(format string, args ...interface{}) {
	j.log(fmt.Sprintf(format, args...))
}
