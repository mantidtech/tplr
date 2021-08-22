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
	RunTemplateTest(t, []TestSet{
		{
			name:     "empty",
			template: `{{ when .D .S }}`,
			args: TestArgs{
				"D": "x",
				"S": "",
			},
			want: "",
		},
		{
			name:     "not empty",
			template: `{{ when .D .S }}`,
			args: TestArgs{
				"D": "x",
				"S": "y",
			},
			want: "x",
		},
		{
			name:     "default also empty",
			template: `{{ when .D .S }}`,
			args: TestArgs{
				"D": "",
				"S": "",
			},
			want: "",
		},
		{
			name:     "int, not empty",
			template: `{{ when .D .S }}`,
			args: TestArgs{
				"D": "x",
				"S": 9,
			},
			want: "x",
		},
		{
			name:     "int, empty",
			template: `{{ when .D .S }}`,
			args: TestArgs{
				"D": "x",
				"S": 0,
			},
			want: "",
		},
	})
}

// TestWhenEmpty provides unit test coverage for WhenEmpty()
func TestWhenEmpty(t *testing.T) {
	RunTemplateTest(t, []TestSet{
		{
			name:     "empty",
			template: `{{ whenEmpty .D .S }}`,
			args: TestArgs{
				"D": "x",
				"S": "",
			},
			want: "x",
		},
		{
			name:     "not empty",
			template: `{{ whenEmpty .D .S }}`,
			args: TestArgs{
				"D": "x",
				"S": "y",
			},
			want: "y",
		},
		{
			name:     "default also empty",
			template: `{{ whenEmpty .D .S }}`,
			args: TestArgs{
				"D": "",
				"S": "",
			},
			want: "",
		},
		{
			name:     "int, not empty",
			template: `{{ whenEmpty .D .S }}`,
			args: TestArgs{
				"D": "x",
				"S": 9,
			},
			want: "9",
		},
		{
			name:     "int, empty",
			template: `{{ whenEmpty .D .S }}`,
			args: TestArgs{
				"D": "x",
				"S": 0,
			},
			want: "x",
		},
	})
}

// TestIsZero provides unit test coverage for IsZero()
func TestIsZero(t *testing.T) {
	RunTemplateTest(t, []TestSet{
		{
			name:     "nil",
			template: `{{ isZero .Val }}`,
			args: TestArgs{
				"Val": nil,
			},
			want: "true",
		},
		{
			name:     "bool true",
			template: `{{ isZero .Val }}`,
			args: TestArgs{
				"Val": true,
			},
			want: "false",
		},
		{
			name:     "bool false",
			template: `{{ isZero .Val }}`,
			args: TestArgs{
				"Val": false,
			},
			want: "true",
		},
		{
			name:     "zero int",
			template: `{{ isZero .Val }}`,
			args: TestArgs{
				"Val": 0,
			},
			want: "true",
		},
		{
			name:     "non-zero int",
			template: `{{ isZero .Val }}`,
			args: TestArgs{
				"Val": 10,
			},
			want: "false",
		},
		{
			name:     "pointer zero int",
			template: `{{ isZero .Val }}`,
			args: TestArgs{
				"Val": helperPtrToInt(0),
			},
			want: "false",
		},
		{
			name:     "pointer non-zero int",
			template: `{{ isZero .Val }}`,
			args: TestArgs{
				"Val": helperPtrToInt(-82),
			},
			want: "false",
		},
		{
			name:     "non-zero int",
			template: `{{ isZero .Val }}`,
			args: TestArgs{
				"Val": 10,
			},
			want: "false",
		},
		{
			name:     "non-zero int",
			template: `{{ isZero .Val }}`,
			args: TestArgs{
				"Val": 10,
			},
			want: "false",
		},
		{
			name:     "empty string",
			template: `{{ isZero .Val }}`,
			args: TestArgs{
				"Val": "",
			},
			want: "true",
		},
		{
			name:     "non-empty string",
			template: `{{ isZero .Val }}`,
			args: TestArgs{
				"Val": "foo",
			},
			want: "false",
		},
		{
			name:     "empty array",
			template: `{{ isZero .Val }}`,
			args: TestArgs{
				"Val": []int{},
			},
			want: "true",
		},
		{
			name:     "nil array",
			template: `{{ isZero .Val }}`,
			args: TestArgs{
				"Val": []float64(nil),
			},
			want: "true",
		},
		{
			name:     "non-empty array",
			template: `{{ isZero .Val }}`,
			args: TestArgs{
				"Val": []string{"bar"},
			},
			want: "false",
		},
		{
			name:     "less simple & true",
			template: `{{- if isZero .Val -}}one{{- else -}}two{{- end -}}`,
			args: TestArgs{
				"Val": 0,
			},
			want: "one",
		},
		{
			name:     "less simple & false",
			template: `{{- if isZero .Val -}}one{{- else -}}two{{- end -}}`,
			args: TestArgs{
				"Val": 2,
			},
			want: "two",
		},
	})
}

// TestAnd provides unit test coverage for And()
func TestAnd(t *testing.T) {
	RunTemplateTest(t, []TestSet{
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
	})
}

// TestOr provides unit test coverage for Or()
func TestOr(t *testing.T) {
	RunTemplateTest(t, []TestSet{
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
	})
}
