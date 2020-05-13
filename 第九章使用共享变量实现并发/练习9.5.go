package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	c1 := make(chan int)
	c2 := make(chan int)
	defer func() {
		fmt.Println(count)
	}()

	go func() {
		c1 <- count
		for {
			count = <-c2
			count++
			c1 <- count
		}
	}()

	go func() {
		for {
			count = <-c1
			count++
			c2 <- count
		}
	}()

	time.Sleep(1 * time.Second)
}
