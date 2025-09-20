package main

import "fmt"

// main 演示 fmt.Print、fmt.Println 和 fmt.Printf 的基本区别。
func main() {
	// Print 不会附加空格或换行，需要手动补齐
	fmt.Print("Hello")
	fmt.Print(" Go")
	fmt.Print("!\n")

	// Println 会在参数之间加入空格并自动换行
	fmt.Println("Hello", "Go", 2025)

	// Printf 使用格式化动词控制输出内容和精度
	name := "Gopher"
	score := 99.5
	fmt.Printf("%s scored %.1f%% on the test.\n", name, score)
}
