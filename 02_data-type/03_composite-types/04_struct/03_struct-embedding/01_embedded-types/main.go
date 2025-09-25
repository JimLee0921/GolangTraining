package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}
type killer struct {
	person        // 匿名字段 / 嵌入类型
	licenseToKill bool
}

// main go语言中的嵌入
func main() {
	/*
		在 Go 里，嵌入 (embedding) 就是把一个类型（通常是 struct）直接写在另一个 struct 里，而不给它起字段名
		类似于继承但不是继承
		传统 OOP 里的继承（Java/C++）
			子类自动拥有父类的字段和方法
			支持多态、重写（override）、继承链

		Go 的 struct 嵌入
			并不是继承，而是把一个类型嵌入到另一个类型里
			嵌入类型的 字段和方法会提升 (promotion)，可以像直接属于外层 struct 一样访问
			Go 没有继承链，强调的是 组合优于继承
	*/
	killer1 := killer{
		person: person{
			firstName: "James",
			lastName:  "Bond",
			age:       20,
		},
		licenseToKill: true,
	}

	killer2 := killer{
		person: person{
			firstName: "Bruce",
			lastName:  "Lee",
			age:       18,
		}, licenseToKill: false,
	}
	// 字段提升 等同于 killers1.person.firstName
	fmt.Println(killer1.firstName, killer1.lastName, killer2.firstName, killer2.lastName)
	fmt.Println(killer1.licenseToKill, killer2.licenseToKill)

}
