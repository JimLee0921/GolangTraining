package main

import (
	"fmt"
	"net/http"
	"time"
)

// 生产级启动方式的基础版本
func main() {
	// 1. 先创建一个路由器（推荐：不用全局 DefaultServeMux）
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	// 2. 手动创建 server 可以指定一些参数（必须用 & 取地址使用指针方法）
	myServer := &http.Server{
		Addr:         ":8080",          // 监听端口
		Handler:      mux,              // 路由器或 Handler，如果传入 nil 或者不写都是使用默认 http.DefaultServeMux
		ReadTimeout:  5 * time.Second,  // 客户端读取超时时间
		WriteTimeout: 5 * time.Second,  // 写回响应的超时时间
		IdleTimeout:  30 * time.Second, // 长连接空闲超时
	}

	// 3. 启动 serve
	fmt.Println("Server running on http://localhost:8080")
	// 正常的错误处理方式
	if err := myServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("server error: ", err)
	}
}
