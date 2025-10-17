package main

import "fmt"

type Meter float64

func (m Meter) ToCentimeter() float64 {
	return float64(m * 100) // 这里需要手动类型转化
}

func main() {
	/*
		不能直接给内建类型（如 float64）直接加方法，但是可以使用命名类型进行方法挂载
	*/

	var m Meter = 1.75
	fmt.Println(m.ToCentimeter())
}
