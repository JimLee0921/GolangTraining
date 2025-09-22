package main

import "fmt"

// main Sprint 系列的辅助函数，它们用于构建字符串而不是直接打印
func main() {
	name := "Gopher"
	level := 7
	/*
		Sprint -> 拼接成字符串（无换行）
		Sprintln -> 拼接成字符串（结尾带换行）
		Sprintf -> 格式化拼接成字符串（最灵活）
		这些方法都不会直接打印，而是返回字符串 -> 可以存起来或再加工
	*/
	sprint := fmt.Sprint("Player:", " ", name, " level ", level)
	fmt.Println(sprint)

	sprln := fmt.Sprintln("Row", 1, "=>", 2)
	fmt.Printf("Sprintln keeps trailing newline: %q\n", sprln)

	formatted := fmt.Sprintf("Pi ~= %.3f", 3.14159)
	fmt.Println(formatted)

	report := fmt.Sprintf("%s | %s", sprint, formatted)
	fmt.Println(report)
}
