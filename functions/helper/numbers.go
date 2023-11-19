package helper

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func AnyNumberToFloat(a any) float64 {
	var f float64
	switch v := a.(type) {
	case bool:
		f = 0
		if v {
			f = 1
		}
	case int:
		f = float64(v)
	case int8:
		f = float64(v)
	case int16:
		f = float64(v)
	case int32:
		f = float64(v)
	case float32:
		f = float64(v)
	case float64:
		f = v
	default:
		fmt.Printf("toFloat: can't convert from %T to int\n", a)
	}
	return f
}
