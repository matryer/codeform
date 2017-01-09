{{- range .Packages }}
{{- range .Interfaces }}{{ $interface := . }}
{{- range .Methods }}
{{ $interface.Name }}.{{ .Name }}{{ . | Signature }}
{{- end }}
{{- end }}
{{- end }}
