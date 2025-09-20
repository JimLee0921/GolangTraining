package main

import (
	"fmt"
	"os"
)

// main 展示了 Fprint 的不同变体如何写入到不同的 io.Writer 目标
func main() {
	s := "This is a log"
	/*
		Create 函数
		参数：name -> 创建的文件名，可以带路径
		返回值：os.File -> 文件句柄，可读可写。error -> 是否出错，正常为空
		如果文件 不存在，会创建一个新文件
		如果文件 已存在，会清空原内容（覆盖写）
		创建后的文件默认是 可写、可读 的
	*/
	file, err := os.Create("temp_files/log.txt")
	fmt.Print(file, err)
	// 在 main 函数结束后再关闭文件
	defer file.Close()
	fmt.Fprintln(file, "File Log: 写入到文件中")
	fmt.Fprintf(file, "File log: %s - %d", s, 2025)
}
