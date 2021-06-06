package functions

import (
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func helperPtrToInt(i int) *int {
	r := new(int)
	*r = i
	return r
}

const testTemplateName = "test template"

func helperNewTemplate(t *testing.T, tpl string) *template.Template {
	var err error
	tSet := template.New(testTemplateName)
	tSet.Funcs(All(tSet))
	tSet, err = tSet.Parse(tpl)
	require.NoError(t, err)
	return tSet
}

// TestAll provides unit test coverage for All()
func TestAll(t *testing.T) {
	fn := All(nil)
	assert.Len(t, fn, 56, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestCombineFunctionLists provides unit test coverage for CombineFunctionLists
func TestCombineFunctionLists(t *testing.T) {
	type Args struct {
		fnList []template.FuncMap
	}

	one := func() {}
	two := func() {}

	tests := []struct {
		name string
		args Args
		want template.FuncMap
	}{
		{
			name: "nil",
			args: Args{
				fnList: nil,
			},
			want: template.FuncMap{},
		},
		{
			name: "empty",
			args: Args{
				fnList: []template.FuncMap{},
			},
			want: template.FuncMap{},
		},
		{
			name: "one",
			args: Args{
				fnList: []template.FuncMap{
					{
						"one": one,
					},
				},
			},
			want: template.FuncMap{
				"one": one,
			},
		},
		{
			name: "one-one",
			args: Args{
				fnList: []template.FuncMap{
					{
						"one": one,
					},
					{
						"two": two,
					},
				},
			},
			want: template.FuncMap{
				"one": one,
				"two": two,
			},
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := CombineFunctionLists(tt.args.fnList...)
			for k := range got {
				_, exists := tt.want[k]
				assert.Truef(t, exists, "expected to find key %s", k)
			}
			for k := range tt.want {
				_, exists := got[k]
				assert.Truef(t, exists, "unexpected key %s", k)
			}
		})
	}
}
