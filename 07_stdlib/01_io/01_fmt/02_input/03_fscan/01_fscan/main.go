package main

import (
	"fmt"
	"os"
)

/*
文件内容
Tom 18
Jerry 20
*/
func main() {
	file, err := os.Open("temp_files/log.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	var name string
	var age int
	// 按空白符（空格、换行、Tab）分隔，可以跨行读取
	for {
		n, err := fmt.Fscan(file, &name, &age)
		if n == 0 || err != nil {
			break
		}
		fmt.Printf("name: %s, age: %d\n", name, age)
	}
}
