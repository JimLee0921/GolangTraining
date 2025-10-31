package main

import "fmt"

func main() {
	var name string
	var age int
	input := "Tom 18\nJim 29"
	// 遇到 \n 换行符立即停止
	n, err := fmt.Sscanln(input, &name, &age)
	fmt.Println("read params count:", n, "err:", err)
	fmt.Printf("result：%s, %d\n", name, age)
}
