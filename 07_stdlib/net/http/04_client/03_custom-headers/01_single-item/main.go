package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	req, _ := http.NewRequest("GET", "https://httpbin.org/headers", nil)
	// 设置常见的自定义 Header 单条进行添加
	req.Header.Set("User-Agent", "Go-Client/999.0")
	req.Header.Set("X-Api-key", "1234567890")
	req.Header.Set("X-Request-ID", "abc-xyz-001")
	req.Header.Set("Accept-Encoding", "gzip")
	// Add 进行追加
	req.Header.Add("Accept-Encoding", "br")
	// Del 删除
	//req.Header.Del("X-API-Key")
	// Get / Values 获取值
	fmt.Println(req.Header.Values("Accept-Encoding")) // 如果有多个获取全部
	fmt.Println(req.Header.Get("Accept-Encoding"))    // get 只能获取第一个

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
