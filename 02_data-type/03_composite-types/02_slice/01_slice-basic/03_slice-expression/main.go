package main

import "fmt"

// main 切片表达式
func main() {
	/*

	 */
	// 普通切片
	originalNumberArr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sonSlice1 := originalNumberArr[1:9]                                          // 从下标 1 到 下标 9
	sonSlice2 := originalNumberArr[:5]                                           // 省略 low 默认为下标 0
	fmt.Println(sonSlice1)                                                       // [1 2 3 4 5 6 7 8]
	fmt.Println(sonSlice2)                                                       // [0 1 2 3 4]
	fmt.Println("sonSlice1 -> ", "长度为:", len(sonSlice1), "容量为:", cap(sonSlice1)) // 8 9
	fmt.Println("sonSlice2 -> ", "长度为:", len(sonSlice2), "容量为:", cap(sonSlice2)) // 5 10
	// 全切片
	originalStringSlice := []string{"Java", "Python", "Go", "Java", "Ruby"}
	sonSlice3 := originalStringSlice[1:3:4]
	fmt.Println(sonSlice3)                                        // [1, 2]
	fmt.Println("长度为: ", len(sonSlice3), "容量为: ", cap(sonSlice3)) // 2	3
}
