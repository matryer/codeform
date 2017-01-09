# Codeform Examples

## Create a struct implementation for each interface

```
{{ range .Packages }}package {{.Name}}
{{- range .Interfaces }}
{{ $interface := . }}
type My{{.Name}} struct{}
{{ range .Methods }}
func (m *My{{.Name}}) {{.Name}}{{ . | Signature}} {
	panic("TODO: implement")
}
{{ end }}
{{- end }}
{{ end }}
```