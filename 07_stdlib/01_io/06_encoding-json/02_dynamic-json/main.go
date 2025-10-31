package main

import (
	"encoding/json"
	"fmt"
)

func BasicDemo() {
	jsonStr := `{"name": "Tom", "age": 18, "scores": [90, 88, 95]}`
	var m map[string]any
	json.Unmarshal([]byte(jsonStr), &m)
	fmt.Println(m) // map[age:18 name:Tom scores:[90 88 95]]
	// 数字会被解析成 float64，需要强转
	fmt.Printf("%T\n", m["age"]) // float64
	age := int(m["age"].(float64))
	fmt.Println(age)
}

func StructNestedDemo() {
	// 对象嵌套
	jsonStr := `{"user":{"name":"Alice","age":20}}`
	var m map[string]any
	json.Unmarshal([]byte(jsonStr), &m)

	user := m["user"].(map[string]interface{})
	fmt.Println(user["name"]) // Alice

}

func main() {
	BasicDemo()
	StructNestedDemo()

}
