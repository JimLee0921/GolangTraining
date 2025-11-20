package main

import (
	"fmt"
	"net/url"
)

func main() {
	// PathEscape / PathUnescape 路径拼接转义
	escaped := url.PathEscape("我的文件+file.pdf")
	fmt.Println(escaped) // %E6%88%91%E7%9A%84%E6%96%87%E4%BB%B6+file.pdf

	orig, _ := url.PathUnescape(escaped)
	fmt.Println(orig) // 我的文件+file.pdf
}
