package main

import "fmt"

// main 使用 make 创建切片
func main() {
	/*
		make 内置函数用来创建并初始化 slice、map、chan 这三种引用类型
		对于切片（slice），make 会 分配底层数组，并返回一个切片对象
		语法：make([]Type, len, cap)：
			len：切片的长度（初始化后能直接用下标访问的元素个数）
			cap：切片的容量（底层数组大小）
		规则：必须满足 0 <= len <= cap
			只指定 len 时 len = cap
			如果 len > cap，会直接编译报错
			如果 len == cap，切片刚好占满底层数组。
			如果 len < cap，切片“预留”了一些空间，方便后续 append
			省略 cap 时，cap = len
			len=0 但 cap>0 时，切片为空，但有预留容量，只有在 0 <= index < len(slice) 时，下标赋值才合法
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
