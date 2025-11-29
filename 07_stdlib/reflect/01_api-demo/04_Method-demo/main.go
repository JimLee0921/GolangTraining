package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
}

// Hello 给 User 绑定方法
func (u User) Hello(name string) string {
	return "hello," + name + ", i'm JimLee."
}

func (u User) Eat(food string) string {
	return "i'm eating" + food + "."
}

func main() {
	// reflect.TypeOf 取方法信息
	t := reflect.TypeOf(User{name: "JimLee"})
	fmt.Println("method num:", t.NumMethod())
	fmt.Println(t.Method(1))             // 根据索引取 Method 结构体
	fmt.Println(t.MethodByName("Hello")) // 按名称取 Method 结构体
	fmt.Println("--------------------")
	// 方法结构体可以访问属性
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println("Method Name:", m.Name)
		fmt.Println("Method Type:", m.Type)
		fmt.Println("PkgPath:", m.PkgPath)
		fmt.Println("Index:", m.Index)
		fmt.Println("--------------------")
	}

	// reflect.ValueOf 取方法 Value 可以用于访问
	v := reflect.ValueOf(User{name: "JimLee"})
	m := v.MethodByName("Hello") // 和 t.MethodByName 不一样。这个可以调用 Call
	result := m.Call([]reflect.Value{
		reflect.ValueOf("JimLee"),
	})

	// 取值
	fmt.Println(result[0].String()) // hello,JimLee, i'm JimLee.

	m2 := v.MethodByName("fuck")
	fmt.Println(m2.IsValid()) // false
	fmt.Println(m2)           // 没有就是 invalid

	m3 := v.MethodByName("Eat")
	fmt.Println(m3.IsValid()) // true
	result2 := m3.Call([]reflect.Value{
		reflect.ValueOf("Banana"),
	})
	fmt.Println(result2) // [i'm eatingBanana.]
}
