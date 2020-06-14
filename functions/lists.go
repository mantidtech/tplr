package functions

import (
	"fmt"
	"reflect"
)

// listInfo returns basic info that we need for most list processing,
// ie it's reflect.Value (converted to the generic []interface{} type), length,
// or error message if it's not a list
func listInfo(list interface{}) (reflect.Value, int, error) {
	var v reflect.Value
	if list == nil {
		return v, 0, fmt.Errorf("list is nil")
	}

	t := reflect.TypeOf(list).Kind()
	if t != reflect.Slice && t != reflect.Array {
		return v, 0, fmt.Errorf("type %s is not a list", t)
	}

	listType := reflect.TypeOf(([]interface{})(nil))

	v = reflect.ValueOf(list)
	if reflect.TypeOf(list) == listType {
		return v, v.Len(), nil
	}

	n := reflect.MakeSlice(listType, v.Len(), v.Cap())
	for i := 0; i < v.Len(); i++ {
		n.Index(i).Set(v.Index(i))
	}

	return n, n.Len(), nil
}

// itemForList returns the item as a reflect.Value that can be inserted into the list.
// This is because nil requires special casing
func itemForList(item interface{}) reflect.Value {
	if item == nil {
		return reflect.Zero(reflect.TypeOf((*interface{})(nil)).Elem()) // heh heh, ugh
	}
	return reflect.ValueOf(item)
}

// List returns a new list comprised of the given elements
func List(s ...interface{}) (interface{}, error) {
	return s, nil
}

// First returns the head of a list
func First(list interface{}) (interface{}, error) {
	a, l, err := listInfo(list)
	if err != nil || l == 0 {
		return nil, err
	}

	return a.Index(0).Interface(), nil
}

// Rest / Shift returns the tail of a list
func Rest(list interface{}) (interface{}, error) {
	a, l, err := listInfo(list)
	if err != nil || l < 2 {
		return nil, err
	}

	return a.Slice(1, l).Interface(), nil
}

// Last returns the last item of a list
func Last(list interface{}) (interface{}, error) {
	a, l, err := listInfo(list)
	if err != nil || l == 0 {
		return nil, err
	}

	return a.Index(l - 1).Interface(), nil
}

// Pop removes the first element of the list, returning the list
func Pop(list interface{}) (interface{}, error) {
	a, l, err := listInfo(list)
	if err != nil || l < 2 {
		return nil, err
	}

	return a.Slice(0, l-1).Interface(), nil
}

// Slice returns a slice of a list
// where i is the lower index (inclusive) and j is the upper index (exclusive) to extract
func Slice(i, j int, list interface{}) (interface{}, error) {
	a, l, err := listInfo(list)
	if err != nil {
		return list, err
	} else if i < 0 {
		return nil, fmt.Errorf("index '%d' out of bounds (min 0)", i)
	} else if j > l {
		return nil, fmt.Errorf("index '%d' out of bounds (max %d)", j, l)
	}

	return a.Slice(i, j).Interface(), nil
}

// Contains returns true if the item is present in the list
func Contains(list interface{}, item interface{}) (bool, error) {
	a, l, err := listInfo(list)
	if err != nil {
		return false, err
	}

	for i := 0; i < l; i++ {
		if item == a.Index(i).Interface() {
			return true, nil
		}
	}

	return false, nil
}

// Filter returns list with all instances of item removed
func Filter(list interface{}, item interface{}) (interface{}, error) {
	a, l, err := listInfo(list)
	if err != nil || l == 0 {
		return list, err
	}
	s := reflect.MakeSlice(a.Type(), 0, l)

	for c := 0; c < l; c++ {
		v := a.Index(c)
		if item != v.Interface() {
			s = reflect.Append(s, v)
		}
	}

	return s.Interface(), nil
}

// Push returns the list with item appended
func Push(list interface{}, item interface{}) (interface{}, error) {
	a, _, err := listInfo(list)
	if err != nil {
		return nil, err
	}
	i := itemForList(item)
	a = reflect.Append(a, i)
	return a.Interface(), nil
}

// Unshift returns the list with item prepended
func Unshift(list interface{}, item interface{}) (interface{}, error) {
	a, l, err := listInfo(list)
	if err != nil {
		return nil, err
	}

	i := itemForList(item)
	s := reflect.MakeSlice(a.Type(), 1, l+1)
	s.Index(0).Set(i)

	s = reflect.AppendSlice(s, a)

	return s.Interface(), nil
}
