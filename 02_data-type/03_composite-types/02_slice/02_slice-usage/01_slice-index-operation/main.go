package main

import "fmt"

// main slice下标
func main() {
	/*
		slice 中下标也是从 0 开始的，可以通过下标访问和修改元素
	*/

	staff := []string{"JimLee", "Jane", "Tom", "Jane"}

	// 通过下标访问元素
	fmt.Println(staff[0])
	fmt.Println(staff[1])
	fmt.Println(staff[2])
	fmt.Println(staff[3])
	//fmt.Println(staff[4])  超出会报错

	// 通过下标修改元素
	staff[0] = "Jim"
	fmt.Println(staff)
}
