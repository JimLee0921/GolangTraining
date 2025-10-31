# `encoding/json`

encoding/json 是 Go 官方标准库，用来做 JSON 与 Go 数据结构之间的相互转换。

| 方向         | 方法                                    | 描述                                 |
|------------|---------------------------------------|------------------------------------|
| Go -> JSON | `json.Marshal` / `json.MarshalIndent` | 把 Go 值编码成 JSON 字符串（byte slice）     |
| JSON -> Go | `json.Unmarshal`                      | 解析 JSON 字符串到 Go 变量、结构体、map、slice 等 |

> 官方文档：https://pkg.go.dev/encoding