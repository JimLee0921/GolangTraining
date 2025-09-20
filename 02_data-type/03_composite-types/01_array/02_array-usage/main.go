package main

import "fmt"

func main() {
	// 定义一个数组
	numberArray := [10]int{0: 55, 2: 22, 6: 2}
	stringArray := [...]string{"go", "python", "java", "c#"}
	fmt.Println(numberArray)

	// 1. 访问数组下标获取指定下标的值（下标从 0 开始）
	fmt.Println(numberArray[2])

	// 2. 使用 len 获取数组长度
	fmt.Printf("stringArray的长度为: %d\n", len(stringArray))

	// 3. 通过下标修改数组的某一个值
	numberArray[5] = 123
	fmt.Println(numberArray)

	// 4. 遍历数组：用 for 循环或 for ... range
	for i := 0; i < len(stringArray); i++ {
		fmt.Printf("stringArray的第 %d 个值为: %v\n", i+1, stringArray[i])
	}
	// 可配合 break
	for i, v := range numberArray {
		fmt.Printf("numberArray[%d]: %v\n", i, v)
		if i > 6 {
			break
		}
	}
}
