package main

import (
	"fmt"
)

// Example 测试函数名以 Example 开头，里面有一次调用，结果也需要有一个
func ExampleAdd() {
	fmt.Println(Add(10, 10))
	// Output:
	// 100
}

// 多次测试多个 output
func ExampleSub() {
	fmt.Println(Sub(10, 10))
	fmt.Println(Sub(1, 10))
	// Output:
	// 0
	// -9
}
