package main

import "fmt"

func main() {
	x := 2
	defer fmt.Println(x)
	if x == 1 {
		return
	} else {
		panic("宕机了吗")
	}

}
