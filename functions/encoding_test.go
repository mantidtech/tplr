package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// // TestEncodingFunctions provides unit test coverage for EncodingFunctions
func TestEncodingFunctions(t *testing.T) {
	fn := EncodingFunctions()
	assert.Len(t, fn, 5, "weakly ensuring functions haven't been added/removed without updating tests")
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

// TestToBase64 provides unit test coverage for ToBase64()
func TestToBase64(t *testing.T) {
	t.Parallel()
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
				s: "A basic string",
			},
			want: "QSBiYXNpYyBzdHJpbmc=",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := ToBase64(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestFromBase64 provides unit test coverage for FromBase64()
func TestFromBase64(t *testing.T) {
	t.Parallel()
	type Args struct {
		s string
	}

	tests := []struct {
		name       string
		args       Args
		wantString string
		wantError  bool
	}{
		{
			name: "empty",
			args: Args{
				s: "",
			},
			wantString: "",
		},
		{
			name: "simple",
			args: Args{
				s: "QSBiYXNpYyBzdHJpbmc=",
			},
			wantString: "A basic string",
		},
		{
			name: "bad string",
			args: Args{
				s: "QSBiYXNpYyBzdHapbmc",
			},
			wantError: true,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotString, gotError := FromBase64(tt.args.s)
			if tt.wantError {
				require.Error(t, gotError)
				return
			}

			require.NoError(t, gotError)
			assert.Equal(t, tt.wantString, gotString)
		})
	}
}
