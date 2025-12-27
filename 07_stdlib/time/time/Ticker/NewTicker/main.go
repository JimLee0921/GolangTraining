package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)

	defer ticker.Stop() // 逻辑结束时停止 ticker

	go func() {
		for t := range ticker.C {
			fmt.Println("tick at", t)
		}
	}()

	time.Sleep(10 * time.Second)
	ticker.Stop()
}
