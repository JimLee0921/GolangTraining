package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "this is index, hello!")
}

func main() {
	http.HandleFunc("/", hello) // 首页兜底（未找到路径就会显示这个路由内容）
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api subtree:", r.URL.Path)
	}) // 子树匹配：可以在 http://localhost:8080/api/ 后面加任何内容
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok") // /api/health 精确匹配
	})
	// 传入 nil 使用全局默认路由器但是使用包级函数或注册路由（这个必须放在HandleFunc之后，不然路由全是404）
	log.Fatal(http.ListenAndServe(":8080", nil)) // 如果服务器启动或运行出错，会打印错误并退出

}
