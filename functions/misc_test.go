package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMiscellaneousFunctions provides unit test coverage for MiscellaneousFunctions
func TestMiscellaneousFunctions(t *testing.T) {
	fn := MiscellaneousFunctions()
	assert.Len(t, fn, 1, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestTerminalWidth provides unit test coverage for TerminalWidth()
func TestTerminalWidth(t *testing.T) {
	RunTemplateTest(t, []TestSet{
		{
			name:     "basic",
			template: "{{ terminalWidth }}",
			want:     "0", // probably - depends on how/where the test is run
			wantErr:  false,
		},
	})
}
