package main

import "fmt"

func half(n int) (float64, bool) {
	return float64(n) / 2, n%2 == 0
}
func main() {
	/*
		half 函数传入一个 int n
			返回：
				float64(n) / 2 把 int 转成 float64再 / 2
				n%2 == 0（一个布尔值，判断 n 是否是偶数）
	*/
	h, even := half(5)
	fmt.Println(h, even) // 2.5 false
}
