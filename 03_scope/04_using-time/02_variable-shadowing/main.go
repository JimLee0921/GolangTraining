package main

import "fmt"

func max(x int) int {
	return x + 1
}

func main() {
	// 重名造成变量遮蔽
	max := max(3)    // 这里的 max 变量遮蔽了 max 函数
	fmt.Println(max) // 此时 max 是 int 变量 不是函数
}
