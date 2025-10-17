package main

import "fmt"

type Meter float64

func (m Meter) ToCentimeter() float64 {
	return float64(m * 100)
}

//func (n int)()  {}	这样写会报错

func main() {
	/*
		底层类型不能直接添加方法，只能通过类型定义间接扩展
		这里定义一个基于 float64 的新类型 Meter
			绑定了方法 ToCentimeter()
			只有新定义的类型才能绑定方法，别名类型不行
	*/
	var length Meter = 1.5
	fmt.Println(length.ToCentimeter()) // 150
}
