package main

import "fmt"

func main() {
	/*
		使用切片展开式传参给可变参数
	*/
	data := []int{1, 2, 4, 5, 6, 4, 8, 9, 0}
	res := sum(data...)
	fmt.Println("result:", res)
}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
