package main

import "fmt"

// main 切片常用操作
func main() {
	/*
		slice... 语法：切片展开式（variadic argument expansion），主要用在 函数调用（可变参数函数） 或 append（拼接切片） 里
	*/
	books := []string{"python", "go", "rust"}

	// 1. 通过下标访问数据（即使 slice 的长度会变化，访问时还是不能下标越界）
	fmt.Println(books[0])

	// 2. 通过下标修改数据
	books[1] = "java"
	fmt.Println(books)

	// 3. 访问当前切片长度和容量（len()和cap()方法）
	fmt.Printf("长度为: %d, 容量为: %d\n", len(books), cap(books))

	/*
		3.追加新的元素
		使用 append 方法，触发扩容返回一个新的切片需要接收（长度和容量都会变化）
		扩容策略是实现细节，容量不够会分配更大数组并复制旧数据，append 返回的新切片要接回变量
		切片的容量扩容规则不是固定翻倍，而是由底层实现（runtime）决定的，跟当前容量和需要的大小有关
		扩容规则大致是：
			如果 append 后的长度 ≤ 原来容量的 2 倍，通常就直接 翻倍
			如果 append 后的长度大于 2 倍容量，就会直接扩容到所需长度
			当切片容量比较大时（比如 >1024），扩容增量会逐渐变小（不是严格翻倍，而是接近 1.25 倍）
		append 默认在末尾进行插入
		开头插入：
			思路：先把新元素放前面，再拼接原切片：
			s = append([]int{0}, s...)
			fmt.Println(s) // [0 1 2 3 4]
			注意：这里会分配新切片。
		中间插入一个或多个元素（中间追加多个元素把 99 换为 多个元素即可）
			比如在索引 2 位置插入 99：
				i := 2
				s = append(s[:i], append([]int{99}, s[i:]...)...)
				fmt.Println(s) // [1 2 99 3 4]
			步骤：
				s[:i] → 插入点之前的部分
				[]int{99} → 要插入的元素
				s[i:] → 插入点及之后的部分
				append 两次拼接
		中间只插一个元素，也可以利用切片容量移动后半部分
			s := []int{1, 2, 3, 4}
			i := 2
			s = append(s, 0)              // 扩展一位，这里是补一个类型的空值即可，string就是 ""
			copy(s[i+1:], s[i:])          // 后半部分整体向后移动
			s[i] = 99                     // 插入新值
			fmt.Println(s) // [1 2 99 3 4]
			特点：
				只申请一次 append 的扩容
				更高效，适合频繁插入
	*/
	books = append(books, "C++") // 默认追加到结尾
	fmt.Println(books)
	books = append([]string{"C#"}, books...) // 追加到开头
	fmt.Println(books)
	books = append(books[:3], append([]string{"javascript", "docker"}, books[3:]...)...) // 追加到中间指定索引位置上多个元素
	books = append(books, "")
	fmt.Println(books)
	copy(books[3:], books[4:])
	books[3] = "react"
	fmt.Println(books)
	fmt.Printf("长度为: %d, 容量为: %d\n", len(books), cap(books)) // 这里容量为 6

	/*
		4. 切片表达式
			普通切片表达式：
				s[low : high]（结果是一个新的切片）
				low：起始索引（包含）
				high：结束索引（不包含）
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
				1. 控制容量：限制新切片的可用容量，避免影响原切片后续的数据。
				2. 防止 append 时覆盖原数据。
	*/
	// 普通切片
	originalArr := [5]int{0, 1, 2, 3, 4}
	sonArr := originalArr[1:3]
	fmt.Println(sonArr)                                           // [1 2]
	fmt.Println("长度为: ", len(sonArr), "容量为: ", cap(sonArr)) // 2	4
	// 全切片
	fullOriginalArr := [5]int{0, 1, 2, 3, 4}
	fullSonArr := fullOriginalArr[1:3:4]
	fmt.Println(fullSonArr)                                               // [1, 2]
	fmt.Println("长度为: ", len(fullSonArr), "容量为: ", cap(fullSonArr)) // 2	3

	/*
		5. 拷贝与克隆
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
				另外一种克隆写法（不需要关注长度/容量细节）
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
