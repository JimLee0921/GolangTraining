package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlStr := "https://example.com:8080/v1/user?id=10&debug=true#sec"
	u1, err1 := url.Parse(urlStr)
	if err1 != nil {
		panic(err1)
	}
	// 1. 使用 url.Parse 解析完整 URL
	fmt.Println("Scheme:", u1.Scheme)     // https
	fmt.Println("Host:", u1.Host)         // example.com:8080
	fmt.Println("Path:", u1.Path)         // /v1/user
	fmt.Println("Query:", u1.RawQuery)    // id=10&debug=true
	fmt.Println("Fragment:", u1.Fragment) // sec
	fmt.Printf("urlObj:%#v\n", *u1)

	// 2. 使用 url.ParseRequestURI
	u2, err2 := url.ParseRequestURI(urlStr)
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("urlObj:%#v\n", *u2) // 这里的 Fragment 是空，#sec 被解析到了 RawQuery 的位置

	// 3. 使用 url.ParseQuery（直解析查询字符串也就是 ? 之后的字典）
	queryStr := "id=10&id=20&debug=true"
	values, err3 := url.ParseQuery(queryStr)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(values["id"])        // [10 20]
	fmt.Println(values.Get("debug")) // true
}
