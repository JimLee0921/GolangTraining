package main

import "fmt"

func main() {
	// new() 是语法糖，常用于结构体或内置类型的指针初始化
	p := new(int)   // *int
	fmt.Println(*p) //这里是 0（零值）
	*p = 199
	fmt.Println(*p) // 100

	//上面代码等价于
	//var x int
	//p := &x
}
