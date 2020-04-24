package main

import "fmt"
import "github.com/structLearn"

func main() {
	w := Wheel{
		circle: circle{
			point:  point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // 注意尾部的逗号是必须的
	}
	e := EmployeeByID(0)
	fmt.Println(w, e)
}
