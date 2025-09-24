package main

import "fmt"

// main 切片删除元素
func main() {
	/*
		删除元素（需要删除指定元素需要使用 for range 进行遍历判断）
			Go 中需要用 切片拼接 或 copy 来实现删除切片元素
			1. 删除下标为 i 到 j 的元素（最常用，j = i +1 则就是只删除下标为 i 的元素）
				s = append(s[:i], s[j:]...) -> i < j
				s[:i] -> 要保留的前半部分
				s[j:] -> 要保留的后半部分
				拼接在一起就相当于跳过了下标为 i 到 j 的元素，注意不包含下标为 j 的元素
			2. 删除开头或结尾元素
				删除开头元素：s = s[i:] 直接把切片向右偏移 i 格，也就是舍弃前 i 个元素
				删除结尾元素：s = s[:len(s)-i] 直接把切片截断去除后 i 个元素
			3. 使用 copy 删除多个元素（更高效）
				copy(s[i:], s[j:])	-> i < j
				把 s[j:] 挪到 s[i:]，覆盖掉要删除的部分
				s = s[:len(s)-(j-i)]
				缩短切片长度，删除的就是从 i 到 j 的元素
	*/
	numberSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 删除下标从 3 到 5 的元素
	numberSlice = append(numberSlice[:3], numberSlice[5:]...)
	fmt.Println(numberSlice)

	// 删除开头一个元素和结尾两个元素
	numberSlice = numberSlice[1:]
	numberSlice = numberSlice[:len(numberSlice)-2]
	fmt.Println(numberSlice)

	// 使用 copy 删除元素
	copy(numberSlice[1:], numberSlice[2:])
	numberSlice = numberSlice[:len(numberSlice)-1]
	fmt.Println(numberSlice)
}
