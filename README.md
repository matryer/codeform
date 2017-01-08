# codeform
Go code generation framework.

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
