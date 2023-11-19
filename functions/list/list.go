// Package list provides methods for manipulating lists in templates
package list

import (
	"reflect"
	"strings"
	"text/template"

	"github.com/mantidtech/tplr/functions/helper"
)

// Functions operate on collections of items
func Functions() template.FuncMap {
	return template.FuncMap{
		"list":     List,
		"first":    First,
		"last":     Last,
		"rest":     Rest,
		"pop":      Pop,
		"push":     Push,
		"shift":    Rest,
		"unshift":  Unshift,
		"contains": Contains,
		"filter":   Filter,
		"join":     Join,
		"joinWith": JoinWith,
	}
}

// List returns a new list comprised of the given elements
func List(items ...any) (any, error) {
	return items, nil
}

// First returns the head of a list
func First(list any) (any, error) {
	a, l, err := helper.ListInfo(list)
	if err != nil || l == 0 {
		return nil, err
	}

	return a.Index(0).Interface(), nil
}

// Rest / Shift returns the tail of a list
func Rest(list any) (any, error) {
	a, l, err := helper.ListInfo(list)
	if err != nil || l < 2 {
		return nil, err
	}

	return a.Slice(1, l).Interface(), nil
}

// Last returns the last item of a list
func Last(list any) (any, error) {
	a, l, err := helper.ListInfo(list)
	if err != nil || l == 0 {
		return nil, err
	}

	return a.Index(l - 1).Interface(), nil
}

// Pop removes the last element of the list, returning the list
func Pop(list any) (any, error) {
	a, l, err := helper.ListInfo(list)
	if err != nil || l < 2 {
		return nil, err
	}

	return a.Slice(0, l-1).Interface(), nil
}

// Contains returns true if the item is present in the list
func Contains(list any, item any) (bool, error) {
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
func Filter(list any, item any) (any, error) {
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
func Push(list any, item any) (any, error) {
	a, _, err := helper.ListInfo(list)
	if err != nil {
		return nil, err
	}
	i := helper.ItemForList(item)
	a = reflect.Append(a, i)
	return a.Interface(), nil
}

// Unshift returns the list with item prepended
func Unshift(list any, item any) (any, error) {
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
func Join(list any) (string, error) {
	return JoinWith("", list)
}

// JoinWith joins the given strings together using the given string as glue
func JoinWith(glue string, list any) (string, error) {
	s, err := helper.AsStringList(list)

	return strings.Join(s, glue), err
}
