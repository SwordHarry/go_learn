package main

import (
	"fmt"
	"sort"
)

type palindrome string

func (p *palindrome) Len() int {
	return len(*p)
}

func (p *palindrome) Less(i, j int) bool {
	return (*p)[i] < (*p)[j]
}

func (p *palindrome) Swap(i, j int) {
}

func IsPalindrome(s sort.Interface) bool {
	length := s.Len()
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		}
		return false
	}
	return true
}

func main() {
	p := palindrome("123454321")
	fmt.Println(IsPalindrome(&p))
}
