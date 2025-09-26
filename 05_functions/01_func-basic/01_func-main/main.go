package main

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("init one ...")
}

func init() {
	fmt.Println("init two ...")
}

func main() {
	/*
		在 Go 里，程序的入口点就是 main 函数，但它必须满足两个条件：
			它所在的 包名 必须是 package main，Go 的规定：只有 package main 才能编译为可执行文件
			函数名必须是 main，没有参数，没有返回值，如果要获取命令行参数，需要使用 os.Args
		package main -> 表示这是一个独立的可执行程序（executable program），不是库（library）
		func main() -> 表示这个程序的入口点，编译后运行时会先执行这里
		退出程序要用 os.Exit(code)code 为 0 表示正常退出，非 0 表示异常
		init 函数：
			在 main 包里，还可以定义 init() 函数
			它会在 main() 之前自动执行，可以有多个 init，按声明顺序运行
			常用于初始化操作
	*/

	fmt.Println("func main is the entry point to your program")
	fmt.Println("命令行参数:", os.Args)
	os.Exit(1) // 携带问题退出

}
