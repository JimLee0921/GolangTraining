package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.AfterFunc(3*time.Second, func() {
		fmt.Println("after three seconds call func")
	})

	// 1 秒后又决定取消
	time.Sleep(1 * time.Second)
	stopped := timer.Stop()
	fmt.Println(stopped)

	// 等一会看是否被执行
	time.Sleep(3 * time.Second)
}
