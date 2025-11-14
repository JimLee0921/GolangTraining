package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUserReq struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreateUserResp struct {
	Json map[string]interface{} `json:"json"` // httpbin 返回结构示例
}

// 使用 json.NewDecoder 流式读取，写服务端/高并发客户端，优先用这个
func main() {
	// 1. 构造请求 JSON 数据
	reqBody := CreateUserReq{
		Name: "TuoLee",
		Age:  22,
	}

	// 2. 编码为 JSON
	jsonBytes, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	// 3. 创建 Request（推荐使用 WithContext）
	req, err := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewReader(jsonBytes))
	if err != nil {
		panic(err)
	}

	// 4. 设置 Content-Type
	req.Header.Set("Content-Type", "application/json")

	// 5. 发送请求（client 可以是自定义的）
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 6. 使用 json.NewDecoder 流式读取
	var result CreateUserResp

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println(result)
}
