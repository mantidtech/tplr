package wordcase

import (
	"strings"
)

// Tokens represents a string broken up by boundaries defined by IsRuneSeparator functions
type Tokens []string

// String returns all tokens separated by a space as a string
func (t Tokens) String() string {
	return strings.Join(t, " ")
}

// Format applies to the given function to the tokens specified by the 'indexes' list
func (t Tokens) Format(fn Formatter, items TokenSelector) Tokens {
	var r Tokens
	idx := map[int]bool{}
	for _, i := range items(t) {
		idx[i] = true
	}
	for i, x := range t {
		if idx[i] {
			r = append(r, fn(x))
		} else {
			r = append(r, x)
		}
	}
	return r
}

// FormatAll applies to the given function to all tokens.
//
//	convenience function that does the same as t.Format(fn, ToAll(t)) (just faster and less verbosely)
func (t Tokens) FormatAll(fn Formatter) Tokens {
	var r Tokens
	for _, x := range t {
		r = append(r, fn(x))
	}
	return r
}

// Tokenize applies the given IsRuneSeparator function to each token, returning a new token set
func (t Tokens) Tokenize(test SeparatorTest, isSeparator IsRuneSeparator, rmSep bool) Tokens {
	var r Tokens
	for _, x := range t {
		r = append(r, TokenizeString(x, test, isSeparator, rmSep)...)
	}
	return r
}

// TokenizeString breaks a string into a set of tokens using the rules supplied
func TokenizeString(s string, test SeparatorTest, sepRune IsRuneSeparator, rmSep bool) Tokens {
	var b strings.Builder
	res := Tokens{}

	r := []rune(s)
	for i, n := range r {
		isSep := test(s, i, sepRune)
		if isSep && b.Len() > 0 {
			res = append(res, b.String())
			b.Reset()
		}
		if !rmSep || !sepRune(n) {
			b.WriteRune(n)
		}
	}
	if b.Len() > 0 {
		res = append(res, b.String())
	}
	return res
}
