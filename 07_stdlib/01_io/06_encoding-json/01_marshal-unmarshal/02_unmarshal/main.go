package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string // 正常写入
	Age    int    // 正常写入
	secret string // 首字母小写无法写入
}

func main() {
	// 结构体序列化 结构体字段必须首字母大写才能参与反序列化
	jsonStr := `{"Name":"Bob","Age":25, "secret": "my secret"}`
	var u User
	err := json.Unmarshal([]byte(jsonStr), &u)

	if err != nil {
		panic(err)
	}
	fmt.Println(u) // {Bob 25 }

	// JSON 数组序列化
	var users []User
	usersJsonStr := `[{"name": "A", "age": 10}, {"name": "B", "age": 20}]`
	json.Unmarshal([]byte(usersJsonStr), &users)
	fmt.Println(users) // [{A 10 } {B 20 }]

}
