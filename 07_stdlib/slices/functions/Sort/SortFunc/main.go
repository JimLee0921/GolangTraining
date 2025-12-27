package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func main() {
	names := []string{"Bob", "alice", "VERA"}
	slices.SortFunc(names, func(a, b string) int {
		// 不区分大小写
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	fmt.Println(names)

	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Gopher", 13},
		{"Alice", 55},
		{"Bob", 24},
		{"Alice", 20},
	}
	slices.SortFunc(people, func(a, b Person) int {
		// 多字段比较
		if n := strings.Compare(a.Name, b.Name); n != 0 {
			return n
		}
		return cmp.Compare(a.Age, b.Age)
	})
	fmt.Println(people)
}
