package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	u, _ := url.Parse("https://httpbin.org/get")

	// 构造 query 参数
	q := u.Query() // 返回 url.Values（已存在的参数也会被取出）
	q.Set("name", "JimLee")
	q.Set("lang", "go")
	q.Add("hobby", "music")
	q.Add("hobby", "reading")
	//q.Add("hobby", "大傻逼")
	// 编码回 URL
	u.RawQuery = q.Encode()
	req, _ := http.NewRequest("GET", u.String(), nil)
	fmt.Println("final url: ", u.String())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	fmt.Println(string(data))

}
