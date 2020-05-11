package main

import "fmt"

func main() {
	fmt.Println(recoverFromDefer())
	fmt.Println(recover())
}

func recoverFromDefer() (x int) {
	// var err error
	fmt.Println("panic 之前")
	defer func() {
		fmt.Println("我是比recover 更早的 defer")
	}()
	defer func() {
		if p := recover(); p != nil {
			err := fmt.Errorf("internal error: %v", p)
			fmt.Println(err)
			x = p.(int)
		}
		panic(5)
	}()
	defer func() {
		fmt.Println("我是recover 后的 defer")
		panic(4)
	}()
	panic(3)
	// fmt.Println("panic 之后")
	// x++
}
