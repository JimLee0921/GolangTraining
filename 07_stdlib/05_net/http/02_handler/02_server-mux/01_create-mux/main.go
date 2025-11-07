package main

import (
	"fmt"
	"net/http"
)

type HiHandler struct {
}

func (h HiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hi")
}

func main() {

	// 1. 自定义路由器
	mux := http.NewServeMux()
	// 2. HandleFunc 绑定路由 直接函数
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "hello")
	})

	// 3.Handle 绑定路由，必须是一个实现了 ServeHttp 的类型示例
	mux.Handle("/hi", HiHandler{})

	// 4. 启动服务
	http.ListenAndServe(":8080", mux)
}
