package functions

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestListFunctions provides unit test coverage for ListFunctions
func TestListFunctions(t *testing.T) {
	fn := ListFunctions()
	assert.Len(t, fn, 13, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestList provides unit test coverage for List()
func TestList(t *testing.T) {
	tests := []TestSet{
		{
			name:     "empty",
			template: `{{ list }}`,
			args:     TestArgs{},
			want:     "[]",
		},
		{
			name:     "single int",
			template: `{{ list .A }}`,
			args: TestArgs{
				"A": 5,
			},
			want: "[5]",
		},
		{
			name:     "int + string",
			template: `{{ list .A .B }}`,
			args: TestArgs{
				"A": 5,
				"B": "rawr",
			},
			want: "[5 rawr]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestFirst provides unit test coverage for First()
func TestFirst(t *testing.T) {
	tests := []TestSet{
		{
			name:     "not a list",
			template: `{{ $x := "not a list" }}{{ first $x }}`,
			wantErr:  true,
		},
		{
			name:     "from zero",
			template: `{{ $x := list }}{{ first $x }}`,
			want:     "<no value>",
		},
		{
			name:     "from two",
			template: `{{ $x := list .A .B }}{{ first $x }}`,
			args: TestArgs{
				"A": "A",
				"B": "B",
			},
			want: "A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestRest provides unit test coverage for Rest()
func TestRest(t *testing.T) {
	tests := []TestSet{
		{
			name:     "not a list",
			template: `{{ $x := "not a list" }}{{ rest $x }}`,
			wantErr:  true,
		},
		{
			name:     "from zero",
			template: `{{ $x := list }}{{ rest $x }}`,
			want:     "<no value>",
		},
		{
			name:     "from two",
			template: `{{ $x := list .A .B }}{{ rest $x }}`,
			args: TestArgs{
				"A": "A",
				"B": "B",
			},
			want: "[B]",
		},
		{
			name:     "with nils",
			template: `{{ $x := list .A .B .C .D }}{{ rest $x }}`,
			args: TestArgs{
				"A": "one",
				"B": "two",
				"C": nil,
				"D": "four",
			},
			want: "[two <nil> four]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestRest provides unit test coverage for Rest()
func TestPop(t *testing.T) {
	tests := []TestSet{
		{
			name:     "not a list",
			template: `{{ $x := "not a list" }}{{ pop $x }}`,
			wantErr:  true,
		},
		{
			name:     "from zero",
			template: `{{ $x := list }}{{ pop $x }}`,
			want:     "<no value>",
		},
		{
			name:     "from two",
			template: `{{ $x := list .A .B }}{{ pop $x }}`,
			args: TestArgs{
				"A": "A",
				"B": "B",
			},
			want: "[A]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestLast provides unit test coverage for Last()
func TestLast(t *testing.T) {
	tests := []TestSet{
		{
			name:     "not a list",
			template: `{{ $x := "not a list" }}{{ last $x }}`,
			wantErr:  true,
		},
		{
			name:     "from zero",
			template: `{{ $x := list }}{{ last $x }}`,
			want:     "<no value>",
		},
		{
			name:     "from two",
			template: `{{ $x := list .A .B }}{{ last $x }}`,
			args: TestArgs{
				"A": "A",
				"B": "B",
			},
			want: "B",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestContains provides unit test coverage for Contains()
func TestContains(t *testing.T) {
	tests := []TestSet{
		{
			name:     "test against empty",
			template: `{{ contains .Haystack .Needle }}`,
			args: TestArgs{
				"Needle":   "A",
				"Haystack": []string{},
			},
			want: "false",
		},
		{
			name:     "not a list",
			template: `{{ $x := "not a list" }}{{ contains $x }}`,
			wantErr:  true,
		},
		{
			name:     "in list",
			template: `{{ contains .Haystack .Needle }}`,
			args: TestArgs{
				"Needle":   "A",
				"Haystack": []string{"A", "B"},
			},
			want: "true",
		},
		{
			name:     "not in list",
			template: `{{ contains .Haystack .Needle }}`,
			args: TestArgs{
				"Needle":   "C",
				"Haystack": []string{"A", "B"},
			},
			want: "false",
		},
		{
			name:     "different types",
			template: `{{ contains .Haystack .Needle }}`,
			args: TestArgs{
				"Needle":   1,
				"Haystack": []interface{}{"1", 2},
			},
			want: "false",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestFilter provides unit test coverage for Filter()
func TestFilter(t *testing.T) {
	tests := []TestSet{
		{
			name:     "filter from empty",
			template: `{{ filter .List .Filter }}`,
			args: TestArgs{
				"List":   []string{},
				"Filter": "A",
			},
			want: "[]",
		},
		{
			name:     "not a list",
			template: `{{ filter .NotList .Filter }}`,
			args: TestArgs{
				"List":   "not a list",
				"Filter": "A",
			},
			wantErr: true,
		},
		{
			name:     "in list",
			template: `{{ filter .List .Filter }}`,
			args: TestArgs{
				"List":   []string{"A", "B"},
				"Filter": "A",
			},
			want: "[B]",
		},
		{
			name:     "not in list",
			template: `{{ filter .List .Filter }}`,
			args: TestArgs{
				"List":   []string{"A", "B"},
				"Filter": "C",
			},
			want: "[A B]",
		},
		{
			name:     "different type",
			template: `{{ filter .List .Filter }}`,
			args: TestArgs{
				"List":   []int{1, 2},
				"Filter": "1",
			},
			want: "[1 2]",
		},
		{
			name:     "remove multiple",
			template: `{{ filter .List .Filter }}`,
			args: TestArgs{
				"List":   []int{1, 2, 1, 2},
				"Filter": 1,
			},
			want: "[2 2]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestPush provides unit test coverage for Push()
func TestPush(t *testing.T) {
	tests := []TestSet{
		{
			name:     "not a list",
			template: `{{ push .NotList .Item }}`,
			args: TestArgs{
				"NotList": "not a list",
				"Item":    "A",
			},
			wantErr: true,
		},
		{
			name:     "push to empty",
			template: `{{ push .List .Item }}`,
			args: TestArgs{
				"List": []interface{}{},
				"Item": "A",
			},
			want: "[A]",
		},
		{
			name:     "push to existing",
			template: `{{ push .List .Item }}`,
			args: TestArgs{
				"List": []interface{}{"A", "B"},
				"Item": "C",
			},
			want: "[A B C]",
		},
		{
			name:     "mixed types",
			template: `{{ push .List .Item }}`,
			args: TestArgs{
				"List": []interface{}{"A", 3},
				"Item": "1.2",
			},
			want: "[A 3 1.2]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestUnshift provides unit test coverage for Unshift()
func TestUnshift(t *testing.T) {
	tests := []TestSet{
		{
			name:     "not a list",
			template: `{{ unshift .NotList .Item }}`,
			args: TestArgs{
				"NotList": "not a list",
				"Item":    "A",
			},
			wantErr: true,
		},
		{
			name:     "unshift to empty",
			template: `{{ unshift .List .Item }}`,
			args: TestArgs{
				"List": []interface{}{},
				"Item": "A",
			},
			want: "[A]",
		},
		{
			name:     "unshift to existing",
			template: `{{ unshift .List .Item }}`,
			args: TestArgs{
				"List": []interface{}{"A", "B"},
				"Item": "C",
			},
			want: "[C A B]",
		},
		{
			name:     "mixed types",
			template: `{{ unshift .List .Item }}`,
			args: TestArgs{
				"List": []interface{}{"A", 3},
				"Item": "1.2",
			},
			want: "[1.2 A 3]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestSlice provides unit test coverage for TestSlice()
func TestSlice(t *testing.T) {
	tests := []TestSet{
		{
			name:     "not a list",
			template: `{{ slice .I .J .NotList }}`,
			args: TestArgs{
				"NotList": "not a list",
				"I":       0,
				"J":       0,
			},
			wantErr: true,
		},
		{
			name:     "slice on empty",
			template: `{{ slice .I .J .List }}`,
			args: TestArgs{
				"List": []int{},
				"I":    0,
				"J":    0,
			},
			want: "[]",
		},
		{
			name:     "middle slice",
			template: `{{ slice .I .J .List }}`,
			args: TestArgs{
				"List": []string{"A", "B", "C", "D"},
				"I":    1,
				"J":    3,
			},
			want: "[B C]",
		},
		{
			name:     "out of bounds - leading",
			template: `{{ slice .I .J .List }}`,
			args: TestArgs{
				"List": []interface{}{"A", "B", "C", "D"},
				"I":    -1,
				"J":    2,
			},
			wantErr: true,
		},
		{
			name:     "out of bounds - trailing",
			template: `{{ slice .I .J .List }}`,
			args: TestArgs{
				"List": []interface{}{"A", "B", "C", "D"},
				"I":    3,
				"J":    5,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
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
