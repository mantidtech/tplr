package tplr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestVersion provides somewhat cynical but otherwise complete unit test coverage for Version()
func TestVersion(t *testing.T) {
	want := "unknown"
	got := Version()
	assert.Equal(t, want, got)
}
