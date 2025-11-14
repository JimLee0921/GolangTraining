package main

import (
	"fmt"
	"log"
	"net/http"
)

// indexHandler 返回 r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "url.path = %q\n", req.URL.Path)
}

// helloHandler 返回请求头 r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "header[%q] = %q\n", k, v)
	}
}

func main() {
	// 注册路由
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	// 启动 HTTP 服务监听本机 9999 端口
	log.Fatal(http.ListenAndServe(":9999", nil))
}
