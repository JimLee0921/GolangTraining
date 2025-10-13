package main

import "fmt"

// main 使用 make 创建切片
func main() {
	/*

	 */
	// 分配长度和容量
	sized := make([]int, 3, 5)
	sized[0] = 42
	fmt.Printf("sized: %#v | len=%d cap=%d\n", sized, len(sized), cap(sized))

	// 仅分配容量，后续通过 append 填充
	buffered := make([]int, 0, 4)
	buffered = append(buffered, 10, 20, 30)
	fmt.Printf("buffered: %#v | len=%d cap=%d\n", buffered, len(buffered), cap(buffered))

	// 只指定长度（容量 = 长度）
	stringSlice := make([]string, 2)
	stringSlice[0] = "Hello"
	stringSlice[1] = "World"
	//stringSlice[2] = "!" 这里由于超出长度，会编译错误，必须使用 append
	stringSlice = append(stringSlice, "!")
	fmt.Println(stringSlice)

}
