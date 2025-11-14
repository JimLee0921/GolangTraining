package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, _ := url.Parse("https://example.com/files/%E6%96%87%E4%BB%B6")
	fmt.Println(u.Path)          // /files/文件
	fmt.Println(u.EscapedPath()) // /files/%E6%96%87%E4%BB%B6
}
