package main

import "fmt"

// main 引用数据 map 作为参数传递
func main() {
	/*
		引用数据类型在赋值/传参时复制的是引用（类似指针）
		多个变量共享同一份底层数据，修改一处会影响另一处
	*/

	score := map[string]int{
		"JimLee": 20,
		"Bruce":  29,
	}
	fmt.Println(score) // map[Bruce:29 JimLee:20]
	changeMap(score)
	fmt.Println(score) // map[Bruce:29 JimLee:20 Ted:20]

}

func changeMap(m map[string]int) {
	m["Ted"] = 20
}
