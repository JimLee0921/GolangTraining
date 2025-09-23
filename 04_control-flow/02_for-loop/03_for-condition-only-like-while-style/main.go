package main

import "fmt"

// main 类似于 while 写法
func main() {
	/*
		在 go 语言中是没有 while 循环的，但是可以通过修改 for 循环的写法来模拟 while 循环
		将 条件变动放入 循环体中
	*/
	i := 0
	for i <= 100 {
		fmt.Printf("第 %d 次循环\n", i)
		i++
	}
}
