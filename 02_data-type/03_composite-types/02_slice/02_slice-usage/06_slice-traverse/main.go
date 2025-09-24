package main

import "fmt"

// main 切片遍历
func main() {
	/*
		可以使用 for 或 for range 语法进行遍历
	*/
	stringSlice := []string{"Hello", "World", "DSB", "DSB", "JimLee", "Jane"}
	for i := 0; i < len(stringSlice); i++ {
		fmt.Println(i, stringSlice[i])
	}

	for i, v := range stringSlice {
		fmt.Println(i, v)
	}

}
