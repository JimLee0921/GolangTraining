package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello JimLee\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"JimLee"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}

/*
实现错误处理机制
*/
