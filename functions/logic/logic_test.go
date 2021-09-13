package logic

import (
	"testing"

	"github.com/mantidtech/tplr/functions/helper"
	"github.com/stretchr/testify/assert"
)

// TestLogicFunctions provides unit test coverage for LogicFunctions
func TestLogicFunctions(t *testing.T) {
	fn := Functions()
	assert.Len(t, fn, 5, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestWhen provides unit test coverage for When()
func TestWhen(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "empty",
			Template: `{{ when .return .notZero }}`,
			Args: helper.TestArgs{
				"return":  "x",
				"notZero": "",
			},
			Want: "",
		},
		{
			Name:     "not empty",
			Template: `{{ when .return .notZero }}`,
			Args: helper.TestArgs{
				"return":  "x",
				"notZero": "y",
			},
			Want: "x",
		},
		{
			Name:     "default also empty",
			Template: `{{ when .return .notZero }}`,
			Args: helper.TestArgs{
				"return":  "",
				"notZero": "",
			},
			Want: "",
		},
		{
			Name:     "int, not empty",
			Template: `{{ when .return .notZero }}`,
			Args: helper.TestArgs{
				"return":  "x",
				"notZero": 9,
			},
			Want: "x",
		},
		{
			Name:     "int, empty",
			Template: `{{ when .return .notZero }}`,
			Args: helper.TestArgs{
				"return":  "x",
				"notZero": 0,
			},
			Want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestWhenEmpty provides unit test coverage for WhenEmpty()
func TestWhenEmpty(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "empty",
			Template: `{{ whenEmpty .else .if }}`,
			Args: helper.TestArgs{
				"else": "x",
				"if":   "",
			},
			Want: "x",
		},
		{
			Name:     "not empty",
			Template: `{{ whenEmpty .else .if }}`,
			Args: helper.TestArgs{
				"else": "x",
				"if":   "y",
			},
			Want: "y",
		},
		{
			Name:     "default also empty",
			Template: `{{ whenEmpty .else .if }}`,
			Args: helper.TestArgs{
				"else": "",
				"if":   "",
			},
			Want: "",
		},
		{
			Name:     "int, not empty",
			Template: `{{ whenEmpty .else .if }}`,
			Args: helper.TestArgs{
				"else": "x",
				"if":   9,
			},
			Want: "9",
		},
		{
			Name:     "int, empty",
			Template: `{{ whenEmpty .else .if }}`,
			Args: helper.TestArgs{
				"else": "x",
				"if":   0,
			},
			Want: "x",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestIsZero provides unit test coverage for IsZero()
func TestIsZero(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "nil",
			Template: `{{ isZero .value }}`,
			Args: helper.TestArgs{
				"value": nil,
			},
			Want: "true",
		},
		{
			Name:     "bool true",
			Template: `{{ isZero .value }}`,
			Args: helper.TestArgs{
				"value": true,
			},
			Want: "false",
		},
		{
			Name:     "bool false",
			Template: `{{ isZero .value }}`,
			Args: helper.TestArgs{
				"value": false,
			},
			Want: "true",
		},
		{
			Name:     "zero int",
			Template: `{{ isZero .value }}`,
			Args: helper.TestArgs{
				"value": 0,
			},
			Want: "true",
		},
		{
			Name:     "non-zero int",
			Template: `{{ isZero .value }}`,
			Args: helper.TestArgs{
				"value": 10,
			},
			Want: "false",
		},
		{
			Name:     "pointer zero int",
			Template: `{{ isZero .value }}`,
			Args: helper.TestArgs{
				"value": helper.PtrToInt(0),
			},
			Want: "false",
		},
		{
			Name:     "pointer non-zero int",
			Template: `{{ isZero .value }}`,
			Args: helper.TestArgs{
				"value": helper.PtrToInt(-82),
			},
			Want: "false",
		},
		{
			Name:     "non-zero int",
			Template: `{{ isZero .value }}`,
			Args: helper.TestArgs{
				"value": 10,
			},
			Want: "false",
		},
		{
			Name:     "non-zero int",
			Template: `{{ isZero .value }}`,
			Args: helper.TestArgs{
				"value": 10,
			},
			Want: "false",
		},
		{
			Name:     "empty string",
			Template: `{{ isZero .value }}`,
			Args: helper.TestArgs{
				"value": "",
			},
			Want: "true",
		},
		{
			Name:     "non-empty string",
			Template: `{{ isZero .value }}`,
			Args: helper.TestArgs{
				"value": "foo",
			},
			Want: "false",
		},
		{
			Name:     "empty array",
			Template: `{{ isZero .value }}`,
			Args: helper.TestArgs{
				"value": []int{},
			},
			Want: "true",
		},
		{
			Name:     "nil array",
			Template: `{{ isZero .value }}`,
			Args: helper.TestArgs{
				"value": []float64(nil),
			},
			Want: "true",
		},
		{
			Name:     "non-empty array",
			Template: `{{ isZero .value }}`,
			Args: helper.TestArgs{
				"value": []string{"bar"},
			},
			Want: "false",
		},
		{
			Name:     "less simple & true",
			Template: `{{- if isZero .value -}}one{{- else -}}two{{- end -}}`,
			Args: helper.TestArgs{
				"value": 0,
			},
			Want: "one",
		},
		{
			Name:     "less simple & false",
			Template: `{{- if isZero .value -}}one{{- else -}}two{{- end -}}`,
			Args: helper.TestArgs{
				"value": 2,
			},
			Want: "two",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestAnd provides unit test coverage for And()
func TestAnd(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "no args",
			Template: `{{ and }}`,
			Args:     helper.TestArgs{},
			Want:     "",
		},
		{
			Name:     "bool false",
			Template: `{{ and .A }}`,
			Args: helper.TestArgs{
				"A": false,
			},
			Want: "",
		},
		{
			Name:     "bool true",
			Template: `{{ and .A }}`,
			Args: helper.TestArgs{
				"A": true,
			},
			Want: "true",
		},
		{
			Name:     "int 9",
			Template: `{{ and .A }}`,
			Args: helper.TestArgs{
				"A": 9,
			},
			Want: "9",
		},
		{
			Name:     "2 empty args",
			Template: `{{ and .A .B }}`,
			Args: helper.TestArgs{
				"A": "",
				"B": 0,
			},
			Want: "",
		},
		{
			Name:     "3 true args",
			Template: `{{ and .A .B .C }}`,
			Args: helper.TestArgs{
				"A": 2,
				"B": 3,
				"C": "X",
			},
			Want: "X",
		},
		{
			Name:     "3 args, 1 false",
			Template: `{{ and .A .B .C }}`,
			Args: helper.TestArgs{
				"A": 2,
				"B": 0,
				"C": "X",
			},
			Want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestOr provides unit test coverage for Or()
func TestOr(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "no args",
			Template: `{{ or }}`,
			Args:     helper.TestArgs{},
			Want:     "",
		},
		{
			Name:     "bool false",
			Template: `{{ or .A }}`,
			Args: helper.TestArgs{
				"A": false,
			},
			Want: "",
		},
		{
			Name:     "bool true",
			Template: `{{ or .A }}`,
			Args: helper.TestArgs{
				"A": true,
			},
			Want: "true",
		},
		{
			Name:     "int 9",
			Template: `{{ or .A }}`,
			Args: helper.TestArgs{
				"A": 9,
			},
			Want: "9",
		},
		{
			Name:     "2 empty args",
			Template: `{{ or .A .B }}`,
			Args: helper.TestArgs{
				"A": "",
				"B": 0,
			},
			Want: "",
		},
		{
			Name:     "3 true args",
			Template: `{{ or .A .B .C }}`,
			Args: helper.TestArgs{
				"A": 2,
				"B": 3,
				"C": "X",
			},
			Want: "2",
		},
		{
			Name:     "3 args, first false",
			Template: `{{ or .A .B .C }}`,
			Args: helper.TestArgs{
				"A": 0,
				"B": 2,
				"C": "X",
			},
			Want: "2",
		},
		{
			Name:     "3 args, middle false",
			Template: `{{ or .A .B .C }}`,
			Args: helper.TestArgs{
				"A": 2,
				"B": 0,
				"C": "X",
			},
			Want: "2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}
