package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	d := reflect.ValueOf(&x).Elem()
	px := d.Addr().Interface().(*int)
	*px = 3
	fmt.Println(x) // 3
}
