package main

import "fmt"

func main() {
	s := []string{"a", "a", "a", "b", "b", "c", "c", "c", "d"}
	s = noRepeat(s)
	fmt.Println(s)
}

func noRepeat(s []string) []string {
	length := len(s)
	i := 0
	for j := 1; j < length; j++ {
		if s[i] != s[j] {
			i++
			s[i] = s[j]
		}
	}
	return s[:i+1]
}
