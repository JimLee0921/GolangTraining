package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启 goroutine 将 0-100 的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启 goroutine 从 ch1 中接收值，并将该值的平方发送到 ch2 中
	go func() {
		// 通过 ok=false 判断管道是否关闭
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主 goroutine 中从 ch2 中接收值打印，使用 for range
	for i := range ch2 { // 通道关闭后会自动退出 for range 循环
		fmt.Println(i)
	}
}
