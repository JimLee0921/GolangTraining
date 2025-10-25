package main

import "fmt"

func main() {
	// 空格，换行，TAB 等空白符作为分隔符
	var name string
	var age int
	input := "Tom 19"
	n, err := fmt.Sscan(input, &name, &age)
	fmt.Println("read params count: ", n, "err: ", err)
	fmt.Println(name, age)
}
