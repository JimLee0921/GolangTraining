package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []float64{1.0, 2.0, 3.3, 4.6, 6.1, 7.2, 8.0}
	x := 2.0
	i := sort.SearchFloat64s(a, x)
	if i < len(a) && a[i] == x {
		fmt.Printf("found %g at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("not found %g in %v\n", x, a)
	}

	x = 0.2
	i = sort.SearchFloat64s(a, x)
	if i < len(a) && a[i] == x {
		fmt.Printf("found %g at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("not found %g in %v, can insert in %v\n", x, a, i)
	}

}
