# strings 顶层方法

## 内存与比较基础类

- `Clone`：返回 `s` 的一个拷贝
- `Compare`：字典序比较
- `EqualFold`：Unicode case-folding 的不区分大小写比较

### `strings.Clone`

用于获取一个与 `s` 内容相同的新字符串，实现上通常会分配新内存并拷贝字节。

```
func Clone(s string) string
```

- 过度使用可能会导致程序占用更多内存
- 对于长度为 0 的字符串，Clone 会返回字符串 `""` 而不会分配内存

### `strings.Compare`

Compare 函数用于排序或有序比较，比较依据是字典顺序，输出有三种值：

- 如果 `a == b` 则结果为 0
- 如果 `a < b` 则结果为 -1
- 如果 `a > b` 则结果为 +1

```
func Compare(a, b string) int
```

> 通常情况下，最常用的还是使用 `>`、`==`、`<` 进行字符串比较，会更清晰并且速度更快，Compare 的优势是返回三态整数，适用于某些接口或想写得更统一

### `strings.EqualFold`

用于大小写不敏感的比较，而且是 Unicode 语义的大小写折叠，返回 `s` 和 `t` 在 `Unicode case-folding` 规则下是否相等，不是简单的
`ASCII A-Z / a-z` 映射，而是更广泛的 Unicode 规则。

```
EqualFold(s, t string) bool
```

> 类似于使用 `strings.ToLower(s) == strings.ToLower(t)`，但使用 EqualFold 不会产生新的字符串，性能更高并且更符合语义

## 包含/匹配判定类

这组函数主要是查看是否存在某种匹配，返回 bool：

| 需求            | 选哪个            |
|---------------|----------------|
| 固定子串是否存在      | `Contains`     |
| 是否包含某几个字符之一   | `ContainsAny`  |
| 是否包含某个指定 rune | `ContainsRune` |
| 是否存在满足复杂条件的字符 | `ContainsFunc` |

### `strings.Contains`

用于判断 `s` 中是否包含了一个完整的子字符串 `substr`，可以看作判断是否存在某个索引 i 使得 `s[i:i+len()substr] == substr`
，本质上等价于 `strings.Index(s, substr) >= 0`

```
func Contains(s, substr string) bool
```

- 如果 `substr == ""` 则永远返回 true
- 如果 substr 长度大于 s 返回 false

### `strings.ContainsAny`

判断 s 中是否包含 chars 中的任意一个字符，这里的 chars 不是子串集合，而是字符集合。对于 s 中的每一个 rune，只要属于 chars
中的任意 rune 就返回 true，注意 `strings.ContainsAny("hello", "ll")` 返回 true 是因为 `l` 在 `ll` 字符集合中，而不是因为
`ll` 字串存。

```
func ContainsAny(s, chars string) bool
```

如果 chars 为 `""` 则永远返回 false

> 主要用于判断是否包含任一非法字符 / 任一分隔符

### `strings.ContainsRune`

判断 s 中是否包含某一个指定的 Unicode 字符，ContainsAny 的单字符特化版，也就是按照 rune 进行匹配

```
func ContainsRune(s string, r rune) bool
```

### `strings.ContainsFunc`

最灵活的 Contains 判断，判断 s 中是否存在某个 rune，使得 `f(rune) == true`，Contains 的函数版本

```
ContainsFunc(s string, f func(rune) bool) bool
```

**原理**

1. 遍历 s 中的每一个 rune
2. 对每个 rune 调用 `f()` 方法
3. 只要有一个返回 true 则返回 true

> 经常用于判断字符串是否包含数字、空白字符等

## 计数与定位类

这组函数主要用于指定元素和统计数量

- 计数：`Count`
- 正向查找：`Index`、`IndexAny`、`IndexByte`、`IndexRune`、`IndexFunc`
- 反向查找：`LastIndex`、`LastIndexAny`、`LastIndexByte`、`LastIndexFunc`

