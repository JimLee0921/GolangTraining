package main

import (
	"encoding/json"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func BasicDemo() {
	u := User{"Alice", 20}
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(u) // 写到控制台 {"name":"Alice","age":20}
}

func multiJsonDemo() {
	enc := json.NewEncoder(os.Stdout)

	enc.Encode(User{
		Name: "JimLee",
		Age:  20,
	})

	enc.Encode(User{
		Name: "BruceLee",
		Age:  33,
	})

	enc.Encode(User{
		Name: "FrankStan",
		Age:  12,
	})
	/*
		{"name":"JimLee","age":20}
		{"name":"BruceLee","age":33}
		{"name":"FrankStan","age":12}
	*/
}

func SetIndentDemo() {
	// 使用 SetIndex 美化输出
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t")
	enc.Encode(User{
		Name: "JimLee",
		Age:  22,
	})
	/*
		{
			"name": "JimLee",
			"age": 22
		}
	*/
}
func SetEscapeHTMLDemo() {
	// HTML 自动转义
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(User{Name: "<Tom & Jerry>"}) // {"name":"\u003cTom \u0026 Jerry\u003e","age":0}
	enc.SetEscapeHTML(false)
	enc.Encode(User{Name: "<Tom & Jerry>"}) // {"name":"<Tom & Jerry>","age":0}

}

func main() {
	//BasicDemo()
	//multiJsonDemo()
	//SetIndentDemo()
	SetEscapeHTMLDemo()
}
