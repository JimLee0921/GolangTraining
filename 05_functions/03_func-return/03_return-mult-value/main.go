package main

import "fmt"

// main 函数有多个返回值
func main() {
	/*
		Go 中可以有多个返回值
		在函数定义中返回值部分写多个类型即可
		接收函数返回值时需要多个变量进行接收
		不需要使用的返回值可以使用 _ 进行接收
	*/
	x, y := swap(1, 2)
	fmt.Println(x, y)

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("出错啦:", err)
	} else {
		fmt.Println("结果:", result)
	}

	name, age, _ := getUser() // 第三个参数不需要接收使用 _ 占位
	fmt.Println(name, age)
}

func swap(a, b int) (int, int) {
	return b, a
}

// 错误处理模式应用场景（无错误时 error 返回 nil）
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide by zero")
	}
	return a / b, nil
}

// 返回多个不同类型
func getUser() (string, int, bool) {
	return "Tom", 20, true
}
