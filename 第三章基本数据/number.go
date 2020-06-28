package main

import "fmt"

func main() {
	var u uint8 = 255
	o := 0666
	fmt.Println(u, u+1, u*u, o)
	var s uint8 = 2
	fmt.Println(u ^ s)

	t := -3
	fmt.Println(t >> 6)
}
