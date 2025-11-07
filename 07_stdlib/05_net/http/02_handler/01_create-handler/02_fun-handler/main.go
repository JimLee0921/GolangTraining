package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {
	http.HandleFunc("/hello", Hello) // 这一步就是适配器，没有创建路由所以绑定到默认
	http.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "hi")
	})
	http.ListenAndServe(":8080", nil) // 这里 handler 传入 nil 就使用默认的路由
}
