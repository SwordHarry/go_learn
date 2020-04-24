package structLearn

import "time"

type Employee struct {
	ID        int
	Name      string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

type point struct {
	X, Y int
}

type circle struct {
	point
	Radius int
}

type Wheel struct {
	circle
	Spokes int
}

// func main() {

// 	println(dilbert.ID)

// 	var employeeOfTheMonth *Employee = &dilbert
// 	(*employeeOfTheMonth).Position += " (procative team player)"
// 	println(dilbert.Position)
// 	EmployeeByID(0).Salary = 0

// 	w := Wheel{
// 		Circle: Circle{
// 			Point:  Point{X: 8, Y: 8},
// 			Radius: 5,
// 		},
// 		Spokes: 20, // 注意尾部的逗号是必须的
// 	}
// 	w.X = 9
// 	fmt.Printf("%#v\n", w)
// }

func EmployeeByID(id int) *Employee {
	return &dilbert
}
