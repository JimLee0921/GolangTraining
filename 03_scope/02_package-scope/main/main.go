package main

import (
	"fmt"

	vis "GolangTraining/03_scope/02_package-scope/visibility"
)

// main 演示如何在包作用域下使用可导出的标识符。
func main() {
	fmt.Println("main中的 vis.MyName", vis.MyName)
	//fmt.Println(vis.secret) 这里不能使用 secret
	vis.PrintVar()
}
