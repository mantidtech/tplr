package functions

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"
	"time"
	"unicode"
	"unicode/utf8"
)

// All returns all of the templating functions
func All(t *template.Template) template.FuncMap {
	return template.FuncMap{
		"first":     First,
		"include":   GenerateIncludeFn(t),
		"indent":    Indent,
		"last":      Last,
		"nl":        Newline,
		"now":       Now,
		"padLeft":   PadLeft,
		"padRight":  PadRight,
		"rep":       Rep,
		"rest":      Rest,
		"space":     Space,
		"spIndent":  IndentSpace,
		"tab":       Tab,
		"toLower":   strings.ToLower,
		"toUpper":   strings.ToUpper,
		"ucFirst":   UppercaseFirst,
		"whenEmpty": WhenEmpty,
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

//func listSize(list interface{}) (int, error) {
//	if list == nil {
//		return 0, nil
//	}
//
//	t := reflect.TypeOf(list).Kind()
//	if t != reflect.Slice && t != reflect.Array {
//		return 0, fmt.Errorf("type %s is not a list", t)
//	}
//
//	return reflect.ValueOf(list).Len(), nil
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
