package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
)

func main() {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	// 第一次请求设置 cookie
	client.Get("https://httpbin.org/cookies/set?token=abc123")

	// 第二次请求自动带上 cookie
	resp, _ := client.Get("https://httpbin.org/cookies")
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	fmt.Println(string(data))

}
