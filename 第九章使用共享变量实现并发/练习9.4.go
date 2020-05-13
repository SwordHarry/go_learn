package main

import (
	"fmt"
	"runtime"
)

// 获得目前内存使用情况
func getMemSys() uint64 {
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	return mem.Sys
}

func main() {

	before := getMemSys()

	chs := make([]chan int, 0)
	start := make(chan int)
	chs = append(chs, start)
	count := 0
	go func() {
		for {
			start <- count
		}
	}()

	for n := 0; ; n++ {
		out := chs[n]
		in := make(chan int)
		chs = append(chs, in)
		go func(n int, in, out chan int) {
			for {
				count = <-out
				in <- count + 1
			}
		}(n, in, out)

		// 获得程序使用内存情况
		memAlloc := getMemSys() - before
		if memAlloc > 1024*1024*1024 { // 消耗1G内存时终止程序
			fmt.Println(n, "Goroutines", count) // 大概11~12万，每个goroutine消耗内存9k左右
			break
		}
	}
}
