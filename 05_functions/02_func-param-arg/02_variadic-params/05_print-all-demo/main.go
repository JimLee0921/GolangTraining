package main

import "fmt"

// PrintAll 可变参数 + interface{}：可以接受任意数量、任意类型的参数
func PrintAll(values ...any) {
	for i, v := range values {
		fmt.Printf("[%d] (%T): %v\n", i, v, v)
	}
}

func main() {
	PrintAll("hello", 123, 3.14, true, []string{"a", "b"}, map[string]int{"x": 1})
}
