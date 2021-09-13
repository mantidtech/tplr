package functions

import (
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

// TestAll provides unit test coverage for All()
func TestFunctionCount(t *testing.T) {
	fn := All(nil)
	assert.Len(t, fn, 59, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestCombineFunctionLists provides unit test coverage for CombineFunctionLists
func TestCombineFunctionLists(t *testing.T) {
	type Args struct {
		fnList []template.FuncMap
	}

	one := func() {}
	two := func() {}

	tests := []struct {
		Name string
		Args Args
		Want template.FuncMap
	}{
		{
			Name: "nil",
			Args: Args{
				fnList: nil,
			},
			Want: template.FuncMap{},
		},
		{
			Name: "empty",
			Args: Args{
				fnList: []template.FuncMap{},
			},
			Want: template.FuncMap{},
		},
		{
			Name: "one",
			Args: Args{
				fnList: []template.FuncMap{
					{
						"one": one,
					},
				},
			},
			Want: template.FuncMap{
				"one": one,
			},
		},
		{
			Name: "one-one",
			Args: Args{
				fnList: []template.FuncMap{
					{
						"one": one,
					},
					{
						"two": two,
					},
				},
			},
			Want: template.FuncMap{
				"one": one,
				"two": two,
			},
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			got := CombineFunctionLists(tt.Args.fnList...)
			for k := range got {
				_, exists := tt.Want[k]
				assert.Truef(t, exists, "expected to find key %s", k)
			}
			for k := range tt.Want {
				_, exists := got[k]
				assert.Truef(t, exists, "unexpected key %s", k)
			}
		})
	}
}
