package main

import (
	"fmt"
	"strconv"
)

func FormatBoolDemo() {
	s1 := strconv.FormatBool(true)  // "true"
	s2 := strconv.FormatBool(false) // "false"
	fmt.Printf("%T: %v\n", s1, s1)
	fmt.Printf("%T: %v\n", s2, s2)
}

func FormatInt() {
	s1 := strconv.FormatInt(255, 10) // "255"
	s2 := strconv.FormatInt(255, 2)  // "11111111"
	s3 := strconv.FormatInt(255, 16) // "ff"

	fmt.Printf("%T: %v\n", s1, s1)
	fmt.Printf("%T: %v\n", s2, s2)
	fmt.Printf("%T: %v\n", s3, s3)
}

func FormatUnit() {
	s := strconv.FormatUint(255, 16) // "ff"
	fmt.Printf("%T: %v\n", s, s)
}

func FormatFloatDemo() {
	s1 := strconv.FormatFloat(3.1415926, 'f', 2, 64) // "3.14"
	s2 := strconv.FormatFloat(3.1415926, 'e', 4, 64) // "3.1416e+00"
	s3 := strconv.FormatFloat(3.1415926, 'g', 6, 64) // "3.14159"

	fmt.Printf("%T: %v\n", s1, s1)
	fmt.Printf("%T: %v\n", s2, s2)
	fmt.Printf("%T: %v\n", s3, s3)
}

func main() {
	FormatBoolDemo()
	FormatInt()
	FormatUnit()
	FormatFloatDemo()
}
