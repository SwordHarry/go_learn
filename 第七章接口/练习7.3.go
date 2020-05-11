// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package main

import (
	"fmt"
)

//!+
type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	res := "{"
	res += tree2Str(t)
	res += " }"
	return res
}

func tree2Str(t *tree) string {
	if t == nil {
		return ""
	}
	res := ""

	res += tree2Str(t.left)
	res = fmt.Sprintf("%s %d", res, t.value)
	res += tree2Str(t.right)
	return res
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

//!-
func main() {
	t := tree{
		value: 0,
		left: &tree{
			value: 1,
			left:  nil,
			right: nil,
		},
		right: &tree{
			value: 2,
			left:  nil,
			right: nil,
		},
	}
	fmt.Println(t.String())

	str := "123"
	fmt.Println(str[1])
}
