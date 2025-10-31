# `RawMessage`

RawMessage 是原始编码的 JSON 值，表示原始 JSON 片段，在第一次反序列化时保留原样，不解析。
实现了 Marshaler 和 Unmarshaler 接口，可用于延迟 JSON 解码或预先计算 JSON 编码

```
type RawMessage []byte
``` 

> 可以先把 data 读成 RawMessage，然后根据 type 决定如何再解析它

## 使用场景

当 JSON 的 某个字段的结构不固定，需要根据 类型字段 决定后续解析方式。

```
{
  "type": "text",
  "data": { "content": "hello" }
}

{
  "type": "image",
  "data": { "url": "xxx.jpg", "width": 100 }
}

{
  "type": "video",
  "data": { "url": "xxx.mp4", "duration": 10 }
}
```

如果不用 RawMessage：

- data 必须预先设成一个确定类型（无法实现）
- 只能走 `map[string]any` -> 丢失类型信息、需要大量断言

使用 RawMessage：

- 第一次只解析固定字段（如 type）
- 动态字段原封不动保存
- 根据 type 再精准反序列化一次（类型安全 ）

> data 字段到底是什么结构，需要看 type 决定，这就是多态 JSON

## 核心逻辑

1. 定义外层结构，其中可变字段用 RawMessage
2. 先 Unmarshal 外层，提取类型标识字
3. 根据类型进行 switch 判断，用第二次 Unmarshal 精准解析 RawMessage
