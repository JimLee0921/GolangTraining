package math

import (
	"fmt"
)

// Example 测试函数名
func ExampleAdd() {
	fmt.Println(Add(10, 10))
	// Output:
	// 20
}

// 多次测试多个 output
func ExampleSub() {
	fmt.Println(Sub(10, 10))
	fmt.Println(Sub(1, 10))
	// Output:
	// 0
	// -9
}
