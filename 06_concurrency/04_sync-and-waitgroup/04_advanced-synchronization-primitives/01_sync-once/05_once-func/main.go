package main

import (
	"fmt"
	"sync"
)

var initServer = sync.OnceFunc(func() {
	fmt.Println("start server")
})

func main() {
	// 所有调用都安全，且只会执行一次
	initServer() // 只有第一次会被执行
	initServer() // 被忽略
	initServer() // 被忽略
	initServer() // 被忽略
}
