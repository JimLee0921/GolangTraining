package main

import "fmt"

func send(ch chan int) {
	ch <- 42
}

func main() {
	ch := make(chan int, 1)
	send(ch)
	fmt.Println(<-ch) // 42
}
