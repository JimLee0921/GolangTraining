package main

import (
	"fmt"
	"net/http"
)

type UserHandler struct {
	DBName string
}

func (h UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "using database: %s", h.DBName)
}

func main() {
	// 结构体可以 带状态（保存变量、依赖、配置、数据库连接等）
	// 意味着 Handler 可以和业务上下文绑定，这就是可维护的工程代码风格
	handler := UserHandler{DBName: "users.db"}
	http.ListenAndServe(":8080", handler)
}
