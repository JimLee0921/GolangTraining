package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
)

func main() {
	// 使用 CookieJar 保持状态
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	// 第一次请求会接受 set-cookie
	client.Get("https://httpbin.org/cookies/set?token=abc123")

	// 第二次请求自动携带 Cookie
	resp, _ := client.Get("https://httpbin.org/cookies")
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
