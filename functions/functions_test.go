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

// TestAll provides unit test coverage for All()
func TestAll(t *testing.T) {
	fn := All(nil)
	assert.Len(t, fn, 43, "weakly ensuring functions haven't been added/removed without updating tests")
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
		s string
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "empty",
			args: Args{
				s: "",
			},
			want: "",
		},
		{
			name: "simple",
			args: Args{
				s: "simple",
			},
			want: "Simple",
		},
		{
			name: "same",
			args: Args{
				s: "Same",
			},
			want: "Same",
		},
		{
			name: "number",
			args: Args{
				s: "3rd",
			},
			want: "3rd",
		},
		{
			name: "multiple words",
			args: Args{
				s: "spam test",
			},
			want: "Spam test",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := UppercaseFirst(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestNewline provides unit test coverage for Newline()
func TestNewline(t *testing.T) {
	type Args struct {
		c []int
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "no params",
			args: Args{
				c: nil,
			},
			want: "\n",
		}, {
			name: "params",
			args: Args{
				c: []int{3},
			},
			want: "\n\n\n",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Newline(tt.args.c...)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestRep provides unit test coverage for Rep()
func TestRep(t *testing.T) {
	type Args struct {
		n int
		s []string
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "none",
			args: Args{
				n: 0,
				s: []string{"foo"},
			},
			want: "",
		},
		{
			name: "repeated empty",
			args: Args{
				n: 2,
				s: []string{""},
			},
			want: "",
		},
		{
			name: "one",
			args: Args{
				n: 1,
				s: []string{"x"},
			},
			want: "x",
		},
		{
			name: "two",
			args: Args{
				n: 2,
				s: []string{"foo"},
			},
			want: "foofoo",
		},
		{
			name: "negative one",
			args: Args{
				n: -1,
				s: []string{"foo"},
			},
			want: "",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Rep(tt.args.n, tt.args.s...)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestWhenEmpty provides unit test coverage for WhenEmpty()
func TestWhenEmpty(t *testing.T) {
	type Args struct {
		d string
		s string
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "empty",
			args: Args{
				d: "x",
				s: "",
			},
			want: "x",
		},
		{
			name: "not empty",
			args: Args{
				d: "x",
				s: "y",
			},
			want: "y",
		},
		{
			name: "default also empty",
			args: Args{
				d: "",
				s: "",
			},
			want: "",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := WhenEmpty(tt.args.d, tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestIndent provides unit test coverage for Indent()
func TestIndent(t *testing.T) {
	type Args struct {
		t       int
		content string
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "none",
			args: Args{
				t:       0,
				content: "foo",
			},
			want: "foo",
		},
		{
			name: "one",
			args: Args{
				t:       1,
				content: "foo",
			},
			want: "\tfoo",
		},
		{
			name: "two",
			args: Args{
				t:       2,
				content: "foo",
			},
			want: "\t\tfoo",
		},
		{
			name: "negative one",
			args: Args{
				t:       -1,
				content: "",
			},
			want: "",
		},
		{
			name: "multi line",
			args: Args{
				t:       1,
				content: "foo\nbar",
			},
			want: "\tfoo\n\tbar",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Indent(tt.args.t, tt.args.content)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestIndentSpace provides unit test coverage for Indent()
func TestIndentSpace(t *testing.T) {
	type Args struct {
		t       int
		content string
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "none",
			args: Args{
				t:       0,
				content: "foo",
			},
			want: "foo",
		},
		{
			name: "one",
			args: Args{
				t:       1,
				content: "foo",
			},
			want: " foo",
		},
		{
			name: "two",
			args: Args{
				t:       2,
				content: "foo",
			},
			want: "  foo",
		},
		{
			name: "negative one",
			args: Args{
				t:       -1,
				content: "",
			},
			want: "",
		},
		{
			name: "multi line",
			args: Args{
				t:       1,
				content: "foo\nbar",
			},
			want: " foo\n bar",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := IndentSpace(tt.args.t, tt.args.content)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestSpace provides unit test coverage for Space()
func TestSpace(t *testing.T) {
	type Args struct {
		n int
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "none",
			args: Args{
				n: 0,
			},
			want: "",
		},
		{
			name: "one",
			args: Args{
				n: 1,
			},
			want: " ",
		},
		{
			name: "two",
			args: Args{
				n: 2,
			},
			want: "  ",
		},
		{
			name: "negative one",
			args: Args{
				n: -1,
			},
			want: "",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Space(tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestTab provides unit test coverage for Tab()
func TestTab(t *testing.T) {
	type Args struct {
		n int
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "none",
			args: Args{
				n: 0,
			},
			want: "",
		},
		{
			name: "one",
			args: Args{
				n: 1,
			},
			want: "\t",
		},
		{
			name: "two",
			args: Args{
				n: 2,
			},
			want: "\t\t",
		},
		{
			name: "negative one",
			args: Args{
				n: -1,
			},
			want: "",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Tab(tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestPadRight provides unit test coverage for PadRight()
func TestPadRight(t *testing.T) {
	type Args struct {
		n int
		s string
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "basic",
			args: Args{
				n: 10,
				s: "basic",
			},
			want: "basic     ",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := PadRight(tt.args.n, tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestPadLeft provides unit test coverage for PadLeft()
func TestPadLeft(t *testing.T) {
	type Args struct {
		n int
		s string
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "basic",
			args: Args{
				n: 10,
				s: "basic",
			},
			want: "     basic",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := PadLeft(tt.args.n, tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestNow shows that there's testing, and there just keeping yourself amused
func TestNow(t *testing.T) {
	n := Now()
	_, err := time.Parse(time.RFC3339, n)
	assert.NoError(t, err)
}

// TestPadLeft provides unit test coverage for PadLeft()
func TestIsZero(t *testing.T) {
	type Args struct {
		val interface{}
	}

	tests := []struct {
		name string
		args Args
		want bool
	}{
		{
			name: "nil",
			args: Args{
				val: nil,
			},
			want: true,
		},
		{
			name: "zero int",
			args: Args{
				val: 0,
			},
			want: true,
		},
		{
			name: "non-zero int",
			args: Args{
				val: 10,
			},
			want: false,
		},
		{
			name: "pointer zero int",
			args: Args{
				val: helperPtrToInt(0),
			},
			want: false,
		},
		{
			name: "pointer non-zero int",
			args: Args{
				val: helperPtrToInt(-82),
			},
			want: false,
		},
		{
			name: "non-zero int",
			args: Args{
				val: 10,
			},
			want: false,
		},
		{
			name: "non-zero int",
			args: Args{
				val: 10,
			},
			want: false,
		},
		{
			name: "empty string",
			args: Args{
				val: "",
			},
			want: true,
		},
		{
			name: "non-empty string",
			args: Args{
				val: "foo",
			},
			want: false,
		},
		{
			name: "empty array",
			args: Args{
				val: []int{},
			},
			want: true,
		},
		{
			name: "nil array",
			args: Args{
				val: []float64(nil),
			},
			want: true,
		},
		{
			name: "non-empty array",
			args: Args{
				val: []string{"bar"},
			},
			want: false,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := IsZero(tt.args.val)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestBracket provides unit test coverage for Bracket()
func TestBracket(t *testing.T) {
	type Args struct {
		s string
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "empty",
			args: Args{
				s: "",
			},
			want: "()",
		},
		{
			name: "word",
			args: Args{
				s: "foo",
			},
			want: "(foo)",
		},
		{
			name: "words",
			args: Args{
				s: "foo bar",
			},
			want: "(foo bar)",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Bracket(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestBracketWith provides unit test coverage for BracketWith()
func TestBracketWith(t *testing.T) {
	type Args struct {
		b string
		s string
	}

	tests := []struct {
		name       string
		args       Args
		wantString string
		wantError  bool
	}{
		{
			name: "none",
			args: Args{
				b: "",
				s: "",
			},
			wantString: "",
		},
		{
			name: "basic",
			args: Args{
				b: "()",
				s: "",
			},
			wantString: "()",
		},
		{
			name: "word",
			args: Args{
				b: "<>",
				s: "foo",
			},
			wantString: "<foo>",
		},
		{
			name: "words",
			args: Args{
				b: "{{-  -}}",
				s: "foo bar",
			},
			wantString: "{{- foo bar -}}",
		},
		{
			name: "mismatched",
			args: Args{
				b: ")",
				s: "baz",
			},
			wantError: true,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotString, gotError := BracketWith(tt.args.b, tt.args.s)
			if tt.wantError {
				require.Error(t, gotError)
			} else {
				require.NoError(t, gotError)
			}
			assert.Equal(t, tt.wantString, gotString)
		})
	}
}

// TestJoin provides unit test coverage for Join()
func TestJoin(t *testing.T) {
	type Args struct {
		s []string
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "nil",
			args: Args{
				s: nil,
			},
			want: "",
		},
		{
			name: "empty",
			args: Args{
				s: []string{},
			},
			want: "",
		},
		{
			name: "one",
			args: Args{
				s: []string{"one"},
			},
			want: "one",
		},
		{
			name: "two",
			args: Args{
				s: []string{"one", "two"},
			},
			want: "onetwo",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Join(tt.args.s...)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestTypeName provides unit test coverage for TypeName()
func TestTypeName(t *testing.T) {
	type Args struct {
		val interface{}
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "nil",
			args: Args{
				val: nil,
			},
			want: "nil",
		},
		{
			name: "int",
			args: Args{
				val: 3,
			},
			want: "int",
		},
		{
			name: "time.Duration",
			args: Args{
				val: 10 * time.Second,
			},
			want: "time.Duration",
		},
		{
			name: "*int",
			args: Args{
				val: helperPtrToInt(10),
			},
			want: "*int",
		},
		{
			name: "[]int",
			args: Args{
				val: []int{4},
			},
			want: "[]int",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := TypeName(tt.args.val)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestJoinWith provides unit test coverage for JoinWith()
func TestJoinWith(t *testing.T) {
	type Args struct {
		glue string
		s    []string
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "nil",
			args: Args{
				glue: "",
				s:    nil,
			},
			want: "",
		},
		{
			name: "empty",
			args: Args{
				glue: "",
				s:    []string{},
			},
			want: "",
		},
		{
			name: "one",
			args: Args{
				glue: "*",
				s:    []string{"one"},
			},
			want: "one",
		},
		{
			name: "two",
			args: Args{
				glue: "^",
				s:    []string{"one", "two"},
			},
			want: "one^two",
		},
		{
			name: "three",
			args: Args{
				glue: " - ",
				s:    []string{"one", "two", "three"},
			},
			want: "one - two - three",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := JoinWith(tt.args.glue, tt.args.s...)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestSplitOn provides unit test coverage for SplitOn()
func TestSplitOn(t *testing.T) {
	type Args struct {
		glue string
		s    string
	}

	tests := []struct {
		name string
		args Args
		want []string
	}{
		{
			name: "one",
			args: Args{
				glue: " ",
				s:    "one",
			},
			want: []string{"one"},
		},
		{
			name: "two",
			args: Args{
				glue: " ",
				s:    "one two",
			},
			want: []string{"one", "two"},
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := SplitOn(tt.args.glue, tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestToJSON provides unit test coverage for ToJSON()
func TestToJSON(t *testing.T) {
	type Args struct {
		val interface{}
	}

	tests := []struct {
		name       string
		args       Args
		wantString string
		wantError  bool
	}{
		{
			name: "simple object",
			args: Args{
				val: map[string]string{
					"one": "foo",
					"two": "bar",
				},
			},
			wantString: `{"one":"foo","two":"bar"}`,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotString, gotError := ToJSON(tt.args.val)
			if tt.wantError {
				require.Error(t, gotError)
			} else {
				require.NoError(t, gotError)
			}
			assert.Equal(t, tt.wantString, gotString)
		})
	}
}

// TestFormatJSON provides unit test coverage for FormatJSON()
func TestFormatJSON(t *testing.T) {
	type Args struct {
		j      string
		indent string
	}

	tests := []struct {
		name       string
		args       Args
		wantString string
		wantError  bool
	}{
		{
			name: "simple object",
			args: Args{
				indent: "\t",
				j:      `{"one":"foo","two":"bar"}`,
			},
			wantString: "{\n\t\"one\": \"foo\",\n\t\"two\": \"bar\"\n}",
		},
		{
			name: "bad json",
			args: Args{
				indent: "\t",
				j:      `{"one":"foo","two":"forgot end brace..."`,
			},
			wantError: true,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotString, gotError := FormatJSON(tt.args.indent, tt.args.j)
			if tt.wantError {
				require.Error(t, gotError)
			} else {
				require.NoError(t, gotError)
			}
			assert.Equal(t, tt.wantString, gotString)
		})
	}
}

// TestToYAML provides unit test coverage for ToYAML()
func TestToYAML(t *testing.T) {
	type Args struct {
		val interface{}
	}

	tests := []struct {
		name       string
		args       Args
		wantString string
		wantError  bool
	}{
		{
			name: "simple object",
			args: Args{
				val: map[string]string{
					"one": "foo",
					"two": "bar",
				},
			},
			wantString: "one: foo\ntwo: bar\n",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotString, gotError := ToYAML(tt.args.val)
			if tt.wantError {
				require.Error(t, gotError)
			} else {
				require.NoError(t, gotError)
			}
			assert.Equal(t, tt.wantString, gotString)
		})
	}
}
