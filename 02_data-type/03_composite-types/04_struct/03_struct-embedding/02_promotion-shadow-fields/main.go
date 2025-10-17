package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

/*
killer 嵌入了 person 并且内部又定义了一个同名的字段 firstName
结果就是 killer 中存在两个不同的 firstName

	外层的 First（直接属于 killer）
	内层的 person.First
*/
type killer struct {
	person        // 匿名字段 / 嵌入类型
	firstName     string
	licenseToKill bool
}

// main 嵌入时由于字段提升造成字段遮蔽
func main() {
	/*
		嵌入 struct 的字段会提升，可以像直接属于外层一样访问
		但如果外层定义了一个同名字段，它会 遮蔽 (shadow) 内层的字段
		内层字段仍然存在，只是要通过完整路径（p1.person.First）才能访问
	*/

	killer1 := killer{
		person: person{
			firstName: "James",
			lastName:  "Bond",
			age:       20,
		},
		firstName:     "Double Zero Seven",
		licenseToKill: true,
	}
	killer2 := killer{
		person: person{
			firstName: "Bruce",
			lastName:  "Lee",
			age:       30,
		},
		firstName:     "China Dragon",
		licenseToKill: false,
	}
	// 想访问被遮蔽的内层字段，需要显式写：killer1.person.firstName
	fmt.Println(killer1.firstName, killer1.person.firstName)
	fmt.Println(killer2.firstName, killer2.person.firstName)
}
