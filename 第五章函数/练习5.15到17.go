package main

import (
	"fmt"
	"strings"
)

func max(vals ...int) (int, bool) {
	if len(vals) == 0 {
		return 0, false
	}
	result := vals[0]
	for _, val := range vals {
		if result < val {
			result = val
		}
	}
	return result, true
}

func strJoin(sep string, strs ...string) string {
	return strings.Join(strs, sep)
}

func main() {
	vals := []int{2, 4, 5, 1, 8}
	fmt.Println(max(vals...))
	fmt.Println(strJoin(" ", "haha", "yes"))
}
