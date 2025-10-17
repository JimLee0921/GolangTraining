package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type person struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Age         int
	notExported int
}

// main 流式解码
func main() {
	/*
		json.Decoder：Encoder 的“反操作”，用来从 流 (io.Reader) 中读取 JSON 并解码到 Go 的数据结构
		dec := json.NewDecoder(r) // r 是 io.Reader，比如 strings.NewReader, os.Stdin, 文件, http.Request.Body
		err := dec.Decode(&v)     // 把 JSON 数据解到 v 中（必须传指针）
			json.NewDecoder(r)：创建一个绑定到 io.Reader 的解码器
			Decode(&v)：从 r 中读出 JSON，把数据填充进 v
		适用场景
			一次解码：类似 json.Unmarshal，但从 io.Reader 里直接读数据
			流式解码：可以在循环里连续 Decode 多个 JSON 值，非常适合日志流、大文件、HTTP 长连接等
		注意事项：
			必须传 指针，否则解码填充不了数据
			流结束时返回 io.EOF，这是正常现象，不是错误。
			默认 JSON key 必须和字段名大小写一致；可用 struct tag 调整
			多余的 key 会被忽略，缺少的字段会保持零值
	*/
	var p1 person
	reader := strings.NewReader(`{"firstName":"James", "lastName":"Bond", "Age":20}`)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&p1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", p1)
}
