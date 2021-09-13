package console

import (
	"testing"

	"github.com/mantidtech/tplr/functions/helper"
	"github.com/stretchr/testify/assert"
)

// TestMiscellaneousFunctions provides unit test coverage for MiscellaneousFunctions
func TestMiscellaneousFunctions(t *testing.T) {
	fn := Functions()
	assert.Len(t, fn, 1, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestTerminalWidth provides unit test coverage for TerminalWidth()
func TestTerminalWidth(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "basic",
			Template: "{{ terminalWidth }}",
			Want:     "0", // probably - depends on how/where the test is run
			WantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}
