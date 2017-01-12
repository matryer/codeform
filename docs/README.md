<iframe src="https://ghbtns.com/github-btn.html?user=matryer&repo=codeform&type=star&count=true" frameborder="0" scrolling="0" width="170px" height="20px"></iframe> <iframe src="https://ghbtns.com/github-btn.html?user=matryer&type=follow&count=true" frameborder="0" scrolling="0" width="170px" height="20px"></iframe>

---

# Welcome

Welcome to Codeform, easy Go code generation using templates. 

* Familiar [Go templates](https://golang.org/pkg/text/template/)
* [Simple data model](https://godoc.org/github.com/matryer/codeform/model)
* [Online editor](http://editor.codeform.in) with live preview

We are working towards a stable 1.0 release, but the project is entirely usable today.
Please get involved.

# How it works

Write [templates](https://golang.org/pkg/text/template/) using the 
[simple data model](https://godoc.org/github.com/matryer/codeform/model)
and use the `codeform` tool to generate code.

# Tutorials

To play with Codeform, head over to the [Online Editor](http://editor.codeform.in)
where you can modify the template, and see a live preview.

## Codeform command line tool

### Install Codeform

To install Codeform, ensure you have [Go installed](https://golang.org/dl/) and in a terminal:

```bash
go get github.com/matryer/codeform/...
```

This will install Codeform from source, including adding the `codeform` command to your `$GOBIN`
directory. 

### Ensure Codeform is installed

In a terminal, do:

```bash
codeform -version
```

You should see the latest Codeform version printed in a terminal. If you get an error,
ensure your `$GOBIN` is added to your `$PATH`.

### Find example code

Navigate to the `$GOPATH/src/github.com/matryer/codeform/tutorial/greeter` folder and look
at the `greeter.go` file:

```bash
cd $GOPATH/src/github.com/matryer/codeform/tutorial/greeter
cat greeter.go
```

The code contains two interfaces:

```go
package greeter

type Greeter interface {
	Greet(name string) (string, error)
	Reset()
}

type Signoff interface {
	Signoff(name string) string
}
```

### Generate code using hosted template

The [Codeform Templates repository](https://github.com/matryer/codeform-templates) contains some shared
templates that you are free to use. To access them, provide a template source beginning with the prefix
`template:` followed by the path to the template.

In a terminal, let's use the interface inspection template:

```bash
codeform -src greeter.go -templatesrc template:inspect/interfaces
```

You should get the following output:

```
greeter.Greeter
greeter.Signoff
```

### Providing our own template

We are going to provide a simple template that outputs the package name.

The template (expanded) will look like this:

{% raw %}
```liquid
{{ range .Packages }}
{{ .Name }}
{{ end }}
```
{% endraw %}

We can provide this template as a flag to the `codeform` command:

{% raw %}
```bash
codeform -src greeter.go -template "{{ range .Packages }}{{ .Name }}{{ end }}"
```
{% endraw %}

You should see the following output:

```
greeter
```

### Using a template file

Create a new file called `methods.tpl` and populate it with the following code:

{% raw %}
```liquid
{{- range .Packages }}
{{- range .Interfaces }}{{ $interface := . }}
{{- range .Methods }}
{{ $interface.Name }}.{{ .Name }}{{ . | Signature }}
{{- end }}
{{- end }}
{{- end }}
```
{% endraw %}

Now provide the file via the `templatesrc` flag:

```bash
codeform -src greeter.go -templatesrc /path/to/methods.tpl
```

You should see a list of interface methods:

```
Greeter.Greet(name string) (string, error)
Greeter.Reset()
Signoff.Signoff(name string) string
```

### What next?

Do `codeform -help` for a complete list of available flags.
