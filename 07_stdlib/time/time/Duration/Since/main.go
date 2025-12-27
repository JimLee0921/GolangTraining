package main

import (
	"fmt"
	"time"
)

func Working() {
	time.Sleep(time.Second)
}

func main() {
	start := time.Now()
	Working()
	elapsed := time.Since(start)
	fmt.Printf("working consume %v\n", elapsed)
}
