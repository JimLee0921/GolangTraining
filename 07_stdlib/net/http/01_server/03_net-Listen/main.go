package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {
	// 1. 自定义路由
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello World!")
	})

	// 2. 手动监听端口（获得一个 net.Listener）
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	// 3. 构造 Server（但是不指定端口）
	srv := &http.Server{
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	fmt.Println("Server running on http://localhost:8080")
	// 4. 使用 Serve() 启动服务而不是 ListenAndServe()
	if err = srv.Serve(ln); err != nil && err != http.ErrServerClosed {
		fmt.Println("server error:", err)

	}
}
