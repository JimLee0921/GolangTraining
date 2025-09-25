package main

import "fmt"

// 定义结构体类型 person
type person struct {
	firstName string
	LastName  string
	age       int
}

/*
给 person 类型定义了一个方法 fullName
这里的 (p person) 称为 接收者 (receiver，其实就是一个方法的一个参数)，表示这个方法属于 person 类型
在 Go 中，不习惯像其他语言一样用 this 或 self 作为名字。
fullName 方法会返回 first 和 last 拼接的结果
因为接收者是 值接收者（而不是 *person），它不会修改原来的 person 对象
*/
func (p person) fullName() string {
	return p.firstName + " " + p.LastName
}

func main() {
	/*
		结构体 + 方法绑定
		是把状态filed和行为methods进行组合，接近面向对象编程的效果
	*/
	p1 := person{"Jim", "Lee", 30}
	p2 := person{"James", "Bond", 50}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}

/*
“很多人不知道，其实方法调用的语法糖 v.Method() 在 Go 里本质上等价于解糖后的写法 (T).Method(v)。
可以在这里看到例子。把接收者像普通参数一样命名，正好体现了它本质上就是一个参数。
这也意味着，在方法内部，接收者参数是可能为 nil 的。而在 Java 等语言里，this 是绝不可能为 nil 的。”
来源: https://www.reddit.com/r/golang/comments/3qoo36/question_why_is_self_or_this_not_considered_a/?utm_source=golangweekly&utm_medium=email
*/
