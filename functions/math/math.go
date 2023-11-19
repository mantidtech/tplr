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

func Add(o ...any) float64 {
	var t float64
	for _, i := range o {
		t += helper.AnyNumberToFloat(i)
	}
	return t
}

func Subtract(o ...any) float64 {
	if len(o) == 0 {
		return 0
	}
	t := helper.AnyNumberToFloat(o[0])
	for _, i := range o[1:] {
		t -= helper.AnyNumberToFloat(i)
	}
	return t
}

func Multiply(o ...any) float64 {
	if len(o) == 0 {
		return 0
	}
	t := helper.AnyNumberToFloat(o[0])
	for _, i := range o[1:] {
		t *= helper.AnyNumberToFloat(i)
	}
	return t
}

func Divide(o ...any) float64 {
	if len(o) == 0 {
		return 0
	}
	t := helper.AnyNumberToFloat(o[0])
	for _, i := range o[1:] {
		t /= helper.AnyNumberToFloat(i)
	}
	return t
}

func AbsoluteValue(o any) float64 {
	return math.Abs(helper.AnyNumberToFloat(o))
}
