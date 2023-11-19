package helper

import (
	"errors"
	"fmt"
	"reflect"
)

// Apply a function to each element of a slice, returning a translated slice
func Apply[T, U any](a []T, fn func(T) (U, error)) ([]U, error) {
	var errs []error
	res := make([]U, len(a))
	for i, o := range a {
		var e error
		res[i], e = fn(o)
		errs = append(errs, e)
	}
	return res, errors.Join(errs...)
}

// Reduce an array to a single value through successive calls to a function
// that takes the previous value and the next element of the array
func Reduce[T, U any](a []T, initial U, fn Reducer[T, U]) U {
	acc := initial
	for _, i := range a {
		acc = fn(acc, i)
	}
	return acc
}

// Reducer is the signature for a reduction method used by Reduce
type Reducer[T, U any] func(a U, b T) U

// ListInfo returns basic info that we need for most list processing,
// ie it's reflect.Value (converted to the generic []any type), length,
// or error message if it's not a list
func ListInfo(list any) (slice reflect.Value, length int, err error) {
	if list == nil {
		return slice, 0, fmt.Errorf("list is nil")
	}

	t := reflect.TypeOf(list).Kind()
	if t != reflect.Slice && t != reflect.Array {
		return slice, 0, fmt.Errorf("type %s is not a list", t)
	}

	listType := reflect.TypeOf(([]any)(nil))

	slice = reflect.ValueOf(list)
	if reflect.TypeOf(list) == listType {
		return slice, slice.Len(), nil
	}

	n := reflect.MakeSlice(listType, slice.Len(), slice.Cap())
	for i := 0; i < slice.Len(); i++ {
		n.Index(i).Set(slice.Index(i))
	}

	return n, n.Len(), nil
}

// ItemForList returns the item as a reflect.Value that can be inserted into the list.
// This is because nil requires special casing
func ItemForList(item any) reflect.Value {
	if item == nil {
		return reflect.Zero(reflect.TypeOf((*any)(nil)).Elem()) // heh heh, ugh
	}
	return reflect.ValueOf(item)
}

// AsStringList converts a list to a list of strings
func AsStringList(list any) ([]string, error) {
	a, l, err := ListInfo(list)
	if err != nil {
		return nil, err
	}

	s := make([]string, l)
	for c := 0; c < l; c++ {
		v := a.Index(c).Interface()
		s[c] = fmt.Sprintf("%v", v)
	}
	return s, nil
}

// func flatten(things ...any) any {
//	if things == nil {
//		return nil
//	}
//
//	list := reflect.MakeSlice(reflect.TypeOf(([]any)(nil)), 0, 25)
//	for _, i := range things {
//		t := reflect.TypeOf(i).Kind()
//		if t == reflect.Slice || t == reflect.Array {
//			list = reflect.AppendSlice(list, reflect.ValueOf(i))
//		} else {
//			list = reflect.Append(list, reflect.ValueOf(i))
//		}
//	}
//
//	return list
// }
