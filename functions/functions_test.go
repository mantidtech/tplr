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
	assert.Len(t, fn, 47, "weakly ensuring functions haven't been added/removed without updating tests")
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

// TestSuffix provides unit test coverage for Indent()
func TestSuffix(t *testing.T) {
	type Args struct {
		suffix  string
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
				suffix:  "",
				t:       0,
				content: "foo",
			},
			want: "foo",
		},
		{
			name: "one",
			args: Args{
				suffix:  "X",
				t:       1,
				content: "foo",
			},
			want: "fooX",
		},
		{
			name: "two",
			args: Args{
				suffix:  "X",
				t:       2,
				content: "foo",
			},
			want: "fooXX",
		},
		{
			name: "negative one",
			args: Args{
				suffix:  "X",
				t:       -1,
				content: "",
			},
			want: "",
		},
		{
			name: "multi line",
			args: Args{
				suffix:  "X",
				t:       1,
				content: "foo\nbar",
			},
			want: "fooX\nbarX",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Suffix(tt.args.suffix, tt.args.t, tt.args.content)
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

// TestPrefix provides unit test coverage for Prefix()
func TestPrefix(t *testing.T) {
	t.Parallel()
	type Args struct {
		prefix  string
		t       int
		content string
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		// tests go here
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Prefix(tt.args.prefix, tt.args.t, tt.args.content)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestToColumns provides unit test coverage for ToColumns()
func TestToColumns(t *testing.T) {
	//t.Parallel()
	type Args struct {
		w int
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
				w: 10,
				s: "",
			},
			want: "",
		},
		{
			name: "too small",
			args: Args{
				w: 10,
				s: "foo",
			},
			want: "foo",
		},
		{
			name: "simple",
			args: Args{
				w: 3,
				s: "foo bar",
			},
			want: "foo\nbar",
		},
		{
			name: "find space",
			args: Args{
				w: 4,
				s: "foo bar",
			},
			want: "foo\nbar",
		},
		{
			name: "long word",
			args: Args{
				w: 4,
				s: "foobar baz",
			},
			want: "foobar\nbaz",
		},
		{
			name: "four lines",
			args: Args{
				w: 3,
				s: "foo bar baz snk",
			},
			want: "foo\nbar\nbaz\nsnk",
		},
		{
			name: "possible off by one",
			args: Args{
				w: 5,
				s: "a b c d e f g",
			},
			want: "a b c\nd e f\ng",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()
			got := ToColumns(tt.args.w, tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
