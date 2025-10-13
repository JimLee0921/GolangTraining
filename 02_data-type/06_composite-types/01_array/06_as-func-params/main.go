package main

import "fmt"

func sum(arr [3]int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	arr[1] = 3 // 这里修改不会对原数组生效，因为这里是复制了一份值，而不是原数组的地址
	return total
}

func main() {
	/*
		还是值类型的特性
	*/
	a := [3]int{1, 2, 3}
	fmt.Println(sum(a)) // 6
	fmt.Println(a)      // [1 2 3]
}
