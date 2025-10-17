package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	// 发送
	ch <- 20
	ch <- 50

	// 接收
	fmt.Println(<-ch) // 20
	fmt.Println(<-ch) // 50

	// 再发送
	ch <- 59
	close(ch) // 关闭 channel

	// 接收 + ok
	val, ok := <-ch
	fmt.Println(val, ok) // 59 true

	val, ok = <-ch
	fmt.Println(val, ok) // 0 false （零值 + 关闭状态）
}
