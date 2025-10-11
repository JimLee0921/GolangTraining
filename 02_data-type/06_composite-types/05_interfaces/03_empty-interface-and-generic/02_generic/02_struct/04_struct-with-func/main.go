package main

import "fmt"

type Point[T int | float64] struct {
	X, Y T
}

func (p Point[T]) Sum() T {
	return p.Y + p.X
}

func main() {
	/*
		泛型结构体可以有方法
		方法不能再单独声明新类型参数，使用结构体已有的 [T]
		泛型方法编译期会为不同类型生成专用版本
	*/
	p1 := Point[int]{X: 2, Y: 3}
	p2 := Point[float64]{X: 1.52, Y: 2.5}

	fmt.Println(p1.Sum()) // 5
	fmt.Println(p2.Sum()) // 4.02
}
