package main

import "fmt"

// main string 指针传递
func main() {
	name := "JimLee"
	fmt.Println(&name) // 0xc000022070
	changeStr(name)
	fmt.Println(name) // JimLee
	changeStrByPointer(&name)
	fmt.Println(name) // JimLee????
}

func changeStr(str string) {
	fmt.Println(str) // JimLee
	str += "????"
	fmt.Println(str) // JimLee????
}

func changeStrByPointer(str *string) {
	fmt.Println(str) // 0xc000022070
	*str += "????"
	fmt.Println(str) // 0xc000022070
}
