package templates

import (
	"fmt"
	"strings"
	"text/template"
)

// The recursion depth that we allow self-referring nested templates to go to
const includedTemplateRecursionLimit = 100

// Functions that operate on templates themselves
func Functions(t *template.Template) template.FuncMap {
	return template.FuncMap{
		"include": GenerateIncludeFn(t),
	}
}

// GenerateIncludeFn creates a function to be used as an "include" function in templates
func GenerateIncludeFn(t *template.Template) func(string, interface{}) (string, error) {
	inc := make(map[string]int) // keep track of how many times each template has been nested
	return func(name string, data interface{}) (string, error) {
		var buf strings.Builder
		if inc[name] > includedTemplateRecursionLimit {
			return "", fmt.Errorf("recursion limit (%d) hit rendering template: %s", includedTemplateRecursionLimit, name)
		}
		inc[name]++
		err := t.ExecuteTemplate(&buf, name, data)
		inc[name]--
		return buf.String(), err
	}
}
