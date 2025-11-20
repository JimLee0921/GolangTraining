package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, _ := url.Parse("https://example.com:8080/path")
	fmt.Println(u.Hostname()) // example.com
	fmt.Println(u.Port())     // 8080
}
