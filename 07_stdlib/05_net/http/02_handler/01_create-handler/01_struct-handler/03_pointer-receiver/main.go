package main

import (
	"fmt"
	"net/http"
)

type CounterHandler struct {
	Count int
}

// 指针接收者来修改内部状态
func (c *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.Count++
	fmt.Fprintf(w, "Visited %d times\n", c.Count)
}

func main() {
	// 使用指针接收者这里也需要取地址传入
	http.ListenAndServe(":8080", &CounterHandler{Count: 10})
}