### `strings.Count`

统计子串在字符串中出现的次数（不重叠）

```
func Count(s, substr string) int
```

- 统计非重叠出现次数，`aaaa` 中如果统计 `aa` 则结果是2
- 如果 `substr == ""` 则返回 `1+len(s)`

### `strings.Index` / `strings.LastIndex`

`strings.Index` 用于查找子串 substr 在 s 中第一次出现的位置，返回满足条件的最小 i，也就是找到就返回，如果不存在返回 -1

- 如果 `substr == ""` 则返回 0
- 大小写敏感
- 返回的是 bytes index 字节索引

`strings.LastIndex` 用于从右往左查找 substr 最后一次出现的位置，返回满足条件的最大 i 规则与 index 完全对称，主要用于判断文件后缀等

- 如果 `substr == ""` 则直接返回 `s.len`
- 大小写敏感
- 返回的是 bytes index 字节索引

```
func Index(s, substr string) int
func LastIndexAny(s, chars string) int
```

### `strings.IndexAny` / `strings.LastIndexAny`

查找 s 中第一个 / 最后一个属于 chars 集合的字符的下标索引，chars 是字符集合，而不是子串集合，如果没找到返回 -1

```
func IndexAny(s, chars string) int
func LastIndexAny(s, chars string) int
```

### `strings.IndexByte` / `strings.LastIndexByte`

为 `ASCII / byte` 场景提供最快路径，只匹配单个 byte，不考虑 UTF-8，返回 第一个/ 最后一个匹配到的 byte 的索引，通常比
`IndexRene / IndexAny` 更快，明确是 `ASCII / byte` 使用 `IndexByte / LastIndexByte`

```
func IndexByte(s string, c byte) int
func LastIndexByte(s string, c byte) int
```

> 如果 c 不存在于 s 中则返回 -1，不能处理中文等

### `strings.IndexRune`

按 rune 查找 Unicode 字符首次出现的位置，注意返回的是 byte index，但是如果中文等一个字符占三个字节等，所以返回的并不是字符索引

```
func IndexRune(s string, r rune) int
```

> 如果 s 中不存在 r，则返回 -1，可以处理中文

### `strings.IndexFunc` / `strings.LastIndexFunc`

自定义规则进行查找，查找第一个/最后一个满足条件的 rune 的 byte index，主要原因复杂规则进行下标查找，如果没有满足 `f(c)` 的
rune 则返回 -1

```
func IndexFunc(s string, f func(rune) bool) int
func LastIndexFunc(s string, f func(rune) bool) int

```

## 前后缀与切割类

这组 cut 函数特点为只切一次且显示的返回是否命中，适合解析 `key=value`、`Header: value` 这类结构。

- 前后缀判定：`HasPrefix`、`HasSuffix`
- Cut切割（Go 1.18+）：`Cut`、`CutPrefix`、`CutPrefix`

### `strings.HasPrefix`

用于判断字符串是否以某个前缀开头

```
func HasPrefix(s, prefix string) bool
```

- 如果 `prefix == ""` 则返回 true
- 比较的是字节级别，大消息敏感
- 不分配内存，`O(len(prefix))`

### `strings.HasSuffix`

判断字符串是否以某个后缀结尾

```
HasSuffix(s, suffix string) bool
```

- 如果 `suffix == ""` 则返回 true
- 本质是尾部对齐比较
- 比 LastIndex 更安全（避免中间误命中）

### `strings.Cut`

以 sep 为分隔符，把字符串切为前后两部分，只切割一次

```
func Cut(s, sep string) (before, after string, found bool)
```

- 返回的 found 会告知是否切割成功
- 如果没有找到：
    - `before = s`
    - `after = ""`
    - `found = false`

> Cut 比 Split 好用，Cut 不分配 `[]string`，不需要判断长度，常用于解析 `key=value`、`Header: value`、URL 参数处理等

### `strings.CutPrefix`

如果字符串 s 以 prefix 开头，则切去 prefix

