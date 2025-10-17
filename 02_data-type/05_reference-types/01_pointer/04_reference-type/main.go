package main

import "fmt"

// main 引用数据 map 作为参数传递
func main() {
	/*
		引用数据类型在赋值/传参时复制的是引用（类似指针）
		多个变量共享同一份底层数据，修改一处会影响另一处
		引用类型（reference types）：slice, map, channel, function, interface
		即使不用 &，函数内部也能通过这个引用修改底层数据
	*/

	score := map[string]int{
		"JimLee": 20,
		"Bruce":  29,
	}
	fmt.Println(score) // map[Bruce:29 JimLee:20]
	changeMap(score)
	fmt.Println(score) // map[Bruce:29 JimLee:20 Ted:20]

	// 切片本身是引用类型，不需要指针
	s := []int{1, 2, 3}
	changeSlice(s)
	fmt.Println(s) // [9 2 3] 修改成功
}

func changeMap(m map[string]int) {
	m["Ted"] = 20
}

func changeSlice(a []int) {
	a[0] = 9 // 因为切片底层是引用语义
}
