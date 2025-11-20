package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "https://www.baidu.com/", nil)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	for _, c := range resp.Cookies() {
		fmt.Printf("%s = %s\n", c.Name, c.Value)
	}
}
