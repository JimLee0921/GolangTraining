package main

import "fmt"

// main int 指针传递
func main() {
	/*
		int bool 指针传递才能修改原值
	*/
	num := 20
	fmt.Println(&num)
	changeNum(num)
	fmt.Println(num)         // 20
	changeNumByPointer(&num) // 0xc00000a088
	fmt.Println(num)         // 40

	isTrue := false
	flip(&isTrue)
	fmt.Println(isTrue) // true

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

func flip(pb *bool) {
	*pb = !*pb
}