```
func CutPrefix(s, prefix string) (after string, found bool)
```

- 如果 s 不以 prefix 开头，则返回 `s, false`
- 如果 `prefix == ""` 则返回 `s, true`

语义等价于：

```
if strings.HasPrefix(s, prefix){
    return s[len(prefix):], true
}
```

### `strings.CutSuffix`

如果字符串 s 以 suffix 结尾，则切去 suffix

```
func CutSuffix(s, suffix string) (before string, found bool)
```

- 如果 s 不以 suffix 结尾，则返回 `s, false`
- 如果 `suffix == ""`，则返回 `s, true`

## 分割/聚合类

这组函数由于把字符串拆成段/把段拼接回去，`Seq` 和 `Lines` 更适合大字符串、流式写入或更希望避免一次性分配 `[]strings` 的场景

- 立即分隔（返回 `[]string`）：`Split`、`SplitN`、`SplitAfter`、`SplitAfterN`、`Fields`、`FieldsFunc`
- 惰性分割（iterator，返回 `iter,Seq[string]`，Go 1.24+引入）：`SplitSeq`、`SplitAfterSeq`、`FieldsSeq`、`FieldsFuncSeqs`、
  `Lines`（这几个不立即拆分、不分配切片、按需产生字符串片段的迭代器），iter 使用见 [iter](../../iter)
- 聚合：`Join`

> Spilt 系列是基分隔符字符串，Fields 系列是基于空白规则

### `strings.Split`

通用拆分，按照分隔符 sep 字符串对 s 字符串进行拆分，返回拆分后的字符串切片并舍弃分隔符

```
func Split(s, sep string) []string
```

- 先在 s 中查找所有 sep，再按照出现位置进行拆分，sep 不出现在结果中
- 如果 `sep == ""` 则按照 UTF-8 rune 进行拆分
- 如果 `s == ""` 则返回 `[]string{""}`

### `strings.SplitN`

将字符串 s 切成以分隔符 sep 分隔的子字符串，并返回分隔符之间的子字符串切片，n 决定了要返回的字符串的数量

```
func SplitN(s, sep string, n int) []string
```

- `n > 0`：最多 n 个子串，最后一个子串是未分隔的剩余部分
- `n == 0`：结果为 nil（0 个子字符串）
- `n < 0`：所有子字符串，等同于 `Split` 函数，通常传入 -1

### `strings.SplitAfter`

将字符串 s 分成每次 sep 出现后的所有子字符串，并返回这些子字符串的切片，也就是在拆分的时候会把 sep 分隔符保留在前一段的末尾而不是像
Split 一样舍弃分隔符，比如 `strings.SplitAfter("a,b,c", ",")` 结果为 `["a,", "b,", "c"]`

```
func SplitAfter(s, sep string) []string
```

- 如果 s 不包含 sep 且 sep 不为空，则返回长度唯一的切片且切片唯一元素为 s
- 如果 sep 为空，会在每个 UTF-8 序列后进行分割
- 如果 s 和 seq 都为空，则返回一个空切片
- 等价于 `strings.SplitAfterN(s, seq, -1)`

### `strings.SplitAfterN`

`SplitAfter` + 数量限制，把 s 在每个 sep 字符串之后切成子字符串，并返回这些子字符串的切片

```
func SplitAfterN(s, sep string, n int) []string
```

- `n > 0`：最多 n 个子串；最后一个子串将是未分割的剩余部分
- `n == 0`：结果为 nil（零个子字符串）
- `n < 0`：所有子字符串，等同于 `SplitAfter`

### `strings.Fields`

按照 Unicode 空白字符拆分字段，和 `Split` 系列不同，会舍弃字符串开头和结尾的连续空格

```
func Fields(s string) []string
```

- 连续空白视为一个分隔
- 如果 s 中只包含空白字符则返回一个空切片
- 这种情况下返回的切片中的每个元素都非空
- 空白由 `unicode.IsSpace` 定义

