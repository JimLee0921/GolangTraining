package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

func main() {
	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,

		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 50,
		IdleConnTimeout:     90 * time.Second,
	}

	client := &http.Client{
		Timeout:   15 * time.Second,
		Transport: tr,
	}
	form := url.Values{}
	form.Set("Name", "JimLee")
	form.Add("Name", "BruceLee")
	form.Add("Lang", "go")
	// PostForm 默认 "Content-Type": "application/x-www-form-urlencoded"
	resp, err := client.PostForm("https://httpbin.org/post", form)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
