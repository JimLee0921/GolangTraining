package main

import (
	"fmt"
	"net/http"
	"time"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ping")
}
func now(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, time.Now())
}

// 将一整个子路由交给另一个处理器，使用 Handle 和 http.StripPrefix
func main() {
	api := http.NewServeMux()
	api.HandleFunc("/v1/ping", ping)
	api.HandleFunc("/v1/time", now)

	root := http.NewServeMux()
	root.Handle("/api/", http.StripPrefix("/api", api)) // /api/v1/* -> api 的 /v1/*
	http.ListenAndServe(":8080", root)

}
