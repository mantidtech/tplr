package functions

import (
	"reflect"
	"text/template"
)

// LogicFunctions perform logical operations
func LogicFunctions() template.FuncMap {
	return template.FuncMap{
		"and":       And,
		"isZero":    IsZero,
		"or":        Or,
		"when":      When,
		"whenEmpty": WhenEmpty,
	}
}

// When returns the second argument if the first is not "empty", otherwise it returns an empty string
func When(d, s interface{}) interface{} {
	if !IsZero(s) {
		return d
	}
	return ""
}

// WhenEmpty returns the second argument if the first is "empty", otherwise it returns the first
func WhenEmpty(d, s interface{}) interface{} {
	if IsZero(s) {
		return d
	}
	return s
}

// IsZero returns true if the value given corresponds to its types zero value,
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

// And returns the value of the last expression if all expressions evaluate to non-zero, or empty string otherwise
func And(expr ...interface{}) interface{} {
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
func Or(expr ...interface{}) interface{} {
	for _, e := range expr {
		if !IsZero(e) {
			return e
		}
	}
	return ""
}
