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
	assert.Len(t, fn, 21, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestGenerateIncludeFn provides unit test coverage for GenerateIncludeFn()
func TestGenerateIncludeFn(t *testing.T) {

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

// TestFirst provides unit test coverage for First()
func TestFirst(t *testing.T) {
	type Args struct {
		list interface{}
	}

	tests := []struct {
		name          string
		args          Args
		wantInterface interface{}
		wantError     bool
	}{
		{
			name: "nil",
			args: Args{
				list: nil,
			},
			wantInterface: nil,
			wantError:     false,
		},
		{
			name: "not a list",
			args: Args{
				list: "actually a string",
			},
			wantError: true,
		},
		{
			name: "from zero",
			args: Args{
				list: []int{},
			},
			wantInterface: nil,
			wantError:     false,
		},
		{
			name: "from two",
			args: Args{
				list: []string{"one", "two"},
			},
			wantInterface: "one",
			wantError:     false,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotInterface, gotError := First(tt.args.list)
			if tt.wantError {
				require.Error(t, gotError)
			} else {
				require.NoError(t, gotError)
			}
			assert.Equal(t, tt.wantInterface, gotInterface)
		})
	}
}

// TestRest provides unit test coverage for Rest()
func TestRest(t *testing.T) {
	type Args struct {
		list interface{}
	}

	tests := []struct {
		name          string
		args          Args
		wantInterface interface{}
		wantError     bool
	}{
		{
			name: "nil",
			args: Args{
				list: nil,
			},
			wantInterface: nil,
			wantError:     false,
		},
		{
			name: "not a list",
			args: Args{
				list: "actually a string",
			},
			wantError: true,
		},
		{
			name: "from zero",
			args: Args{
				list: []int{},
			},
			wantInterface: nil,
			wantError:     false,
		},
		{
			name: "from two",
			args: Args{
				list: []string{"one", "two"},
			},
			wantInterface: []string{"two"},
			wantError:     false,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotInterface, gotError := Rest(tt.args.list)
			if tt.wantError {
				require.Error(t, gotError)
			} else {
				require.NoError(t, gotError)
			}
			assert.Equal(t, tt.wantInterface, gotInterface)
		})
	}
}

// TestLast provides unit test coverage for Last()
func TestLast(t *testing.T) {
	type Args struct {
		list interface{}
	}

	tests := []struct {
		name          string
		args          Args
		wantInterface interface{}
		wantError     bool
	}{
		{
			name: "nil",
			args: Args{
				list: nil,
			},
			wantInterface: nil,
			wantError:     false,
		},
		{
			name: "not a list",
			args: Args{
				list: "actually a string",
			},
			wantError: true,
		},
		{
			name: "from zero",
			args: Args{
				list: []int{},
			},
			wantInterface: nil,
			wantError:     false,
		},
		{
			name: "from two",
			args: Args{
				list: []string{"one", "two"},
			},
			wantInterface: "two",
			wantError:     false,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotInterface, gotError := Last(tt.args.list)
			if tt.wantError {
				require.Error(t, gotError)
			} else {
				require.NoError(t, gotError)
			}
			assert.Equal(t, tt.wantInterface, gotInterface)
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

// TestSpaces provides unit test coverage for Space()
func TestSpaces(t *testing.T) {
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

// TestTabs provides unit test coverage for Tab()
func TestTabs(t *testing.T) {
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

// TestConcat provides unit test coverage for Concat()
func TestConcat(t *testing.T) {
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
			got := Concat(tt.args.s...)
			assert.Equal(t, tt.want, got)
		})
	}
}
