package functions

import (
	"os"
	"text/template"
)

func ExampleAll() {
	htmlTemplate := `
<html>
	<head>
		<title>{{ .title }}</title>
	<head>
	<body>
		<ul>
		{{- range .items }}
			{{- include "listItem" . | tabIndent 3 -}}
		{{- end }}
		</ul>
	</body>
</html>

{{- define "listItem" }}
<li>
	{{.}}
</li>
{{- end -}}`

	data := map[string]any{
		"title": "My List Of Doom",
		"items": []string{
			"Acquire lair",
			"Acquire henchmen",
			"Acquire doomsday weapon",
		},
	}

	t := template.New("example")
	t.Funcs(All(t))
	t, _ = t.Parse(htmlTemplate)
	_ = t.ExecuteTemplate(os.Stdout, "example", data)
	// Output:
	// <html>
	// 	<head>
	// 		<title>My List Of Doom</title>
	// 	<head>
	// 	<body>
	// 		<ul>
	// 			<li>
	// 				Acquire lair
	// 			</li>
	// 			<li>
	// 				Acquire henchmen
	// 			</li>
	// 			<li>
	// 				Acquire doomsday weapon
	// 			</li>
	// 		</ul>
	// 	</body>
	// </html>
}
