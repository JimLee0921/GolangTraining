package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		defer close(ch1)
		for i := 1; i <= 10; i++ {
			ch1 <- fmt.Sprintf("ch1 —— %d", i)
			time.Sleep(400 * time.Millisecond)
		}
	}()

	go func() {
		defer close(ch2)
		for i := 1; i <= 8; i++ {
			ch2 <- fmt.Sprintf("ch2 —— %d", i)
			time.Sleep(700 * time.Millisecond)
		}
	}()

	for {
		select {
		case msg, ok := <-ch1:
			if !ok {
				ch1 = nil // 防止再次触发该通道
				fmt.Println("ch1 is closed")
			} else {
				fmt.Println("ch1 received:", msg)
			}
		case msg, ok := <-ch2:
			if !ok {
				ch2 = nil
				fmt.Println("ch2 is closed")
			} else {
				fmt.Println("ch2 received", msg)
			}
		default:
			// 所有通道都没数据
			if ch1 == nil && ch2 == nil {
				fmt.Println("all channel closed")
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}
