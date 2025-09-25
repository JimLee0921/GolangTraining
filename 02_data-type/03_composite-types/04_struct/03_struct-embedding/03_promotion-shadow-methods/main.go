package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type killer struct {
	person        // 匿名字段 / 嵌入类型
	firstName     string
	licenseToKill bool
}

/*
person 类型有一个方法 greeting()
killer 类型也有一个同名方法 greeting()
在 Go 里，外层类型的方法会 覆盖（遮蔽） 内层嵌入类型的方法
*/
func (p person) greeting() string {
	return "haha, i'm a person"
}

func (k killer) greeting() string {
	return "haha, i'm a killer"
}

// main 嵌入时由于字段提升造成方法遮蔽
func main() {
	/*
		方法提升：嵌入类型的方法也会被提升到外层类型
		仍可访问：内层方法没有消失，可以通过 k.person.Greeting() 调用
	*/
	p := person{firstName: "Jim"}
	k := killer{
		person: person{
			firstName: "James",
		},
		firstName:     "007",
		licenseToKill: false,
	}
	fmt.Println(p.greeting())
	fmt.Println(k.greeting())
	fmt.Println(k.person.greeting())
}
