{{- range $k, $v := . }}
{{ $k }}: The type of '{{ $v }}' is '{{ typeName $v }}'
{{- end}}
