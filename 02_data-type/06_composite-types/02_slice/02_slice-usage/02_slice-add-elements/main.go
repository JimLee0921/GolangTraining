package main

import "fmt"

// main slice 添加元素
func main() {
	/*
		切片 slice 添加元素
			使用 append 方法进行元素添加
			使用 append 触发扩容时方法会返回一个新的切片需要接收（长度和容量都会变化）
			扩容策略是实现细节，容量不够会分配更大数组并复制旧数据，append 返回的新切片必须要接回变量
			切片的容量扩容规则不是固定翻倍，而是由底层实现（runtime）决定的，跟当前容量和需要的大小有关
			扩容规则大致是：
				如果 append 后的长度小于等于 原来容量的 2 倍，通常就直接 翻倍
				如果 append 后的长度大于 2 倍容量，就会直接扩容到所需长度
				当切片容量比较大时（比如 >1024），扩容增量会逐渐变小（不是严格翻倍，而是接近 1.25 倍）
		1. append 默认在末尾进行插入一个或多个元素
			s = append(s, newElements)	// newElements 是可以同时插入多个元素
		2. append 开头插入一个或多个元素
			思路：先把新元素放前面，再拼接原切片：
			s = append([]type{newElements}, s...)
		3. append 两次拼接在中间插入一个或多个元素
			s = append(s[:i], append([]type{newElements}, s[i:]...)...)
			s[:i] -> 插入点之前的部分
			[]type{newElements} -> 要插入的元素
			s[i:] -> 插入点及之后的部分

		4. copy 中间插入一个或多个元素（手动使用 append 添加切片长度）
			newVals := []type{newElements} // 要插入的多个元素
			s = append(s, 0)              // 扩展一位，这里是补一个类型的空值即可，string就是 ""
			copy(s[i+1:], s[i:])          // 后半部分整体向后移动
			s[i] = 99                     // 插入新值
			fmt.Println(s) // [1 2 99 3 4]
			特点：
				只申请一次 append 的扩容
				更高效，适合频繁插入
	*/
	books := []string{"python", "go", "rust"}

	books = append(books, "C++", "PHP") // 默认追加到结尾
	fmt.Println(books)
	books = append([]string{"C#"}, books...) // 追加到开头
	fmt.Println(books)
	books = append(books[:3], append([]string{"javascript", "docker"}, books[3:]...)...) // 追加到中间指定索引位置上多个元素
	fmt.Println(books)
	books = append(books, "")  // 让切片长度 +1，腾出一个空位，以便后面把元素挪动
	copy(books[3:], books[2:]) // 原来第 2 个位置之后的所有元素整体往后挪一格。 这样在索引 2 的位置就空出来了。
	books[2] = "react"         // 把新元素 "react" 填到刚才空出来的位置 2
	fmt.Println(books)
	fmt.Printf("长度为: %d, 容量为: %d\n", len(books), cap(books)) // 这里容量为 6

}
