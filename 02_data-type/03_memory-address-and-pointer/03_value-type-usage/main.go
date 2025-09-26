package main

import "fmt"

// main 值类型指针传递
func main() {
	/*
		函数传值时默认时拷贝一份变量的值进行传递，而不是变量的地址
		值类型（int, string, float64, bool, array, struct等）在函数中修改想要影响原数据需要传地址，然后在函数中使用指针进行修改
	*/
	num := 20
	fmt.Println(&num)
	changeNum(num)
	fmt.Println(num)         // 20
	changeNumByPointer(&num) // 0xc00000a088
	fmt.Println(num)         // 40

}

// 传值（这里修改不会影响到原数据）
func changeNum(num int) {
	fmt.Println(num) // 20
	num += num
	fmt.Println(num) // 40

}

// 传址（直接修改原数据）
func changeNumByPointer(num *int) {
	fmt.Println(num) // 0xc00000a088
	*num += *num
	fmt.Println(num) // 0xc00000a088
}
