package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// // TestEncodingFunctions provides unit test coverage for EncodingFunctions
func TestEncodingFunctions(t *testing.T) {
	fn := EncodingFunctions()
	assert.Len(t, fn, 5, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestToJSON provides unit test coverage for ToJSON()
func TestToJSON(t *testing.T) {
	tests := []TestSet{
		{
			name:     "simple object",
			template: `{{ toJSON .val }}`,
			args: TestArgs{
				"val": map[string]string{
					"one": "foo",
					"two": "bar",
				},
			},
			want: `{"one":"foo","two":"bar"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestFormatJSON provides unit test coverage for FormatJSON()
func TestFormatJSON(t *testing.T) {
	tests := []TestSet{
		{
			name:     "simple object",
			template: `{{ formatJSON .indent .j }}`,
			args: TestArgs{
				"indent": "\t",
				"j":      `{"one":"foo","two":"bar"}`,
			},
			want: "{\n\t\"one\": \"foo\",\n\t\"two\": \"bar\"\n}",
		},
		{
			name:     "bad json",
			template: `{{ formatJSON .indent .j }}`,
			args: TestArgs{
				"indent": "\t",
				"j":      `{"one":"foo","two":"forgot end brace..."`,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestToYAML provides unit test coverage for ToYAML()
func TestToYAML(t *testing.T) {
	tests := []TestSet{
		{
			name:     "simple object",
			template: `{{ toYAML .val }}`,
			args: TestArgs{
				"val": map[string]string{
					"one": "foo",
					"two": "bar",
				},
			},
			want: "one: foo\ntwo: bar\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestToBase64 provides unit test coverage for ToBase64()
func TestToBase64(t *testing.T) {
	tests := []TestSet{
		{
			name:     "empty",
			template: `{{ toBase64 .s }}`,
			args: TestArgs{
				"s": "",
			},
			want: "",
		},
		{
			name:     "simple",
			template: `{{ toBase64 .s }}`,
			args: TestArgs{
				"s": "A basic string",
			},
			want: "QSBiYXNpYyBzdHJpbmc=",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}

// TestFromBase64 provides unit test coverage for FromBase64()
func TestFromBase64(t *testing.T) {
	tests := []TestSet{
		{
			name:     "empty",
			template: `{{ fromBase64 .s }}`,
			args: TestArgs{
				"s": "",
			},
			want: "",
		},
		{
			name:     "simple",
			template: `{{ fromBase64 .s }}`,
			args: TestArgs{
				"s": "QSBiYXNpYyBzdHJpbmc=",
			},
			want: "A basic string",
		},
		{
			name:     "bad string",
			template: `{{ fromBase64 .s }}`,
			args: TestArgs{
				"s": "QSBiYXNpYyBzdHapbmc",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, TemplateTest(tt))
	}
}
