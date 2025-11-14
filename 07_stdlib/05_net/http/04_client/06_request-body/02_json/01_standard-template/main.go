package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CreateUserReq struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreateUserResp struct {
	Json map[string]interface{} `json:"json"` // httpbin 返回结构示例
}

// 使用 io.ReadAll 一次性读取全部 body
func main() {
	// 1. 构造请求 JSON 数据
	reqBody := CreateUserReq{
		Name: "JimLee",
		Age:  20,
	}

	// 2. 进行 JSON 编码
	jsonBytes, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	// 3. 创建 Request
	req, err := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewReader(jsonBytes))
	if err != nil {
		panic(err)
	}
	// 4. 设置 content-type
	req.Header.Set("Content-Type", "application/json")

	// 5. 发送请求（这里速度不用自定义请求了）
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// 6. 判断HTTP状态码
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status code: %d", res.StatusCode))
	}

	// 7. 解析 JSON 数据
	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var result CreateUserResp
	if err = json.Unmarshal(data, &result); err != nil {
		panic(err)
	}
	fmt.Println("server get json: ", result.Json)
}
