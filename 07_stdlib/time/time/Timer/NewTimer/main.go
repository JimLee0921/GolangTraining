package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(10 * time.Second)

	fmt.Println("waiting...")

	t := <-timer.C
	fmt.Println("timer fired at:", t)
}
