package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	//var w io.Writer
	//w = os.Stdout
	//f := w.(*os.File)
	//c := w.(*bytes.Buffer)// 崩溃
	//fmt.Sprintf("%T %T", f, c)
	var w io.Writer
	w = os.Stdout
	rw := w.(io.ReadWriter)
	fmt.Printf("%T %T\n", w, rw)
	w = new(ByteCounter)
	rw = w.(io.ReadWriter)
	fmt.Printf("%T %T\n", w, rw)
}
