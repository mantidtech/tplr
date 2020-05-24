package tplr

import (
	"strings"
)

var version = "unknown"
var build = "unk"

// Version returns the version number of this app
func Version() string {
	s := strings.Builder{}
	s.WriteString(version)
	s.WriteString("+b")
	s.WriteString(build)
	return s.String()
}
