package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("12345678"))

}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

// 以下内容为练习3.10
func comma(s string) string {

	length := len(s)
	if length <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i := length - 1; i >= 0; i-- {
		fmt.Fprintf(&buf, "%c", s[i])
		if i != 0 && (length-i)%3 == 0 {
			buf.WriteRune(',')
		}
	}
	return reverseString(buf.String())
}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
