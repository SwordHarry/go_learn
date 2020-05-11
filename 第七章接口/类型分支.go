package main

import (
	"fmt"
)

func main() {
	x := 1
	switch y := interface{}(x).(type) {
	case int:
		fmt.Println("int", y)
	case float32:
		fmt.Println(y)
	case string:
		fmt.Println(y)
	case bool:
		fmt.Println(y)
	default:
		fmt.Println(y)
	}
}
