package main

import "fmt"

type person struct {
	name string
	age  int
}

// SetAgeOne 结构体方法值接收者
func (p person) SetAgeOne(age int) {
	p.age = age
}

// SetAgeTwo 结构体方法指针接收者
func (p *person) SetAgeTwo(age int) {
	p.age = age
}

// setAgeThree 普通函数指针接收者
func setAgeThree(p *person, age int) {
	p.age = age
}

// setAgeFour 普通函数值接收者
func setAgeFour(p person, age int) {
	p.age = age
}

func main() {
	/*
		函数和结构体方法主要区别在于调用和内部语法糖等
	*/
	p1 := person{"Jim", 3}
	p2 := person{"Bruce", 3}
	p3 := person{"James", 3}
	p4 := person{"Tom", 3}

	// 结构体方法值接收者，使用实例调用，但是不能修改
	p1.SetAgeOne(33)
	fmt.Println(p1.age) // 没有变化

	// 结构体方法指针接收者，使用实例调用，自动取址
	p2.SetAgeTwo(33)
	fmt.Println(p2.age) // 33 修改成功

	// 普通函数指针接收者，普通调用，需要手动传入指针，可以修改成功
	setAgeThree(&p3, 33)
	fmt.Println(p3.age) // 33 修改成功

	// 普通函数值接收者，普通调用，不能修改
	setAgeFour(p4, 33)
	fmt.Println(p4.age) // 没有变化

}
