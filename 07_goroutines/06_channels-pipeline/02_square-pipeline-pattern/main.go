package main

import "fmt"

// main pipeline
func main() {
	// 生成 channel
	//c := genChannel(2, 3)
	//res := square(c)
	//for i := range res {
	//	fmt.Println(i) // 4 9
	//}

	// 简写
	for i := range square(genChannel(2, 3, 4, 5, 6, 7, 8, 9)) {
		fmt.Println(i)
	}
}

func genChannel(nums ...int) <-chan int {
	/*
		返回传入的整数生成的channel
	*/
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(c <-chan int) <-chan int {
	/*
		传入 channel 计算里面每一个数字的平方并插入新的 channel 最终返回
	*/
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range c {
			out <- i * i
		}
	}()
	return out
}
