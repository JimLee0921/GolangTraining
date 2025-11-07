package main

import "net/http"

func main() {
	/*
		手动创建一个 HTTP server 并且 handler 传入 nil
		测试使用的就是 全局默认路由器 http.DefaultServeMux
		此时访问 http://localhost:8080/ 就是 404 错误
	*/
	http.ListenAndServe(":8080", nil)
}
