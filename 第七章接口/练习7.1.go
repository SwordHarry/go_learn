package main

import (
	"bufio"
	"fmt"
)

type WordsCounter int
type LinesCounter int

func (w *WordsCounter) Write(b []byte) (int, error) {
	for start := 0; start < len(b); {
		advance, _, err := bufio.ScanWords(b[start:], true)
		if err != nil {
			return 0, err
		}
		start += advance
		*w++
	}
	return int(*w), nil

}

func (l *LinesCounter) Write(b []byte) (int, error) {
	for start := 0; start < len(b); {
		advance, _, err := bufio.ScanLines(b[start:], true)
		if err != nil {
			return 0, err
		}
		start += advance
		*l++
	}
	return int(*l), nil

}

func main() {
	var wc WordsCounter
	wc.Write([]byte("Hello Worlds Test Me"))
	fmt.Println(wc) // 4
	wc.Write([]byte("append something to the end"))
	fmt.Println(wc) // 9

	var lc LinesCounter
	fmt.Fprintf(&lc, "%s\n%s\n%s\n", "Hello World", "Second Line", "Third Line")
	fmt.Println(lc) // 3
	fmt.Fprintf(&lc, "%s\n%s\n%s", "第4行", "第5行", "")
	fmt.Println(lc) // 5
}
