package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func IsCycle(s interface{}) bool {
	seen := make(map[comparison]bool)
	return isCycle(reflect.ValueOf(s), seen)
}

type comparison struct {
	x unsafe.Pointer
	t reflect.Type
}

func isCycle(v reflect.Value, seen map[comparison]bool) bool {
	if v.CanAddr() {
		xptr := unsafe.Pointer(v.UnsafeAddr())
		c := comparison{xptr, v.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}
	switch v.Kind() {
	case reflect.Struct:
		for i, n := 0, v.NumField(); i < n; i++ {
			if isCycle(v.Field(i), seen) {
				return true
			}
		}
	case reflect.Ptr, reflect.Interface:
		return isCycle(v.Elem(), seen)
	}
	return false
}

func main() {
	type Cycle struct {
		x int
		c *Cycle
	}
	c := Cycle{
		x: 0,
		c: nil,
	}
	b := Cycle{
		x: 3,
		c: nil,
	}
	c.c = &b
	fmt.Println(IsCycle(c))
}
