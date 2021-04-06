package functions

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/mantidtech/wordcase"
	"golang.org/x/sys/unix"
)

// All returns all of the templating functions
func All(t *template.Template) template.FuncMap {
	return template.FuncMap{
		"include":            GenerateIncludeFn(t),
		"indent":             Indent,
		"nl":                 Newline,
		"now":                Now,
		"toColumns":          ToColumn,
		"padLeft":            PadLeft,
		"padRight":           PadRight,
		"prefix":             Prefix,
		"suffix":             Suffix,
		"rep":                Rep,
		"space":              Space,
		"sp":                 Space,
		"tab":                Tab,
		"toLower":            strings.ToLower,
		"toUpper":            strings.ToUpper,
		"trim":               strings.TrimSpace,
		"ucFirst":            UppercaseFirst,
		"whenEmpty":          WhenEmpty,
		"isZero":             IsZero,
		"bracket":            Bracket,
		"bracketWith":        BracketWith,
		"join":               Join,
		"joinWith":           JoinWith,
		"splitOn":            SplitOn,
		"typeName":           TypeName,
		"typeKind":           TypeKind,
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
		"titleCaseWithAbbr":  TitleCaseWithAbbr,
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
		"terminalWidth":      TerminalWidth,
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
	r := strings.Join(s, " ")
	return strings.Repeat(r, n)
}

// WhenEmpty returns the second argument if the first is "empty", otherwise it returns the first
func WhenEmpty(d, s interface{}) interface{} {
	if IsZero(s) {
		return d
	}
	return s
}

// Indent prints the given string with the given number of spaces prepended before each line
func Indent(t int, content string) string {
	return Prefix(" ", t, content)
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

// ToColumn formats the given text to not take more than 'w' characters per line,
// splitting on the space before the word that would take the line over.
// If no space can be found, the line isn't split (ie words bigger than the line size are printed unsplit)
func ToColumn(w int, s string) string {
	var b strings.Builder
	tail := ""

	parts := strings.Split(s, "\n")
	for _, p := range parts {
		p := tail + p
		tail = ""

		lines := columnify(w, p)
		if len(lines) > 1 {
			tail = lines[len(lines)-1]
		}

		numLines := len(lines)
		for i, l := range lines {
			if i > 0 && i == numLines-1 {
				tail = l
				break
			}

			b.WriteString(l)
			b.WriteByte('\n')
		}
	}

	if len(tail) > 0 {
		b.WriteString(tail)
		b.WriteByte('\n')
	}

	return b.String()
}

// columnify is a helper method for ToColumn to split lines on spaces
func columnify(w int, s string) []string {
	var lines []string

	var at, i int
	for at < len(s) {
		i = at + w
		if i >= len(s) {
			lines = append(lines, s[at:])
			break
		}

		// look backwards for a space
		for ; i > at && s[i] != ' '; i-- {
			// just keep stepping
		}

		if i == at { // didn't find one
			// look forwards for a space
			for i = at + w; i < len(s) && s[i] != ' '; i++ {
				// just keep stepping
			}
		}

		lines = append(lines, s[at:i])
		at = i + 1
	}

	return lines
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

// TypeKind returns the 'kind'' of the given value as a string
func TypeKind(val interface{}) string {
	if val == nil {
		return "nil"
	}
	return reflect.ValueOf(val).Kind().String()
}

// TerminalWidth returns the number of columns that the terminal currently has,
// or 0 if the program isn't run in one, or it can't otherwise be determined
func TerminalWidth() int {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return 0
	}
	return int(ws.Col)
}

// titleCaseWithAbbrHelper uppercases the first letter of each word, or the whole word if it matches a given abbreviation
func titleCaseWithAbbrHelper(abbrv []string) wordcase.Combiner {
	selector := wordcase.KeyWordFn(abbrv)
	return wordcase.NewPipeline().
		TokenizeUsing(wordcase.LookAroundCategorizer, wordcase.NotLetterOrDigit, true).
		TokenizeUsing(wordcase.LookAroundCategorizer, wordcase.NotLowerOrDigit, false).
		WithAllFormatter(strings.ToLower).
		WithFormatter(strings.ToUpper, selector).
		WithAllFormatter(wordcase.UppercaseFirst).
		JoinWith(" ")
}

// TitleCaseWithAbbr uppercases the first letter of each word, or the whole word if it matches a given abbreviation
func TitleCaseWithAbbr(abbrv interface{}, word string) (string, error) {
	a, l, err := listInfo(abbrv)
	if err != nil {
		return "", err
	}

	strList := make([]string, l)
	for c := 0; c < l; c++ {
		strList[c] = fmt.Sprintf("%v", a.Index(c).Interface())
	}

	tc := titleCaseWithAbbrHelper(strList)
	return tc(word), nil
}
