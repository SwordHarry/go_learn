package main

import "fmt"

func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	z := x * y
	fmt.Println(z)
	fmt.Println(real(z))
	fmt.Println(imag(z))

}
