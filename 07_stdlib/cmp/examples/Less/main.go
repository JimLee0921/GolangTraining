package main

import (
	"cmp"
	"fmt"
	"math"
)

func main() {
	fmt.Println(cmp.Less(1, 2))            // true
	fmt.Println(cmp.Less("a", "aa"))       // true
	fmt.Println(cmp.Less(1.0, math.NaN())) // false
	fmt.Println(cmp.Less(math.NaN(), 1.0)) // true
	fmt.Println(cmp.Less(10, 2))           // false
}
