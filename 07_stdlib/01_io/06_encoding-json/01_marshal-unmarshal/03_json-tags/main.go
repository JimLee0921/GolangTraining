package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string `json:"user_name"` // 序列化与反序列化时使用 user_name 对照
	Sex    string `json:"sex"`
	Money  int    `json:"money,omitempty"` // 如果是零值就不显示
	Secret string `json:"-"`               // 永远不显示
}

func main() {
	// 序列化
	u := User{
		Name:   "JimLee",
		Sex:    "男",
		Money:  0,
		Secret: "my secret",
	}

	data, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data)) // {"user_name":"JimLee","sex":"男"}

	// 反序列化

	jsonStr := `{"user_name":"JamesBond","sex":"男", "money": 2}`
	var u2 User
	json.Unmarshal([]byte(jsonStr), &u2)
	fmt.Println(u2) // {JamesBond 男 2 }

}
