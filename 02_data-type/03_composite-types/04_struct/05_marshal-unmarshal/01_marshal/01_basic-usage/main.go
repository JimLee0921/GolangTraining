package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string
	Age     int
	Sex     string
	isAdmin bool
}

// main json 序列化
func main() {
	/*
		在 Go 里，把结构体转成 JSON（序列化）最常用的就是标准库 encoding/json 包
		json.Marshal()：把结构体转成 []byte 或字符串形式的 JSON
		json.MarshalIndent()：带缩进的 JSON 形式
		但是注意在 go 中只有导出的字段（首字母大写）才能被 encoding/json 包访问到并序列化
	*/
	p := Person{
		Name:    "Bob",
		Age:     18,
		isAdmin: true,
	}

	jsonBytes, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonBytes))

	indentJsonBytes, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(indentJsonBytes))
}
