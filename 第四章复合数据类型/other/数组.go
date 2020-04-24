package main

import "fmt"

func main() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // true false false
	// d := [3]int{1, 2}
	// fmt.Println(a == d) // 编译错误，[2]int 不能与 [3]int 比较
}
