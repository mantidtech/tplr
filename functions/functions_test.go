package functions

import (
	"bytes"
	"testing"
	"text/template"
	"time"

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
	assert.Len(t, fn, 55, "weakly ensuring functions haven't been added/removed without updating tests")
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

// TestUppercaseFirst provides unit test coverage for UppercaseFirst()
func TestUppercaseFirst(t *testing.T) {
	type Args struct {
		S string
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
			template: `{{ ucFirst .S }}`,
			args: Args{
				S: "",
			},
			want: "",
		},
		{
			name:     "simple",
			template: `{{ ucFirst .S }}`,
			args: Args{
				S: "simple",
			},
			want: "Simple",
		},
		{
			name:     "same",
			template: `{{ ucFirst .S }}`,
			args: Args{
				S: "Same",
			},
			want: "Same",
		},
		{
			name:     "number",
			template: `{{ ucFirst .S }}`,
			args: Args{
				S: "3rd",
			},
			want: "3rd",
		},
		{
			name:     "multiple words",
			template: `{{ ucFirst .S }}`,
			args: Args{
				S: "spam test",
			},
			want: "Spam test",
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

// TestNewline provides unit test coverage for Newline()
func TestNewline(t *testing.T) {
	type Args struct {
		C int
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "no params",
			template: `{{ nl }}`,
			args:     Args{},
			want:     "\n",
		},
		{
			name:     "params",
			template: `{{ nl .C }}`,
			args: Args{
				C: 3,
			},
			want: "\n\n\n",
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

// TestRep provides unit test coverage for Rep()
func TestRep(t *testing.T) {
	type Args struct {
		N  int
		S  string
		S2 string
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "none",
			template: `{{ rep .N .S }}`,
			args: Args{
				N: 0,
				S: "foo",
			},
			want: "",
		},
		{
			name:     "repeated empty",
			template: `{{ rep .N .S }}`,
			args: Args{
				N: 2,
				S: "",
			},
			want: "",
		},
		{
			name:     "one",
			template: `{{ rep .N .S }}`,
			args: Args{
				N: 1,
				S: "x",
			},
			want: "x",
		},
		{
			name:     "two",
			template: `{{ rep .N .S }}`,
			args: Args{
				N: 2,
				S: "foo",
			},
			want: "foofoo",
		},
		{
			name:     "multiple args",
			template: `{{ rep .N .S .S2 }}`,
			args: Args{
				N:  1,
				S:  "one",
				S2: "two",
			},
			want: "one two",
		},
		{
			name:     "multiple args, twice",
			template: `{{ rep .N .S .S2 }}`,
			args: Args{
				N:  2,
				S:  "one",
				S2: "two",
			},
			want: "one twoone two",
		},
		{
			name:     "negative one",
			template: `{{ rep .N .S }}`,
			args: Args{
				N: -1,
				S: "foo",
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

// TestIndent provides unit test coverage for Indent()
func TestIndent(t *testing.T) {
	type Args struct {
		T       int
		Content string
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "none",
			template: `{{ indent .T .Content }}`,
			args: Args{
				T:       0,
				Content: "foo",
			},
			want: "foo",
		},
		{
			name:     "one",
			template: `{{ indent .T .Content }}`,
			args: Args{
				T:       1,
				Content: "foo",
			},
			want: " foo",
		},
		{
			name:     "two",
			template: `{{ indent .T .Content }}`,
			args: Args{
				T:       2,
				Content: "foo",
			},
			want: "  foo",
		},
		{
			name:     "negative one",
			template: `{{ indent .T .Content }}`,
			args: Args{
				T:       -1,
				Content: "",
			},
			want: "",
		},
		{
			name:     "multi line",
			template: `{{ indent .T .Content }}`,
			args: Args{
				T:       1,
				Content: "foo\nbar",
			},
			want: " foo\n bar",
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

// TestSuffix provides unit test coverage for Suffix()
func TestSuffix(t *testing.T) {
	type Args struct {
		Suffix  string
		T       int
		Content string
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "none",
			template: `{{ suffix .Suffix .T .Content }}`,
			args: Args{
				Suffix:  "",
				T:       0,
				Content: "foo",
			},
			want: "foo",
		},
		{
			name:     "one",
			template: `{{ suffix .Suffix .T .Content }}`,
			args: Args{
				Suffix:  "X",
				T:       1,
				Content: "foo",
			},
			want: "fooX",
		},
		{
			name:     "two",
			template: `{{ suffix .Suffix .T .Content }}`,
			args: Args{
				Suffix:  "X",
				T:       2,
				Content: "foo",
			},
			want: "fooXX",
		},
		{
			name:     "negative one",
			template: `{{ suffix .Suffix .T .Content }}`,
			args: Args{
				Suffix:  "X",
				T:       -1,
				Content: "",
			},
			want: "",
		},
		{
			name:     "multi line",
			template: `{{ suffix .Suffix .T .Content }}`,
			args: Args{
				Suffix:  "X",
				T:       1,
				Content: "foo\nbar",
			},
			want: "fooX\nbarX",
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

// TestSpace provides unit test coverage for Space()
func TestSpace(t *testing.T) {
	type Args struct {
		N int
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "none",
			template: `{{ sp .N }}`,
			args: Args{
				N: 0,
			},
			want: "",
		},
		{
			name:     "one",
			template: `{{ sp .N }}`,
			args: Args{
				N: 1,
			},
			want: " ",
		},
		{
			name:     "two",
			template: `{{ sp .N }}`,
			args: Args{
				N: 2,
			},
			want: "  ",
		},
		{
			name:     "negative one",
			template: `{{ sp .N }}`,
			args: Args{
				N: -1,
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

// TestTab provides unit test coverage for Tab()
func TestTab(t *testing.T) {
	type Args struct {
		N int
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "none",
			template: `{{ tab .N }}`,
			args: Args{
				N: 0,
			},
			want: "",
		},
		{
			name:     "one",
			template: `{{ tab .N }}`,
			args: Args{
				N: 1,
			},
			want: "\t",
		},
		{
			name:     "two",
			template: `{{ tab .N }}`,
			args: Args{
				N: 2,
			},
			want: "\t\t",
		},
		{
			name:     "negative one",
			template: `{{ tab .N }}`,
			args: Args{
				N: -1,
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

// TestPadRight provides unit test coverage for PadRight()
func TestPadRight(t *testing.T) {
	type Args struct {
		N int
		S string
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "basic",
			template: `{{ padRight .N .S }}`,
			args: Args{
				N: 10,
				S: "basic",
			},
			want: "basic     ",
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

// TestPadLeft provides unit test coverage for PadLeft()
func TestPadLeft(t *testing.T) {
	type Args struct {
		N int
		S string
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "basic",
			template: `{{ padLeft .N .S }}`,
			args: Args{
				N: 10,
				S: "basic",
			},
			want: "     basic",
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

// TestNow shows that there's testing, and just keeping yourself amused
func TestNow(t *testing.T) {
	n := Now()
	_, err := time.Parse(time.RFC3339, n)
	assert.NoError(t, err)
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

// TestBracket provides unit test coverage for Bracket()
func TestBracket(t *testing.T) {
	type Args struct {
		S string
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
			template: `{{ bracket .S }}`,
			args: Args{
				S: "",
			},
			want: "()",
		},
		{
			name:     "word",
			template: `{{ bracket .S }}`,
			args: Args{
				S: "foo",
			},
			want: "(foo)",
		},
		{
			name:     "words",
			template: `{{ bracket .S }}`,
			args: Args{
				S: "foo bar",
			},
			want: "(foo bar)",
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

// TestBracketWith provides unit test coverage for BracketWith()
func TestBracketWith(t *testing.T) {
	type Args struct {
		B string
		S string
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "none",
			template: `{{ bracketWith .B .S }}`,
			args: Args{
				B: "",
				S: "",
			},
			want: "",
		},
		{
			name:     "basic",
			template: `{{ bracketWith .B .S }}`,
			args: Args{
				B: "()",
				S: "",
			},
			want: "()",
		},
		{
			name:     "word",
			template: `{{ bracketWith .B .S }}`,
			args: Args{
				B: "<>",
				S: "foo",
			},
			want: "<foo>",
		},
		{
			name:     "words",
			template: `{{ bracketWith .B .S }}`,
			args: Args{
				B: "{{-  -}}",
				S: "foo bar",
			},
			want: "{{- foo bar -}}",
		},
		{
			name:     "mismatched",
			template: `{{ bracketWith .B .S }}`,
			args: Args{
				B: ")",
				S: "baz",
			},
			wantErr: true,
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

// TestJoin provides unit test coverage for Join()
func TestJoin(t *testing.T) {
	type Args struct {
		A []interface{}
		B string
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
			template: `{{ join .A }}`,
			args: Args{
				A: nil,
			},
			want: "",
		},
		{
			name:     "empty",
			template: `{{ join .A }}`,
			args: Args{
				A: []interface{}{},
			},
			want: "",
		},
		{
			name:     "one",
			template: `{{ join .A }}`,
			args: Args{
				A: []interface{}{"one"},
			},
			want: "one",
		},
		{
			name:     "two",
			template: `{{ join .A }}`,
			args: Args{
				A: []interface{}{"one", "two"},
			},
			want: "onetwo",
		},
		{
			name:     "2",
			template: `{{ join .A }}`,
			args: Args{
				A: []interface{}{1, 2},
			},
			want: "12",
		},
		{
			name:     "bad list",
			template: `{{ join .B }}`,
			args: Args{
				B: "Fail",
			},
			wantErr: true,
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

// TestTypeName provides unit test coverage for TypeName()
func TestTypeName(t *testing.T) {
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
			template: `{{ typeName .Val }}`,
			args: Args{
				Val: nil,
			},
			want: "nil",
		},
		{
			name:     "int",
			template: `{{ typeName .Val }}`,
			args: Args{
				Val: 3,
			},
			want: "int",
		},
		{
			name:     "time.Duration",
			template: `{{ typeName .Val }}`,
			args: Args{
				Val: 10 * time.Second,
			},
			want: "time.Duration",
		},
		{
			name:     "*int",
			template: `{{ typeName .Val }}`,
			args: Args{
				Val: helperPtrToInt(10),
			},
			want: "*int",
		},
		{
			name:     "[]int",
			template: `{{ typeName .Val }}`,
			args: Args{
				Val: []int{4},
			},
			want: "[]int",
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

// TestTypeKind provides unit test coverage for TypeKind()
func TestTypeKind(t *testing.T) {
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
			template: `{{ typeKind .Val }}`,
			args: Args{
				Val: nil,
			},
			want: "nil",
		},
		{
			name:     "int",
			template: `{{ typeKind .Val }}`,
			args: Args{
				Val: 3,
			},
			want: "int",
		},
		{
			name:     "time.Duration",
			template: `{{ typeKind .Val }}`,
			args: Args{
				Val: 10 * time.Second,
			},
			want: "int64",
		},
		{
			name:     "*int",
			template: `{{ typeKind .Val }}`,
			args: Args{
				Val: helperPtrToInt(10),
			},
			want: "ptr",
		},
		{
			name:     "[]int",
			template: `{{ typeKind .Val }}`,
			args: Args{
				Val: []int{4},
			},
			want: "slice",
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

// TestJoinWith provides unit test coverage for JoinWith()
func TestJoinWith(t *testing.T) {
	type Args struct {
		Glue string
		A    []interface{}
		B    string
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
			template: `{{ joinWith .Glue .A }}`,
			args: Args{
				Glue: "",
				A:    nil,
			},
			want: "",
		},
		{
			name:     "empty",
			template: `{{ joinWith .Glue .A }}`,
			args: Args{
				Glue: "",
				A:    []interface{}{},
			},
			want: "",
		},
		{
			name:     "one",
			template: `{{ joinWith .Glue .A }}`,
			args: Args{
				Glue: "*",
				A:    []interface{}{"one"},
			},
			want: "one",
		},
		{
			name:     "two",
			template: `{{ joinWith .Glue .A }}`,
			args: Args{
				Glue: "^",
				A:    []interface{}{"one", "two"},
			},
			want: "one^two",
		},
		{
			name:     "three",
			template: `{{ joinWith .Glue .A }}`,
			args: Args{
				Glue: " - ",
				A:    []interface{}{"one", "two", "three"},
			},
			want: "one - two - three",
		},
		{
			name:     "bad list",
			template: `{{ joinWith .Glue .B }}`,
			args: Args{
				B: "Fail",
			},
			wantErr: true,
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

// TestSplitOn provides unit test coverage for SplitOn()
func TestSplitOn(t *testing.T) {
	type Args struct {
		Glue string
		S    string
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "one",
			template: `{{ splitOn .Glue .S | toJSON }}`,
			args: Args{
				Glue: " ",
				S:    "one",
			},
			want: `["one"]`,
		},
		{
			name:     "two",
			template: `{{ splitOn .Glue .S | toJSON }}`,
			args: Args{
				Glue: " ",
				S:    "one two",
			},
			want: `["one","two"]`,
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

// TestPrefix provides unit test coverage for Prefix()
func TestPrefix(t *testing.T) {
	type Args struct {
		Prefix  string
		T       int
		Content string
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{}

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

// TestToColumn provides unit test coverage for ToColumn()
func TestToColumn(t *testing.T) {
	type Args struct {
		W int
		S string
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
			template: "{{ toColumns .W .S }}",
			args: Args{
				W: 10,
				S: "",
			},
			want: "",
		},
		{
			name:     "too small",
			template: "{{ toColumns .W .S }}",
			args: Args{
				W: 10,
				S: "foo",
			},
			want: "foo\n",
		},
		{
			name:     "simple",
			template: "{{ toColumns .W .S }}",
			args: Args{
				W: 3,
				S: "foo bar",
			},
			want: "foo\nbar\n",
		},
		{
			name:     "find space",
			template: "{{ toColumns .W .S }}",
			args: Args{
				W: 4,
				S: "foo bar",
			},
			want: "foo\nbar\n",
		},
		{
			name:     "long word",
			template: "{{ toColumns .W .S }}",
			args: Args{
				W: 4,
				S: "foobar baz",
			},
			want: "foobar\nbaz\n",
		},
		{
			name:     "four lines",
			template: "{{ toColumns .W .S }}",
			args: Args{
				W: 3,
				S: "foo bar baz snk",
			},
			want: "foo\nbar\nbaz\nsnk\n",
		},
		{
			name:     "possible off by one",
			template: "{{ toColumns .W .S }}",
			args: Args{
				W: 5,
				S: "a b c d e f g",
			},
			want: "a b c\nd e f\ng\n",
		},
		{
			name:     "many newlines",
			template: "{{ toColumns .W .S }}",
			args: Args{
				W: 3,
				S: "foo\n\n\n\nbar\n\n\n\n\n",
			},
			want: "foo\nbar\n",
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

// TestTerminalWidth provides unit test coverage for TerminalWidth()
func TestTerminalWidth(t *testing.T) {
	type Args struct {
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "basic",
			template: "{{ terminalWidth }}",
			want:     "0", // probably - depends how/where the test is run
			wantErr:  false,
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

// TestTitleCaseWithAbbr provides unit test coverage for TitleCaseWithAbbr()
func TestTitleCaseWithAbbr(t *testing.T) {
	type Args struct {
		Abbrv []string
		Word  string
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "no abbreviations",
			template: `{{ titleCaseWithAbbr .Abbrv .Word }}`,
			args: Args{
				Abbrv: []string{},
				Word:  "nz all blacks",
			},
			want:    "Nz All Blacks",
			wantErr: false,
		},
		{
			name:     "basic",
			template: `{{ titleCaseWithAbbr .Abbrv .Word }}`,
			args: Args{
				Abbrv: []string{"nz"},
				Word:  "nz all blacks",
			},
			want:    "NZ All Blacks",
			wantErr: false,
		},
		{
			name:     "in-line list",
			template: `{{ titleCaseWithAbbr (list "nz") .Word }}`,
			args: Args{
				Word: "nz all blacks",
			},
			want:    "NZ All Blacks",
			wantErr: false,
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

// TestQuoteSingle provides unit test coverage for QuoteSingle
func TestQuoteSingle(t *testing.T) {
	type Args struct {
		S string
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
			template: "{{- q .S -}}",
			args: Args{
				S: "",
			},
			want:    "''",
			wantErr: false,
		},
		{
			name:     "basic",
			template: "{{- q .S -}}",
			args: Args{
				S: "rawr",
			},
			want:    "'rawr'",
			wantErr: false,
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

// TestQuoteDouble provides unit test coverage for QuoteDouble
func TestQuoteDouble(t *testing.T) {
	type Args struct {
		S string
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
			template: "{{- qq .S -}}",
			args: Args{
				S: "",
			},
			want:    "\"\"",
			wantErr: false,
		},
		{
			name:     "basic",
			template: "{{- qq .S -}}",
			args: Args{
				S: "rawr",
			},
			want:    "\"rawr\"",
			wantErr: false,
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

// TestQuoteBack provides unit test coverage for QuoteBack
func TestQuoteBack(t *testing.T) {
	type Args struct {
		S string
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
			template: "{{- bq .S -}}",
			args: Args{
				S: "",
			},
			want:    "``",
			wantErr: false,
		},
		{
			name:     "basic",
			template: "{{- bq .S -}}",
			args: Args{
				S: "rawr",
			},
			want:    "`rawr`",
			wantErr: false,
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
