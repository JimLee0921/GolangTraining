## 字符串分割与拼接

处理字符串数据时最常用的能力之一

- 配置解析
- 文本预处理
- HTTP / CSV / 日志分析
- Web 表单和 URL 参数处理

### Spilt

```
func Split(s, sep string) []string
```

按照指定的分隔符切割，返回 string 切片

- 不保存分隔符
- 分隔符是严格匹配的

```
strings.Split("a,b,c", ",") // ["a", "b", "c"]
```

### SpiltAfter

```
func SplitAfter(s, sep string) [] string
```

和 Split 一样，但保留分隔符，分隔符出现在每段的末尾

```
strings.SplitAfter("a,b,c", ",") // ["a," "b," "c"]
```

### SplitN

SplitN 将字符串 s 切成以分隔符 sep 分隔的子字符串，并返回分隔符之间子字符串的切片。

```
func SplitN(s, sep string , n int) [] string
```

控制分割逻辑，避免把后面内容拆碎，常用于只切第一刀

计数 n 决定了要返回的子字符串的数量：

- n > 0：最多 n 个子串；最后一个子串将是未分割的剩余部分
- n == 0：结果为 nil（零个子字符串）
- n < 0：所有子字符串

```
strings.SplitN("key=value=another", "=", 2) // ["key", "value=another"]
```

### Fields

按空白符切割（自动清除多空格、tab、换行）

```
func Fields(s string ) [] string
```

**能识别的空白符**：

- 空格 " "
- 制表符 \t
- 换行 \n

**适用于**：

- 去除多余空格的文本分词
- 处理用户输入不规范的情况
- 处理日志内容

> Split 是指定分隔符，Fields 是自动按空白切

### Join

join 函数会将第一个参数的元素连接起来，生成一个字符串。分隔符 sep 会放置在结果字符串的各个元素之间。

```
func Join(elems [] string , sep string ) string
```

**适用于**：

- 构造 path、URL
- 输出日志信息
- 将数组转换为字符串返回

## 补充

| 方法                       | 作用                    | 参数                      | 返回值类型      | 特点                        |
|--------------------------|-----------------------|-------------------------|------------|---------------------------|
| `Split(s, sep)`          | 按 **指定子串** 分割         | string, string          | `[]string` | sep 精确匹配                  |
| `SplitN(s, sep, n)`      | 按子串分割，并限制返回的 **切片数量** | string, string, int     | `[]string` | `n` 控制切割次数                |
| `SplitAfter(s, sep)`     | 按子串分割，**保留分隔符**       | string, string          | `[]string` | 分隔符放在片段末尾                 |
| `SplitAfterN(s, sep, n)` | 分割并保留分隔符 + **限制片段数**  | string, string, int     | `[]string` | `SplitN + SplitAfter` 的组合 |
| `Fields(s)`              | 按 **空白字符** 分割         | string                  | `[]string` | 会自动识别空格、制表符、换行……          |
| `FieldsFunc(s, f)`       | 按 **自定义逻辑规则** 分割      | string, func(rune) bool | `[]string` | 最灵活，可按符号/中文等分隔            |
