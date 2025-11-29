package main

import (
	"fmt"
	"reflect"
)

type User struct {
	// 存在 json tag ，有值
	Name string `json:"name"`
	// 存在 json tag，但 value为空字符串
	Nickname string `json:""`
	// 完全没有 json tag
	Age int
}

func main() {
	t := reflect.TypeOf(User{})

	// --- Name ---
	f1 := t.Field(0)
	fmt.Println("[Name]")
	fmt.Println("Get      ->", f1.Tag.Get("json")) // name
	v1, ok1 := f1.Tag.Lookup("json")
	fmt.Println("Lookup   ->", v1, ok1) // name true

	// --- Nickname ---
	f2 := t.Field(1)
	fmt.Println("\n[Nickname]")
	fmt.Println("Get      ->", f2.Tag.Get("json")) // ""
	v2, ok2 := f2.Tag.Lookup("json")
	fmt.Println("Lookup   ->", v2, ok2) // "" true tag存在但值为空

	// --- Age ---
	f3 := t.Field(2)
	fmt.Println("\n[Age]")
	fmt.Println("Get      ->", f3.Tag.Get("json")) // ""
	v3, ok3 := f3.Tag.Lookup("json")
	fmt.Println("Lookup   ->", v3, ok3) // "" false tag不存在
}
