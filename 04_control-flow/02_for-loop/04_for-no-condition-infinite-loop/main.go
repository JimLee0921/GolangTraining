package main

import "fmt"

// main for 不设置条件语句，就是 无限循环
func main() {
	i := 0
	for {
		i++
		fmt.Printf("第 %d 次循环\n", i)
	}
}
