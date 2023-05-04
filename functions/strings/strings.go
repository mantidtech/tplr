package strings

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"text/template"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/mantidtech/tplr/functions/helper"
	"github.com/mantidtech/wordcase"
)

// Functions that primarily operate on strings
func Functions() template.FuncMap {
	return template.FuncMap{
		"bracket":            Bracket,
		"bracketWith":        BracketWith,
		"camelCase":          wordcase.CamelCase,
		"dotCase":            wordcase.DotCase,
		"indent":             Indent,
		"tabIndent":          TabIndent,
		"unindent":           Unindent,
		"kebabCase":          wordcase.KebabCase,
		"nl":                 Newline,
		"now":                Now,
		"padLeft":            PadLeft,
		"padRight":           PadRight,
		"pascalCase":         wordcase.PascalCase,
		"prefix":             Prefix,
		"q":                  QuoteSingle,
		"qq":                 QuoteDouble,
		"qb":                 QuoteBack,
		"rep":                Rep,
		"screamingSnakeCase": wordcase.ScreamingSnakeCase,
		"snakeCase":          wordcase.SnakeCase,
		"sp":                 Space,
		"space":              Space,
		"splitOn":            SplitOn,
		"suffix":             Suffix,
		"tab":                Tab,
		"titleCase":          wordcase.TitleCase,
		"titleCaseWithAbbr":  TitleCaseWithAbbr,
		"toColumns":          ToColumn,
		"toLower":            strings.ToLower,
		"toUpper":            strings.ToUpper,
		"toWords":            wordcase.Words,
		"trim":               strings.TrimSpace,
		"typeKind":           TypeKind,
		"typeName":           TypeName,
		"ucFirst":            UppercaseFirst,
	}
}

// titleCaseWithAbbrHelper converts the first letter of each word to uppercase, or the whole word if it matches a given abbreviation
func titleCaseWithAbbrHelper(abbrev []string) wordcase.Combiner {
	selector := wordcase.KeyWordFn(abbrev)
	return wordcase.NewPipeline().
		TokenizeUsing(wordcase.LookAroundCategorizer, wordcase.NotLetterOrDigit, true).
		TokenizeUsing(wordcase.LookAroundCategorizer, wordcase.NotLowerOrDigit, false).
		WithAllFormatter(strings.ToLower).
		WithFormatter(strings.ToUpper, selector).
		WithAllFormatter(wordcase.UppercaseFirst).
		JoinWith(" ")
}

// TitleCaseWithAbbr upper-cases the first letter of each word, or the whole word if it matches a given abbreviation
func TitleCaseWithAbbr(abbrev any, word string) (string, error) {
	a, l, err := helper.ListInfo(abbrev)
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

// UppercaseFirst converts the first character in a string to uppercase
func UppercaseFirst(s any) string {
	str := fmt.Sprintf("%v", s)
	if str == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(str)
	return string(unicode.ToUpper(r)) + str[n:]
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
func PadRight(n int, s any) string {
	f := fmt.Sprintf("%%-%dv", n)
	return fmt.Sprintf(f, s)
}

// PadLeft prints the given string in the given number of columns, left aligned
func PadLeft(n int, s any) string {
	f := fmt.Sprintf("%%%dv", n)
	return fmt.Sprintf(f, s)
}

// Bracket adds brackets around the given string
func Bracket(item any) string {
	return fmt.Sprintf("(%v)", item)
}

// QuoteSingle adds single quote around the given string
func QuoteSingle(item any) string {
	return fmt.Sprintf("'%v'", item)
}

// QuoteDouble adds double quote around the given string
func QuoteDouble(item any) string {
	return fmt.Sprintf("\"%v\"", item)
}

// QuoteBack adds back-quotes around the given string
func QuoteBack(item any) string {
	return fmt.Sprintf("`%v`", item)
}

// BracketWith adds brackets of a given type around the given string
func BracketWith(bracketPair string, item any) (string, error) {
	if len(bracketPair)%2 != 0 {
		return "", fmt.Errorf("expected a set of brackets with matching left and right sizes")
	}
	h := len(bracketPair) / 2
	l, r := bracketPair[:h], bracketPair[h:]
	return fmt.Sprintf("%s%v%s", l, item, r), nil
}

// Newline prints a newline (handy for trying to format templates)
func Newline(count ...int) string {
	if len(count) == 0 {
		return "\n"
	}
	return strings.Repeat("\n", count[0])
}

// Rep repeats the given string(s) the given number of times
func Rep(count int, item ...string) string {
	if count < 0 {
		return ""
	}
	r := strings.Join(item, " ")
	return strings.Repeat(r, count)
}

// Indent prints the given string with the given number of spaces prepended before each line
func Indent(count int, content string) string {
	return Prefix(" ", count, content)
}

// TabIndent prints the given string with the given number of spaces prepended before each line
func TabIndent(count int, content string) string {
	return Prefix("\t", count, content)
}

// Unindent removes up to 'count' spaces from the start of all lines within 'content'
func Unindent(count int, content string) (string, error) {
	if count == 0 {
		return content, nil
	} else if count < 0 {
		return "", errors.New("cannot unindent by an negative amount")
	}
	parts := strings.Split(content, "\n")
	re := fmt.Sprintf("^\\s{1,%d}", count)

	matcher := regexp.MustCompile(re)
	for i, p := range parts {
		parts[i] = matcher.ReplaceAllString(p, "")
	}

	return strings.Join(parts, "\n"), nil
}

// Prefix prints the given string with the given number of 'prefix' prepended before each line
func Prefix(prefix string, count int, content string) string {
	if count < 0 {
		return ""
	}

	parts := strings.Split(content, "\n")
	tab := strings.Repeat(prefix, count)

	newParts := make([]string, len(parts))

	for i, p := range parts {
		if p != "" {
			newParts[i] = tab + p
		}
	}

	return strings.Join(newParts, "\n")
}

// Suffix prints the given string with the given number of 'suffix' appended to each line
func Suffix(suffix string, count int, content string) string {
	if count < 0 {
		return ""
	}

	parts := strings.Split(content, "\n")
	tab := strings.Repeat(suffix, count)

	newParts := make([]string, len(parts))

	for i, p := range parts {
		if p != "" {
			newParts[i] = p + tab
		}
	}

	return strings.Join(newParts, "\n")
}

// ToColumn formats the given text to not take more than 'w' characters per line,
// splitting on the space before the word that would take the line over.
// If no space can be found, the line isn't split (ie words bigger than the line size are printed un-split)
func ToColumn(width int, content string) string {
	var b strings.Builder
	tail := ""

	parts := strings.Split(content, "\n")
	for _, p := range parts {
		str := tail + p
		tail = ""

		lines := columnify(width, str)
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

var nowActual = time.Now // use an alias, so we can redefine it in testing
// Now returns the current time in the format "2006-01-02T15:04:05Z07:00"
func Now(format ...string) string {
	f := time.RFC3339
	if len(format) > 0 {
		f = format[0]
	}
	return nowActual().Format(f)
}

// SplitOn creates an array from the given string by separating it by the glue string
func SplitOn(glue string, content string) []string {
	return strings.Split(content, glue)
}

// TypeName returns the type of the given value as a string
func TypeName(value any) string {
	if value == nil {
		return "nil"
	}
	return reflect.TypeOf(value).String()
}

// TypeKind returns the "kind" of the given value as a string
func TypeKind(value any) string {
	if value == nil {
		return "nil"
	}
	return reflect.ValueOf(value).Kind().String()
}
