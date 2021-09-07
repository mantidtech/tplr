package functions

import (
	"testing"
	"time"

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
	fn := StringFunctions()
	assert.Len(t, fn, 34, "weakly ensuring functions haven't been added/removed without updating tests")
}

func TestUppercaseFirst(t *testing.T) {
	tests := []TestSet{
		{
			name:     "empty",
			template: `{{ ucFirst .S }}`,
			args: TestArgs{
				"S": "",
			},
			want: "",
		},
		{
			name:     "simple",
			template: `{{ ucFirst .S }}`,
			args: TestArgs{
				"S": "simple",
			},
			want: "Simple",
		},
		{
			name:     "same",
			template: `{{ ucFirst .S }}`,
			args: TestArgs{
				"S": "Same",
			},
			want: "Same",
		},
		{
			name:     "number",
			template: `{{ ucFirst .S }}`,
			args: TestArgs{
				"S": "3rd",
			},
			want: "3rd",
		},
		{
			name:     "multiple words",
			template: `{{ ucFirst .S }}`,
			args: TestArgs{
				"S": "spam test",
			},
			want: "Spam test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestNewline(t *testing.T) {
	tests := []TestSet{
		{
			name:     "no params",
			template: `{{ nl }}`,
			args:     TestArgs{},
			want:     "\n",
		},
		{
			name:     "params",
			template: `{{ nl .C }}`,
			args: TestArgs{
				"C": 3,
			},
			want: "\n\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestRep provides unit test coverage for Rep()
func TestRep(t *testing.T) {
	tests := []TestSet{
		{
			name:     "none",
			template: `{{ rep .N .S }}`,
			args: TestArgs{
				"N": 0,
				"S": "foo",
			},
			want: "",
		},
		{
			name:     "repeated empty",
			template: `{{ rep .N .S }}`,
			args: TestArgs{
				"N": 2,
				"S": "",
			},
			want: "",
		},
		{
			name:     "one",
			template: `{{ rep .N .S }}`,
			args: TestArgs{
				"N": 1,
				"S": "x",
			},
			want: "x",
		},
		{
			name:     "two",
			template: `{{ rep .N .S }}`,
			args: TestArgs{
				"N": 2,
				"S": "foo",
			},
			want: "foofoo",
		},
		{
			name:     "multiple args",
			template: `{{ rep .N .S .S2 }}`,
			args: TestArgs{
				"N":  1,
				"S":  "one",
				"S2": "two",
			},
			want: "one two",
		},
		{
			name:     "multiple args, twice",
			template: `{{ rep .N .S .S2 }}`,
			args: TestArgs{
				"N":  2,
				"S":  "one",
				"S2": "two",
			},
			want: "one twoone two",
		},
		{
			name:     "negative one",
			template: `{{ rep .N .S }}`,
			args: TestArgs{
				"N": -1,
				"S": "foo",
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestIndent provides unit test coverage for Indent()
func TestIndent(t *testing.T) {
	tests := []TestSet{
		{
			name:     "none",
			template: `{{ indent .T .Content }}`,
			args: TestArgs{
				"T":       0,
				"Content": "foo",
			},
			want: "foo",
		},
		{
			name:     "one",
			template: `{{ indent .T .Content }}`,
			args: TestArgs{
				"T":       1,
				"Content": "foo",
			},
			want: " foo",
		},
		{
			name:     "two",
			template: `{{ indent .T .Content }}`,
			args: TestArgs{
				"T":       2,
				"Content": "foo",
			},
			want: "  foo",
		},
		{
			name:     "negative one",
			template: `{{ indent .T .Content }}`,
			args: TestArgs{
				"T":       -1,
				"Content": "",
			},
			want: "",
		},
		{
			name:     "multi line",
			template: `{{ indent .T .Content }}`,
			args: TestArgs{
				"T":       1,
				"Content": "foo\nbar",
			},
			want: " foo\n bar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestIndent provides unit test coverage for Indent()
func TestUnindent(t *testing.T) {
	tests := []TestSet{
		{
			name:     "none",
			template: `{{ unindent .T .Content }}`,
			args: TestArgs{
				"T":       0,
				"Content": "foo",
			},
			want: "foo",
		},
		{
			name:     "one",
			template: `{{ unindent .T .Content }}`,
			args: TestArgs{
				"T":       1,
				"Content": "  foo",
			},
			want: " foo",
		},
		{
			name:     "two",
			template: `{{ unindent .T .Content }}`,
			args: TestArgs{
				"T":       2,
				"Content": "  foo",
			},
			want: "foo",
		},
		{
			name:     "negative one",
			template: `{{ unindent .T .Content }}`,
			args: TestArgs{
				"T":       -1,
				"Content": "",
			},
			wantErr: true,
		},
		{
			name:     "multi line",
			template: `{{ unindent .T .Content }}`,
			args: TestArgs{
				"T":       1,
				"Content": "  foo\n  bar",
			},
			want: " foo\n bar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestSuffix provides unit test coverage for Suffix()
func TestSuffix(t *testing.T) {
	tests := []TestSet{
		{
			name:     "none",
			template: `{{ suffix .Suffix .T .Content }}`,
			args: TestArgs{
				"Suffix":  "",
				"T":       0,
				"Content": "foo",
			},
			want: "foo",
		},
		{
			name:     "one",
			template: `{{ suffix .Suffix .T .Content }}`,
			args: TestArgs{
				"Suffix":  "X",
				"T":       1,
				"Content": "foo",
			},
			want: "fooX",
		},
		{
			name:     "two",
			template: `{{ suffix .Suffix .T .Content }}`,
			args: TestArgs{
				"Suffix":  "X",
				"T":       2,
				"Content": "foo",
			},
			want: "fooXX",
		},
		{
			name:     "negative one",
			template: `{{ suffix .Suffix .T .Content }}`,
			args: TestArgs{
				"Suffix":  "X",
				"T":       -1,
				"Content": "",
			},
			want: "",
		},
		{
			name:     "multi line",
			template: `{{ suffix .Suffix .T .Content }}`,
			args: TestArgs{
				"Suffix":  "X",
				"T":       1,
				"Content": "foo\nbar",
			},
			want: "fooX\nbarX",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestSpace(t *testing.T) {
	tests := []TestSet{
		{
			name:     "none",
			template: `{{ sp .N }}`,
			args: TestArgs{
				"N": 0,
			},
			want: "",
		},
		{
			name:     "one",
			template: `{{ sp .N }}`,
			args: TestArgs{
				"N": 1,
			},
			want: " ",
		},
		{
			name:     "two",
			template: `{{ sp .N }}`,
			args: TestArgs{
				"N": 2,
			},
			want: "  ",
		},
		{
			name:     "negative one",
			template: `{{ sp .N }}`,
			args: TestArgs{
				"N": -1,
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestTab(t *testing.T) {
	tests := []TestSet{
		{
			name:     "none",
			template: `{{ tab .N }}`,
			args: TestArgs{
				"N": 0,
			},
			want: "",
		},
		{
			name:     "one",
			template: `{{ tab .N }}`,
			args: TestArgs{
				"N": 1,
			},
			want: "\t",
		},
		{
			name:     "two",
			template: `{{ tab .N }}`,
			args: TestArgs{
				"N": 2,
			},
			want: "\t\t",
		},
		{
			name:     "negative one",
			template: `{{ tab .N }}`,
			args: TestArgs{
				"N": -1,
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestPadRight(t *testing.T) {
	tests := []TestSet{
		{
			name:     "basic",
			template: `{{ padRight .N .S }}`,
			args: TestArgs{
				"N": 10,
				"S": "basic",
			},
			want: "basic     ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestPadLeft(t *testing.T) {
	tests := []TestSet{
		{
			name:     "basic",
			template: `{{ padLeft .N .S }}`,
			args: TestArgs{
				"N": 10,
				"S": "basic",
			},
			want: "     basic",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestNow(t *testing.T) {
	tests := []TestSet{
		{
			name:     "basic",
			template: `{{ now }}`,
			args:     TestArgs{},
			want:     "1997-08-29T02:14:00-04:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestBracket(t *testing.T) {
	tests := []TestSet{
		{
			name:     "empty",
			template: `{{ bracket .S }}`,
			args: TestArgs{
				"S": "",
			},
			want: "()",
		},
		{
			name:     "word",
			template: `{{ bracket .S }}`,
			args: TestArgs{
				"S": "foo",
			},
			want: "(foo)",
		},
		{
			name:     "words",
			template: `{{ bracket .S }}`,
			args: TestArgs{
				"S": "foo bar",
			},
			want: "(foo bar)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestBracketWith(t *testing.T) {
	tests := []TestSet{
		{
			name:     "none",
			template: `{{ bracketWith .B .S }}`,
			args: TestArgs{
				"B": "",
				"S": "",
			},
			want: "",
		},
		{
			name:     "basic",
			template: `{{ bracketWith .B .S }}`,
			args: TestArgs{
				"B": "()",
				"S": "",
			},
			want: "()",
		},
		{
			name:     "word",
			template: `{{ bracketWith .B .S }}`,
			args: TestArgs{
				"B": "<>",
				"S": "foo",
			},
			want: "<foo>",
		},
		{
			name:     "words",
			template: `{{ bracketWith .B .S }}`,
			args: TestArgs{
				"B": "{{-  -}}",
				"S": "foo bar",
			},
			want: "{{- foo bar -}}",
		},
		{
			name:     "mismatched",
			template: `{{ bracketWith .B .S }}`,
			args: TestArgs{
				"B": ")",
				"S": "baz",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestTypeName(t *testing.T) {
	tests := []TestSet{
		{
			name:     "nil",
			template: `{{ typeName .Val }}`,
			args: TestArgs{
				"Val": nil,
			},
			want: "nil",
		},
		{
			name:     "int",
			template: `{{ typeName .Val }}`,
			args: TestArgs{
				"Val": 3,
			},
			want: "int",
		},
		{
			name:     "time.Duration",
			template: `{{ typeName .Val }}`,
			args: TestArgs{
				"Val": 10 * time.Second,
			},
			want: "time.Duration",
		},
		{
			name:     "*int",
			template: `{{ typeName .Val }}`,
			args: TestArgs{
				"Val": helperPtrToInt(10),
			},
			want: "*int",
		},
		{
			name:     "[]int",
			template: `{{ typeName .Val }}`,
			args: TestArgs{
				"Val": []int{4},
			},
			want: "[]int",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestTypeKind(t *testing.T) {
	tests := []TestSet{
		{
			name:     "nil",
			template: `{{ typeKind .Val }}`,
			args: TestArgs{
				"Val": nil,
			},
			want: "nil",
		},
		{
			name:     "int",
			template: `{{ typeKind .Val }}`,
			args: TestArgs{
				"Val": 3,
			},
			want: "int",
		},
		{
			name:     "time.Duration",
			template: `{{ typeKind .Val }}`,
			args: TestArgs{
				"Val": 10 * time.Second,
			},
			want: "int64",
		},
		{
			name:     "*int",
			template: `{{ typeKind .Val }}`,
			args: TestArgs{
				"Val": helperPtrToInt(10),
			},
			want: "ptr",
		},
		{
			name:     "[]int",
			template: `{{ typeKind .Val }}`,
			args: TestArgs{
				"Val": []int{4},
			},
			want: "slice",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestSplitOn(t *testing.T) {
	tests := []TestSet{
		{
			name:     "one",
			template: `{{ splitOn .Glue .S | toJSON }}`,
			args: TestArgs{
				"Glue": " ",
				"S":    "one",
			},
			want: `["one"]`,
		},
		{
			name:     "two",
			template: `{{ splitOn .Glue .S | toJSON }}`,
			args: TestArgs{
				"Glue": " ",
				"S":    "one two",
			},
			want: `["one","two"]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestPrefix(t *testing.T) {
	tests := []TestSet{}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestToColumn(t *testing.T) {
	tests := []TestSet{
		{
			name:     "empty",
			template: "{{ toColumns .W .S }}",
			args: TestArgs{
				"W": 10,
				"S": "",
			},
			want: "",
		},
		{
			name:     "too small",
			template: "{{ toColumns .W .S }}",
			args: TestArgs{
				"W": 10,
				"S": "foo",
			},
			want: "foo\n",
		},
		{
			name:     "simple",
			template: "{{ toColumns .W .S }}",
			args: TestArgs{
				"W": 3,
				"S": "foo bar",
			},
			want: "foo\nbar\n",
		},
		{
			name:     "find space",
			template: "{{ toColumns .W .S }}",
			args: TestArgs{
				"W": 4,
				"S": "foo bar",
			},
			want: "foo\nbar\n",
		},
		{
			name:     "long word",
			template: "{{ toColumns .W .S }}",
			args: TestArgs{
				"W": 4,
				"S": "foobar baz",
			},
			want: "foobar\nbaz\n",
		},
		{
			name:     "four lines",
			template: "{{ toColumns .W .S }}",
			args: TestArgs{
				"W": 3,
				"S": "foo bar baz snk",
			},
			want: "foo\nbar\nbaz\nsnk\n",
		},
		{
			name:     "possible off by one",
			template: "{{ toColumns .W .S }}",
			args: TestArgs{
				"W": 5,
				"S": "a b c d e f g",
			},
			want: "a b c\nd e f\ng\n",
		},
		{
			name:     "many newlines",
			template: "{{ toColumns .W .S }}",
			args: TestArgs{
				"W": 3,
				"S": "foo\n\n\n\nbar\n\n\n\n\n",
			},
			want: "foo\nbar\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestTitleCaseWithAbbr(t *testing.T) {
	tests := []TestSet{
		{
			name:     "no abbreviations",
			template: `{{ titleCaseWithAbbr .Abbrev .Word }}`,
			args: TestArgs{
				"Abbrev": []string{},
				"Word":   "nz all blacks",
			},
			want:    "Nz All Blacks",
			wantErr: false,
		},
		{
			name:     "basic",
			template: `{{ titleCaseWithAbbr .Abbrev .Word }}`,
			args: TestArgs{
				"Abbrev": []string{"nz"},
				"Word":   "nz all blacks",
			},
			want:    "NZ All Blacks",
			wantErr: false,
		},
		{
			name:     "in-line list",
			template: `{{ titleCaseWithAbbr (list "nz") .Word }}`,
			args: TestArgs{
				"Word": "nz all blacks",
			},
			want:    "NZ All Blacks",
			wantErr: false,
		},
		{
			name:     "not a list",
			template: `{{ titleCaseWithAbbr "nz" .Word }}`,
			args: TestArgs{
				"Word": "nz all blacks",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestQuoteSingle(t *testing.T) {
	tests := []TestSet{
		{
			name:     "empty",
			template: "{{- q .S -}}",
			args: TestArgs{
				"S": "",
			},
			want:    "''",
			wantErr: false,
		},
		{
			name:     "basic",
			template: "{{- q .S -}}",
			args: TestArgs{
				"S": "rawr",
			},
			want:    "'rawr'",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestQuoteDouble(t *testing.T) {
	tests := []TestSet{
		{
			name:     "empty",
			template: "{{- qq .S -}}",
			args: TestArgs{
				"S": "",
			},
			want:    "\"\"",
			wantErr: false,
		},
		{
			name:     "basic",
			template: "{{- qq .S -}}",
			args: TestArgs{
				"S": "rawr",
			},
			want:    "\"rawr\"",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

func TestQuoteBack(t *testing.T) {
	tests := []TestSet{
		{
			name:     "empty",
			template: "{{- bq .S -}}",
			args: TestArgs{
				"S": "",
			},
			want:    "``",
			wantErr: false,
		},
		{
			name:     "basic",
			template: "{{- bq .S -}}",
			args: TestArgs{
				"S": "rawr",
			},
			want:    "`rawr`",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}
