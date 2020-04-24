package main

import "bufio"

import "os"

import "io"

import "fmt"

import "unicode"

func main() {
	s := [3]int{}
	in := bufio.NewReader(os.Stdin)
	invalid := 0
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			s[0]++
		} else if unicode.IsNumber(r) {
			s[1]++
		} else {
			s[2]++
		}
	}
	for _, count := range s {
		fmt.Println(count)
	}
}
