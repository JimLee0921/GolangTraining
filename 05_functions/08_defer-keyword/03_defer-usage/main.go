package main

import "fmt"

func main() {
	num := 3
	numPower(&num)
	fmt.Println(num)     // 9
	defer numPower(&num) // 函数结束前执行，所以后续打印的还是之前的 num
	fmt.Println(num)     // 9
}

func numPower(n *int) {
	*n = *n * *n
}
