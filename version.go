package tplr

import (
	"strings"
)

var version = "unknown"

// Version returns the version number of this app
func Version() string {
	s := strings.Builder{}
	s.WriteString(version)
	return s.String()
}
