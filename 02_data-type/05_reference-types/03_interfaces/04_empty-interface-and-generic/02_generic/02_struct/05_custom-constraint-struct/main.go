package main

import "fmt"

// Number 自定义约束
type Number interface {
	int | float64
}

// Calculator 使用自定义 Number 约束定义结构体
type Calculator[T Number] struct {
	data []T
}

func (c Calculator[T]) Sum() T {
	var total T
	for _, v := range c.data {
		total += v
	}
	return total
}

func main() {
	c1 := Calculator[int]{data: []int{1, 2, 3}}
	c2 := Calculator[float64]{data: []float64{1.1, 2.2, 3.3}}

	fmt.Println(c1.Sum()) // 6
	fmt.Println(c2.Sum()) // 6.6
}
