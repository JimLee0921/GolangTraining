package main

import "fmt"

type User struct {
	Name    string
	Age     int
	Hobbies []string
	Tags    map[string]string
}

// 函数1：传值
func modifyByValue(u User) {
	u.Name = "Alice"                        // 修改基础类型
	u.Age++                                 // 修改基础类型
	u.Hobbies = append(u.Hobbies, "coding") // 修改 slice
	u.Tags["role"] = "pro"                  // 修改 map
}

// 函数2：传指针
func modifyByPointer(u *User) {
	u.Name = "Bob"
	u.Age++
	u.Hobbies = append(u.Hobbies, "music")
	u.Tags["level"] = "senior"
}

func main() {
	u := User{
		Name:    "Tom",
		Age:     20,
		Hobbies: []string{"reading"},
		Tags:    map[string]string{"role": "dev"},
	}

	modifyByValue(u)
	fmt.Println("After modifyByValue:", u)
	// => 基础类型没变，map 改了（共享），slice 可能没变（取决于是否扩容）

	modifyByPointer(&u)
	fmt.Println("After modifyByPointer:", u)
	// => 所有字段都被修改
}
