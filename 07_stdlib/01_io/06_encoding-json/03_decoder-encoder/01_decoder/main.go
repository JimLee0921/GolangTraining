package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func BasicDemo() {
	r := strings.NewReader(`{"name":"Alice","age":20}`)
	dec := json.NewDecoder(r)

	var u User
	if err := dec.Decode(&u); err != nil {
		panic(err)
	}

	fmt.Println(u) // {Alice 20}
}

func MultiJsonDemo() {
	r := strings.NewReader(`{"name":"A","age":1}{"name":"B","age":2}{"name":"C","age":3}`)
	dec := json.NewDecoder(r)
	for {
		var u User
		if err := dec.Decode(&u); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println(u)
	}
}

func NDJsonDemo() {
	f, _ := os.Open("temp_files/users.ndjson")
	dec := json.NewDecoder(f)
	for {
		var u User
		if err := dec.Decode(&u); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println(u)
	}
}

func NoUseNumberDemo() {
	r := strings.NewReader(`{"name":"Alice","age":20}`)
	dec := json.NewDecoder(r)
	var m map[string]any
	dec.Decode(&m)

	fmt.Printf("age type: %T, value: %v\n", m["age"], m["age"])
}

func UseNumberDemo() {
	r := strings.NewReader(`{"name":"Alice","age":20}`)
	dec := json.NewDecoder(r)
	dec.UseNumber()

	var m map[string]any
	dec.Decode(&m)

	fmt.Printf("age type: %T, value: %v\n", m["age"], m["age"])

	age := m["age"].(json.Number)
	i, _ := age.Int64()
	fmt.Printf("after convert -> type: %T, value: %v\n", i, i)
}

func main() {
	//BasicDemo()
	//MultiJsonDemo()
	//NDJsonDemo()
	// 这两个进行对比查看 Json 数字转换的类型
	NoUseNumberDemo()

	UseNumberDemo()
}
