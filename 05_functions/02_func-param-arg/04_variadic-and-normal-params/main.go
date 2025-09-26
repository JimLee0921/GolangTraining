package main

import "fmt"

func greeting(me string, age int, names ...string) {
	fmt.Println("hello!", names, "i am", me, "and i am", age, "years old")
}
func main() {
	/*
		普通参数和可变参数同时出现时可变参数必须放在普通参数之后
	*/
	greeting("JimLee", 29)
	greeting("JimLee", 20, "Skrillex", "Bond", "007")
}
