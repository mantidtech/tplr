package functions

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"text/template"

	"gopkg.in/yaml.v2"
)

// EncodingFunctions for encoding and decoding various formats
func EncodingFunctions() template.FuncMap {
	return template.FuncMap{
		"formatJSON": FormatJSON,
		"fromBase64": FromBase64,
		"toBase64":   ToBase64,
		"toJSON":     ToJSON,
		"toYAML":     ToYAML,
	}
}

// ToJSON returns the given value as a json string
func ToJSON(val interface{}) (string, error) {
	b, err := json.Marshal(val)
	return string(b), err
}

// FormatJSON returns the given json string, formatted with the given indent string
func FormatJSON(indent string, j string) (string, error) {
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(j), "", indent)
	if err != nil {
		return "", fmt.Errorf("failed to format json string %s: %s", j, err.Error())
	}
	return buf.String(), nil
}

// ToYAML returns the given value as a yaml string
func ToYAML(val interface{}) (string, error) {
	b, err := yaml.Marshal(val)
	return string(b), err
}

// ToBase64 converts the given string to a base64 encoding
func ToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// FromBase64 decodes the given encoded string to plain
func FromBase64(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	return string(decoded), err
}
