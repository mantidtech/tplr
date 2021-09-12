package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLogicFunctions provides unit test coverage for LogicFunctions
func TestLogicFunctions(t *testing.T) {
	fn := LogicFunctions()
	assert.Len(t, fn, 5, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestWhen provides unit test coverage for When()
func TestWhen(t *testing.T) {
	tests := []TestSet{
		{
			name:     "empty",
			template: `{{ when .return .notZero }}`,
			args: TestArgs{
				"return":  "x",
				"notZero": "",
			},
			want: "",
		},
		{
			name:     "not empty",
			template: `{{ when .return .notZero }}`,
			args: TestArgs{
				"return":  "x",
				"notZero": "y",
			},
			want: "x",
		},
		{
			name:     "default also empty",
			template: `{{ when .return .notZero }}`,
			args: TestArgs{
				"return":  "",
				"notZero": "",
			},
			want: "",
		},
		{
			name:     "int, not empty",
			template: `{{ when .return .notZero }}`,
			args: TestArgs{
				"return":  "x",
				"notZero": 9,
			},
			want: "x",
		},
		{
			name:     "int, empty",
			template: `{{ when .return .notZero }}`,
			args: TestArgs{
				"return":  "x",
				"notZero": 0,
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestWhenEmpty provides unit test coverage for WhenEmpty()
func TestWhenEmpty(t *testing.T) {
	tests := []TestSet{
		{
			name:     "empty",
			template: `{{ whenEmpty .else .if }}`,
			args: TestArgs{
				"else": "x",
				"if":   "",
			},
			want: "x",
		},
		{
			name:     "not empty",
			template: `{{ whenEmpty .else .if }}`,
			args: TestArgs{
				"else": "x",
				"if":   "y",
			},
			want: "y",
		},
		{
			name:     "default also empty",
			template: `{{ whenEmpty .else .if }}`,
			args: TestArgs{
				"else": "",
				"if":   "",
			},
			want: "",
		},
		{
			name:     "int, not empty",
			template: `{{ whenEmpty .else .if }}`,
			args: TestArgs{
				"else": "x",
				"if":   9,
			},
			want: "9",
		},
		{
			name:     "int, empty",
			template: `{{ whenEmpty .else .if }}`,
			args: TestArgs{
				"else": "x",
				"if":   0,
			},
			want: "x",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestIsZero provides unit test coverage for IsZero()
func TestIsZero(t *testing.T) {
	tests := []TestSet{
		{
			name:     "nil",
			template: `{{ isZero .value }}`,
			args: TestArgs{
				"value": nil,
			},
			want: "true",
		},
		{
			name:     "bool true",
			template: `{{ isZero .value }}`,
			args: TestArgs{
				"value": true,
			},
			want: "false",
		},
		{
			name:     "bool false",
			template: `{{ isZero .value }}`,
			args: TestArgs{
				"value": false,
			},
			want: "true",
		},
		{
			name:     "zero int",
			template: `{{ isZero .value }}`,
			args: TestArgs{
				"value": 0,
			},
			want: "true",
		},
		{
			name:     "non-zero int",
			template: `{{ isZero .value }}`,
			args: TestArgs{
				"value": 10,
			},
			want: "false",
		},
		{
			name:     "pointer zero int",
			template: `{{ isZero .value }}`,
			args: TestArgs{
				"value": helperPtrToInt(0),
			},
			want: "false",
		},
		{
			name:     "pointer non-zero int",
			template: `{{ isZero .value }}`,
			args: TestArgs{
				"value": helperPtrToInt(-82),
			},
			want: "false",
		},
		{
			name:     "non-zero int",
			template: `{{ isZero .value }}`,
			args: TestArgs{
				"value": 10,
			},
			want: "false",
		},
		{
			name:     "non-zero int",
			template: `{{ isZero .value }}`,
			args: TestArgs{
				"value": 10,
			},
			want: "false",
		},
		{
			name:     "empty string",
			template: `{{ isZero .value }}`,
			args: TestArgs{
				"value": "",
			},
			want: "true",
		},
		{
			name:     "non-empty string",
			template: `{{ isZero .value }}`,
			args: TestArgs{
				"value": "foo",
			},
			want: "false",
		},
		{
			name:     "empty array",
			template: `{{ isZero .value }}`,
			args: TestArgs{
				"value": []int{},
			},
			want: "true",
		},
		{
			name:     "nil array",
			template: `{{ isZero .value }}`,
			args: TestArgs{
				"value": []float64(nil),
			},
			want: "true",
		},
		{
			name:     "non-empty array",
			template: `{{ isZero .value }}`,
			args: TestArgs{
				"value": []string{"bar"},
			},
			want: "false",
		},
		{
			name:     "less simple & true",
			template: `{{- if isZero .value -}}one{{- else -}}two{{- end -}}`,
			args: TestArgs{
				"value": 0,
			},
			want: "one",
		},
		{
			name:     "less simple & false",
			template: `{{- if isZero .value -}}one{{- else -}}two{{- end -}}`,
			args: TestArgs{
				"value": 2,
			},
			want: "two",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestAnd provides unit test coverage for And()
func TestAnd(t *testing.T) {
	tests := []TestSet{
		{
			name:     "no args",
			template: `{{ and }}`,
			args:     TestArgs{},
			want:     "",
		},
		{
			name:     "bool false",
			template: `{{ and .A }}`,
			args: TestArgs{
				"A": false,
			},
			want: "",
		},
		{
			name:     "bool true",
			template: `{{ and .A }}`,
			args: TestArgs{
				"A": true,
			},
			want: "true",
		},
		{
			name:     "int 9",
			template: `{{ and .A }}`,
			args: TestArgs{
				"A": 9,
			},
			want: "9",
		},
		{
			name:     "2 empty args",
			template: `{{ and .A .B }}`,
			args: TestArgs{
				"A": "",
				"B": 0,
			},
			want: "",
		},
		{
			name:     "3 true args",
			template: `{{ and .A .B .C }}`,
			args: TestArgs{
				"A": 2,
				"B": 3,
				"C": "X",
			},
			want: "X",
		},
		{
			name:     "3 args, 1 false",
			template: `{{ and .A .B .C }}`,
			args: TestArgs{
				"A": 2,
				"B": 0,
				"C": "X",
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestOr provides unit test coverage for Or()
func TestOr(t *testing.T) {
	tests := []TestSet{
		{
			name:     "no args",
			template: `{{ or }}`,
			args:     TestArgs{},
			want:     "",
		},
		{
			name:     "bool false",
			template: `{{ or .A }}`,
			args: TestArgs{
				"A": false,
			},
			want: "",
		},
		{
			name:     "bool true",
			template: `{{ or .A }}`,
			args: TestArgs{
				"A": true,
			},
			want: "true",
		},
		{
			name:     "int 9",
			template: `{{ or .A }}`,
			args: TestArgs{
				"A": 9,
			},
			want: "9",
		},
		{
			name:     "2 empty args",
			template: `{{ or .A .B }}`,
			args: TestArgs{
				"A": "",
				"B": 0,
			},
			want: "",
		},
		{
			name:     "3 true args",
			template: `{{ or .A .B .C }}`,
			args: TestArgs{
				"A": 2,
				"B": 3,
				"C": "X",
			},
			want: "2",
		},
		{
			name:     "3 args, first false",
			template: `{{ or .A .B .C }}`,
			args: TestArgs{
				"A": 0,
				"B": 2,
				"C": "X",
			},
			want: "2",
		},
		{
			name:     "3 args, middle false",
			template: `{{ or .A .B .C }}`,
			args: TestArgs{
				"A": 2,
				"B": 0,
				"C": "X",
			},
			want: "2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}
