package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// main json 反序列化
func main() {
	/*
		使用 json.Unmarshal() 方法
		用于 把 JSON 字节数据解析（解码）成 Go 的数据结构
		func Unmarshal(data []byte, v any) error
			data []byte → JSON 格式的字节切片（常用 []byte 或 string 转过来）
			v any → 必须传指针，这样函数才能修改它指向的值
			返回 error，如果 JSON 格式不合法或者类型不匹配，会返回错误
	*/
	data := []byte(`{"name":"Bob","age":18}`)
	var p Person
	_ = json.Unmarshal(data, &p)
	fmt.Printf("%+v\n", p)
	fmt.Println(p)

}
