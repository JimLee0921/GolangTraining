package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 在已有 URL 上解析相对或绝对路径
	base, _ := url.Parse("https://example.com/api/v1/")

	// 解析相对路径，会继承 base 的 scheme 与 host
	fmt.Println(base.Parse("../product?id=100")) // https://example.com/api/product?id=100 <nil>
	fmt.Println(base.Parse("/product?id=100"))   // https://example.com/product?id=100 <nil>
	fmt.Println(base.Parse("./product?id=100"))  // https://example.com/api/v1/product?id=100 <nil>

	// 解析绝对路径，不继承 base
	abs, _ := base.Parse("https://foo.com/x")
	fmt.Println(abs.String()) // https://foo.com/x
}
