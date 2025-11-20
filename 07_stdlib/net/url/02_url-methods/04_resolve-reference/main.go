package main

import (
	"fmt"
	"net/url"
)

func main() {
	base, _ := url.Parse("https://example.com/api/")
	ref, _ := url.Parse("v1/user")

	fmt.Println(base.ResolveReference(ref)) // https://example.com/api/v1/user
}
