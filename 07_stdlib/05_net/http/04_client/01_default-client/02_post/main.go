package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	form := url.Values{}
	form.Set("name", "ChatGPT")
	form.Set("Lang", "Golang")
	form.Add("Lang", "Java")

	resp, err := http.Post(
		"https://httpbin.org/post",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
