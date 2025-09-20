package main

import "fmt"

func main() {
	/*
		这里如果运行 `go run .\xxx\xxx\main.go` 会报错 x 未定义
		但是如果运行 `go run .\xxx\xxx\` 就可以打印成功
			这里和 Go 的编译顺序和包作用域有关：
				包级作用域（package scope） 是整个包的范围，不限于某一个源文件
				当编译器编译一个包时，它会把同一个包下面的所有 .go 文件一起编译，并且按作用域规则解析标识符
				所以虽然把 var x = 1 放在另一个文件里，但它仍然在 main 包的作用域内，main.go 自然就能访问它
	*/
	fmt.Println(x)
}
