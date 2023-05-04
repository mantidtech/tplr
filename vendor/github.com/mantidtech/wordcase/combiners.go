package wordcase

import (
	"strings"
)

// Combiner is function that joins tokens together to create the final output
type Combiner func(string) string

// Join concatenates tokens into a string with the given separator
func (t Tokens) Join(sep string) string {
	return strings.Join(t, sep)
}
