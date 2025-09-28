package main

import "math"

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	r float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.r
}
