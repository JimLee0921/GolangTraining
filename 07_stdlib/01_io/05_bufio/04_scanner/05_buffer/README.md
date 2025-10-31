## Buffer

让 Scanner 能处理超长行或大块数据

### 方法签名

```
func (s *Scanner) Buffer(buf []byte, max int)
```
参数：
- `buf []byte`：初始缓冲区（可以是 nil）
- `max int`：最大 token 大小，（必须 > len(buf)）


### 功能说明

`Buffer()` 用于自定义 Scanner 的内部缓冲区大小。

默认情况下，Scanner 每次扫描时最多能处理 64KB（65536 字节） 的数据。
如果扫描到的token（例如一行文本或一个词）超过这个大小，就会报错。


1. 设置初始缓冲区（buf）：提供一个自己的缓冲区（可以是 nil）
2. 设置最大 token 大小（max）：决定单个 token（如一行或一词）允许的最大字节数