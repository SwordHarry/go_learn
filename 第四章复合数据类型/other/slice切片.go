package main

import "fmt"

func main() {
	arr := []string{0: "a", 1: "b", 2: "c", 3: "d", 4: "e", 5: "f", 6: "g", 7: "h", 8: "i", 9: "j", 10: "k", 11: "l", 12: "m", 13: "n"}
	s := arr[4:7]
	s2 := arr[6:9]
	fmt.Println(s)
	fmt.Println((s2))
	s3 := s2[:5]
	fmt.Println(s3)
	s4 := []int{0, 1, 2, 3, 4, 5} // cap 为 6
	// s5 := s4[:8] // 报错，超出 cap
	fmt.Println(s4)
	m := map[string]bool{}
	// m = nil
	fmt.Println(m["a"])
	m["a"] = true
	fmt.Println(m["a"])
}
