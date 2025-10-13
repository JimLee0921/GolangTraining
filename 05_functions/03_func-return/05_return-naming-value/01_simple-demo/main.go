package main

import "fmt"

// main 返回值进行命名
func main() {
	name := concatName("Jim", "Lee")
	fmt.Println(name)

	area, perimeter := rectangle(4, 5)
	fmt.Println(area, perimeter)
}

// 单个命名返回值
func concatName(firstName, lastName string) (fullName string) {
	fullName = firstName + " " + lastName
	return // 等同于 return fullName
}

// 多个命名返回值
func rectangle(width, height int) (area, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}
