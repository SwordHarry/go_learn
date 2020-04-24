package main

// dup2 打印输入中多次出现的行的个数和文本
// 它 从 stdin 或指定的文件列表读取
import "os"

import "bufio"

import "fmt"

func main() {
	files := os.Args[1:]
	for _, arg := range files {
		counts := make(map[string]int)
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2; %v\n", err)
			continue
		}
		countLines(f, counts)
		f.Close()
		for _, n := range counts {
			if n > 1 {
				fmt.Println(arg)
				break
			}
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
