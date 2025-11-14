package main

import (
	"fmt"
	"net/url"
)

func main() {
	u := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		User:   url.UserPassword("root", "123456"),
	}

	fmt.Println(u.String())   // https://root:123456@example.com
	fmt.Println(u.Redacted()) // https://root:xxxxx@example.com

}
