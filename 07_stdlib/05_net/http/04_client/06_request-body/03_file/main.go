package main

import (
	"bytes"
	"mime/multipart"
	"net/http"
)

func main() {
	// 服务端使用 r.FormFile("file") 或 r.MultipartForm 可读取。
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// 添加普通字段
	writer.WriteField("desc", "my file")

	// 添加文件
	fileWriter, _ := writer.CreateFormFile("file", "test.txt")
	fileWriter.Write([]byte("Hello file!"))
	writer.Close() // 结束 boundary

	req, _ := http.NewRequest("POST", "https://httpbin.org/post", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

}
