package main

import "fmt"

func call(f func()) {
	f()
}

func main() {
	fn := func() { fmt.Println("run") }
	call(fn) // 输出 run
}
