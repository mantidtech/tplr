package wordcase

import (
	"unicode"
	"unicode/utf8"
)

// empty is a zero-sized empty constant
var empty = struct{}{}

// WordSet is a collection of words
type WordSet map[string]struct{}

// GoLintKeywords are words that golint considers to be common initialisms that should always be same-cased
// (this list is https://github.com/golang/lint/blob/master/lint.go#L740 @ 2020-05-17)
var GoLintKeywords = []string{
	"acl",
	"api",
	"ascii",
	"cpu",
	"css",
	"dns",
	"eof",
	"guid",
	"html",
	"http",
	"https",
	"id",
	"ip",
	"json",
	"lhs",
	"qps",
	"ram",
	"rhs",
	"rpc",
	"sla",
	"smtp",
	"sql",
	"ssh",
	"tcp",
	"tls",
	"ttl",
	"udp",
	"ui",
	"uid",
	"uuid",
	"uri",
	"url",
	"utf8",
	"vm",
	"xml",
	"xmpp",
	"xsrf",
	"xss",
}

// UsefulKeyWords are words that you may want to be same-cased
var UsefulKeyWords = append(
	GoLintKeywords,
	"grpc",
	"tml",
	"toml",
	"yaml",
	"yml",
)

// A Formatter is used to format a token
type Formatter func(string) string

// UppercaseFirst returns the given string with the first rune converted to uppercase
func UppercaseFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}
