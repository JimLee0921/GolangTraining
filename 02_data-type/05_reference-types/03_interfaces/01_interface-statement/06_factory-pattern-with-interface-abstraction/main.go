package main

import "fmt"

/*
接口封装 + 工厂构造（Factory Pattern with Interface Abstraction）
也就是抽象工厂模式

API 对外返回接口（interface），但内部实际持有未导出的具体类型实现（struct）
用工厂函数创建对象，并隐藏实现细节，用工厂函数创建对象，并隐藏实现细节

主要是参考 go reflect 包中的 reflect.Type
*/

// Animal 接口
type Animal interface {
	Speak() string
	Eat() string
}

// dog 实现 Animal 接口，但是小写不导出
type dog struct {
}

func (d dog) Speak() string {
	return "woof!"
}

func (d dog) Eat() string {
	return "yummy!"
}

// NewAnimal 构造工厂，返回的是接口，而不是实现
func NewAnimal() Animal {
	return dog{} // 返回内部实现
}

func main() {
	// 在外部使用时
	a := NewAnimal() // 返回的是 Animal 接口
	// 这里可以直接使用内部实现的方法，但是并不知道实现的是 dog 后续修改其他实现而不需要修改代码
	fmt.Println(a.Eat())   // woof!
	fmt.Println(a.Speak()) // yummy!
}
