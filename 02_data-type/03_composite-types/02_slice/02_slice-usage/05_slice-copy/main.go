package main

import "fmt"

// main 切片常用操作
func main() {
	/*
		拷贝与克隆
			切片（slice）是一个 轻量级的数据结构，底层包含三个字段：
				指针：指向底层数组的数据起始位置
				长度（len）：切片中实际元素个数
				容量（cap）：从起始位置到底层数组末尾的容量
			浅拷贝（直接赋值）
				因为切片共享底层数组，所以直接赋值时不会产生真正的拷贝
				切片的赋值本质上是浅拷贝，所以直接赋值当一个切片被修改另外一个也会被修改
			深拷贝（copy函数）
				Go 提供了内置函数 copy，用来把一个切片的元素复制到另一个切片中
				创建新的底层数组，两个切片互不影响
				使用 copy 函数：copy(dst, src)
					返回值是 实际复制的元素个数 = min(len(dst), len(src))
					dst 必须是一个切片（slice），而且 要有足够的长度（len）来容纳元素
					新切片的容量不需要关注，但是如果 len 为 0 则什么都不会拷贝
			使用 append 进行拷贝
				clone := append([]int(type), src...)
				append 会自动分配足够的空间，写法简洁，不容易出错
				[]type(nil)：
					表示一个 nil 切片，类型是 []type，它没有底层数组，长度和容量都是 0
				src...
					... 是 切片展开（unpack） 语法，把 src 切片里的每个元素依次作为 append 的参数
				append 用于把元素追加到一个新切片里
				因为最开始的切片是 []int(nil)，里面没有元素，容量为 0，所以 append 时必须分配新的底层数组
				这样一来，新切片的底层数组和 src 的底层数组完全不同。

	*/
	// 浅拷贝：slice1 和 slice1 指向同一个底层数组，改变其中一个，另外一个也会被修改
	slice1 := []int{1, 2, 3}
	slice2 := slice1 // 赋值（引用拷贝）
	slice2[0] = 100
	fmt.Println(slice1) // [100 2 3]
	fmt.Println(slice2) // [100 2 3]
	slice1[1] = 666
	fmt.Println(slice1) // [100 666 3]
	fmt.Println(slice2) // [100 666 3]

	// copy 深拷贝
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, len(slice3)) // 分配新切片，长度必须至少相同
	fmt.Println(copy(slice4, slice3))  // 把 slice3 的元素拷贝到 slice4 返回实际元素的个数为 3
	fmt.Println(slice3)
	fmt.Println(slice4)
	slice3 = append(slice3, 4, 5, 6)
	fmt.Println(slice3)
	fmt.Println(slice4) // 不受影响

	// append 深拷贝
	stringSlice := []string{"hello", "bro", "shu"}
	clonedStringSlice := append([]string{}, stringSlice...)
	fmt.Println(stringSlice)
	fmt.Println(clonedStringSlice)
	stringSlice[2] = "Lee"
	fmt.Println(stringSlice)
	fmt.Println(clonedStringSlice)
}
