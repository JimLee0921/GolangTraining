package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 不用 url 创建 query 最后于 url 进行拼接，中文和特殊符号自动 URL 编码
	params := url.Values{}
	params.Set("page", "1")
	params.Set("sort", "desc")
	fullURL := "https://api.example.com/list?" + params.Encode()
	fmt.Println(fullURL)
}
