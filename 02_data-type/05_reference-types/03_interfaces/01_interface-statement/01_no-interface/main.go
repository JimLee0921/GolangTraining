package main

import "fmt"

// 定义 square 结构体
type square struct {
	side float64
}

// 定义 circle 结构体
type circle struct {
	radius float64
}

// 方法接收方为 circle 计算其面积
func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// 方法接收方为 square 计算其面积
func (z square) area() float64 {
	return z.side * z.side
}

// 不用接口时，必须写两个函数进行打印
func printSquareArea(s square) {
	fmt.Println("Square area:", s.area())
}
func printCircleArea(c circle) {
	fmt.Println("Circle area:", c.area())
}

func main() {
	/*
		上面定义了两个结构体 square 和 circle
		都定义了计算面积的 area 方法
		并且
	*/
	s := square{10}
	c := circle{5}
	printSquareArea(s)
	printCircleArea(c)
}
