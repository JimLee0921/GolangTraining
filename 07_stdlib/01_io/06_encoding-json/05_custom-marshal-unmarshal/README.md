# 自定义 Marshal/Unmarshal

自定义 MarshalJSON / UnmarshalJSON。
用它可以优雅地解决枚举值映射、时间格式/时间戳、以及对接外部协议的特殊字段编码问题。

## 核心接口

```
type Marshaler interface {
    MarshalJSON() ([]byte, error)
}

type Unmarshaler interface {
    UnmarshalJSON([]byte) error
}
```

> 只要类型本身或其指针实现了以上接口，`encoding/json` 就会调用自定义的方法来自定义该类型的 JSON 表现

## 常见用途

- 时间格式
- 金额/精度
- 枚举校验
- 敏感信息脱敏/加密
- 对外协议兼容（字段别名/多格式）