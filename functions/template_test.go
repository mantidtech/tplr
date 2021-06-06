package functions

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestTemplateFunctions provides unit test coverage for TemplateFunctions
func TestTemplateFunctions(t *testing.T) {
	fn := TemplateFunctions(nil)
	assert.Len(t, fn, 1, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestGenerateIncludeFn provides unit test coverage for GenerateIncludeFn()
func TestGenerateIncludeFn(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		template       string
		vars           map[string]interface{}
		want           string
		wantParseErr   bool
		wantExecuteErr bool
	}{
		{
			name: "simple",
			template: `
				{{- define "testMain" -}}
					[ {{- include "testInclude" . -}} ]
				{{- end -}}
			
				{{- define "testInclude" -}}
					included
				{{- end -}}
			`,
			want: "[included]",
		},
		{
			name: "infinite recursion",
			template: `
				{{- define "testMain" -}}
					{{- include "testInclude" . -}}
				{{- end -}}
			
				{{- define "testInclude" -}}
					{{- include "testInclude" . -}}
				{{- end -}}
			`,
			wantExecuteErr: true,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var err error

			tpl := template.New("")
			includeFn := GenerateIncludeFn(tpl)
			tpl.Funcs(template.FuncMap{"include": includeFn})
			tpl, err = tpl.Parse(tt.template)

			if tt.wantParseErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			var f bytes.Buffer
			err = tpl.ExecuteTemplate(&f, "testMain", tt.vars)
			if tt.wantExecuteErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			assert.Equal(t, tt.want, f.String())
		})
	}
}
