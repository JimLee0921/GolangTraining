## Split

`Scanner` 默认的分割规则是：`bufio.ScanLines`
也就是：每次遇到一个换行符（\n），就算作一条记录。
`Split()` 方法用来设置扫描的分割规则。

### 方法定义

```
func (s *Scanner) Split(split SplitFunc)
```

SplitFunc 也就是分隔函数，Go 中内置了四个常用的 SplitFunc：

| 函数名               | 作用             | 分割依据         | 返回内容                         |
|-------------------|----------------|--------------|------------------------------|
| `bufio.ScanLines` | 按行分割（默认）       | 换行符 `\n`     | 每一行文本（不含换行符）                 |
| `bufio.ScanWords` | 按空白字符分割        | 空格、制表符、换行符等  | 每个单词                         |
| `bufio.ScanBytes` | 按字节分割          | 每个字节         | 单个 `byte`（如 `"A"` → `["A"]`） |
| `bufio.ScanRunes` | 按 Unicode 字符分割 | UTF-8 解码字符边界 | 每个字符（包括中文、emoji 等）           |

### 自定义 SpiltFunc

可以自定义 `SplitFunc` 来使用自定义的分隔符

**注意事项**

1. 如果不设置 `scanner.Split()`，就一直是默认的按行扫描
2. 自定义 `SplitFunc` 必须保证进展：否则容易卡死（如果返回 advance=0 且没返回 token，Scanner 会一直读不到新内容）
3. 若扫描到的 `token` 超过 64KB，需要用 `scanner.Buffer()` 扩容，但更建议使用 `bufio.Reader`
