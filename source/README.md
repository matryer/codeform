# Codeform sources

Sources are how Codeform reads the Go code (and templates) in order to generate output.

A source is a string that contains a reference to a valid local or online source.

## Valid sources

A source may be:

1. A local path to a file
1. A local path to a package (directory)
1. A `go get` style path to a package (e.g. `github.com/matryer/codeform`)
1. A `go get` style path to a nested package (e.g. `github.com/matryer/codeform/render/testdata/types`)
1. A `go get` style path to a file (e.g. `github.com/matryer/codeform/render/testdata/types/interfaces.go`)
1. A URL to a file (e.g. `https://raw.githubusercontent.com/matryer/codeform/master/render/testdata/types/interfaces.go`)
1. `"default"` string indicating the `default/source.go` file
1. `template:{path}` template from the [github.com/matryer/codeform-templates](https://github.com/matryer/codeform-templates) repository
