package main

import "fmt"

func main() {
	isLogin := true
	hasPermission := false

	if isLogin && hasPermission {
		fmt.Println("允许访问")
	} else {
		fmt.Println("拒绝访问")
	}
}
