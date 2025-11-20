package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 路径拼接
	result, err := url.JoinPath("https://baidu.com/api", "v1", "user")
	if err != nil {
		panic(err)
	}
	fmt.Println(result) // https://baidu.com/api/v1/user
}
