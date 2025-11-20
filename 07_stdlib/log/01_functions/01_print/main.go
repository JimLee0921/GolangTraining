package main

import (
	"log"
)

// 普通输出
func main() {
	log.Print("hello", " ", "world")
	log.Println("server started")
	log.Printf("listen on %s", ":8080")
}

/*
2025/11/20 09:50:25 hello world
2025/11/20 09:50:25 server started
2025/11/20 09:50:25 listen on :8080
*/
