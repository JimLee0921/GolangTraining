package main

import (
	"fmt"
	"reflect"
)

type User struct {
}

// Hello User struct 导出方法
func (u User) Hello() {

}

// Greet User struct 导出方法
func (u User) Greet() {

}

// eat User struct 未导出方法
func (u User) eat() {

}

// Behavior 接口定义方法中包含导出和未导出的清空
type Behavior interface {
	Hello()
	eat()
	Greet()
}

func main() {
	t1 := reflect.TypeOf(User{}) // 结构体
	var b Behavior
	t2 := reflect.TypeOf(&b).Elem() // 获取结构类型
	fmt.Println(t1.NumMethod())     // 2 说明结构体只包含到处方法
	fmt.Println(t2.NumMethod())     // 3 接口包含所有方法

	// Method 根据下标获取到 reflect.Method 对象
	m1 := t1.Method(0)
	fmt.Println(m1)

	// MethodByName 根据方法名进行获取，并会返回是否获取到了
	m2, ok := t2.MethodByName("Hello")
	if ok {
		fmt.Println(m2)
	} else {
		fmt.Println("not found Hello method from Behavior")
	}

	m3, ok := t2.MethodByName("wtf")
	if ok {
		fmt.Println(m3)
	} else {
		fmt.Println("not found wtf method from Behavior")
	}

}
