package main

import (
	"fmt"
	"time"
)

func main() {
	m := map[string]string{}
	timer := time.AfterFunc(10, func() {
		m["JimLee"] = "JimLee"
	})
	time.Sleep(20 * time.Second)
	ok := timer.Stop()
	fmt.Println("timer stop:", ok)

	fmt.Println(m)
}
