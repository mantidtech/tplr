package functions

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestLogicFunctions provides unit test coverage for LogicFunctions
func TestLogicFunctions(t *testing.T) {
	fn := LogicFunctions()
	assert.Len(t, fn, 5, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestWhen provides unit test coverage for When()
func TestWhen(t *testing.T) {
	type Args struct {
		D interface{}
		S interface{}
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "empty",
			template: `{{ when .D .S }}`,
			args: Args{
				D: "x",
				S: "",
			},
			want: "",
		},
		{
			name:     "not empty",
			template: `{{ when .D .S }}`,
			args: Args{
				D: "x",
				S: "y",
			},
			want: "x",
		},
		{
			name:     "default also empty",
			template: `{{ when .D .S }}`,
			args: Args{
				D: "",
				S: "",
			},
			want: "",
		},
		{
			name:     "int, not empty",
			template: `{{ when .D .S }}`,
			args: Args{
				D: "x",
				S: 9,
			},
			want: "x",
		},
		{
			name:     "int, empty",
			template: `{{ when .D .S }}`,
			args: Args{
				D: "x",
				S: 0,
			},
			want: "",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got bytes.Buffer

			tpl := helperNewTemplate(t, tt.template)
			err := tpl.ExecuteTemplate(&got, testTemplateName, tt.args)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got.String())
		})
	}
}

// TestWhenEmpty provides unit test coverage for WhenEmpty()
func TestWhenEmpty(t *testing.T) {
	type Args struct {
		D interface{}
		S interface{}
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "empty",
			template: `{{ whenEmpty .D .S }}`,
			args: Args{
				D: "x",
				S: "",
			},
			want: "x",
		},
		{
			name:     "not empty",
			template: `{{ whenEmpty .D .S }}`,
			args: Args{
				D: "x",
				S: "y",
			},
			want: "y",
		},
		{
			name:     "default also empty",
			template: `{{ whenEmpty .D .S }}`,
			args: Args{
				D: "",
				S: "",
			},
			want: "",
		},
		{
			name:     "int, not empty",
			template: `{{ whenEmpty .D .S }}`,
			args: Args{
				D: "x",
				S: 9,
			},
			want: "9",
		},
		{
			name:     "int, empty",
			template: `{{ whenEmpty .D .S }}`,
			args: Args{
				D: "x",
				S: 0,
			},
			want: "x",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got bytes.Buffer

			tpl := helperNewTemplate(t, tt.template)
			err := tpl.ExecuteTemplate(&got, testTemplateName, tt.args)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got.String())
		})
	}
}

// TestIsZero provides unit test coverage for IsZero()
func TestIsZero(t *testing.T) {
	type Args struct {
		Val interface{}
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "nil",
			template: `{{ isZero .Val }}`,
			args: Args{
				Val: nil,
			},
			want: "true",
		},
		{
			name:     "bool true",
			template: `{{ isZero .Val }}`,
			args: Args{
				Val: true,
			},
			want: "false",
		},
		{
			name:     "bool false",
			template: `{{ isZero .Val }}`,
			args: Args{
				Val: false,
			},
			want: "true",
		},
		{
			name:     "zero int",
			template: `{{ isZero .Val }}`,
			args: Args{
				Val: 0,
			},
			want: "true",
		},
		{
			name:     "non-zero int",
			template: `{{ isZero .Val }}`,
			args: Args{
				Val: 10,
			},
			want: "false",
		},
		{
			name:     "pointer zero int",
			template: `{{ isZero .Val }}`,
			args: Args{
				Val: helperPtrToInt(0),
			},
			want: "false",
		},
		{
			name:     "pointer non-zero int",
			template: `{{ isZero .Val }}`,
			args: Args{
				Val: helperPtrToInt(-82),
			},
			want: "false",
		},
		{
			name:     "non-zero int",
			template: `{{ isZero .Val }}`,
			args: Args{
				Val: 10,
			},
			want: "false",
		},
		{
			name:     "non-zero int",
			template: `{{ isZero .Val }}`,
			args: Args{
				Val: 10,
			},
			want: "false",
		},
		{
			name:     "empty string",
			template: `{{ isZero .Val }}`,
			args: Args{
				Val: "",
			},
			want: "true",
		},
		{
			name:     "non-empty string",
			template: `{{ isZero .Val }}`,
			args: Args{
				Val: "foo",
			},
			want: "false",
		},
		{
			name:     "empty array",
			template: `{{ isZero .Val }}`,
			args: Args{
				Val: []int{},
			},
			want: "true",
		},
		{
			name:     "nil array",
			template: `{{ isZero .Val }}`,
			args: Args{
				Val: []float64(nil),
			},
			want: "true",
		},
		{
			name:     "non-empty array",
			template: `{{ isZero .Val }}`,
			args: Args{
				Val: []string{"bar"},
			},
			want: "false",
		},
		{
			name:     "less simple & true",
			template: `{{- if isZero .Val -}}one{{- else -}}two{{- end -}}`,
			args: Args{
				Val: 0,
			},
			want: "one",
		},
		{
			name:     "less simple & false",
			template: `{{- if isZero .Val -}}one{{- else -}}two{{- end -}}`,
			args: Args{
				Val: 2,
			},
			want: "two",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got bytes.Buffer

			tpl := helperNewTemplate(t, tt.template)
			err := tpl.ExecuteTemplate(&got, testTemplateName, tt.args)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got.String())
		})
	}
}

