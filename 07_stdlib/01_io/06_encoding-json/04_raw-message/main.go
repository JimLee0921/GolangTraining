package main

import (
	"encoding/json"
	"fmt"
)

// 1. 定义外层结构，data 使用 json.RawMessage
type Message struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// 2. 定义具体类型，用于匹配 Data 类型
type TextData struct {
	Content string `json:"content"`
}
type ImageData struct {
	URL   string `json:"url"`
	Width int    `json:"width"`
}

func handleMessage(b []byte) {
	var msg Message
	// 3. 先反序列化外层结构
	if err := json.Unmarshal(b, &msg); err != nil {
		panic(err)
	}
	// 4. 配合 switch 根据 type 分发解析 data
	switch msg.Type {
	case "text":
		var td TextData
		if err := json.Unmarshal(msg.Data, &td); err != nil {
			panic(err)
		}
		fmt.Printf("TEXT MESSAGE → content=%q\n", td.Content)
	case "image":
		var id ImageData
		if err := json.Unmarshal(msg.Data, &id); err != nil {
			panic(err)
		}
		fmt.Printf("IMAGE MESSAGE → url=%s width=%d\n", id.URL, id.Width)
	default:
		fmt.Println("Unknown message type", msg.Type)
	}
}

func main() {
	handleMessage([]byte(`{"type":"text","data":{"content":"hello"}}`))
	handleMessage([]byte(`{"type":"image","data":{"url":"pic.png","width":300}}`))
	handleMessage([]byte(`{"type":"video","data":{"url":"video.mp4","size":50000}}`))
}
