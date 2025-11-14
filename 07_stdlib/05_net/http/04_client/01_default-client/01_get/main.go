package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpBinResp struct {
	Args   map[string]string `json:"args"`
	Origin string            `json:"origin"`
	URL    string            `json:"url"`
}

func main() {
	// 使用 http.DefaultClient 创建 get 请求
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	// 读取完毕关闭 body 防止连接泄露
	defer resp.Body.Close()

	//// 读取响应体
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("status: ", resp.Status)
	//fmt.Println("status: ", resp.StatusCode)
	//// 1. 简单字符串解析打印
	//fmt.Println("body: ", string(body))
	//// 2. 不知道结构，解析 JSON 为 map[string]interface{}
	//var res map[string]any
	//json.Unmarshal(body, &res)
	//fmt.Println("args: ", res["arys"])
	//fmt.Println("headers: ", res["headers"])
	//fmt.Println("origin", res["origin"])
	//fmt.Println("url: ", res["url"])
	//fmt.Println("something: ", res["something"]) // 解析不到就是 nil
	// 3. 知道 body 接口，并且使用 json.NewDecoder(resp.Body) 进行流式解析
	dec := json.NewDecoder(resp.Body)
	var result HttpBinResp
	if err := dec.Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println(result.URL)
	fmt.Println(result.Origin)
	fmt.Println(result.Args)
}
