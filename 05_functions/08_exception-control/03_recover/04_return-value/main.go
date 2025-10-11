package main

import "fmt"

func main() {
	/*
		如果 panic 程序会从 defer 执行完毕后正常返回
		所以可以在 defer 中修改命名返回值
	*/
	fmt.Println("result:", f())
}

func f() (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
			result = -1 // 修改返回值
		}
	}()
	panic("failed")
	return 1
}
