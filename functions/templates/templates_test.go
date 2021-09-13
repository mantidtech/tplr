package templates

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestTemplateFunctions provides unit test coverage for TemplateFunctions
func TestTemplateFunctions(t *testing.T) {
	fn := Functions(nil)
	assert.Len(t, fn, 1, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestGenerateIncludeFn provides unit test coverage for GenerateIncludeFn()
func TestGenerateIncludeFn(t *testing.T) {
	t.Parallel()
	tests := []struct {
		Name           string
		Template       string
		Vars           map[string]interface{}
		Want           string
		WantParseErr   bool
		WantExecuteErr bool
	}{
		{
			Name: "simple",
			Template: `
				{{- define "testMain" -}}
					[ {{- include "testInclude" . -}} ]
				{{- end -}}
			
				{{- define "testInclude" -}}
					included
				{{- end -}}
			`,
			Want: "[included]",
		},
		{
			Name: "infinite recursion",
			Template: `
				{{- define "testMain" -}}
					{{- include "testInclude" . -}}
				{{- end -}}
			
				{{- define "testInclude" -}}
					{{- include "testInclude" . -}}
				{{- end -}}
			`,
			WantExecuteErr: true,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			var err error

			tpl := template.New("")
			includeFn := GenerateIncludeFn(tpl)
			tpl.Funcs(template.FuncMap{"include": includeFn})
			tpl, err = tpl.Parse(tt.Template)

			if tt.WantParseErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			var f bytes.Buffer
			err = tpl.ExecuteTemplate(&f, "testMain", tt.Vars)
			if tt.WantExecuteErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			assert.Equal(t, tt.Want, f.String())
		})
	}
}
