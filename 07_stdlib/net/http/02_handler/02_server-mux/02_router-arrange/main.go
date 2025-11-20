package main

import (
	"fmt"
	"net/http"
)

// --- User Handlers ---

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "user login")
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "user logout")
}

func profile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "user profile")
}

// --- Order Handlers --

func create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "order create")
}

func detail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "order detail")
}

// --- Route Register Functions ---

func registerUser(mux *http.ServeMux) {
	mux.HandleFunc("/user/login", login)
	mux.HandleFunc("/user/logout", logout)
	mux.HandleFunc("/user/profile", profile)
}

func registerOrder(mux *http.ServeMux) {
	mux.HandleFunc("/order/create", create)
	mux.HandleFunc("/order/detail", detail)
}

// 使用路径前缀进行分组
func main() {
	// 1. 创建 Mux
	mux := http.NewServeMux()

	// 2. 使用路由注册函数进行路由注册
	registerUser(mux)
	registerOrder(mux)

	// 3. 启动服务
	http.ListenAndServe(":8080", mux)
}
