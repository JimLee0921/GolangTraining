package main

import "fmt"

func TypeSwitch(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Println("int:", val)
	case string:
		fmt.Println("string:", val)
	case bool:
		fmt.Println("bool:", val)
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	TypeSwitch(42)
	TypeSwitch("hi")
	TypeSwitch(true)
}
