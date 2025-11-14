package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 修改 query 标准流程，注意 Query() 返回的是副本，所以修改后必须写回 RawQuery
	u, _ := url.Parse("https://example.com/search?q=go&lang=en")

	// 第 1 步：取 Query 副本
	q := u.Query()

	// 第 2 步：修改 Query
	q.Set("lang", "zh") // 覆盖 lang 参数
	q.Add("tag", "backend")
	q.Add("tag", "api")

	// 第 3 步：Encode 回 RawQuery
	u.RawQuery = q.Encode()

	fmt.Println(u.String())
}
