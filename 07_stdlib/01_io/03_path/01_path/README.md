# Path 模块

path 包提供了对斜杠分隔路径字符串（如 `a/b/c`）的操作函数。
与操作系统无关，不会使用 Windows 的 \。
主要给 URL 或类 Unix 路径设计。
适用于：

- URL 路径（/api/v1/user）
- Web 资源路径（/static/css/main.css）
- 内存中虚拟路径（非真实文件路径）

> 如果在操作真实文件系统应该用 `path/filepath` 模块

## 常用函数总览

| 函数                            | 作用                      | 示例                                       |
|-------------------------------|-------------------------|------------------------------------------|
| `Base(path string)`           | 获取路径最后一段                | `Base("/a/b/c.txt") → "c.txt"`           |
| `Dir(path string)`            | 获取路径所在目录                | `Dir("/a/b/c.txt") → "/a/b"`             |
| `Ext(path string)`            | 获取扩展名                   | `Ext("/a/b/c.txt") → ".txt"`             |
| `Join(elem ...string)`        | 拼接路径                    | `Join("a","b","c") → "a/b/c"`            |
| `Clean(path string)`          | 规范化路径（去掉 `.`、`..`、多余斜杠） | `Clean("/a/b/../c") → "/a/c"`            |
| `Split(path string)`          | 拆分为目录和文件名               | `Split("/a/b/c.txt") → "/a/b/", "c.txt"` |
| `IsAbs(path string)`          | 判断是否为绝对路径               | `IsAbs("/a/b") → true`                   |
| `Match(pattern, name string)` | 通配符匹配                   | `Match("*.go","main.go") → true`         |

## 长于详解

### 1. 拼接路径 `path.Join`

Join 将任意数量的路径元素连接成一个路径，并用斜杠分隔。空元素将被忽略。
结果为 Cleaned（去掉冗余路径符号，如 `.、..、` 重复斜杠）。但是，如果参数列表为空或其所有元素均为空，则 Join 返回空字符串。

语法：`Join(elem ...string) string`

> 安全地将多个路径段拼起来，自动插入 `/` 并去重

### 2. 规范化路径 `path.Clean()`

去掉冗余路径符号，如 `.、..、` 重复斜杠。
语法：`Clean(path string) string`

```
fmt.Println(path.Clean("/a/b/../c/./d//")) // "/a/c/d"
fmt.Println(path.Clean("a/b/../../c"))     // "c"
fmt.Println(path.Clean("/../"))            // "/"
```

> 常用于服务器中防止目录越界（安全路径处理）
> Clean 通过纯词法处理返回与 path 等价的最短路径名。它会迭代应用以下规则，直到无法进行进一步处理：
> 1. 用一个斜杠替换多个斜杠
> 2. 消除每个 . 路径名元素（当前目录）
> 3. 消除每个内部 `..` 路径名元素（父目录）以及其前面的非 `..` 元素
> 4. 消除以 `..` 开头的元素：即，将路径开头的`/..`替换为`/`。
     > 仅当返回的路径是根`/`时，才以斜杠结尾，如果此过程的结果为空字符串，则 Clean 返回字符串`.`

### 3. 拆分目录和文件名 `path.Split()`

语法：`path.Split(path string) (dir, file string)`

- Split 会在最后一个斜杠后立即拆分 path，将其分为目录和文件名部分
- 如果 path 中没有斜杠，Split 则返回一个空目录，并将文件设置为 path
- 返回值具有 `path = dir+file` 的属性

> 最后一个必须时文件，不能带 `/`，否则 file 为空

### 4. 获取路径最后一段 `path.Base()`

语法：`path.Base(path string) string`

- Base 返回路径的最后一个元素。提取最后一个元素之前，会删除路径尾部的斜杠
- 如果路径为空，Base 返回`.`，如果路径完全由斜杠组成，Base 返回`/`

> 末尾 / 会被忽略，除非路径只有 /

### 5. 获取目录部分 `path.Dir()`

Dir 返回路径中除最后一个元素之外的所有元素，通常是路径的目录。
语法：`path.Dir(path string) string`

- 使用Split删除最后一个元素后，路径将被 Cleaned 并删除尾部的斜杠
- 如果路径为空，Dir 返回`.`；如果路径完全由斜杠后跟非斜杠字节组成，Dir 返回单个斜杠
- 在其他情况下，返回的路径不以斜杠结尾。

### 6. 获取扩展名 `path.Ext()`

语法：`path.Ext(path string) string`

- Ext 返回 path 使用的文件扩展名
- 扩展名是从 path 中最后一个以斜杠分隔的元素的最后一个点开始的后缀
- 如果没有点，则为空

### 7. 是否为绝对路径 `path.IsAbs()`

path 包是设计给 URL 或类 Unix 路径 用的，它不识别 Windows 的反斜杠 `\`，也不识别盘符 `C:`。
所以在 Windows 上使用这个方法判断路径永远都是 false，可以使用 filepath.IsAbs 进行判断

语法：`path.IsAbs(path string) bool`

> path 包认为 以 "/" 开头 就是绝对路径（不会判断盘符，如 Windows 的 C:\）

### 8. 通配符匹配 `path.Match()`

类似于 shell 匹配模式（glob pattern）。

语法：`path.Match(pattern, name string) (matched bool, err error)`

| 符号       | 含义             |
|----------|----------------|
| `*`      | 匹配任意字符（不含 `/`） |
| `?`      | 匹配单个字符         |
| `[...]`  | 匹配字符集          |
| `[^abc]` | 不匹配集合中的字符      |
