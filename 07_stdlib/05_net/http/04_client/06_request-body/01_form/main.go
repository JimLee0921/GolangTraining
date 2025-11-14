package main

import (
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// 服务器端用 r.FormValue("username") 就能直接解析。
	form := url.Values{}
	form.Set("username", "JimLee")
	form.Set("password", "123456")
	form.Set("hobby", "music")
	form.Add("hobby", "coding")

	req, _ := http.NewRequest("POST", "https://httpbin.org/post", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}
