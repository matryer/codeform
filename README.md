# codeform # silk [![Build Status](https://travis-ci.org/matryer/codeform.svg?branch=master)](https://travis-ci.org/matryer/codeform) [![Go Report Card](https://goreportcard.com/badge/github.com/matryer/codeform)](https://goreportcard.com/report/github.com/matryer/codeform)
Go code generation framework.

* Generate output from Go code using templates
* Interact with a simple model (see [Model documentation](https://godoc.org/github.com/matryer/codeform/model))
* Write templates using the [online editor](http://editor.codeform.in/) (with live preview)
* Contribute to a [repository of shared templates](https://github.com/matryer/codeform-templates)

## Get started

To install Codeform, get it:

```bash
go get github.com/matryer/codeform/...
```

This will install the `codeform` tool, as well as make the codeform
packages available to you.

## `codeform` comamnd line tool

The `codeform` tool allows you to generate output using Go templates.

```
-src string
	code source (default ".")
-srcin
	take source from standard in
-template string
	template verbatim
-templatesrc string
	template source
-out string
	output file (defaults to standard out)
	
-timeout duration
	timeout for HTTP requests (default 2s)
-v	verbose logging
-n  suppress final line feed
-version
	print version and exit

-func string
	comma separated list of funcs to include
-interface string
	comma separated list of interfaces to include
-name string
	comma separated list of things to include
-package string
	comma separated list of packages to include
-struct string
	comma separated list of structs to include
```

* See the [sources documentation](https://github.com/matryer/codeform/tree/master/source) for an overview of acceptable values for `src` and `templatesrc`

# Advanced

If you want to generate output using Go code instead of templates, you can
import the `github.com/matryer/codeform/parser` package directly and interact
with the model yourself (see [Model documentation](https://godoc.org/github.com/matryer/codeform/model)).
