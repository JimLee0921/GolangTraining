package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 手动创建 Values
	v := url.Values{}

	// Add：追加多个值
	v.Add("tag", "go")
	v.Add("tag", "web")

	// Set：覆盖掉原有所有值（变成单值）
	v.Set("page", "1")

	// Get：取第一个值
	fmt.Println("tag Get:", v.Get("tag"))   // go
	fmt.Println("page Get:", v.Get("page")) // 1

	// Has：检查 key 是否存在
	fmt.Println("Has(tag):", v.Has("tag"))     // true
	fmt.Println("Has(debug):", v.Has("debug")) // false

	// Del：删除某个 key
	v.Del("page")
	fmt.Println("Has(page):", v.Has("page")) // false

	// Encode：编码回查询字符串
	fmt.Println("Encode:", v.Encode()) // tag=go&tag=web
}
