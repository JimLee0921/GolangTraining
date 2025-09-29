package main

import "fmt"

func main() {
	/*
		如果我们列出所有小于 10 的自然数中能被 3 或 5 整除的数，它们是 3、5、6 和 9，这些数的和是 23。
		问题：
			请找出所有小于 1000 的自然数中能被 3 或 5 整除的数的和
	*/
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println("the sum of all the multiples of 3 or 5 below 1000 is: ", sum)
}
