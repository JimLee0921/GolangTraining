package main

import "fmt"

// main 函数返回基础类型
func main() {
	/*
		如果不需要使用返回值可以直接使用 _ = 进行接收
	*/
	total := sum(10, 20)
	fmt.Println("total:", total)
	name := concatName("Jim", "Lee")
	fmt.Println(name)
}

func sum(x, y int) int {
	return x + y
}

func concatName(firstName, lastName string) string {
	return fmt.Sprint(firstName, " ", lastName) // 不会打印，而是返回一个格式化后的 字符串

}
