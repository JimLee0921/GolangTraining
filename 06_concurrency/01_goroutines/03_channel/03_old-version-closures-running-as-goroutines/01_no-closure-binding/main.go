package main

import "fmt"

// main 1.22 之前版本打印的可能是 ccc
func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}

	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// 等待所有 routines 完成
	for _ = range values {
		<-done
	}
}
