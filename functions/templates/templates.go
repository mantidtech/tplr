// Package templates provides methods for operating on templates
package templates

import (
	"errors"
	"fmt"
	"strings"
	"text/template"

	"github.com/mantidtech/tplr/functions/helper"
)

// The recursion depth that we allow self-referring nested templates to go to
const includedTemplateRecursionLimit = 100

// Functions that operate on templates themselves
func Functions(t *template.Template) template.FuncMap {
	return template.FuncMap{
		"include":      GenerateIncludeFn(t),
		"applyInclude": GenerateIncludeApplyFn(t),
	}
}

// GenerateIncludeFn creates a function to be used as an "include" function in templates
func GenerateIncludeFn(t *template.Template) func(string, any) (string, error) {
	inc := make(map[string]int) // keep track of how many times each template has been nested
	return func(name string, data any) (string, error) {
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

// GenerateIncludeApplyFn creates a function to be used as an "include" function in templates, which allows
func GenerateIncludeApplyFn(t *template.Template) func(name string, list any) (any, error) {
	return func(name string, list any) (any, error) {
		var errs []string
		a, l, err := helper.ListInfo(list)
		if err != nil || l == 0 {
			return list, err
		}
		s := make([]string, l)

		for c := 0; c < l; c++ {
			var buf strings.Builder
			v := a.Index(c)
			errI := t.ExecuteTemplate(&buf, name, v)
			if errI != nil {
				errs = append(errs, errI.Error())
			}
			s[c] = buf.String()
		}

		if len(errs) > 0 {
			return s, errors.New(strings.Join(errs, ": "))
		}

		return s, nil
	}
}
