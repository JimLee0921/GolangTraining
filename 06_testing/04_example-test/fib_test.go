package main

import (
	"fmt"
)

// Example 测试函数名以 Example 开头
func ExampleFib() {
	fmt.Println(Fib(0))
	fmt.Println(Fib(1))
	fmt.Println(Fib(5))
	fmt.Println(Fib(10))
	// Output:
	// 0
	// 1
	// 5
	// 55
}

/*
	go test -v ./06_testing/04_example-test/
		=== RUN   ExampleFib
		--- PASS: ExampleFib (0.00s)
		PASS
		ok      github.com/JimLee0921/GolangTraining/06_testing/04_example-test 0.008s
*/
