package strings

import (
	"testing"
	"time"

	"github.com/mantidtech/tplr/functions/helper"
	"github.com/mantidtech/tplr/functions/list"
	"github.com/stretchr/testify/assert"
)

func init() {
	// set to return a constant for testing
	nowActual = func() time.Time {
		loc, err := time.LoadLocation("US/Eastern")
		if err != nil {
			panic(err)
		}
		return time.Date(1997, 8, 29, 2, 14, 0, 133_700_000, loc)
	}
}

// TestStringFunctions provides unit test coverage for StringFunctions
func TestStringFunctions(t *testing.T) {
	fn := Functions()
	assert.Len(t, fn, 34, "weakly ensuring functions haven't been added/removed without updating tests")
}

func TestUppercaseFirst(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "empty",
			Template: `{{ ucFirst .S }}`,
			Args: helper.TestArgs{
				"S": "",
			},
			Want: "",
		},
		{
			Name:     "simple",
			Template: `{{ ucFirst .S }}`,
			Args: helper.TestArgs{
				"S": "simple",
			},
			Want: "Simple",
		},
		{
			Name:     "same",
			Template: `{{ ucFirst .S }}`,
			Args: helper.TestArgs{
				"S": "Same",
			},
			Want: "Same",
		},
		{
			Name:     "number",
			Template: `{{ ucFirst .S }}`,
			Args: helper.TestArgs{
				"S": "3rd",
			},
			Want: "3rd",
		},
		{
			Name:     "multiple words",
			Template: `{{ ucFirst .S }}`,
			Args: helper.TestArgs{
				"S": "spam test",
			},
			Want: "Spam test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

func TestNewline(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "no params",
			Template: `{{ nl }}`,
			Args:     helper.TestArgs{},
			Want:     "\n",
		},
		{
			Name:     "params",
			Template: `{{ nl .C }}`,
			Args: helper.TestArgs{
				"C": 3,
			},
			Want: "\n\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestRep provides unit test coverage for Rep()
func TestRep(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "none",
			Template: `{{ rep .N .S }}`,
			Args: helper.TestArgs{
				"N": 0,
				"S": "foo",
			},
			Want: "",
		},
		{
			Name:     "repeated empty",
			Template: `{{ rep .N .S }}`,
			Args: helper.TestArgs{
				"N": 2,
				"S": "",
			},
			Want: "",
		},
		{
			Name:     "one",
			Template: `{{ rep .N .S }}`,
			Args: helper.TestArgs{
				"N": 1,
				"S": "x",
			},
			Want: "x",
		},
		{
			Name:     "two",
			Template: `{{ rep .N .S }}`,
			Args: helper.TestArgs{
				"N": 2,
				"S": "foo",
			},
			Want: "foofoo",
		},
		{
			Name:     "multiple args",
			Template: `{{ rep .N .S .S2 }}`,
			Args: helper.TestArgs{
				"N":  1,
				"S":  "one",
				"S2": "two",
			},
			Want: "one two",
		},
		{
			Name:     "multiple args, twice",
			Template: `{{ rep .N .S .S2 }}`,
			Args: helper.TestArgs{
				"N":  2,
				"S":  "one",
				"S2": "two",
			},
			Want: "one twoone two",
		},
		{
			Name:     "negative one",
			Template: `{{ rep .N .S }}`,
			Args: helper.TestArgs{
				"N": -1,
				"S": "foo",
			},
			Want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestIndent provides unit test coverage for Indent()
func TestIndent(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "none",
			Template: `{{ indent .T .Content }}`,
			Args: helper.TestArgs{
				"T":       0,
				"Content": "foo",
			},
			Want: "foo",
		},
		{
			Name:     "one",
			Template: `{{ indent .T .Content }}`,
			Args: helper.TestArgs{
				"T":       1,
				"Content": "foo",
			},
			Want: " foo",
		},
		{
			Name:     "two",
			Template: `{{ indent .T .Content }}`,
			Args: helper.TestArgs{
				"T":       2,
				"Content": "foo",
			},
			Want: "  foo",
		},
		{
			Name:     "negative one",
			Template: `{{ indent .T .Content }}`,
			Args: helper.TestArgs{
				"T":       -1,
				"Content": "",
			},
			Want: "",
		},
		{
			Name:     "multi line",
			Template: `{{ indent .T .Content }}`,
			Args: helper.TestArgs{
				"T":       1,
				"Content": "foo\nbar",
			},
			Want: " foo\n bar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestIndent provides unit test coverage for Indent()
func TestUnindent(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "none",
			Template: `{{ unindent .T .Content }}`,
			Args: helper.TestArgs{
				"T":       0,
				"Content": "foo",
			},
			Want: "foo",
		},
		{
			Name:     "one",
			Template: `{{ unindent .T .Content }}`,
			Args: helper.TestArgs{
				"T":       1,
				"Content": "  foo",
			},
			Want: " foo",
		},
		{
			Name:     "two",
			Template: `{{ unindent .T .Content }}`,
			Args: helper.TestArgs{
				"T":       2,
				"Content": "  foo",
			},
			Want: "foo",
		},
		{
			Name:     "negative one",
			Template: `{{ unindent .T .Content }}`,
			Args: helper.TestArgs{
				"T":       -1,
				"Content": "",
			},
			WantErr: true,
		},
		{
			Name:     "multi line",
			Template: `{{ unindent .T .Content }}`,
			Args: helper.TestArgs{
				"T":       1,
				"Content": "  foo\n  bar",
			},
			Want: " foo\n bar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestSuffix provides unit test coverage for Suffix()
func TestSuffix(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "none",
			Template: `{{ suffix .Suffix .T .Content }}`,
			Args: helper.TestArgs{
				"Suffix":  "",
				"T":       0,
				"Content": "foo",
			},
			Want: "foo",
		},
		{
			Name:     "one",
			Template: `{{ suffix .Suffix .T .Content }}`,
			Args: helper.TestArgs{
				"Suffix":  "X",
				"T":       1,
				"Content": "foo",
			},
			Want: "fooX",
		},
		{
			Name:     "two",
			Template: `{{ suffix .Suffix .T .Content }}`,
			Args: helper.TestArgs{
				"Suffix":  "X",
				"T":       2,
				"Content": "foo",
			},
			Want: "fooXX",
		},
		{
			Name:     "negative one",
			Template: `{{ suffix .Suffix .T .Content }}`,
			Args: helper.TestArgs{
				"Suffix":  "X",
				"T":       -1,
				"Content": "",
			},
			Want: "",
		},
		{
			Name:     "multi line",
			Template: `{{ suffix .Suffix .T .Content }}`,
			Args: helper.TestArgs{
				"Suffix":  "X",
				"T":       1,
				"Content": "foo\nbar",
			},
			Want: "fooX\nbarX",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

func TestSpace(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "none",
			Template: `{{ sp .N }}`,
			Args: helper.TestArgs{
				"N": 0,
			},
			Want: "",
		},
		{
			Name:     "one",
			Template: `{{ sp .N }}`,
			Args: helper.TestArgs{
				"N": 1,
			},
			Want: " ",
		},
		{
			Name:     "two",
			Template: `{{ sp .N }}`,
			Args: helper.TestArgs{
				"N": 2,
			},
			Want: "  ",
		},
		{
			Name:     "negative one",
			Template: `{{ sp .N }}`,
			Args: helper.TestArgs{
				"N": -1,
			},
			Want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

func TestTab(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "none",
			Template: `{{ tab .N }}`,
			Args: helper.TestArgs{
				"N": 0,
			},
			Want: "",
		},
		{
			Name:     "one",
			Template: `{{ tab .N }}`,
			Args: helper.TestArgs{
				"N": 1,
			},
			Want: "\t",
		},
		{
			Name:     "two",
			Template: `{{ tab .N }}`,
			Args: helper.TestArgs{
				"N": 2,
			},
			Want: "\t\t",
		},
		{
			Name:     "negative one",
			Template: `{{ tab .N }}`,
			Args: helper.TestArgs{
				"N": -1,
			},
			Want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

func TestPadRight(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "basic",
			Template: `{{ padRight .N .S }}`,
			Args: helper.TestArgs{
				"N": 10,
				"S": "basic",
			},
			Want: "basic     ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

func TestPadLeft(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "basic",
			Template: `{{ padLeft .N .S }}`,
			Args: helper.TestArgs{
				"N": 10,
				"S": "basic",
			},
			Want: "     basic",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

func TestNow(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "default",
			Template: `{{ now }}`,
			Args:     helper.TestArgs{},
			Want:     "1997-08-29T02:14:00-04:00",
		},
		{
			Name:     "date",
			Template: `{{ now "Mon Jan _2" }}`,
			Args:     helper.TestArgs{},
			Want:     "Fri Aug 29",
		},
		{
			Name:     "time",
			Template: `{{ now "15:04:05 MST" }}`,
			Args:     helper.TestArgs{},
			Want:     "02:14:00 EDT",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

func TestBracket(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "empty",
			Template: `{{ bracket .S }}`,
			Args: helper.TestArgs{
				"S": "",
			},
			Want: "()",
		},
		{
			Name:     "word",
			Template: `{{ bracket .S }}`,
			Args: helper.TestArgs{
				"S": "foo",
			},
			Want: "(foo)",
		},
		{
			Name:     "words",
			Template: `{{ bracket .S }}`,
			Args: helper.TestArgs{
				"S": "foo bar",
			},
			Want: "(foo bar)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

func TestBracketWith(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "none",
			Template: `{{ bracketWith .B .S }}`,
			Args: helper.TestArgs{
				"B": "",
				"S": "",
			},
			Want: "",
		},
		{
			Name:     "basic",
			Template: `{{ bracketWith .B .S }}`,
			Args: helper.TestArgs{
				"B": "()",
				"S": "",
			},
			Want: "()",
		},
		{
			Name:     "word",
			Template: `{{ bracketWith .B .S }}`,
			Args: helper.TestArgs{
				"B": "<>",
				"S": "foo",
			},
			Want: "<foo>",
		},
		{
			Name:     "words",
			Template: `{{ bracketWith .B .S }}`,
			Args: helper.TestArgs{
				"B": "{{-  -}}",
				"S": "foo bar",
			},
			Want: "{{- foo bar -}}",
		},
		{
			Name:     "mismatched",
			Template: `{{ bracketWith .B .S }}`,
			Args: helper.TestArgs{
				"B": ")",
				"S": "baz",
			},
			WantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

func TestTypeName(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "nil",
			Template: `{{ typeName .Val }}`,
			Args: helper.TestArgs{
				"Val": nil,
			},
			Want: "nil",
		},
		{
			Name:     "int",
			Template: `{{ typeName .Val }}`,
			Args: helper.TestArgs{
				"Val": 3,
			},
			Want: "int",
		},
		{
			Name:     "time.Duration",
			Template: `{{ typeName .Val }}`,
			Args: helper.TestArgs{
				"Val": 10 * time.Second,
			},
			Want: "time.Duration",
		},
		{
			Name:     "*int",
			Template: `{{ typeName .Val }}`,
			Args: helper.TestArgs{
				"Val": helper.PtrToInt(10),
			},
			Want: "*int",
		},
		{
			Name:     "[]int",
			Template: `{{ typeName .Val }}`,
			Args: helper.TestArgs{
				"Val": []int{4},
			},
			Want: "[]int",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

func TestTypeKind(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "nil",
			Template: `{{ typeKind .Val }}`,
			Args: helper.TestArgs{
				"Val": nil,
			},
			Want: "nil",
		},
		{
			Name:     "int",
			Template: `{{ typeKind .Val }}`,
			Args: helper.TestArgs{
				"Val": 3,
			},
			Want: "int",
		}, {
			Name:     "string",
			Template: `{{ typeKind .Val }}`,
			Args: helper.TestArgs{
				"Val": "A",
			},
			Want: "string",
		},
		{
			Name:     "time.Duration",
			Template: `{{ typeKind .Val }}`,
			Args: helper.TestArgs{
				"Val": 10 * time.Second,
			},
			Want: "int64",
		},
		{
			Name:     "*int",
			Template: `{{ typeKind .Val }}`,
			Args: helper.TestArgs{
				"Val": helper.PtrToInt(10),
			},
			Want: "ptr",
		},
		{
			Name:     "[]int",
			Template: `{{ typeKind .Val }}`,
			Args: helper.TestArgs{
				"Val": []int{4},
			},
			Want: "slice",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

func TestSplitOn(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "one",
			Template: `{{ splitOn .Glue .S }}`,
			Args: helper.TestArgs{
				"Glue": " ",
				"S":    "one",
			},
			Want: `[one]`,
		},
		{
			Name:     "two",
			Template: `{{ splitOn .Glue .S }}`,
			Args: helper.TestArgs{
				"Glue": " ",
				"S":    "one two",
			},
			Want: `[one two]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

func TestPrefix(t *testing.T) {
	tests := []helper.TestSet{}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

func TestToColumn(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "empty",
			Template: "{{ toColumns .W .S }}",
			Args: helper.TestArgs{
				"W": 10,
				"S": "",
			},
			Want: "",
		},
		{
			Name:     "too small",
			Template: "{{ toColumns .W .S }}",
			Args: helper.TestArgs{
				"W": 10,
				"S": "foo",
			},
			Want: "foo\n",
		},
		{
			Name:     "simple",
			Template: "{{ toColumns .W .S }}",
			Args: helper.TestArgs{
				"W": 3,
				"S": "foo bar",
			},
			Want: "foo\nbar\n",
		},
		{
			Name:     "find space",
			Template: "{{ toColumns .W .S }}",
			Args: helper.TestArgs{
				"W": 4,
				"S": "foo bar",
			},
			Want: "foo\nbar\n",
		},
		{
			Name:     "long word",
			Template: "{{ toColumns .W .S }}",
			Args: helper.TestArgs{
				"W": 4,
				"S": "foobar baz",
			},
			Want: "foobar\nbaz\n",
		},
		{
			Name:     "four lines",
			Template: "{{ toColumns .W .S }}",
			Args: helper.TestArgs{
				"W": 3,
				"S": "foo bar baz snk",
			},
			Want: "foo\nbar\nbaz\nsnk\n",
		},
		{
			Name:     "possible off by one",
			Template: "{{ toColumns .W .S }}",
			Args: helper.TestArgs{
				"W": 5,
				"S": "a b c d e f g",
			},
			Want: "a b c\nd e f\ng\n",
		},
		{
			Name:     "many newlines",
			Template: "{{ toColumns .W .S }}",
			Args: helper.TestArgs{
				"W": 3,
				"S": "foo\n\n\n\nbar\n\n\n\n\n",
			},
			Want: "foo\nbar\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

func TestTitleCaseWithAbbr(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "no abbreviations",
			Template: `{{ titleCaseWithAbbr .Abbrev .Word }}`,
			Args: helper.TestArgs{
				"Abbrev": []string{},
				"Word":   "nz all blacks",
			},
			Want:    "Nz All Blacks",
			WantErr: false,
		},
		{
			Name:     "basic",
			Template: `{{ titleCaseWithAbbr .Abbrev .Word }}`,
			Args: helper.TestArgs{
				"Abbrev": []string{"nz"},
				"Word":   "nz all blacks",
			},
			Want:    "NZ All Blacks",
			WantErr: false,
		},
		{
			Name:     "in-line list",
			Template: `{{ titleCaseWithAbbr (list "nz") .Word }}`,
			Args: helper.TestArgs{
				"Word": "nz all blacks",
			},
			Want:    "NZ All Blacks",
			WantErr: false,
		},
		{
			Name:     "not a list",
			Template: `{{ titleCaseWithAbbr "nz" .Word }}`,
			Args: helper.TestArgs{
				"Word": "nz all blacks",
			},
			WantErr: true,
		},
	}

	fns := Functions()
	fns["list"] = list.List

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, fns))
	}
}

func TestQuoteSingle(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "empty",
			Template: "{{- q .S -}}",
			Args: helper.TestArgs{
				"S": "",
			},
			Want:    "''",
			WantErr: false,
		},
		{
			Name:     "basic",
			Template: "{{- q .S -}}",
			Args: helper.TestArgs{
				"S": "rawr",
			},
			Want:    "'rawr'",
			WantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

func TestQuoteDouble(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "empty",
			Template: "{{- qq .S -}}",
			Args: helper.TestArgs{
				"S": "",
			},
			Want:    "\"\"",
			WantErr: false,
		},
		{
			Name:     "basic",
			Template: "{{- qq .S -}}",
			Args: helper.TestArgs{
				"S": "rawr",
			},
			Want:    "\"rawr\"",
			WantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

func TestQuoteBack(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "empty",
			Template: "{{- qb .S -}}",
			Args: helper.TestArgs{
				"S": "",
			},
			Want:    "``",
			WantErr: false,
		},
		{
			Name:     "basic",
			Template: "{{- qb .S -}}",
			Args: helper.TestArgs{
				"S": "rawr",
			},
			Want:    "`rawr`",
			WantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}
