package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Location())
	fmt.Println(t.UTC().Location())
}
