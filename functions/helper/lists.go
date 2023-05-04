package helper

import (
	"fmt"
	"reflect"
)

// ListInfo returns basic info that we need for most list processing,
// ie it's reflect.Value (converted to the generic []any type), length,
// or error message if it's not a list
func ListInfo(list any) (reflect.Value, int, error) {
	var v reflect.Value
	if list == nil {
		return v, 0, fmt.Errorf("list is nil")
	}

	t := reflect.TypeOf(list).Kind()
	if t != reflect.Slice && t != reflect.Array {
		return v, 0, fmt.Errorf("type %s is not a list", t)
	}

	listType := reflect.TypeOf(([]any)(nil))

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

//func flatten(things ...any) any {
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
//}
