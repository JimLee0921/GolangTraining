package main

import (
	"fmt"
	"io"
	"net/http"
)

// main http 忽略错误
func main() {
	res, _ := http.Get("https://www.baidu.com/")
	page, _ := io.ReadAll(res.Body)
	_ = res.Body.Close()
	fmt.Printf("%s", page)
}
