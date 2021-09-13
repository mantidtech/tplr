package encoding

import (
	"bytes"
	"encoding/json"
	"fmt"
)

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
