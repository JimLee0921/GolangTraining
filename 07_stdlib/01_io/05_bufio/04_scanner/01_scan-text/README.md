# Scan + Text

`bufio.Scanner`从一个输入流（文件、标准输入、网络连接等）中，逐段 读取内容。

默认情况下：

- 每次调用 Scan() 会扫描到下一行（以换行符 \n 作为分隔）
- 然后使用 Text() 获取刚刚扫描到的内容

## 方法签名

```
func (s *Scanner) Scan() bool

func (s *Scanner) Text() string
```

**Scan**

- 向前推进扫描器到下一段（通常是一行）
- 成功扫描到内容时返回 true
- 如果到达文件结尾（EOF）或遇到错误时返回 false

> 注意：返回 false 并不一定是错误，也可能是正常的 EOF


**Text**

- 返回最近一次 `Scan()` 调用时扫描到的文本内容（字符串形式）
- 仅在 `Scan()` 返回 true 时有效