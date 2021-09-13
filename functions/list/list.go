package list

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/mantidtech/tplr/functions/helper"
)

// Functions operate on collections of items
func Functions() template.FuncMap {
	return template.FuncMap{
		"contains": Contains,
		"filter":   Filter,
		"first":    First,
		"join":     Join,
		"joinWith": JoinWith,
		"last":     Last,
		"list":     List,
		"pop":      Pop,
		"push":     Push,
		"rest":     Rest,
		"shift":    Rest,
		"slice":    Slice,
		"unshift":  Unshift,
	}
}

// List returns a new list comprised of the given elements
func List(items ...interface{}) (interface{}, error) {
	return items, nil
}

// First returns the head of a list
func First(list interface{}) (interface{}, error) {
	a, l, err := helper.ListInfo(list)
	if err != nil || l == 0 {
		return nil, err
	}

	return a.Index(0).Interface(), nil
}

// Rest / Shift returns the tail of a list
func Rest(list interface{}) (interface{}, error) {
	a, l, err := helper.ListInfo(list)
	if err != nil || l < 2 {
		return nil, err
	}

	return a.Slice(1, l).Interface(), nil
}

// Last returns the last item of a list
func Last(list interface{}) (interface{}, error) {
	a, l, err := helper.ListInfo(list)
	if err != nil || l == 0 {
		return nil, err
	}

	return a.Index(l - 1).Interface(), nil
}

// Pop removes the first element of the list, returning the list
func Pop(list interface{}) (interface{}, error) {
	a, l, err := helper.ListInfo(list)
	if err != nil || l < 2 {
		return nil, err
	}

	return a.Slice(0, l-1).Interface(), nil
}

// Slice returns a slice of a list
// where i is the lower index (inclusive) and j is the upper index (exclusive) to extract
func Slice(i, j int, list interface{}) (interface{}, error) {
	a, l, err := helper.ListInfo(list)
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
	a, l, err := helper.ListInfo(list)
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
	a, l, err := helper.ListInfo(list)
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
	a, _, err := helper.ListInfo(list)
	if err != nil {
		return nil, err
	}
	i := helper.ItemForList(item)
	a = reflect.Append(a, i)
	return a.Interface(), nil
}

// Unshift returns the list with item prepended
func Unshift(list interface{}, item interface{}) (interface{}, error) {
	a, l, err := helper.ListInfo(list)
	if err != nil {
		return nil, err
	}

	i := helper.ItemForList(item)
	s := reflect.MakeSlice(a.Type(), 1, l+1)
	s.Index(0).Set(i)

	s = reflect.AppendSlice(s, a)

	return s.Interface(), nil
}

// Join joins the given strings together
func Join(list interface{}) (string, error) {
	return JoinWith("", list)
}

// JoinWith joins the given strings together using the given string as glue
func JoinWith(glue string, list interface{}) (string, error) {
	s, err := helper.AsStringList(list)

	return strings.Join(s, glue), err
}
