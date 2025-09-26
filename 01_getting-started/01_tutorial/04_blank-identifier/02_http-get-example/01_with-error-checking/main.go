package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// main http 请求接受错误
func main() {
	res, err := http.Get("https://www.baidu.com/")
	if err != nil {
		log.Fatal(err)
	}
	page, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	fmt.Printf("%s", page)
}
