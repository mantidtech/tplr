package encoding

import (
	"testing"

	"github.com/mantidtech/tplr/functions/helper"
	"github.com/stretchr/testify/assert"
)

// // TestEncodingFunctions provides unit test coverage for EncodingFunctions
func TestEncodingFunctions(t *testing.T) {
	fn := Functions()
	assert.Len(t, fn, 5, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestToJSON provides unit test coverage for ToJSON()
func TestToJSON(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "simple object",
			Template: `{{ toJSON .object }}`,
			Args: helper.TestArgs{
				"object": map[string]string{
					"one": "foo",
					"two": "bar",
				},
			},
			Want: `{"one":"foo","two":"bar"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestFormatJSON provides unit test coverage for FormatJSON()
func TestFormatJSON(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "simple object",
			Template: `{{ formatJSON .indent .json }}`,
			Args: helper.TestArgs{
				"indent": "\t",
				"json":   `{"one":"foo","two":"bar"}`,
			},
			Want: "{\n\t\"one\": \"foo\",\n\t\"two\": \"bar\"\n}",
		},
		{
			Name:     "bad json",
			Template: `{{ formatJSON .indent .json }}`,
			Args: helper.TestArgs{
				"indent": "\t",
				"json":   `{"one":"foo","two":"forgot end brace..."`,
			},
			WantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestToYAML provides unit test coverage for ToYAML()
func TestToYAML(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "simple object",
			Template: `{{ toYAML .object }}`,
			Args: helper.TestArgs{
				"object": map[string]string{
					"one": "foo",
					"two": "bar",
				},
			},
			Want: "one: foo\ntwo: bar\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestToBase64 provides unit test coverage for ToBase64()
func TestToBase64(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "empty",
			Template: `{{ toBase64 .string }}`,
			Args: helper.TestArgs{
				"string": "",
			},
			Want: "",
		},
		{
			Name:     "simple",
			Template: `{{ toBase64 .string }}`,
			Args: helper.TestArgs{
				"string": "A basic string",
			},
			Want: "QSBiYXNpYyBzdHJpbmc=",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestFromBase64 provides unit test coverage for FromBase64()
func TestFromBase64(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "empty",
			Template: `{{ fromBase64 .string }}`,
			Args: helper.TestArgs{
				"string": "",
			},
			Want: "",
		},
		{
			Name:     "simple",
			Template: `{{ fromBase64 .string }}`,
			Args: helper.TestArgs{
				"string": "QSBiYXNpYyBzdHJpbmc=",
			},
			Want: "A basic string",
		},
		{
			Name:     "bad string",
			Template: `{{ fromBase64 .string }}`,
			Args: helper.TestArgs{
				"string": "QSBiYXNpYyBzdHapbmc",
			},
			WantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}
