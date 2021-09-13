package helper

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestTemplateName is the template name used for testing
const TestTemplateName = "test template"

// NewTemplate creates a new template
func NewTemplate(t *testing.T, tpl string, fns template.FuncMap) *template.Template {
	var err error
	tSet := template.New(TestTemplateName)
	tSet.Funcs(fns)
	tSet, err = tSet.Parse(tpl)
	require.NoError(t, err)
	return tSet
}

// TemplateTest generates a test function for a given TestSet.
// fns is a set of functions needed by the template to be rendered
func TemplateTest(test TestSet, fns template.FuncMap) func(t *testing.T) {
	return func(t *testing.T) {
		t.Helper()
		var got bytes.Buffer

		tpl := NewTemplate(t, test.Template, fns)
		err := tpl.ExecuteTemplate(&got, TestTemplateName, test.Args)
		if test.WantErr {
			assert.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, test.Want, got.String())
	}
}
