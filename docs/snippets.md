# Snippets

## Create a struct implementation for each interface

{% raw %}
```liquid
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
{% endraw %}

## Package imports

To import packages mentioned in the source code, use the `.Imports` field on
the package:

{% raw %}
```liquid
{{ range .Packages }}package {{.Name}}
import (
	"github.com/explicit/import1"
	"github.com/explicit/import2"
{{- range .Imports }}
	"{{ .Name }}"
{{ end -}}
)
{{ end }}
```
{% endraw %}
