package value

import (
	"fmt"
	"strconv"
	"time"
)

// From creates a new Value variable from a given base value
func From(v any) (string, error) {
	var res string
	switch x := v.(type) {
	case string:
		res = x
	case *string:
		res = *x
	case bool:
		res = strconv.FormatBool(x)
	case *bool:
		res = strconv.FormatBool(*x)
	case byte:
		res = fmt.Sprintf("%c", x)
	case *byte:
		res = fmt.Sprintf("%c", *x)
	case complex64:
		res = strconv.FormatComplex(complex128(x), 'f', 3, 64)
	case *complex64:
		res = strconv.FormatComplex(complex128(*x), 'f', 3, 64)
	case complex128:
		res = strconv.FormatComplex(x, 'f', 3, 128)
	case *complex128:
		res = strconv.FormatComplex(*x, 'f', 3, 128)
	case float32:
		res = strconv.FormatFloat(float64(x), 'f', 3, 32)
	case *float32:
		res = strconv.FormatFloat(float64(*x), 'f', 3, 32)
	case float64:
		res = strconv.FormatFloat(x, 'f', 3, 64)
	case *float64:
		res = strconv.FormatFloat(*x, 'f', 3, 64)
	case int:
		res = strconv.FormatInt(int64(x), 10)
	case *int:
		res = strconv.FormatInt(int64(*x), 10)
	case int8:
		res = strconv.FormatInt(int64(x), 10)
	case *int8:
		res = strconv.FormatInt(int64(*x), 10)
	case int16:
		res = strconv.FormatInt(int64(x), 10)
	case *int16:
		res = strconv.FormatInt(int64(*x), 10)
	case int32:
		res = strconv.FormatInt(int64(x), 10)
	case *int32:
		res = strconv.FormatInt(int64(*x), 10)
	case int64:
		res = strconv.FormatInt(x, 10)
	case *int64:
		res = strconv.FormatInt(*x, 10)
	case uint:
		res = strconv.FormatUint(uint64(x), 10)
	case *uint:
		res = strconv.FormatUint(uint64(*x), 10)
	case uint16:
		res = strconv.FormatUint(uint64(x), 10)
	case *uint16:
		res = strconv.FormatUint(uint64(*x), 10)
	case uint32:
		res = strconv.FormatUint(uint64(x), 10)
	case *uint32:
		res = strconv.FormatUint(uint64(*x), 10)
	case uint64:
		res = strconv.FormatUint(x, 10)
	case *uint64:
		res = strconv.FormatUint(*x, 10)
	case time.Duration:
		res = x.String()
	case *time.Duration:
		res = x.String()
	case time.Time:
		res = x.Format(time.RFC3339)
	case *time.Time:
		res = x.Format(time.RFC3339)
	default:
		return "", fmt.Errorf("cannot convert object type %T to a Value", x)
	}

	return res, nil
}

// ParserFn is the method signature for a value parser (converts a string to a value)
type ParserFn[T any] func(val string) (T, error)

// FormatterFn is the method signature for a value formatter (converts a value to a string)
type FormatterFn[T any] func(value T) string

// Value can store any data type where it can be represented as a string
type Value[T any] struct {
	parser ParserFn[T]
	format FormatterFn[T]
	value  T
}

// New returns a new Value[T], using the given value as the underlying type/value
func New[T any](value T) *Value[T] {
	return &Value[T]{
		value: value,
	}
}

// WithParser adds a parsing function to convert from a string for the value
func (v *Value[T]) WithParser(parser ParserFn[T]) *Value[T] {
	v.parser = parser
	return v
}

// WithFormatter adds a function to format the value as a string
func (v *Value[T]) WithFormatter(format FormatterFn[T]) *Value[T] {
	v.format = format
	return v
}

// Set the value from a string
func (v *Value[T]) Set(value string) error {
	var err error
	if v.parser == nil {
		return fmt.Errorf("parser not set for value instance with type %T", v.value)
	}
	v.value, err = v.parser(value)
	return err
}

// String implements Stringer, returning the value as a string
func (v *Value[T]) String() string {
	if v.format != nil {
		return v.format(v.value)
	}
	return fmt.Sprintf("%v", v.value)
}
