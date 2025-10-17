package main

import "fmt"

func main() {
	/*
		循环体里重新遮蔽一份变量
	*/
	done := make(chan bool)

	values := []string{"a", "b", "c"}

	for _, v := range values {
		v := v // 重新声明一份局部变量
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	for _ = range values {
		<-done
	}

}
