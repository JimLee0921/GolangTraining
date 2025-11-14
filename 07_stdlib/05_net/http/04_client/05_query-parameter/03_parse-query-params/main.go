package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 解析 URL 中的 query
	u, _ := url.Parse("https://api.example.com/search?q=go+lang&page=2")
	q := u.Query()
	fmt.Println(q.Get("q"))
	fmt.Println(q.Get("page"))
}
