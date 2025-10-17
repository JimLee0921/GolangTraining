package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	/*
		定义结构体类型
		type + struct 定义复合类型
			字段可以是不同类型
			首字母大写为导出共有,小写为包内私有
		初始化(普通结构体变量):
			1. 零值初始化
			var instance structType
				struct 的零值 = 每个字段取该类型的零值（数值 0、布尔 false、string ""、指针 nil 等）。
				零值可用，不会是未初始化状态。
			2. 命名字段(推荐)
			instance := structType{field1: value1, field2: value2}
				字段可缺省，缺省部分取零值
				推荐写法，安全、可读性好
			3. 按字段顺序
			instance := structType{value1, value2, ...}
				必须提供所有字段
				不推荐，字段顺序变动会出现问题
		结构体指针见 04_struct-pointer

	*/
	// 1. 零值初始化
	var p1 person
	fmt.Printf("%+v\n", p1) // {firstName: lastName: age:0}

	// 2. 命名字段初始化
	p2 := person{firstName: "Alice", lastName: "Bob"}
	fmt.Printf("%+v\n", p2) // {firstName:Alice lastName:Bob age:0}

	// 3. 按字段初始化
	p3 := person{"Jim", "Lee", 30}
	fmt.Printf("%+v\n", p3) // {firstName:Jim lastName:Lee age:30}

}
