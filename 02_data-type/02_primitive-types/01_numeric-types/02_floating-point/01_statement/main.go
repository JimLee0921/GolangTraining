package main

import "fmt"

func main() {
	var a float32 = 1.23
	var b = 4.56 // 推断为 float64
	c := -7.89   // 同样是 float64

	fmt.Printf("a=%v (%T)\n", a, a)
	fmt.Printf("b=%v (%T)\n", b, b)
	fmt.Printf("c=%v (%T)\n", c, c)
}

/*
a=1.23 (float32)
b=4.56 (float64)
c=-7.89 (float64)
*/