### `strings.FieldsFunc`

自行决定分隔符，当 `f(rune) == true` 时，该 rune 被视为分隔符，连续命中仍只产生一个分割，和 `Fields` 一样会把开头结尾符合规定的
rune 进行舍弃

```
func FieldsFunc(s string, f func(rune) bool) []string
```

- 如果 s 中所有 rune 都满足 `f(c)` 或字符串 s 为空则返回一个空切片
- 正常情况下返回的切片中的每个元素都非空

### `strings.Join`

把字符串切片使用指定分隔符拼接成一个字符串，内部会预计总长度，并且是一次性分配，比循环拼接性能更高

```
func Join(elems []string, sep string) string
```

### `strings.SplitSeq / SplitAfterSeq`

Split 和 SplitAfter 的惰性版本，按照 sep 进行分割，SplitSeq 舍弃分隔符，SplitAfterSql 保留分隔符，返回一个一次性使用迭代器，迭代器中返回的字符串与
`Split(s, seq)` 和 `SplitAfter(s, seq)` 中的字符相同

```
func SplitSeq(s, sep string) iter.Seq[string]
func SplitAfterSeq(s, sep string) iter.Seq[string]
```

### `strings.FieldsSeq / FieldsFuncSeqs`

Fields 的惰性版本，FieldsSeq 按照 Unicode1空白字符进行拆分，FieldsFuncSeq 使用自定义规则进行拆分，都是返回一个一次性使用的迭代器，迭代器中返回的字符串与
`Fields(s)` 和 `FieldsFunc(s)` 返回的字符串相同

```
func FieldsSeq(s string) iter.Seq[string]
func FieldsFuncSeq(s string, f func(rune) bool) iter.Seq[string]
```

### `strings.Lines`

按换行符拆分遍历字符串，返回一个一次性使用的迭代器，用于遍历字符串 s 中所有以换行符结尾的行

```
func Lines(s string) iter.Seq[string]
```

- 迭代器返回的行包含起结尾的换行符，
- 如果 s 为空字符串，则迭代器不返回任何行
- 如果 s 不以换行符结尾，则返回的最后一行也不会以换行符结尾

## 替换/映射/构造字符串类

主要用于内容重写、替换字串、或按 rune 映射改写

- `Repeat`：构造重复文本
- `Replace`：替换指定 n 个子串
- `ReplaceAll`
- `Map`

### `strings.Repeat`

把一个字符串重复 count 次并拼接返回一个新的字符串

```
func Repeat(s string, count int) string
```

- 如果 `count == 0` 会返回 `""` 空字符串
- 如果 `count < 0` 会触发 panic

> 内部会预计最终生成字符串的长度并做一次性分配，比手写循环 `+=` 效率更高

### `strings.Replace`

把字符串 s 中最多 n 个 old 子串替换为 new 并返回一个新的字符串副本

```
func Replace(s, old, new string, n int) string
```

- 如果 old 为 `""` 空字符串则在 s 的开头以及每个 UTF-8 序列之后进行匹配（不推荐）
- 从左到右进行替换，最多替换 n 次，如果 `n < 0` 表示替换全部，等同于 `strings.ReplaceAll`

### `strings.ReplaceAll`

把字符串中所有 old 子串替换为 new 并返回一个新的字符串副本，等价于 `strings.Replace(s, old, new, -1)`

```
func ReplaceAll(s, old, new string) string
```

### `strings.Map`

按 rune 遍历字符串，对每个字符做映射变换，这是字符级别的改写，不是子串级

```
func Map(mapping func(rune) rune, s string) string
```

**底层实现**

- 对 s 中的每个 rune 调用 `mapping(r)`
- 返回值 `>=0` 则写入该 rune，返回值 `<0` 则删除该 rune
- 返回的要是全新的字符串副本

## 大小写与 Unicode 规范化

