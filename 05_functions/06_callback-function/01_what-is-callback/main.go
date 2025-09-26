package main

import "fmt"

// main 回调函数
func main() {
	/*
		回调函数就是 把函数当作参数传给另一个函数，等到某个时机再去调用
		在 Go 里，函数是一等公民（first-class citizen），可以作为参数传递，也可以作为返回值
		所以回调函数本质就是：函数接收另一个函数作为参数，并在内部调用它
	*/
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 8, 5, 65, 2, 2, 2, 504}
	visit(data, func(i int) {
		fmt.Println(i)
	})
}

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}
