package main

import "fmt"

func main() {
	/*
		第一种解决方案是把循环变量作为参数显式传入
	*/
	done := make(chan bool)

	values := []string{"a", "b", "c"}

	for _, v := range values {
		go func(u string) {
			fmt.Println(u)
			done <- true
		}(v)
	}

	for _ = range values {
		<-done
	}

}
