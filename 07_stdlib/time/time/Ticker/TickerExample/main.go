package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	done := make(chan struct{})

	// 5 秒后退出
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("tick at", t)

		case <-done:
			fmt.Println("done, exit loop")
			return
		}
	}
}
