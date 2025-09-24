package main

import "fmt"

// main 切片表达式
func main() {
	/*
		切片表达式可以从数组或切片里截取一段子切片，但是子切片修改会影响原数组或切片因为共享底层存储
			普通切片表达式：
				s[low : high]（结果是一个新的切片）
				low：起始索引（包含），省略就是 0 开始
				high：结束索引（不包含），省略就是到 len(s)
				结果切片的：
					长度 = high - low
					容量 = cap(s) - low
			全切片表达式
				s[low : high : max]（返回的也是一个新的切片）
				low：起始索引（包含）
				high：结束索引（不包含）
				max：容量上限（不包含）
				结果切片的：
					长度 = high - low
					容量 = max - low
			全切片表达式是为了：
				1. 控制容量：限制新切片的可用容量，避免影响原切片后续的数据，
				2. 防止 append 时覆盖原数据
				3. 但并不是彻底“独立拷贝”，它依然共享底层数组，只是 容量(cap) 被限制了。
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
