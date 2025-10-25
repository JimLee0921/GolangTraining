package main

import (
	"fmt"
	"os"
)

func main() {
	// 下面两个是等价的
	fmt.Println("This goes to standard output") // 默认就是输出到 os.Stdout
	fmt.Fprintln(os.Stdout, "This goes to standard output")

	/*
		使用 > 或者 >> 操作符可以重定向 Stdout打指定文件
		不同在于 > 写入文件默认覆盖	>> 是追加到文件末尾（文件不存在会自动创建）
		go run main.go > out.txt
		go run main.go >> out.txt
	*/
	fmt.Fprintln(os.Stdout, "Hello")
	fmt.Fprintln(os.Stdout, "Brother")

	// 写入文件和屏幕的区别
	//f, _ := os.Create("temp_files/output.txt")
	//fmt.Fprintln(f, "Write to file")            // 写入文件
	//fmt.Fprintln(os.Stdout, "Write to console") // 默认会终端显示

}
