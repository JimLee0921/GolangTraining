package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	count := 0

	for {
		select {
		case t := <-ticker.C:
			count++
			fmt.Println("tick", count, "at", t)

			// 第 3 次 tick 后，加快节奏
			if count == 3 {
				fmt.Println("reset ticker to 500ms")
				ticker.Reset(500 * time.Millisecond)
			}

			// 第 8 次 tick 后退出
			if count == 8 {
				fmt.Println("exit")
				return
			}
		}
	}
}
