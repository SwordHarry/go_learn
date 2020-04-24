package main

import (
	"flag"
	"fmt"
	"strings"
)

// 意思是 -s 后面紧跟的字符将作为分隔符使用，-n 将在结尾添加换行符
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
