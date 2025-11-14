package main

import (
	"fmt"
	"net/url"
)

func main() {
	// QueryEscape / QueryUnescape 查询参数转义
	q := url.QueryEscape("张 3")
	fmt.Println(q) // %E5%BC%A0+3 空格转义为 +

	orig, _ := url.QueryUnescape(q)
	fmt.Println(orig) // 张 3
}
