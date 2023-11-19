package helper

import (
	"fmt"
	"strconv"
)

// ToFloat converts the supplied argument to a float, if it can be
func ToFloat(a any) (float64, error) {
	var f float64
	var err error

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
	case string:
		f, err = strconv.ParseFloat(v, 64)
	default:
		err = fmt.Errorf("toFloat: can't convert from %T to float", a)
	}

	return f, err
}
