<html>
	<head>
		<title>{{ .title }}</title>
	<head>
	<body>
		<ul>
			{{- range .items -}}
				{{- include "listItem" . | indent 3 -}}
			{{- end }}
		</ul>
	</body>
</html>
{{- nl -}}

{{- define "listItem" }}
<li>
	{{.}}
</li>
{{- end -}}
