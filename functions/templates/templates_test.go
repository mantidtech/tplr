package templates

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/mantidtech/tplr/functions/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestTemplateFunctions provides unit test coverage for TemplateFunctions
func TestTemplateFunctions(t *testing.T) {
	fn := Functions(nil)
	assert.Len(t, fn, 2, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestGenerateIncludeFn provides unit test coverage for GenerateIncludeFn()
func TestGenerateIncludeFn(t *testing.T) {
	t.Parallel()
	tests := []struct {
		Name           string
		Template       string
		Vars           map[string]any
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

// TestList provides unit test coverage for List()
func TestGenerateIncludeApplyFn(t *testing.T) {
	tests := []struct {
		Name           string
		Template       string
		Vars           map[string]any
		Want           string
		WantParseErr   bool
		WantExecuteErr bool
	}{
		{
			Name: "simple",
			Template: `
				{{- define "testMain" -}}
					{{ applyInclude "q" .L }}
				{{- end -}}

				{{- define "q" -}}
					{{- printf "x%s" . -}}
				{{- end -}}
			`,
			Vars: helper.TestArgs{
				"L": []string{"a", "b"},
			},
			Want: "[xa xb]",
		},
		{
			Name: "empty",
			Template: `
				{{- define "testMain" -}}
					{{ applyInclude "q" .L }}
				{{- end -}}

				{{- define "q" -}}
					{{- printf "x%s" . -}}
				{{- end -}}
			`,
			Vars: helper.TestArgs{
				"L": []string{},
			},
			Want: "[]",
		},
		{
			Name: "nil",
			Template: `
				{{- define "testMain" -}}
					{{ applyInclude "q" .L }}
				{{- end -}}

				{{- define "q" -}}
					{{- printf "x%s" . -}}
				{{- end -}}
			`,
			Vars: helper.TestArgs{
				"L": []string(nil),
			},
			Want: "[]",
		},
		{
			Name: "simple",
			Template: `
				{{- define "testMain" -}}
					{{ applyInclude "q" .L }}
				{{- end -}}

				{{- define "q" -}}
					{{- adder . -}}
				{{- end -}}
			`,
			Vars: helper.TestArgs{
				"L": []string{"a", "b"},
			},
			WantExecuteErr: true,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			var err error

			tpl := template.New("")
			tpl.Funcs(template.FuncMap{
				"applyInclude": GenerateIncludeApplyFn(tpl),
				"adder":        func(i int) int { return i + i },
			})
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
