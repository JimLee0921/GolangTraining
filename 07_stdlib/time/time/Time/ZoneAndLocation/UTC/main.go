package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	utc := t.UTC()
	fmt.Println(t)
	fmt.Println(utc)
}
