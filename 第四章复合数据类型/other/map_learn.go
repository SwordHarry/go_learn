package main

import (
	"fmt"
)

var m = make(map[string]int)

func main() {
	// m := map[string]int{}
	// m["b"] = 2
	// k, v := m["b"]
	// delete(m, "a")
	// println(k, v)
	s1 := []string{"123"}
	s2 := []string{"456"}
	Add(s1)
	Add(s2)
	fmt.Println(Count(s1), Count(s2))
}

func k(list []string) string {
	return fmt.Sprintf("q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}
