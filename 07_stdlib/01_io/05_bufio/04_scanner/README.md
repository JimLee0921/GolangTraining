# Scanner 扫描

Scanner 是一种高层封装，用于从输入流中逐个提取「标记（token）」。

常见应用：

- 一次读取一行（默认行为）
- 按单词、空格分割
- 按自定义分隔符（例如 CSV、JSON token、双换行等）

连续调用 `Scanner.Scan` 方法将逐行遍历文件的标记，并跳过标记之间的字节。
标记的规范由`SplitFunc`类型的拆分函数定义；默认拆分函数会将输入拆分为行，并去除行尾。
扫描会在 `EOF`、第一个` I/O` 错误或令牌过大而无法放入 `Scanner.Buffer` 时不可恢复地停止。
扫描停止时，读取器可能已经任意地超出了最后一个令牌的范围。
对于需要对错误处理或大令牌进行更多控制，或必须在读取器上运行顺序扫描的程序，应改用 `bufio.Reader`。

## 类型定义（简化）

```
type Scanner struct {
    r         io.Reader
    split     SplitFunc
    buf       []byte
    token     []byte
    err       error
}
```

- r: 底层输入源
- split: 用于切分输入的函数（决定如何分词）
- buf: 内部缓冲
- token: 当前扫描到的一个结果片段

## 定义创建

使用 `func NewScanner(r io.Reader) *Scanner` 进行创建

```
scanner := bufio.NewScanner(strings.NewReader("hello\nworld\n"))
```

## SplitFunc

`splitFunc` 决定了 Scanner 如何从底层 `io.Reader` 分割数据成一个个 token。
也就是说：每次调用 `Scan()` 时，Scanner 内部都会调用一次 `SplitFunc` 来决定下一个 token 的边界

例如：

- ScanLines -> 每一行是一个 token
- ScanWords -> 每个单词是一个 token
- ScanBytes -> 每个字节是一个 token
- ScanRunes -> 每个 UTF-8 字符是一个 token
- 也可以自定义，比如“每个逗号分隔的数据”是一个 token

### 内置分割函数

| 函数                | 功能            | 示例              |
|-------------------|---------------|-----------------|
| `bufio.ScanLines` | 按行分割（默认）      | 读取每行文本（去掉 `\n`） |
| `bufio.ScanWords` | 按空格、换行符分割     | 一次返回一个单词        |
| `bufio.ScanBytes` | 每次一个字节        | 类似 `ReadByte()` |
| `bufio.ScanRunes` | 每次一个 UTF-8 字符 | 支持中文、多字节字符      |

### 函数签名说明

```
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
```

| 参数            | 类型       | 含义                   |
|---------------|----------|----------------------|
| `data`        | `[]byte` | 当前可用的输入数据缓冲区         |
| `atEOF`       | `bool`   | 是否到达输入末尾             |
| 返回值 `advance` | `int`    | 表示消费了多少字节（从输入缓冲区中移除） |
| 返回值 `token`   | `[]byte` | 提取出的数据片段（即扫描到的内容）    |
| 返回值 `err`     | `error`  | 错误（非 EOF）            |

> 只要返回 token 非空，Scanner 就会让 Text() 返回该字符串

### 底层原理

- Scanner 内部维护了一个缓冲区 `buf []byte`
- 每次调用 `Scan()`，buf 中的内容会更新为下一个分片
- `Bytes()` 返回的就是指向这块 buf 的切片

> 调用下一次 Scan() 后，之前的 Bytes() 返回值将失效（内容会被覆盖）

### 使用场景

- 处理二进制数据（无需转成字符串）
- 高性能场景，避免 Text() 的内存复制
- 临时解析、即时消费（比如直接写入 socket、文件等）