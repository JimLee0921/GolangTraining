package main

import "fmt"

type sharp interface {
	area() float64
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (s square) area() float64 {
	return s.side * s.side
}

/*
通用聚合
可以一次性传入 circle、square、triangle（只要实现了 area()）
函数里统一用 s.area() 就能计算总面积
代码非常通用，不管有多少种形状，都能直接复用。
*/
func getTotalArea(sharps ...sharp) (totalArea float64) {
	for _, sharp := range sharps {
		totalArea += sharp.area()
	}
	return
}

func info(s sharp) {
	fmt.Println(s.area())
}

func main() {
	/*
		接口优势：
			多态：不同类型 (circle、square) 可以统一当作 shape 使用
			解耦：info、totalArea 不需要关心具体是哪个形状，只依赖接口
			可扩展：以后加个 triangle，只要实现 area() 方法，就能直接传给 totalArea，函数不需要修改
	*/
	c := circle{10}
	s := square{30}
	info(c)
	info(s)
	fmt.Println(getTotalArea(c, s))
}
