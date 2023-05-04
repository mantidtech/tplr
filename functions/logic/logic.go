package logic

import (
	"reflect"
	"text/template"
)

// Functions perform logical operations
func Functions() template.FuncMap {
	return template.FuncMap{
		"and":       And,
		"isZero":    IsZero,
		"or":        Or,
		"when":      When,
		"whenEmpty": WhenEmpty,
	}
}

// When returns 'value' if 'cond' is not a zero value, otherwise it returns an empty string
func When(value, cond any) any {
	if !IsZero(cond) {
		return value
	}
	return ""
}

// WhenEmpty returns 'value' if 'cond' is a zero value, otherwise it returns 'cond'
func WhenEmpty(value, cond any) any {
	if IsZero(cond) {
		return value
	}
	return cond
}

// IsZero returns true if the value given corresponds to its types zero value,
// points to something zero valued, or if it's a type with a length which is 0
func IsZero(val any) bool {
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

// And returns the value of the last expression if all expressions evaluate to non-zero, or empty string otherwise
func And(expr ...any) any {
	if len(expr) == 0 {
		return ""
	}
	for _, e := range expr {
		if IsZero(e) {
			return ""
		}
	}
	return expr[len(expr)-1]
}

// Or returns the first expression that evaluates to non-zero, or empty string if none do
func Or(expr ...any) any {
	for _, e := range expr {
		if !IsZero(e) {
			return e
		}
	}
	return ""
}
