<iframe src="https://ghbtns.com/github-btn.html?user=matryer&repo=codeform&type=star&count=true" frameborder="0" scrolling="0" width="170px" height="20px"></iframe> <iframe src="https://ghbtns.com/github-btn.html?user=matryer&type=follow&count=true" frameborder="0" scrolling="0" width="170px" height="20px"></iframe>

---

Welcome to Codeform, easy Go code generation using templates. 

## Tutorial

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

```
&#123; range .Packages &#125;
&#123; .Name &#125;
&#123; end &#125;
```

We can provide this template as a flag to the `codeform` command:

```bash
codeform -src greeter.go -template "&#123; range .Packages &#125;&#123; .Name &#125;&#123; end &#125;"
```

You should see the following output:

```
greeter
```

### Using a template file

Create a new file called `methods.tpl` and populate it with the following code:

```
&#123;- range .Packages &#125;
&#123;- range .Interfaces &#125;&#123; $interface := . &#125;
&#123;- range .Methods &#125;
&#123; $interface.Name &#125;.&#123; .Name &#125;&#123; . | Signature &#125;
&#123;- end &#125;
&#123;- end &#125;
&#123;- end &#125;
```

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
