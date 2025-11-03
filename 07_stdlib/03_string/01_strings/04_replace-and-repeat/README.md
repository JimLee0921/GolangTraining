## 字符串内容替换与生成

Go 的字符串是不可变的，所以不能直接改，只能生成一个新字符串。
而 `Replace / Repeat` 正是用来构造新字符串的。

### Replace

```
func Replace(s, old, new string , n int) string
```

Replace 函数返回字符串 s 的副本，其中前 n 个不重叠的 old 元素被 new 元素替换。
如果 old 元素为空，则它在字符串的开头以及每个 UTF-8 序列之后进行匹配，对于 k-rune 字符串，最多会生成 k+1 个替换。
如果 n < 0，则替换次数没有限制等同于 ReplaceAll。

> n 为负数时常使用 -1

### ReplaceAll

```
func ReplaceAll(s, old, new string) string
```

替换所有出现的内容。
ReplaceAll 函数返回字符串 s 的副本，其中所有不重叠的 old 元素都被替换为 new 元素。Replace(s, old, new, -1) 完全等价

### Repeat

```
func Repeat(s string , count int) string
```

Repeat 返回一个由字符串 s 的 count 个副本组成的新字符串。
如果 count 为负数或 `(len(s) * count)` 的结果溢出，则会引发 panic。

**常用于**

- 生成分隔符
- 生成测试字符串
- 做简单分隔线