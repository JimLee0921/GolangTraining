package main

import "fmt"

func main() {
	/*
		Go 中使用 取地址符 & 来得到变量的内存地址
		使用解引用符 * 访问和修改变量实际的值
	*/
	x := 10
	p := &x         // p 是 *int 类型，存储 x 的地址
	fmt.Println(p)  // 打印的 x 的内存地址 0xc00000a088
	*p = 20         // 使用 解引用修改 x 的值
	fmt.Println(*p) // 解引用访问 x 的值 20

	fmt.Printf("%T\n", x)    // 普通值类型 int
	fmt.Printf("%T\n", p)    // 指向 x 的指针 *int
	fmt.Printf("%T\n", *(p)) // 解引用回原值 int
}
