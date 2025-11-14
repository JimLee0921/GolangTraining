package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 手动使用 client.Do() 必须使用 http.NewRequest 或 http.NewRequestWithContext 手动构造 request
	req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)
	if err != nil {
		panic(err)
	}

	// 设置一些请求头
	req.Header.Set("User-Agent", "Go-HttpClient/666.0")

	// 使用默认客户端发送
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	fmt.Println(string(data))
}
