package main

import "fmt"

// main 返回值进行命名
func main() {
	/*
		https://www.ardanlabs.com/blog/2013/10/functions-and-naked-returns-in-go.html
		在函数声明里，可以为返回值进行
		这样返回值会在函数体里 自动声明为局部变量，可以直接为它们赋值
		func 函数名(参数) (返回名1 类型1, 返回名2 类型2) {
			// 在函数体内可以直接使用返回名1、返回名2
			return // 不写任何变量，自动返回
		}
		返回值变量在函数体中可直接使用
		不需要 var 声明，Go 会自动生成。
			return 可以省略变量名，如果已经给返回值变量赋过值，直接写 return 就能返回。
		可读性提升，但要注意滥用
			在函数逻辑比较长的时候，命名返回值能让 return 简洁一些
			但如果函数太复杂，不建议使用，可读性会降低
	*/
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
func rectangle(width, height int) (area int, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}
