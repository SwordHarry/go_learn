package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)
	var a, b int
	fmt.Println(&a == &a, &a == &b, &a == nil)
}
