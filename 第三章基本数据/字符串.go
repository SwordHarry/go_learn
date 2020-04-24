package main

import "fmt"

func main() {
	s := "left foot"
	t := s
	s += ", right foot"
	fmt.Println(s, t)
	for i, r := range "Hello，世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
