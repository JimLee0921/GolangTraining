package main

import "fmt"

// main 函数参数传入指针
func main() {
	/*
		指针参数，更详细指针参数见 passing-pointer 章节
	*/
	num := 20
	// 由于传入的是值，无法真正修改 num
	updateValue(num)
	fmt.Println(num) // 20
	// 传入 num 的地址，真正修改了 num
	updateValueUsePoint(&num)
	fmt.Println(num) // 30

}

func updateValueUsePoint(value *int) {
	*value += 10 // 等价于 *value = *value + 10
}

func updateValue(value int) {
	value = value + 10

}
