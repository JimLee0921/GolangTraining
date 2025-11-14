package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "https://httpbin.org/headers", nil)

	// 批量设置 headers
	req.Header = http.Header{
		"User-Agent":    {"Go-HttpClient/2.0"},
		"Authorization": {"Bearer my-token"},
		"Accept":        {"application/json"},
		"X-TAG":         {"v1", "v2"},
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
