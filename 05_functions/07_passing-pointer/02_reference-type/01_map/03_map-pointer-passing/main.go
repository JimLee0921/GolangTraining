package main

import "fmt"

func main() {
	/*
		可以通过传入 *map 进行 map 的重新分配
	*/
	m := map[string]int{"old": 1}
	resetMap(&m)
	fmt.Println(m) // map[new:1]
}

func resetMap(m *map[string]int) {
	*m = make(map[string]int)
	(*m)["new"] = 1
}
