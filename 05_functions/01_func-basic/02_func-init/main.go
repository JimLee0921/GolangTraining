package main

import (
	"fmt"

	"github.com/JimLee0921/GolangTraining/05_functions/01_func-basic/02_func-init/utils"
)

func init() {
	fmt.Println("init() in main.go")
}

func main() {
	/*
		导入包 -> 初始化包
		当 main 导入 utils 时，Go 会先执行：
			所有全局变量初始化
			然后执行 utils 包内所有 init()
			文件顺序 按编译顺序（通常按文件名排序）
		再执行 main 包的 init()
		最后执行 main()
	*/
	fmt.Println("main() started")
	utils.A()
	/*
		init() in utils/a.go
		init() in utils/b.go
		init() in main.go
		main() started
		A() called
	*/
}