// TestAnd provides unit test coverage for And()
func TestAnd(t *testing.T) {
	type Args struct {
		A interface{}
		B interface{}
		C interface{}
		D interface{}
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "no args",
			template: `{{ and }}`,
			args:     Args{},
			want:     "",
		},
		{
			name:     "bool false",
			template: `{{ and .A }}`,
			args: Args{
				A: false,
			},
			want: "",
		},
		{
			name:     "bool true",
			template: `{{ and .A }}`,
			args: Args{
				A: true,
			},
			want: "true",
		},
		{
			name:     "int 9",
			template: `{{ and .A }}`,
			args: Args{
				A: 9,
			},
			want: "9",
		},
		{
			name:     "2 empty args",
			template: `{{ and .A .B }}`,
			args: Args{
				A: "",
				B: 0,
			},
			want: "",
		},
		{
			name:     "3 true args",
			template: `{{ and .A .B .C }}`,
			args: Args{
				A: 2,
				B: 3,
				C: "X",
			},
			want: "X",
		},
		{
			name:     "3 args, 1 false",
			template: `{{ and .A .B .C }}`,
			args: Args{
				A: 2,
				B: 0,
				C: "X",
			},
			want: "",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got bytes.Buffer

			tpl := helperNewTemplate(t, tt.template)
			err := tpl.ExecuteTemplate(&got, testTemplateName, tt.args)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got.String())
		})
	}
}

// TestOr provides unit test coverage for Or()
func TestOr(t *testing.T) {
	type Args struct {
		A interface{}
		B interface{}
		C interface{}
		D interface{}
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "no args",
			template: `{{ or }}`,
			args:     Args{},
			want:     "",
		},
		{
			name:     "bool false",
			template: `{{ or .A }}`,
			args: Args{
				A: false,
			},
			want: "",
		},
		{
			name:     "bool true",
			template: `{{ or .A }}`,
			args: Args{
				A: true,
			},
			want: "true",
		},
		{
			name:     "int 9",
			template: `{{ or .A }}`,
			args: Args{
				A: 9,
			},
			want: "9",
		},
		{
			name:     "2 empty args",
			template: `{{ or .A .B }}`,
			args: Args{
				A: "",
				B: 0,
			},
			want: "",
		},
		{
			name:     "3 true args",
			template: `{{ or .A .B .C }}`,
			args: Args{
				A: 2,
				B: 3,
				C: "X",
			},
			want: "2",
		},
		{
			name:     "3 args, first false",
			template: `{{ or .A .B .C }}`,
			args: Args{
				A: 0,
				B: 2,
				C: "X",
			},
			want: "2",
		},
		{
			name:     "3 args, middle false",
			template: `{{ or .A .B .C }}`,
			args: Args{
				A: 2,
				B: 0,
				C: "X",
			},
			want: "2",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got bytes.Buffer

			tpl := helperNewTemplate(t, tt.template)
			err := tpl.ExecuteTemplate(&got, testTemplateName, tt.args)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got.String())
		})
	}
}
