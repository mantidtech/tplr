package encoding

import "encoding/base64"

// ToBase64 converts the given string to a base64 encoding
func ToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// FromBase64 decodes the given encoded string to plain
func FromBase64(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	return string(decoded), err
}
