package main

import "io"

type MyType struct{}

func (m *MyType) Write(p []byte) (int, error) {
	return len(p), nil
}

type person struct {
}

func main() {
	/*
		var _ 接口类型 = 具体类型值
		编译期断言：
			本质：一次赋值语句
			左边是一个接口类型变量，右边是一个具体类型的值/指针
			但是如果右边的类型没有实现接口的所有方法，编译器会报错
		常见用途
			显式声明实现关系
				Go 的接口实现是隐式的，但有时需要明确表示这个类型实现了某个接口
			防止未来改动破坏接口实现
				比如改动了类型的方法签名，可能导致它不再实现接口
				这种写法能在编译时第一时间报错，而不是运行时报错
			标准库惯用写法
				标准库经常使用这种模式，作为静态检查
	*/
	// 编译期检查：*MyType 必须实现 io.Writer
	var _ io.Writer = (*MyType)(nil)
}
