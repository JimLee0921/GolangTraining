package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{
		"JimLee": 20,
		"James":  123,
	}

	m2 := map[string]int{
		"JimLee": 30,
		"Bond":   22,
	}

	m1Seq := maps.All(m1)
	// Insert 直接原地修改 m2
	maps.Insert(m2, m1Seq)
	fmt.Println(m2)

}
