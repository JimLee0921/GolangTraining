package main

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	/*
		half 函数传入一个 int n
			返回：
				n/2（整除的结果，丢弃小数部分）
				n%2 == 0（一个布尔值，判断 n 是否是偶数）
	*/
	h, even := half(5)
	fmt.Println(h, even) // 2 false
}
