package main

import "fmt"

// main for 循环嵌套
func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("外部第 %d 层 内部第 %d 层\n", i, j)
		}
	}
}
