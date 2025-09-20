package main

import "fmt"

// main 数组的多种声明与初始化方式
func main() {
	// 方式一：只声明长度和类型，元素保持零值
	var zero [3]int
	fmt.Printf("zero: %v，长度: %d\n", zero, len(zero))

	// 方式二：声明时提供完整字面量
	var explicit = [3]int{1, 2, 3}
	fmt.Printf("explicit: %v\n", explicit)

	// 方式三：使用 := 简写，同时初始化
	shorthand := [4]string{"Go", "Rust", "Python", "Java"}
	fmt.Printf("shorthand: %v\n", shorthand)

	// 方式四：使用省略号让编译器推断长度
	inferred := [...]int{10, 20, 30, 40}
	fmt.Printf("inferred: %v，长度: %d\n", inferred, len(inferred))

	// 方式五：指定索引初始化，未列出的索引用零值填充
	sparse := [5]int{0: 99, 3: 42}
	fmt.Printf("sparse: %v\n", sparse)

	// 方式六：借助常量或表达式，长度必须在编译期已知
	const size = 2
	var sized = [size * 2]float64{3.14, 2.71, 1.41}
	fmt.Printf("sized: %v\n", sized)

	// 方式七：多维数组，本质是数组的数组
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("matrix: %v\n", matrix)

	// 方式八：使用 new 创建数组指针，底层数组依旧是零值
	ptr := new([3]bool)
	ptr[1] = true
	fmt.Printf("ptr: %v，类型: %T\n", ptr, ptr)
}
