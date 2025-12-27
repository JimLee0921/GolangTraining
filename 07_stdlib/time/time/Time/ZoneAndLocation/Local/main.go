package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Unix(0, 0)
	fmt.Println(t)
	fmt.Println(t.Local())
}
