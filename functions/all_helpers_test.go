package functions

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testTemplateName = "test template"

type TestSet struct {
	name     string
	template string
	args     TestArgs
	want     string
	wantErr  bool
}

type TestArgs map[string]interface{}

func helperNewTemplate(t *testing.T, tpl string) *template.Template {
	var err error
	tSet := template.New(testTemplateName)
	tSet.Funcs(All(tSet))
	tSet, err = tSet.Parse(tpl)
	require.NoError(t, err)
	return tSet
}

func helperPtrToInt(i int) *int {
	r := new(int)
	*r = i
	return r
}

func TemplateTest(test TestSet) func(t *testing.T) {
	return func(t *testing.T) {
		t.Helper()
		var got bytes.Buffer

		tpl := helperNewTemplate(t, test.template)
		err := tpl.ExecuteTemplate(&got, testTemplateName, test.args)
		if test.wantErr {
			assert.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, test.want, got.String())
	}
}
