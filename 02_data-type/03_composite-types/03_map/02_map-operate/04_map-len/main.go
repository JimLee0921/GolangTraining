package main

import "fmt"

func main() {
	/*
		可以使用 len() 函数查询 map 长度
	*/
	myMap := map[string]int{"go": 1, "rust": 2, "python": 0}
	fmt.Println(len(myMap))
	myMap["java"] = 2
	myMap["javascript"] = 3
	fmt.Println(len(myMap))
	delete(myMap, "go")
	fmt.Println(len(myMap))

}
