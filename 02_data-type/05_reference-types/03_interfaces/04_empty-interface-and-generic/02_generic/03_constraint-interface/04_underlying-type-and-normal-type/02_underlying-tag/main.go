package main

import "fmt"

type Ordered interface {
	~int | ~float64 | ~string
}

type MyInt int

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Max(MyInt(1), MyInt(2))) // 这里是可以的ok
}
