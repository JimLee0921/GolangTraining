package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]int{
		"JimLee":    20,
		"JamesBond": 30,
		"FrankStan": 22,
	}

	//newM := map[string]int{}
	newM := make(map[string]int)
	
	// 有点无意义，但是是这样的
	newM = maps.Collect(maps.All(m))
	fmt.Println(newM)
}
