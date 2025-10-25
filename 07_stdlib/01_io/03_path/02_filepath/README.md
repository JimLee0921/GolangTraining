# path/filepath 模块

filepath 提供用于操作文件系统路径的函数。主要用于处理真实的文件系统路径，而且会自动适配不同操作系统的分隔符。

| 特点                | 说明                          |
|-------------------|-----------------------------|
| 自动使用系统分隔符         | Linux/macOS：`/`，Windows：`\` |
| 真实文件路径操作          | 可与 `os` 文件操作函数配合使用          |
| 支持遍历、匹配、规范化、绝对路径等 |                             |

> 如果是 URL 或 Web 路径用 path 如果是本地文件路径用 filepath

## 常用函数

> Windows linux 等平台不同，结果可能也不同

### 路径拼接 `join`

语法：`filepath.Join(elem ...string) string`

自动插入分隔符、去掉多余的 `/` 或 `\`

### 路径规范化 `Clean()`

语法：`filepath.Clean(path string) string`

### Dir() 与 Base()

`Dir`：`filepath.Dir(path string) string`

- Dir 返回路径中除最后一个元素之外的所有元素，通常是路径的目录
- 删除最后一个元素后，Dir 会对路径调用 Clean 函数，并删除路径尾部的斜杠。如果路径为空，Dir 返回`.`
- 如果路径完全由分隔符组成，Dir 返回单个分隔符。除非路径是根目录，否则返回的路径不会以分隔符结尾

`Base`：`Base(path string) string`

- Base 返回路径的最后一个元素。提取最后一个元素之前，会移除路径尾部的分隔符
- 如果路径为空，Base 返回`.`；如果路径完全由分隔符组成，Base 返回单个分隔符

### 获取扩展名 `Ext()`

语法：`filepath.Ext(path string) string`

只取最后一个点后的部分；如果想取多级扩展，需要自己处理。

### IsAbs() 与 Abs()

`IsAbs`：`filepath.IsAbs(path string) bool`

- 返回路径是否为绝对路径

`Abs`：`filepath.Abs(path string) (string, error)`

- Abs 返回路径的绝对表示形式
- 如果路径不是绝对路径，它将与当前工作目录连接起来，将其转换为绝对路径
- 给定文件的绝对路径名不保证唯一，Abs 会对结果 调用Clean 函数
- 会基于当前工作目录补全相对路径

### `Split()` 与 `SplitList`

`Split`：`Split(path string) (dir, file string)`

- 把一个完整路径拆分成两部分
    - dir：目录部分（会保留结尾分隔符）
    - file：最后一段文件名部分

| 情况      | 结果               |
|---------|------------------|
| 路径为空    | `("", "")`       |
| 路径只有文件名 | `("", "file")`   |
| 末尾有分隔符  | 会把分隔符保留在 `dir` 中 |

`SplitList`：`SplitList(path string) []string`

- 把一个路径列表字符串拆分为切片
- 路径列表指的是系统环境变量中那种多个路径拼接在一起的字符串，比如 PATH
- SplitList 会自动识别当前系统的路径列表分隔符（filepath.ListSeparator），把整个字符串拆成多个路径元素

不同系统用不同字符分隔多个路径

| 系统            | 分隔符 | 示例                                      |
|---------------|-----|-----------------------------------------|
| Linux / macOS | `:` | `/usr/bin:/bin:/usr/local/bin`          |
| Windows       | `;` | `C:\Windows;C:\Program Files;D:\Go\bin` |

### 计算相对路径 `Rel`

语法：`filepath.Rel(basepath, targpath string) (string, error)`

- 如果 `targpath` 不在 `basepath` 下，会返回带 `..` 的路径
- Rel 会对结果 调用Clean 函数

### 路径通配符匹配 `Match`

语法：`filepath.Match(pattern, name string) (matched bool, err error)`

- Match 报告 name 是否与 shell 文件名模式匹配
- 匹配要求模式匹配名称的所有内容，而不仅仅是子字符串
- 当模式格式错误时，唯一可能返回的错误是 ErrBadPattern

> 在 Windows 上，转义被禁用。相反，'\\' 被视为路径分隔符

### 扫描匹配文件路径 `Glob`

返回匹配的路径切片，不会报错但可能返回空切片
语法：`filepath.Glob(pattern string) (matches []string, err error)`

- pattern：要匹配的通配符模式（比如 "*.txt" 或 "data/*.go"）
- matches：返回所有匹配到的文件路径切片
- err：如果模式无效（例如 [ 没有闭合）会返回错误

- 返回所有与 pattern 匹配的文件的名称，如果没有匹配的文件则返回 nil
- pattern 的语法与Match相同
- pattern 可以描述层级结构的名称，例如 /usr/*/bin/ed （假设分隔符为 `/`）

> Glob 会忽略文件系统错误，例如读取目录时的 I/O 错误。当 pattern 格式错误时， 唯一可能返回的错误是ErrBadPattern

Windows 和 Linux 因为文件分隔符不同，推荐做法是使用 `filepath.Join()` 构建路径模式

```
// 使用 filepath.Join 构建模式
pattern := filepath.Join("data", "*.txt")

files, err := filepath.Glob(pattern)
if err != nil {
    fmt.Println("Error:", err)
    return
}

if len(files) == 0 {
    fmt.Println("No matches found.")
    return
}

fmt.Println("Matched files:")
for _, f := range files {
    fmt.Println(" -", f)
}
```

### 遍历目录树 `WalkDir`

> work 和 workdir 区别等见 os 模块 directory-traversal 新版推荐用 WalkDir（性能更高、无多余 Stat 调用）

```
filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
    if err != nil {
        return err
    }
    fmt.Println("Found:", path)
    return nil
})
```

### Separator 与 ListSeparator

等价于 os 模块中的 Separator 和 ListSeparator，见 os 模块补充

```
const (
	Separator     = os.PathSeparator
	ListSeparator = os.PathListSeparator
)
```