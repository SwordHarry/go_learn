package main

import "fmt"

func deferLearn() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

func deferLoop() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func a() {
	i := 0
	defer fmt.Println(i) //输出0，因为i此时就是0
	i++
	defer fmt.Println(i) //输出1，因为i此时就是1
	return
}

func main() {
	a()
}
