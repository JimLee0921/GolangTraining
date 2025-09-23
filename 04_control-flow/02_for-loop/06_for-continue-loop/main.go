package main

import "fmt"

// main for 循环配合 continue 跳过指定满足条件的循环
func main() {
	i := 0
	for {
		// 这里需要把自增放在最前面，不然当 i 为 2 时就会一直陷入 continue 状态
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d 为奇数\n", i)
		if i > 49 {
			break
		}
	}
}
