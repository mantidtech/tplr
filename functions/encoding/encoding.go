// Package encoding provides template methods related to various data encodings
package encoding

import (
	"text/template"
)

// Functions for encoding and decoding various formats
func Functions() template.FuncMap {
	return template.FuncMap{
		"formatJSON": FormatJSON,
		"fromBase64": FromBase64,
		"toBase64":   ToBase64,
		"toJSON":     ToJSON,
		"toYAML":     ToYAML,
	}
}
