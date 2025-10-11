package main

import "fmt"

func resetMap(m map[string]int) {
	m = make(map[string]int) // 新分配，原指针被覆盖
	m["new"] = 1
}

func main() {
	m := map[string]int{"old": 1}
	resetMap(m)
	fmt.Println(m) // map[old:1]
}
