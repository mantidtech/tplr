package functions

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMiscellaneousFunctions provides unit test coverage for MiscellaneousFunctions
func TestMiscellaneousFunctions(t *testing.T) {
	fn := MiscellaneousFunctions()
	assert.Len(t, fn, 1, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestTerminalWidth provides unit test coverage for TerminalWidth()
func TestTerminalWidth(t *testing.T) {
	type Args struct {
	}

	tests := []struct {
		name     string
		template string
		args     Args
		want     string
		wantErr  bool
	}{
		{
			name:     "basic",
			template: "{{ terminalWidth }}",
			want:     "0", // probably - depends how/where the test is run
			wantErr:  false,
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