这类函数强调 Unicode 语义，其中 Title 已弃用，通常用 ToTitle 或更明确的文本处理方式替代

- `ToLower`
- `ToUpper`
- `ToTitle`
- `ToLowerSpecial`
- `ToUpperSpecial`
- `ToTitleSpecial`
- `ToValidUTF8`
- `Title`（deprecated）

### `strings.ToLower`

把字符串转换为 Unicode 小写并返回一个新的字符串副本

```
func ToLower(s string) string
```

### `strings.ToUpper`

把字符串转换为 Unicode 大写并返回一个新的字符串副本

```
func ToUpper(s string) string
```

### `strings.ToTitle`

把字符串转换为 Title Case 也就是标题大小写（不是每个单词的首字母大写，而是 Unicode 定义的 title-case 映射）

```
func ToTitle(s string) string
```

### `strings.ToLowerSpecial / ToUpperSpecial / ToTitleSpecial`

为特定语言/区域应用特殊大消息规则，某些语言的大小写规则不符合通用 Unicode 映射，此时需要使用 Special 类型进行处理

```
func ToTitleSpecial(c unicode.SpecialCase, s string) string
func ToUpperSpecial(c unicode.SpecialCase, s string) string
func ToLowerSpecial(c unicode.SpecialCase, s string) string
```

> 绝大多数业务代码用不到，做国际化或语言学处理才考虑，否则坚持用 `ToLower`/ `ToUpper` / `ToTitle`

### `strings.ToValidUTF8`

把字符串中非法的 UTF-8 字节序列替换为指定内容（健壮性处理函数，而不是格式化函数），同样返回一个新的字符串

```
func ToValidUTF8(s, replacement string) string
```

**底层逻辑**

1. 扫描字符串 s
2. 如果遇到非法 UTF-8 字节序列，使用 replacement 进行替换
3. 保证返回结果是合法的 UTF-8

## 裁剪/清洗类

这一类主要用于去掉边缘噪声（空白/特定字符集/固定前后缀）

- `Trim`
- `TrimLeft`
- `TrimRight`
- `TrimSpace`
- `TrimFunc`
- `TrimLeftFunc`
- `TrimRightFunc`
- `TrimPrefix`
- `TrimSuffix`

### `strings.Trim`

从字符串两端删除属于 cutset （cutset 是字符集合，不是字符串）的字符并返回一个新的字符串副本

```
func Trim(s, cutset string) string
```

### `strings.TrimLeft / TrimRight`

`TrimLeft` 只处理左边的字符，`TrimRight` 只处理右边的字符，同样 cutset 是字符集合而不是字符串

```
func TrimLeft(s, cutset string) string
func TrimRight(s, cutset string) string
```

### `strings.TrimSpace`

从字符串 s 两端删除所有 Unicode 空白字符，包括 ` `、`\t`、`\n` 等

```
func TrimSpace(s string) string
```

常用于处理用户输入，最常用安全

### `strings.TrimFunc / TrimLeftFunc / TrimRightFunc`

这几个 TrimFunc 更为灵活，可以自定义哪些字符应该被裁剪

```
func TrimFunc(s string, f func(rune) bool) string
func TrimLeftFunc(s string, f func(rune) bool) string
func TrimRightFunc(s string, f func(rune) bool) string
```

> 删除的是满足了 `f(c)` 的 Unicode 码点

### `strings.TrimPrefix`

如果 字符串 s 以 prefix 开头，就删除并返回一个新的字符串副本，如果不是则原样返回新的字符串副本

```
func TrimPrefix(s, prefix string) string
```

> 不会返回是否找到，也就是无法直接区分是否执行成功，精准处理建议使用 CutPrefix

### `strings.TrimSuffix`

如果 字符串 s 以 suffix 结尾，就删除并返回一个新的字符串副本，如果不是则原样返回新的字符串副本

```
func TrimSuffix(s, suffix string) string
```

> 同样无法区分是否执行成功，精准处理建议使用 CutSuffix