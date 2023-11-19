// Package math provides mathematical operations in templates
package math

import (
	"math"
	"text/template"

	"github.com/mantidtech/tplr/functions/helper"
)

// Functions operate on numerical data
func Functions() template.FuncMap {
	return template.FuncMap{
		"add":  Add,
		"sub":  Subtract,
		"mult": Multiply,
		"div":  Divide,
		"abs":  AbsoluteValue,
	}
}

func floatOperation(o []any, fn helper.Reducer[float64, float64]) (float64, error) {
	a, err := helper.Apply(o, helper.ToFloat)
	if err != nil {
		return 0, err
	}
	if len(a) == 0 {
		return 0, nil
	}
	f := helper.Reduce(a[1:], a[0], fn)
	return f, nil
}

// Add zero or more operands left to right. The result of an Add with no arguments is 0
func Add(o ...any) (float64, error) {
	return floatOperation(o, func(a, b float64) float64 {
		return a + b
	})
}

// Subtract zero or more operands left to right. The result of a Subtract with no arguments is 0
func Subtract(o ...any) (float64, error) {
	return floatOperation(o, func(a, b float64) float64 {
		return a - b
	})
}

// Multiply zero or more operands left to right. The result of a Multiply with no arguments is 0
func Multiply(o ...any) (float64, error) {
	return floatOperation(o, func(a, b float64) float64 {
		return a * b
	})
}

// Divide zero or more operands left to right. The result of a Divide with no arguments is 0
func Divide(o ...any) (float64, error) {
	return floatOperation(o, func(a, b float64) float64 {
		return a / b
	})
}

// AbsoluteValue returns the magnitude of the given argument
func AbsoluteValue(o any) (float64, error) {
	v, err := helper.ToFloat(o)
	if err != nil {
		return 0, err
	}
	return math.Abs(v), nil
}
