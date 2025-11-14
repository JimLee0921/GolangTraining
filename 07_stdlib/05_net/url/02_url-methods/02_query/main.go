package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 获取查询参数
	u, _ := url.Parse("https://example.com/search?q=go&lang=en")
	q := u.Query() // url.Values = map[string][]string

	fmt.Println(q.Get("q"))    // go
	fmt.Println(q.Get("lang")) // en

}
