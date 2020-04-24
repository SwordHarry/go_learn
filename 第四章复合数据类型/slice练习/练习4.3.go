package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	reverse(&arr)
	fmt.Println(arr)
}

func reverse(arr *[4]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
