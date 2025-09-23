package main

import "fmt"

// main for 配合 break 终止循环（在循环体内进行条件判断）
func main() {
	i := 0
	for {
		if i > 100 {
			break
		}
		fmt.Printf("第 %d 次循环\n", i)
		i++
	}
}
