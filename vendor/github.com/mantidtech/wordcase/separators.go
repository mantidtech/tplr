package wordcase

import (
	"unicode"
)

// IsRuneSeparator determines if the given rune is a token separator for the string it's being applied to
type IsRuneSeparator func(rune) bool

// NotLowerOrDigit returns true if the given rune is neither a lower case letter or a digit
func NotLowerOrDigit(r rune) bool {
	return !unicode.IsLower(r) && !unicode.IsDigit(r)
}

// NotLetterOrDigit returns true if the given rune is neither a letter or a digit
func NotLetterOrDigit(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsDigit(r)
}

// SeparatorTest determines if the given index in the given string is a token separator.
//
//	By having the whole parse text as reference, smarter tests can be performed using lookahead/lookbehind along with the IsRuneSeparator test.
type SeparatorTest func(word string, idx int, test IsRuneSeparator) bool

// LookAroundCategorizer considers a rune to be a separator if it passes test(c),
//
//	but also if both the preceding and succeeding runes are NOT separators
func LookAroundCategorizer(word string, idx int, test IsRuneSeparator) bool {
	w := []rune(word)

	p := true
	if idx > 0 {
		p = test(w[idx-1])
	}

	n := true
	if idx < len(w)-1 {
		n = test(w[idx+1])
	}
	return test(w[idx]) && (!p || !n)
}

// SimpleCategorizer considers only the current rune when deciding on separators.
// Note: no bounds checking because the Tokenize method doesn't do out-of-bounds queries
func SimpleCategorizer(word string, idx int, test IsRuneSeparator) bool {
	w := []rune(word)
	return test(w[idx])
}
