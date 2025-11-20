## 文本请理

> TrimSpace -> 去空白
>
> Trim -> 去所有指定字符
>
> TrimPrefix/Suffix -> 去明确前后缀

### TrimSpace

```
func TrimSpace(s string ) string
```

TrimSpace 返回字符串 s 的一个切片，其中所有前导和尾随的空格都被删除，如 Unicode 所定义

- 空格 " "
- 制表符 \t
- 换行 \n、\r

### Trim

Trim 返回字符串 s 的一个切片，其中删除了 cutset 中包含的所有前导和尾随 Unicode 代码点

```
func Trim(s, cutset string ) string
```

> 不是去子串，而是去字符集合

```
strings.Trim("abchelloabcccc", "abc")  // "hello"
``` 

> 它会把前后所有属于 a b c 这三个字符的部分全部去掉

### TrimPrefix

TrimPrefix 函数返回 s 时会移除提供的前缀字符串。
如果 s 本身不以前缀字符串开头，则返回原样，不会报错。

```
func TrimPrefix(s, prefix string) string
```

- 去掉 URL 前缀
- 去掉日志前缀
- 去掉路径中的特定启动符号

### TrimSuffix

TrimSuffix 函数返回字符串 s，但移除其末尾提供的后缀字符串。
如果字符串 s 本身不以后缀结尾，则返回原字符串 s。

```
func TrimSuffix(s, suffix string) string
```

常用于：

- 去掉文件后缀
- 去掉换行符
- 去掉多余的分隔符

### TrimLeft

```
func TrimLeft(s, cutset string ) string
```

TrimLeft 返回字符串 s 的一个切片，其中删除了 cutset 中包含的所有前导 Unicode 代码点。
要删除前缀，应改用 `TrimPrefix`

### TrimRight

```
func TrimRight(s, cutset string ) string
```

TrimRight 返回字符串 s 的一个切片，其中 cutset 中包含的所有尾随 Unicode 代码点都被删除。
要删除后缀，应改用 `TrimSuffix`

## 补充

| 方法                      | 作用              | 典型用途           |
|-------------------------|-----------------|----------------|
| `Trim(s, cutset)`       | 去掉**两端**出现的任意字符 | 去掉特定字符，如标点、换行  |
| `TrimLeft(s, cutset)`   | 去掉**左侧**字符      | 去掉左前缀式字符       |
| `TrimRight(s, cutset)`  | 去掉**右侧**字符      | 去掉右尾随符号        |
| `TrimSpace(s)`          | 去掉两端的**所有空白字符** | 最常用的去空格        |
| `TrimPrefix(s, prefix)` | 若存在前缀则移除        | 去掉特定前缀         |
| `TrimSuffix(s, suffix)` | 若存在后缀则移除        | 去掉特定后缀         |
| `TrimFunc(s, f)`        | 根据函数判断要不要去掉字符   | 比如自定义是否为标点、数字等 |
| `TrimLeftFunc(s, f)`    | 左侧自定义判定         | 与 TrimFunc 同逻辑 |
| `TrimRightFunc(s, f)`   | 右侧自定义判定         | 与 TrimFunc 同逻辑 |




