package main

import "fmt"

// 定义接口只要有 area() 方法的类型都算实现了 shape 接口
type shape interface {
	area() float64
}
type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// 通用函数：接受任何实现了 shape 接口的类型
func printArea(s shape) {
	fmt.Println(s.area())
}
func main() {
	/*
		type 接口名 interface {
			方法名1(参数列表) 返回值
			方法名2(参数列表) 返回值
		}
		如果一个接口里定义了多个方法，那么类型必须实现接口里所有的方法，才能被认为实现了这个接口
		如果只实现部分方法，Go 编译器会直接报错，不会部分匹配
	*/
	s := square{10}
	c := circle{5}

	printArea(s) // 直接用一个函数处理
	printArea(c) // 不需要写额外的函数
}
