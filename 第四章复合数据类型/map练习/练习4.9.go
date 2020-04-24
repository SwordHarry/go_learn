package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./wordfreq.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	words := make(map[string]int)
	for scanner.Scan() {
		words[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(words)
}
