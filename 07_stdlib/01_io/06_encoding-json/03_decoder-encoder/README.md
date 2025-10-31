# JSON 流式解析

`json.Decoder` / `json.Encoder` 主要用于 JSON 的流式解析，
基于 `io.Reader` / `io.Writer`，特别适合大文件、网络流、多段 JSON 连续输入等场景。

- 数据很大：不想一次性 []byte 全读入内存
- 持续到达：HTTP 长连接/文件流，边到边解
- 多段 JSON：一个连接里连续发多个 JSON 文档或数组元素



