package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		sort.Reverse
		接收一个已经实现了 sort.Interface 的对象（比如 sort.StringSlice、sort.IntSlice 等），返回一个新的 sort.Interface，这个新接口的排序逻辑是反转的
	*/
	s := []string{"JimLee", "Bruce", "Django", "James", "Tom"}

	// 升序
	sort.Strings(s)
	fmt.Println("升序:", s) // [Bruce Django James JimLee Tom]

	// 降序
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println("降序:", s) // [Tom JimLee James Django Bruce]

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	// 升序
	sort.Ints(n)
	fmt.Println("升序:", n) // [2 3 4 7 8 9 12 19 32]

	// 降序
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println("降序:", n) // [32 19 12 9 8 7 4 3 2]

	// 步骤讲解
	//sort.Sort(sort.StringSlice(s))
	//fmt.Println(s)
	//
	//fmt.Printf("just s: %T\n", s)
	//s = sort.StringSlice(s)
	//fmt.Printf("just s: %T\n", s)
	//t := sort.StringSlice(s)
	//fmt.Printf("just t: %T\n", t)
	//
	//fmt.Printf("s converted to StringSlice: %T\n", sort.StringSlice(s))
	////	fmt.Printf("s sorted: %T\n", sort.Sort(sort.StringSlice(s)))
	//fmt.Printf("s reversed: %T\n", sort.Reverse(sort.StringSlice(s)))
	//i := sort.Reverse(sort.StringSlice(s))
	//fmt.Println(i)
	//fmt.Printf("%T\n", i)
	//sort.Sort(i)
	//fmt.Println(s)
}
