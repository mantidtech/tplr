package functions

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/mantidtech/wordcase"
)

// All returns all of the templating functions
func All(t *template.Template) template.FuncMap {
	return template.FuncMap{
		"include":            GenerateIncludeFn(t),
		"indent":             Indent,
		"nl":                 Newline,
		"now":                Now,
		"padLeft":            PadLeft,
		"padRight":           PadRight,
		"prefix":             Prefix,
		"suffix":             Suffix,
		"rep":                Rep,
		"space":              Space,
		"tab":                Tab,
		"toLower":            strings.ToLower,
		"toUpper":            strings.ToUpper,
		"ucFirst":            UppercaseFirst,
		"whenEmpty":          WhenEmpty,
		"isZero":             IsZero,
		"bracket":            Bracket,
		"bracketWith":        BracketWith,
		"join":               Join,
		"joinWith":           JoinWith,
		"splitOn":            SplitOn,
		"typeName":           TypeName,
		"toJSON":             ToJSON,
		"formatJSON":         FormatJSON,
		"toYAML":             ToYAML,
		"camelCase":          wordcase.CamelCase,
		"dotCase":            wordcase.DotCase,
		"kebabCase":          wordcase.KebabCase,
		"pascalCase":         wordcase.PascalCase,
		"screamingSnakeCase": wordcase.ScreamingSnakeCase,
		"snakeCase":          wordcase.SnakeCase,
		"titleCase":          wordcase.TitleCase,
		"toWords":            wordcase.Words,
		"list":               List,
		"first":              First,
		"last":               Last,
		"rest":               Rest,
		"contains":           Contains,
		"filter":             Filter,
		"push":               Push,
		"pop":                Pop,
		"unshift":            Unshift,
		"shift":              Rest,
		"slice":              Slice,
		"toBase64":           ToBase64,
		"fromBase64":         FromBase64,
	}
}

// GenerateIncludeFn creates a function to be used as an "include" function in templates
func GenerateIncludeFn(t *template.Template) func(string, interface{}) (string, error) {
	const recursionLimit = 100
	inc := make(map[string]int) // keep track of how many times each template has been nested
	return func(name string, data interface{}) (string, error) {
		var buf strings.Builder
		if inc[name] > recursionLimit {
			return "", fmt.Errorf("recursion limit hit rendering template: %s", name)
		}
		inc[name]++
		err := t.ExecuteTemplate(&buf, name, data)
		inc[name]--
		return buf.String(), err
	}
}

// UppercaseFirst converts the first character in a string to uppercase
func UppercaseFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

// Newline prints a newline (handy for trying to format templates)
func Newline(c ...int) string {
	if len(c) == 0 {
		return "\n"
	}
	return strings.Repeat("\n", c[0])
}

// Rep repeats the given string(s) the given number of times
func Rep(n int, s ...string) string {
	if n < 0 {
		return ""
	}
	r := strings.Join(s, "")
	return strings.Repeat(r, n)
}

// WhenEmpty returns the second argument if the first is "empty", otherwise it returns the first
func WhenEmpty(d, s string) string {
	if s == "" {
		return d
	}
	return s
}

// Indent prints the given string with the given number of tabs prepended before each line
func Indent(t int, content string) string {
	return Prefix("\t", t, content)
}

// Prefix prints the given string with the given number of 'prefix' prepended before each line
func Prefix(prefix string, t int, content string) string {
	if t < 0 {
		return ""
	}

	parts := strings.Split(content, "\n")
	tab := strings.Repeat(prefix, t)

	newParts := make([]string, len(parts))

	for i, p := range parts {
		if p != "" {
			newParts[i] = tab + p
		}
	}

	return strings.Join(newParts, "\n")
}

// Suffix prints the given string with the given number of 'suffix' appended to each line
func Suffix(suffix string, t int, content string) string {
	if t < 0 {
		return ""
	}

	parts := strings.Split(content, "\n")
	tab := strings.Repeat(suffix, t)

	newParts := make([]string, len(parts))

	for i, p := range parts {
		if p != "" {
			newParts[i] = p + tab
		}
	}

	return strings.Join(newParts, "\n")
}

// Space prints a space character the given number of times
func Space(n int) string {
	if n < 0 {
		return ""
	}
	return strings.Repeat(" ", n)
}

// Tab prints a tab character the given number of times
func Tab(n int) string {
	if n < 0 {
		return ""
	}
	return strings.Repeat("\t", n)
}

// PadRight prints the given string in the given number of columns, right aligned
func PadRight(n int, s string) string {
	f := fmt.Sprintf("%%-%ds", n)
	return fmt.Sprintf(f, s)
}

// PadLeft prints the given string in the given number of columns, left aligned
func PadLeft(n int, s string) string {
	f := fmt.Sprintf("%%%ds", n)
	return fmt.Sprintf(f, s)
}

// Now returns the current time in the format "2006-01-02T15:04:05Z07:00"
func Now() string {
	return time.Now().Format(time.RFC3339)
}

// IsZero returns true if the value given corresponds to it's types zero value,
// points to something zero valued, or if it's a type with a length which is 0
func IsZero(val interface{}) bool {
	if val == nil {
		return true
	}

	t := reflect.TypeOf(val).Kind()
	v := reflect.ValueOf(val)

	switch t {
	case reflect.Slice, reflect.Array, reflect.Chan, reflect.Map, reflect.String:
		return v.Len() == 0
	}
	return v.IsZero()
}

// Bracket adds brackets around the given string
func Bracket(s interface{}) string {
	return "(" + fmt.Sprintf("%v", s) + ")"
}

// BracketWith adds brackets of a given type around the given string
func BracketWith(b string, s interface{}) (string, error) {
	if len(b)%2 != 0 {
		return "", fmt.Errorf("expected a set of brackets with matching left and right sizes")
	}
	h := len(b) / 2
	l, r := b[:h], b[h:]
	return l + fmt.Sprintf("%v", s) + r, nil
}

// Join joins the given strings together
func Join(s ...string) string {
	if s == nil {
		return ""
	}
	return strings.Join(s, "")
}

// JoinWith joins the given strings together using the given string as glue
func JoinWith(glue string, s ...string) string {
	if s == nil {
		return ""
	}
	return strings.Join(s, glue)
}

// SplitOn creates an array from the given string by separating it by the glue string
func SplitOn(glue string, s string) []string {
	return strings.Split(s, glue)
}

// TypeName returns the type of the given value as a string
func TypeName(val interface{}) string {
	if val == nil {
		return "nil"
	}
	return reflect.TypeOf(val).String()
}
