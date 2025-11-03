## 字符串大小写处理

| 函数                   | 说明                   | 作用范围  | 是否按 Unicode 规则 | 示例                               |
|----------------------|----------------------|-------|----------------|----------------------------------|
| `strings.ToUpper(s)` | 转大写                  | 整个字符串 | 是              | `"go语言" -> "GO语言"`               |
| `strings.ToLower(s)` | 转小写                  | 整个字符串 | 是              | `"GO语言" -> "go语言"`               |
| `strings.Title(s)`   | 每个单词首字母大写（不准确）       | 单词首字母 | ⚠已废弃，不推荐       | `"hello world" -> "Hello World"` |
| `strings.ToTitle(s)` | 全大写，适用于 Unicode 上层规则 | 整个字符串 | 更智能            | 比 `ToUpper` 更语义化                 |

### Title vs ToTitle

`strings.Title(s)`

- 已废弃（Deprecated）
- 逻辑简单：遇到分隔符就把下一个字母变大写
- 对 Unicode 和多语言支持差

```
strings.Title("go language") // "Go Language"
strings.Title("hello-world") // "Hello-World"
```

`strings.ToTitle(s)`

- 不改变分词方式
- 直接将每个字符变成对应的标题形式（语言学意义上的 Title Case）
- 更符合 Unicode 规则

```
strings.ToTitle("go language") // "GO LANGUAGE"
```

### 地区化（特殊规则）版本

某些语言（比如土耳其语）中，大写/小写不是简单 1:1 对应。

Go 提供：

```
strings.ToUpperSpecial(c, s)	
strings.ToLowerSpecial(c, s)	
strings.ToTitleSpecial(c, s)
```

其中 c 为 unicode.SpecialCase，比如：

```
import "unicode"

strings.ToLowerSpecial(unicode.TurkishCase, "İSTANBUL")
// 输出："istanbul"
```

> 在土耳其语中，“İ” 的小写不是“i”，而是“ı”，所以需要语言特殊规则

### 忽略大小写比较

Go 标准库中 不建议用 `ToUpper / ToLower` 后再比较，因为：

- 这样会产生额外的字符串分配（性能差）
- 有些 Unicode 字符大小写映射不对等（可能比较不准确）

Go 提供了 专门的大小写无关比较函数：

```
func EqualFold(s, t string) bool
```

忽略大小写进行字符串比较，并且 遵循 Unicode 规则。

```
strings.EqualFold("Go", "go")       // true
strings.EqualFold("Gopher", "gopher") // true
strings.EqualFold("GO语言", "go语言") // true 
```

> 对 非 ASCII 字符也安全，比如中文会被直接按 rune 比较，不受大小写影响

| 比较方式        | 是否区分大小写 | 性能    | 说明           |
|-------------|---------|-------|--------------|
| `==`        | 区分大小写   | 最快    | 完全匹配         |
| `EqualFold` | 忽略大小写   | 稍微慢一点 | 推荐的大小写无关比较方式 |
