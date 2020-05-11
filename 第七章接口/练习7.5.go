package main

import (
	"fmt"
	"io"
	"os"
)

type LimitedReader struct {
	reader io.Reader
	limit  int
}

func (l *LimitedReader) Read(b []byte) (int, error) {
	if l.limit <= 0 {
		return 0, nil
	}
	if len(b) > l.limit {
		b = b[:l.limit]
	}

	n, err := l.reader.Read(b)
	return n, err
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &LimitedReader{
		r, n,
	}
}

func main() {
	file, err := os.Open("./第七章接口/limit.txt") // 1234567890
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lr := LimitReader(file, 5)
	buf := make([]byte, 10)
	n, err := lr.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, buf, string(buf)) // 5 [49 50 51 52 53 0 0 0 0 0] 12345
}
