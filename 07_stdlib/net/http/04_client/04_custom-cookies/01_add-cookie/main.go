package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "https://httpbin.org/cookies", nil)

	// 手动使用 req.AddCookie 添加
	req.AddCookie(&http.Cookie{
		Name:  "session_id",
		Value: "xyz-123",
	})
	req.AddCookie(&http.Cookie{
		Name:  "user",
		Value: "JimLee",
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	fmt.Println(string(data))

}
