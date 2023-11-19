package wordcase

import (
	"sort"
	"strings"
)

// TokenSelector is a function given a set of tokens, returns an array of
//
//	indices to the token matching a condition defined in the function
type TokenSelector func(t Tokens) []int

// ToFirst returns the index of the first token
func ToFirst(_ Tokens) []int {
	return []int{0}
}

// ToLast returns the index of the last token
func ToLast(t Tokens) []int {
	return []int{len(t) - 1} // Note negative indices are ok. The formatter doesn't use them, so bad addressing is avoided
}

// ToRest returns the indices of all but the first token
func ToRest(t Tokens) []int {
	l := len(t) - 1
	var ret []int
	for i := 0; i < l; i++ {
		ret = append(ret, i+1)
	}
	return ret
}

// ToAll returns the indices of all tokens
func ToAll(t Tokens) []int {
	ret := make([]int, len(t))
	for i := range t {
		ret[i] = i
	}
	return ret
}

// KeyWordFn returns a new selector that matches tokens based on the given set of keywords.
// The keywords are assumed to be lowercase.
func KeyWordFn(keywords []string) TokenSelector {
	kw := make(map[string]struct{})
	for _, w := range keywords {
		kw[w] = empty
	}
	return func(t Tokens) []int {
		var ret []int
		for i, s := range t {
			l := strings.ToLower(s)
			if _, isKeyword := kw[l]; isKeyword {
				ret = append(ret, i)
			}
		}
		return ret
	}
}

// KeyWords returns the indices of tokens that match a set of keywords
var KeyWords = KeyWordFn(UsefulKeyWords)

// LintWords returns the indices of tokens that matches keywords used by go lint
var LintWords = KeyWordFn(GoLintKeywords)

// Not inverts the given selector's matches
func Not(sFn TokenSelector) TokenSelector {
	return func(t Tokens) []int {
		m := sFn(t)
		var ret []int

		at := 0
		for _, i := range m {
			for x := at; x < i; x++ {
				ret = append(ret, x)
			}
			at = i + 1
		}
		for x := at; x < len(t); x++ {
			ret = append(ret, x)
		}

		return ret
	}
}

// And returns a selector that matches tokens that are matched by both a & b
func And(a, b TokenSelector) TokenSelector {
	return func(t Tokens) []int {
		one := a(t)
		var ret []int

		two := make(map[int]bool)
		for _, v := range b(t) {
			two[v] = true
		}

		for _, i := range one {
			if two[i] {
				ret = append(ret, i)
			}
		}

		return ret
	}
}

// Or returns a selector that matches tokens that are matched by either a or b
func Or(a, b TokenSelector) TokenSelector {
	return func(t Tokens) []int {
		one := a(t)
		ret := b(t)

		seen := make(map[int]bool)

		for _, i := range ret {
			seen[i] = true
		}

		for _, i := range one {
			if !seen[i] {
				ret = append(ret, i)
			}
		}
		sort.Ints(ret)
		return ret
	}
}
