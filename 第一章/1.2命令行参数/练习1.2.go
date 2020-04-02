package main

import "os"

import "fmt"

func main() {
	// s, sep := "", ""
	for index, val := range os.Args[1:] {
		fmt.Println(index, val)
	}
}
