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

// 不用接口时，必须写两个函数
func printSquareArea(s square) {
	fmt.Println("Square area:", s.area())
}
func printCircleArea(c circle) {
	fmt.Println("Circle area:", c.area())
}

func main() {
	/*
		使用结构体加方法，每个结构体都得写一个 area 方法，重复逻辑很多
	*/
	s := square{10}
	c := circle{5}
	printSquareArea(s)
	printCircleArea(c)
}
