## 请求体 body

在 HTTP 协议里，请求体（Request Body） 就是请求中实际携带的数据。内容类型由 Content-Type 决定。

Body 在 Go 中是一个流（Reader）。它不是字符串、不是字节数组，而是 可读取的数据流。

```
type Request struct {
    ...
    Body io.ReadCloser
    ...
}
```

- Body 可以来自内存、文件、网络、管道
- Body 必须实现 `io.Reader` 接口（通常也加上 `io.Closer`）

### 常见构造方式

| 类型       | Go 构造方式                        | 说明          |
|----------|--------------------------------|-------------|
| **字符串**  | `strings.NewReader("data")`    | 发送纯文本或表单    |
| **字节数组** | `bytes.NewReader([]byte{...})` | 发送二进制数据     |
| **JSON** | `bytes.NewReader(jsonBytes)`   | 发送结构化 JSON  |
| **文件**   | `os.Open("file.txt")`          | 上传文件或发送大数据流 |
| **动态生成** | `io.Pipe()`                    | 实时流式发送      |

### 常见 content-type 总结

> MDN content-type（也叫 MIME Type）网址：https://www.iana.org/assignments/media-types/media-types.xhtml
> Go 中实际处理 MIME 类型用 mime 包：https://pkg.go.dev/mime

| Content-Type                        | Body 格式      | 使用场景             | Go 中写法                     | 示例                                  |
|-------------------------------------|--------------|------------------|----------------------------|-------------------------------------|
| `application/json`                  | JSON 文本      | API 调用（REST、RPC） | `json.Marshal()`           | `{"username": "JimLee", "age": 22}` |
| `application/x-www-form-urlencoded` | URL 编码表单     | 登录、提交表单          | `url.Values{}.Encode()`    | `username=JimLee&password=123456`   |
| `multipart/form-data`               | 多段体（文本 + 文件） | 文件上传             | `multipart.Writer`         | Body 由多个 part 构成，每段有自己头部            |
| `text/plain`                        | 纯文本          | 简单调试、日志          | `strings.NewReader("...")` |                                     |
| `application/octet-stream`          | 原始二进制流       | 上传文件流、下载         | `bytes.NewReader(file)`    | 二进制流                                |

