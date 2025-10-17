package main

import "fmt"

type MySlice []int

func main() {
	/*
		切片的类型包括：元素类型 + 方向（默认是双向）
		因此，不同元素类型的切片不可转换
	*/
	var s1 []int
	var s2 []int64
	// s2 = s1 // 编译错误，不同类型切片
	fmt.Println(s1, s2)

	// 可以通过手动转换每个元素
	s3 := []int{1, 2, 3}
	s4 := make([]int64, len(s3))
	for i, v := range s3 {
		s4[i] = int64(v)
	}
	fmt.Println(s3, s4)

	var s []int = []int{1, 2, 3}

	var ms MySlice = MySlice(s) // 可以底层类型相同
	fmt.Println(s, ms)
}
