package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	s = rotate(s, 8)
	fmt.Println(s)
}

func rotate(s []int, r int) []int {
	length := len(s)
	if r <= 0 {
		r += length
	}
	if r >= length {
		r %= length
	}
	return append(s[length-r:], s[:length-r]...)

}
