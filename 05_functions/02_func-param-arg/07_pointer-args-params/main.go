package main

import "fmt"

// main 函数参数传入指针
func main() {
	/*
		如果传值（参数是普通类型），函数里拿到的是一个拷贝，修改不会影响原来的变量
		如果传地址（参数是指针），函数里操作的就是原始数据，能直接修改外部变量
		为什么使用指针
			Go 里函数参数默认是值传递，会拷贝一份如果传的是很大的结构体，拷贝开销大用指针可以避免复制，直接操作原对象，更高效
			有些场景需要多个地方共享同一份数据，而不是复制
			通过传指针，操作的是同一块内存
			有时需要区分没有值和有值，比如返回 nil 指针，传指针能检查空值，实现一些语义上的区分
		使用：
			param（形参）：写成 *T，例如 *int、*string
			arg（实参）：调用时传入 &变量
	*/
	num := 20
	updateValue(num)
	fmt.Println(num) // 20
	updateValueUsePoint(&num)
	fmt.Println(num) // 30

}

func updateValueUsePoint(value *int) {
	*value = *value + 10

}

func updateValue(value int) {
	value = value + 10

}
