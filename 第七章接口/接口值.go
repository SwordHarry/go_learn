package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	var t io.Writer
	//t = new(bytes.Buffer)
	t = os.Stdout
	fmt.Println(w == t)
}
