package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		// 动态规划
		// 除2相当于右移1位，而零位的1则可以由按位与得到
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var result byte = 0
	for i := 0; i < 8; i++ {
		result += pc[byte(x>>(i*8))]
	}
	return int(result)
}

func main() {
	result := PopCount(255)
	fmt.Println(pc)
	fmt.Println(result, byte(7))
}
