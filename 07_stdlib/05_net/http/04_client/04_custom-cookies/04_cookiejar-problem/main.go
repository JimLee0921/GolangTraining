package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
)

func main() {
	// cookiejar 的缓存问题，如果使用了CookieJar，那么复用一个请求，会出现请求，cookie累积的问题
	jar, _ := cookiejar.New(nil)
	client := http.Client{
		Jar: jar,
	}
	request, _ := http.NewRequest("GET", "https://www.baidu.com/", nil)

	for x := 0; x < 10; x++ {
		res, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("req %+v\n", request.Cookies())
		fmt.Printf("resp %+v\n", res.Cookies())
	}

}
