package functions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"text/template"
	"time"
	"unicode"
	"unicode/utf8"

	"gopkg.in/yaml.v2"

	"github.com/mantidtech/wordcase"
)

// All returns all of the templating functions
func All(t *template.Template) template.FuncMap {
	return template.FuncMap{
		"first":              First,
		"include":            GenerateIncludeFn(t),
		"indent":             Indent,
		"last":               Last,
		"nl":                 Newline,
		"now":                Now,
		"padLeft":            PadLeft,
		"padRight":           PadRight,
		"rep":                Rep,
		"rest":               Rest,
		"space":              Space,
		"spIndent":           IndentSpace,
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
		"snakeCase":          wordcase.SnakeCase,
		"kebabCase":          wordcase.KebabCase,
		"dotCase":            wordcase.DotCase,
		"screamingSnakeCase": wordcase.ScreamingSnakeCase,
		"camelCase":          wordcase.CamelCase,
		"pascalCase":         wordcase.PascalCase,
		"tpoWords":           wordcase.Words,
		"titleCase":          wordcase.TitleCase,
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

// First returns the head of a list
func First(list interface{}) (interface{}, error) {
	if list == nil {
		return list, nil
	}

	t := reflect.TypeOf(list).Kind()
	if t != reflect.Slice && t != reflect.Array {
		return nil, fmt.Errorf("type %s is not a list", t)
	}

	a := reflect.ValueOf(list)
	if a.Len() == 0 {
		return nil, nil
	}

	return a.Index(0).Interface(), nil
}

// Rest returns the tail of a list
func Rest(list interface{}) (interface{}, error) {
	if list == nil {
		return list, nil
	}

	t := reflect.TypeOf(list).Kind()
	if t != reflect.Slice && t != reflect.Array {
		return nil, fmt.Errorf("type %s is not a list", t)
	}

	a := reflect.ValueOf(list)
	l := a.Len()
	if l < 2 {
		return nil, nil
	}

	// Slice is supposed to be a[i:j], so oughta be l-1 for arg 2...
	return a.Slice(1, l).Interface(), nil
}

// Last returns the last item of a list
func Last(list interface{}) (interface{}, error) {
	if list == nil {
		return list, nil
	}

	t := reflect.TypeOf(list).Kind()
	if t != reflect.Slice && t != reflect.Array {
		return nil, fmt.Errorf("type %s is not a list", t)
	}

	a := reflect.ValueOf(list)
	l := a.Len()
	if l == 0 {
		return nil, nil
	}

	return a.Index(l - 1).Interface(), nil
}

//func listHelper(list interface{}) (reflect.Value, error) {
//	var v reflect.Value
//	if list == nil {
//		return v, nil
//	}
//
//	t := reflect.TypeOf(list).Kind()
//	if t != reflect.Slice && t != reflect.Array {
//		return v, fmt.Errorf("type %s is not a list", t)
//	}
//
//	return reflect.ValueOf(list), nil
//}

// WhenEmpty returns the second argument if the first is "empty", otherwise it returns the first
func WhenEmpty(d, s string) string {
	if s == "" {
		return d
	}
	return s
}

// Indent prints the given string with the given number of tabs prepended before each line
func Indent(t int, content string) string {
	if t < 0 {
		return ""
	}

	parts := strings.Split(content, "\n")
	tab := strings.Repeat("\t", t)

	newParts := make([]string, len(parts))

	for i, p := range parts {
		if p != "" {
			newParts[i] = tab + p
		}
	}

	return strings.Join(newParts, "\n")
}

// IndentSpace prints the given string with the given number of spaces prepended before each line
func IndentSpace(t int, content string) string {
	if t < 0 {
		return ""
	}

	parts := strings.Split(content, "\n")
	tab := strings.Repeat(" ", t)

	newParts := make([]string, len(parts))

	for i, p := range parts {
		if p != "" {
			newParts[i] = tab + p
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
func Bracket(s string) string {
	return "(" + s + ")"
}

// BracketWith adds brackets of a given type around the given string
func BracketWith(b string, s string) (string, error) {
	if len(b)%2 != 0 {
		return "", fmt.Errorf("expected a set of brackets with matching left and right sizes")
	}
	h := len(b) / 2
	l, r := b[:h], b[h:]
	return l + s + r, nil
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
