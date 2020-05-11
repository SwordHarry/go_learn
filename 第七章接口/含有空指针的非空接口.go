package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	//var buf *bytes.Buffer
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
}

func f(out io.Writer) {
	fmt.Println(out != nil)
	if out != nil {
		fmt.Println("进来了")
		out.Write([]byte("done!\n"))
	}
}
