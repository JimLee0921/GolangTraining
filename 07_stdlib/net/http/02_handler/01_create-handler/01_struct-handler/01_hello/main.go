package main

import (
	"fmt"
	"net/http"
)

// HelloHandler 1. 定义结构体
type HelloHandler struct{}

// 2. 实现 ServeHTTP 方法
func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Handler")
}

func main() {
	// 3. 传入自定义的 HelloHandler 结构体
	http.ListenAndServe(":8080", HelloHandler{})
}
