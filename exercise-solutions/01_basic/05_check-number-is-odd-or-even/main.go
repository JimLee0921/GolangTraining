package main

import "fmt"

func main() {
	/*
		判断 1 - 1000 所有数是奇数还是偶数
	*/
	for i := 0; i <= 1000; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
