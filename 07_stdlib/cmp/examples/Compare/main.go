package main

import (
	"cmp"
	"fmt"
	"math"
)

func main() {
	fmt.Println(cmp.Compare(1, 2))            // -1
	fmt.Println(cmp.Compare(10, 2))           // 1
	fmt.Println(cmp.Compare("a", "aa"))       // -1
	fmt.Println(cmp.Compare(1.5, 1.5))        // 0
	fmt.Println(cmp.Compare(math.NaN(), 1.0)) // -1
}
