package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string
	Age    int
	secret string
}

func MarshalDemo() {
	u := User{
		Name:   "JimLee",
		Age:    20,
		secret: "this is my secret",
	}
	data, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	// 只会序列化首字母为大写的可导出字段 {"Name":"JimLee","Age":20}
	fmt.Println(string(data))
}

func MarshalIndentDemo() {
	u := User{
		Name:   "Bruce",
		Age:    22,
		secret: "my secret",
	}
	data, err := json.MarshalIndent(u, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	/*
		{
			"Name": "Bruce",
			"Age": 22
		}
		美化输出
	*/
}

func main() {
	//MarshalDemo()
	MarshalIndentDemo()

	data := []string{"Go", "Rust", "Python"}
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData)) // 切片序列化 ["Go","Rust","Python"]
}
