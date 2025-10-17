package main

import "fmt"

type person struct {
	name string
	age  int
}

/*
如果不用指针，而是 func setAge(p Person, age int)，那么传的是副本，修改不会影响外部变量。
*/
func setAge(p *person, age int) {
	p.age = age
}

func main() {
	/*
		结构体指针使用好处
			避免拷贝：传值时如果结构体很大（比如几百个字段），指针能避免复制整个对象
			共享修改：多个函数操作同一个对象时，传指针能直接修改原值
			惯例：很多库/标准库函数都用 *Struct 作为参数，方便修改
		当然更好的是直接编写结构体方法，结构体方法可以自动进行解引用
	*/
	p := person{name: "Bob", age: 18}
	fmt.Println(p.age)
	setAge(&p, 30)
	fmt.Println(p.age) // 30，说明函数内部修改生效

	p2 := person{name: "Jim", age: 33}
	setAge(&p2, 22)
	fmt.Println(p2.age)
}
