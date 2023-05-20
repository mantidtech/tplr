package math

import (
	"testing"

	"github.com/mantidtech/tplr/functions/helper"
	"github.com/stretchr/testify/assert"
)

// TestFunctions provides unit test coverage for Functions.
func TestFunctions(t *testing.T) {
	fn := Functions()
	assert.Len(t, fn, 5, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestAdd provides unit test coverage for Add.
func TestAdd(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "zero",
			Template: "{{- add -}}",
			Args:     helper.TestArgs{},
			Want:     "0",
			WantErr:  false,
		},
		{
			Name:     "one",
			Template: "{{- add .A -}}",
			Args: helper.TestArgs{
				"A": 1.0,
			},
			Want:    "1",
			WantErr: false,
		},
		{
			Name:     "three",
			Template: "{{- add .A .B .C -}}",
			Args: helper.TestArgs{
				"A": 1.0,
				"B": 2.0,
				"C": 3.0,
			},
			Want:    "6",
			WantErr: false,
		},
		{
			Name:     "int",
			Template: "{{- add .A .B .C -}}",
			Args: helper.TestArgs{
				"A": 1,
				"B": 2,
				"C": 3,
			},
			Want:    "6",
			WantErr: false,
		},
		{
			Name:     "mix",
			Template: "{{- add .A .B .C -}}",
			Args: helper.TestArgs{
				"A": 1.5,
				"B": -2,
				"C": 3,
			},
			Want:    "2.5",
			WantErr: false,
		},
		{
			Name:     "rune",
			Template: "{{- add .A -}}",
			Args: helper.TestArgs{
				"A": 'a',
			},
			Want:    "97",
			WantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestSubtract provides unit test coverage for Subtract.
func TestSubtract(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "zero",
			Template: "{{- sub -}}",
			Args:     helper.TestArgs{},
			Want:     "0",
			WantErr:  false,
		},
		{
			Name:     "one",
			Template: "{{- sub .A -}}",
			Args: helper.TestArgs{
				"A": 1.0,
			},
			Want:    "1",
			WantErr: false,
		},
		{
			Name:     "three",
			Template: "{{- sub .A .B .C -}}",
			Args: helper.TestArgs{
				"A": 1.0,
				"B": 2.0,
				"C": 3.0,
			},
			Want:    "-4",
			WantErr: false,
		},
		{
			Name:     "int",
			Template: "{{- sub .A .B .C -}}",
			Args: helper.TestArgs{
				"A": 1,
				"B": 2,
				"C": 3,
			},
			Want:    "-4",
			WantErr: false,
		},
		{
			Name:     "mix",
			Template: "{{- sub .A .B .C -}}",
			Args: helper.TestArgs{
				"A": 1.5,
				"B": -2,
				"C": 3,
			},
			Want:    "0.5",
			WantErr: false,
		},
		{
			Name:     "rune",
			Template: "{{- sub .A -}}",
			Args: helper.TestArgs{
				"A": 'a',
			},
			Want:    "97",
			WantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestMultiply provides unit test coverage for Multiply.
func TestMultiply(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "zero",
			Template: "{{- mult -}}",
			Args:     helper.TestArgs{},
			Want:     "0",
			WantErr:  false,
		},
		{
			Name:     "one",
			Template: "{{- mult .A -}}",
			Args: helper.TestArgs{
				"A": 1.0,
			},
			Want:    "1",
			WantErr: false,
		},
		{
			Name:     "three",
			Template: "{{- mult .A .B .C -}}",
			Args: helper.TestArgs{
				"A": 1.0,
				"B": 2.0,
				"C": 3.0,
			},
			Want:    "6",
			WantErr: false,
		},
		{
			Name:     "int",
			Template: "{{- mult .A .B .C -}}",
			Args: helper.TestArgs{
				"A": 1,
				"B": 2,
				"C": 3,
			},
			Want:    "6",
			WantErr: false,
		},
		{
			Name:     "mix",
			Template: "{{- mult .A .B .C -}}",
			Args: helper.TestArgs{
				"A": 1.5,
				"B": -2,
				"C": 3,
			},
			Want:    "-9",
			WantErr: false,
		},
		{
			Name:     "rune",
			Template: "{{- mult .A -}}",
			Args: helper.TestArgs{
				"A": 'a',
			},
			Want:    "97",
			WantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestDivide provides unit test coverage for Divide.
func TestDivide(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "zero",
			Template: "{{- div -}}",
			Args:     helper.TestArgs{},
			Want:     "0",
			WantErr:  false,
		},
		{
			Name:     "one",
			Template: "{{- div .A -}}",
			Args: helper.TestArgs{
				"A": 1.0,
			},
			Want:    "1",
			WantErr: false,
		},
		{
			Name:     "three",
			Template: "{{- div .A .B .C -}}",
			Args: helper.TestArgs{
				"A": 1.0,
				"B": 2.0,
				"C": 4.0,
			},
			Want:    "0.125",
			WantErr: false,
		},
		{
			Name:     "int",
			Template: "{{- div .A .B .C -}}",
			Args: helper.TestArgs{
				"A": 1,
				"B": 2,
				"C": 4,
			},
			Want:    "0.125",
			WantErr: false,
		},
		{
			Name:     "mix",
			Template: "{{- div .A .B .C -}}",
			Args: helper.TestArgs{
				"A": 4.0,
				"B": -2,
				"C": 4,
			},
			Want:    "-0.5",
			WantErr: false,
		},
		{
			Name:     "rune",
			Template: "{{- div .A -}}",
			Args: helper.TestArgs{
				"A": 'a',
			},
			Want:    "97",
			WantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestAbs provides unit test coverage for Abs.
func TestAbs(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "one",
			Template: "{{- abs .A -}}",
			Args: helper.TestArgs{
				"A": 1.0,
			},
			Want:    "1",
			WantErr: false,
		},
		{
			Name:     "negative 1.0",
			Template: "{{- abs .A -}}",
			Args: helper.TestArgs{
				"A": -1.0,
			},
			Want:    "1",
			WantErr: false,
		},
		{
			Name:     "int",
			Template: "{{- abs .A -}}",
			Args: helper.TestArgs{
				"A": 1,
			},
			Want:    "1",
			WantErr: false,
		},
		{
			Name:     "negative int",
			Template: "{{- abs .A -}}",
			Args: helper.TestArgs{
				"A": -4,
			},
			Want:    "4",
			WantErr: false,
		},
		{
			Name:     "rune",
			Template: "{{- abs .A -}}",
			Args: helper.TestArgs{
				"A": 'a',
			},
			Want:    "97",
			WantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}
