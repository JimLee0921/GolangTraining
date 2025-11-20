## 字符串判断和查找

strings 包中最常用、最基础的能力。

### Contains

判断一个字符串中是否包含另一个字符串

```
func Contains(s, substr string) bool
```

- 区分大小写
- substr 可以是任意长度的字符串（包含空字符串）
- 当 substr 为空字符串时，返回 true

**适用场景**

- 判断用户输入中是否含有某些关键字
- 判断 URL 是否包含某些 path
- 文本过滤 / 搜索

### HasPrefix

```
func HasPrefix(s, prefix string) bool
```

判断字符串是否以某个内容开头。

**适用场景**

- 检查 URL、请求路径、文件格式
- 检查配置字段是否带特定 header
- 判断日志行是否属于某种类型

### HasSuffix

```
func HasSuffix(s, suffix string) bool
```

判断字符串是否 以某个内容结尾。

**适用场景**

- 判断文件类型（图片、文档、视频等）
- 检查域名后缀
- 判断日志行格式

### Index

```
func Index(s, substr string) int
```

返回子串第一次出现的位置（从 0 开始计数）

**适用场景**

- 文本解析
- 切分字符串前先找位置

### LastIndex

```
func LastIndex(s, substr string) int
```

返回子串最后出现的位置

## 补充

### contains 系列

| 方法             | 用途                         | 参数                        | 特点                  |
|----------------|----------------------------|---------------------------|---------------------|
| `Contains`     | 判断是否包含某个**子串**             | `string, string`          | 最常用                 |
| `ContainsAny`  | 判断是否包含**字符集合**中任意一个字符      | `string, string`          | 用于字符集合匹配            |
| `ContainsRune` | 判断是否包含某个 **单个 Unicode 字符** | `string, rune`            | 精确匹配单个字符（含中文/emoji） |
| `ContainsFunc` | 判断是否存在满足 **函数逻辑条件** 的字符    | `string, func(rune) bool` | 可处理复杂规则匹配           |

### index 系列

| 方法                   | 作用                          | 参数                      | 返回值 | 示例                           |
|----------------------|-----------------------------|-------------------------|-----|------------------------------|
| `Index(s, substr)`   | 查找 **子串第一次出现** 的位置          | string, string          | int | `hello` 中找 `l` → 2           |
| `IndexAny(s, chars)` | 查找 **chars 中任一字符** 首次出现的位置  | string, string          | int | `hello` 中找 `aeiou` → 1 (`e`) |
| `IndexByte(s, c)`    | 查找 **单个 byte** 首次出现的位置      | string, byte            | int | 只适用于 ASCII                   |
| `IndexRune(s, r)`    | 查找 **Unicode rune** 首次出现的位置 | string, rune            | int | 可用于中文、emoji                  |
| `IndexFunc(s, f)`    | 查找 **满足函数条件的 rune** 首次出现的位置 | string, func(rune) bool | int | 可按逻辑自定义条件                    |

### lastIndex 系列

| 方法                       | 作用                          | 参数                      | 返回值 | 示例                          |
|--------------------------|-----------------------------|-------------------------|-----|-----------------------------|
| `LastIndex(s, substr)`   | 查找 **子串最后一次出现** 的位置         | string, string          | int | `hello` 找 `l` → 3           |
| `LastIndexAny(s, chars)` | 查找 **chars 中任一字符** 最后出现的位置  | string, string          | int | `hello` 找 `aeiou` → 4 (`o`) |
| `LastIndexByte(s, c)`    | 查找 **单个 byte** 最后出现的位置      | string, byte            | int | ASCII 使用                    |
| `LastIndexFunc(s, f)`    | 查找 **满足函数条件的 rune** 最后出现的位置 | string, func(rune) bool | int | 常用于找最后一个数字/字母等              |
