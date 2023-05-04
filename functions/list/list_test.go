package list

import (
	"testing"

	"github.com/mantidtech/tplr/functions/helper"
	"github.com/stretchr/testify/assert"
)

// TestListFunctions provides unit test coverage for ListFunctions
func TestListFunctions(t *testing.T) {
	fn := Functions()
	assert.Len(t, fn, 12, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestList provides unit test coverage for List()
func TestList(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "empty",
			Template: `{{ list }}`,
			Args:     helper.TestArgs{},
			Want:     "[]",
		},
		{
			Name:     "single int",
			Template: `{{ list .A }}`,
			Args: helper.TestArgs{
				"A": 5,
			},
			Want: "[5]",
		},
		{
			Name:     "int + string",
			Template: `{{ list .A .B }}`,
			Args: helper.TestArgs{
				"A": 5,
				"B": "rawr",
			},
			Want: "[5 rawr]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestFirst provides unit test coverage for First()
func TestFirst(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "not a list",
			Template: `{{ $x := "not a list" }}{{ first $x }}`,
			WantErr:  true,
		},
		{
			Name:     "from zero",
			Template: `{{ $x := list }}{{ first $x }}`,
			Want:     "<no value>",
		},
		{
			Name:     "from two",
			Template: `{{ $x := list .A .B }}{{ first $x }}`,
			Args: helper.TestArgs{
				"A": "A",
				"B": "B",
			},
			Want: "A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestRest provides unit test coverage for Rest()
func TestRest(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "not a list",
			Template: `{{ $x := "not a list" }}{{ rest $x }}`,
			WantErr:  true,
		},
		{
			Name:     "from zero",
			Template: `{{ $x := list }}{{ rest $x }}`,
			Want:     "<no value>",
		},
		{
			Name:     "from two",
			Template: `{{ $x := list .A .B }}{{ rest $x }}`,
			Args: helper.TestArgs{
				"A": "A",
				"B": "B",
			},
			Want: "[B]",
		},
		{
			Name:     "with nils",
			Template: `{{ $x := list .A .B .C .D }}{{ rest $x }}`,
			Args: helper.TestArgs{
				"A": "one",
				"B": "two",
				"C": nil,
				"D": "four",
			},
			Want: "[two <nil> four]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestRest provides unit test coverage for Rest()
func TestPop(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "not a list",
			Template: `{{ $x := "not a list" }}{{ pop $x }}`,
			WantErr:  true,
		},
		{
			Name:     "from zero",
			Template: `{{ $x := list }}{{ pop $x }}`,
			Want:     "<no value>",
		},
		{
			Name:     "from two",
			Template: `{{ $x := list .A .B }}{{ pop $x }}`,
			Args: helper.TestArgs{
				"A": "A",
				"B": "B",
			},
			Want: "[A]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestLast provides unit test coverage for Last()
func TestLast(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "not a list",
			Template: `{{ $x := "not a list" }}{{ last $x }}`,
			WantErr:  true,
		},
		{
			Name:     "from zero",
			Template: `{{ $x := list }}{{ last $x }}`,
			Want:     "<no value>",
		},
		{
			Name:     "from two",
			Template: `{{ $x := list .A .B }}{{ last $x }}`,
			Args: helper.TestArgs{
				"A": "A",
				"B": "B",
			},
			Want: "B",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestContains provides unit test coverage for Contains()
func TestContains(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "nil",
			Template: `{{ contains .Haystack .Needle }}`,
			Args: helper.TestArgs{
				"Needle":   "C",
				"Haystack": nil,
			},
			WantErr: true,
		},
		{
			Name:     "test against empty",
			Template: `{{ contains .Haystack .Needle }}`,
			Args: helper.TestArgs{
				"Needle":   "A",
				"Haystack": []string{},
			},
			Want: "false",
		},
		{
			Name:     "not a list",
			Template: `{{ $x := "not a list" }}{{ contains $x }}`,
			WantErr:  true,
		},
		{
			Name:     "in list",
			Template: `{{ contains .Haystack .Needle }}`,
			Args: helper.TestArgs{
				"Needle":   "A",
				"Haystack": []string{"A", "B"},
			},
			Want: "true",
		},
		{
			Name:     "not in list",
			Template: `{{ contains .Haystack .Needle }}`,
			Args: helper.TestArgs{
				"Needle":   "C",
				"Haystack": []string{"A", "B"},
			},
			Want: "false",
		},
		{
			Name:     "different types",
			Template: `{{ contains .Haystack .Needle }}`,
			Args: helper.TestArgs{
				"Needle":   1,
				"Haystack": []any{"1", 2},
			},
			Want: "false",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestFilter provides unit test coverage for Filter()
func TestFilter(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "filter from empty",
			Template: `{{ filter .List .Filter }}`,
			Args: helper.TestArgs{
				"List":   []string{},
				"Filter": "A",
			},
			Want: "[]",
		},
		{
			Name:     "not a list",
			Template: `{{ filter .NotList .Filter }}`,
			Args: helper.TestArgs{
				"List":   "not a list",
				"Filter": "A",
			},
			WantErr: true,
		},
		{
			Name:     "in list",
			Template: `{{ filter .List .Filter }}`,
			Args: helper.TestArgs{
				"List":   []string{"A", "B"},
				"Filter": "A",
			},
			Want: "[B]",
		},
		{
			Name:     "not in list",
			Template: `{{ filter .List .Filter }}`,
			Args: helper.TestArgs{
				"List":   []string{"A", "B"},
				"Filter": "C",
			},
			Want: "[A B]",
		},
		{
			Name:     "different type",
			Template: `{{ filter .List .Filter }}`,
			Args: helper.TestArgs{
				"List":   []int{1, 2},
				"Filter": "1",
			},
			Want: "[1 2]",
		},
		{
			Name:     "remove multiple",
			Template: `{{ filter .List .Filter }}`,
			Args: helper.TestArgs{
				"List":   []int{1, 2, 1, 2},
				"Filter": 1,
			},
			Want: "[2 2]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestPush provides unit test coverage for Push()
func TestPush(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "not a list",
			Template: `{{ push .NotList .Item }}`,
			Args: helper.TestArgs{
				"NotList": "not a list",
				"Item":    "A",
			},
			WantErr: true,
		},
		{
			Name:     "push to empty",
			Template: `{{ push .List .Item }}`,
			Args: helper.TestArgs{
				"List": []any{},
				"Item": "A",
			},
			Want: "[A]",
		},
		{
			Name:     "push to existing",
			Template: `{{ push .List .Item }}`,
			Args: helper.TestArgs{
				"List": []any{"A", "B"},
				"Item": "C",
			},
			Want: "[A B C]",
		},
		{
			Name:     "mixed types",
			Template: `{{ push .List .Item }}`,
			Args: helper.TestArgs{
				"List": []any{"A", 3},
				"Item": "1.2",
			},
			Want: "[A 3 1.2]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestUnshift provides unit test coverage for Unshift()
func TestUnshift(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "not a list",
			Template: `{{ unshift .NotList .Item }}`,
			Args: helper.TestArgs{
				"NotList": "not a list",
				"Item":    "A",
			},
			WantErr: true,
		},
		{
			Name:     "unshift to empty",
			Template: `{{ unshift .List .Item }}`,
			Args: helper.TestArgs{
				"List": []any{},
				"Item": "A",
			},
			Want: "[A]",
		},
		{
			Name:     "unshift to existing",
			Template: `{{ unshift .List .Item }}`,
			Args: helper.TestArgs{
				"List": []any{"A", "B"},
				"Item": "C",
			},
			Want: "[C A B]",
		},
		{
			Name:     "mixed types",
			Template: `{{ unshift .List .Item }}`,
			Args: helper.TestArgs{
				"List": []any{"A", 3},
				"Item": "1.2",
			},
			Want: "[1.2 A 3]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestJoin provides unit test coverage for Join()
func TestJoin(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "nil",
			Template: `{{ join .A }}`,
			Args: helper.TestArgs{
				"A": nil,
			},
			WantErr: true,
		},
		{
			Name:     "empty",
			Template: `{{ join .A }}`,
			Args: helper.TestArgs{
				"A": []any{},
			},
			Want: "",
		},
		{
			Name:     "one",
			Template: `{{ join .A }}`,
			Args: helper.TestArgs{
				"A": []any{"one"},
			},
			Want: "one",
		},
		{
			Name:     "two",
			Template: `{{ join .A }}`,
			Args: helper.TestArgs{
				"A": []any{"one", "two"},
			},
			Want: "onetwo",
		},
		{
			Name:     "2",
			Template: `{{ join .A }}`,
			Args: helper.TestArgs{
				"A": []any{1, 2},
			},
			Want: "12",
		},
		{
			Name:     "bad list",
			Template: `{{ join .B }}`,
			Args: helper.TestArgs{
				"B": "Fail",
			},
			WantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestJoinWith provides unit test coverage for JoinWith()
func TestJoinWith(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "nil",
			Template: `{{ joinWith .Glue .A }}`,
			Args: helper.TestArgs{
				"Glue": "",
				"A":    nil,
			},
			WantErr: true,
		},
		{
			Name:     "empty",
			Template: `{{ joinWith .Glue .A }}`,
			Args: helper.TestArgs{
				"Glue": "",
				"A":    []any{},
			},
			Want: "",
		},
		{
			Name:     "one",
			Template: `{{ joinWith .Glue .A }}`,
			Args: helper.TestArgs{
				"Glue": "*",
				"A":    []any{"one"},
			},
			Want: "one",
		},
		{
			Name:     "two",
			Template: `{{ joinWith .Glue .A }}`,
			Args: helper.TestArgs{
				"Glue": "^",
				"A":    []any{"one", "two"},
			},
			Want: "one^two",
		},
		{
			Name:     "three",
			Template: `{{ joinWith .Glue .A }}`,
			Args: helper.TestArgs{
				"Glue": " - ",
				"A":    []any{"one", "two", "three"},
			},
			Want: "one - two - three",
		},
		{
			Name:     "bad list",
			Template: `{{ joinWith .Glue .B }}`,
			Args: helper.TestArgs{
				"B": "Fail",
			},
			WantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}
