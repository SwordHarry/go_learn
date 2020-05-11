package main

import "fmt"

// 嵌套同名的情况
type Point struct {
	X, Y float64
}

type Haha struct {
	Y int
}

type ColoredPoint struct {
	Point
	Haha
	X float64
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) ScaleBy2(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{1,2}
	p.ScaleBy(2)
	//等价于
	//(&p).ScaleBy(2)
	fmt.Println(p)
	pptr := &Point{3,3}
	pptr.ScaleBy(2)
	// 等价于
	//(*pptr).ScaleBy(2)
	fmt.Println(*pptr)

	//Point{1,2}.ScaleBy(2) // T 无法 变成 *T
	//pptr.ScaleBy2(2) // *T 可以变成 T

	c := new(ColoredPoint)
	c.Point.X = 1
	c.Y = 1
}
