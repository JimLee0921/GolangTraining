package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// 需要取地址
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
