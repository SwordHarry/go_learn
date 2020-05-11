package main

import "fmt"

type chanNum struct {
	num int
}

func main() {
	c := make(chan int)
	go func() { c <- 1 }()

	done := make(chan chanNum)
	close(done)

	select {
	case x := <-done:
		fmt.Println("done", x)
	case x := <-c:
		fmt.Println("c", x)
	}
}
