package main

import (
	"fmt"
	"sort"
)

// 自定义结构体
type person struct {
	name string
	age  int
}

// 自定义 String 方法，类似于 python 中的 __str__
func (p person) String() string {
	return fmt.Sprintf("name: %s, age: %d", p.name, p.age)
}

// ByAge 实现 []person 的 sort.Interface，基于 Age 字段。
type ByAge []person

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByAge) Less(i, j int) bool {
	return b[i].age < b[j].age
}

func main() {
	/*
		String() string方法：字符串表示接口，相当于 Java 的 toString() 或 Python 的 __str__
			让自定义类型在打印时可以更友好的自定义输出，Go 的 fmt 包会优先使用这个方法
			如果一个类型实现了 String() string 方法，那么在用 fmt.Print* 打印该类型时，会自动调用这个方法
		定义新类型（如 ByAge），为它实现 sort.Interface，在 Less 方法里控制排序逻辑（可以按 Age，也可以按 Name）
		sort.Sort 不关心具体类型，只要满足 sort.Interface 就能排序
	*/
	persons := []person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(persons[0])   // 调用 String()，输出 name: Bob, age: 31
	fmt.Println(persons)      // 打印整个 slice，会调用每个元素的 String()方法：[name: Bob, age: 31 name: John, age: 42 name: Michael, age: 17 name: Jenny, age: 26]
	sort.Sort(ByAge(persons)) // 按 Age 排序，在排序时，临时把 []persons 转成 ByAge（这个类型实现了接口）
	fmt.Println(persons)      // 排序后的结果：[name: Michael, age: 17 name: Jenny, age: 26 name: Bob, age: 31 name: John, age: 42]

}
