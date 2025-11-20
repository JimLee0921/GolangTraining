package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	/*
		除了使用 HandleFunc 传入函数，还可以 http.HandlerFunc 来创建 Handler
	*/
	myHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch {
		case request.URL.Path == "/":
			fmt.Fprintln(writer, "this is home")
		case strings.HasPrefix(request.URL.Path, "/api/"):
			fmt.Fprintln(writer, "api:", request.URL.Path)
		default:
			http.NotFound(writer, request) // 使用快捷函数返回 404
		}
	})
	// 注册使用
	log.Fatal(http.ListenAndServe(":8080", myHandler))
}
