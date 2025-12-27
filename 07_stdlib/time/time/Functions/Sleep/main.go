package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hahaha")
	fmt.Println("sleep 2s")
	time.Sleep(2 * time.Second)
	fmt.Println("done")
}
