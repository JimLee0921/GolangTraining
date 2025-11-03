````# `time.Location`

`time.Location` 表示一个 时区规则，包括：

- 该时区与 UTC 的偏移量（例如 +08:00）
- 是否需要考虑夏令时（DST）
- 时区的名称（如 `Asia/Shanghai`、`America/New_York`）

Go 的时间 `time.Time` 结构中始终 携带一个 Location，否则就无法正确表示本地时间的意义

一个时间并不是 `2025-01-01 10:00` 这么简单，它必须是：

```text
2025-01-01 10:00  在 Asia/Shanghai

2025-01-01 02:00  在 UTC
```

## 常见 Location

| 名称                                   | 说明                   |
|--------------------------------------|----------------------|
| `time.UTC`                           | 世界协调时（零时区）           |
| `time.Local`                         | 当前系统所在本地时区（依赖系统环境变量） |
| `time.LoadLocation("Asia/Shanghai")` | 明确指定某个时区（推荐）         |

但是在 Go 中建议使用 LoadLocation 而不是 Local。
因为 Local 依赖服务器环境：

- Linux 用 /etc/localtime
- Windows 与系统区域设置有关
- Docker 容器里经常是 UTC 导致乱套

所以 跨系统/跨服务器 时，Local 不可靠

## 总结

Location 决定“这个时间在哪个地区被解释和显示”。

- `time.In()`：用于转换显示地区
- `ParseInLocation()`：用于正确解析时间字符串
- `LoadLocation()`：是获取时区的主要方式
