# `strings.Replacer`

`strings.Replacer` 是一个可复用的，多对字符串替换引擎，用于高效，一次遍历的完成多组替换。

```
type Replacer struct {
	// contains filtered or unexported fields
}
```

常规使用 `strings.ReplaceAll` 进行替换时每一次都需要完整扫描字符串并且生成新的 string，时间复杂度约等于 `O(k*n)`，
使用 `strings.Replacer` 用于把欧茨的替换+多次扫描压缩为一次构建规则+一次扫描。

## 构造函数

`strings.NewReplacer` 用于用一组 `old -> new` 的映射规则来构建 `strings.Replacer` 对象

```
func NewReplacer(oldnew ...string) *Replacer
```

- 参数形式为：`old1, new1, old2, new2, ...`
- 注意参数个数必须为偶数，否则直接 panic

## 方法

`Replacer` 一共有 `Replace` 和 `WriteString` 两个执行类方法

### `Replace`

对字符串 s 应用所有的替换规则并返回替换后的新字符串（最直接且常用的执行方式）

```
func (r *Replacer) Replace(s string) string
```

**行为特征**

- 不修改源字符串，而是返回一个新的 string
- 内部使用一次扫描
- 按 `r` 的规则并行匹配而不是顺序替换

Replacer 的替换不是级联的，而是基于原输入的一次性匹配：

```
r := strings.NewReplacer("a", "b", "b", "c")
r.Replace("a")
```

结果是 `"b"`，而不是 `"c"`

### `WriteString`

IO 优化接口，用于把替换后的结果直接写入 writer 中，而不是返回 string

```
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
```

如果使用：

```
out := r.Replace(s)
w.Write([]byte(out))
```

会先生成一个新的字符串，然后再一次性拷贝到 write 中，而是用 `WriteString` 可以做到：

1. 边扫描，边替换，边写出
2. 减少中间分配
3. 更便于大文本的替换输出

## Replacer 对比 ReplaceAll

| 场景        | 推荐                     |
|-----------|------------------------|
| 只替换 1 种规则 | `strings.ReplaceAll`   |
| 替换规则很多    | `strings.Replacer`     |
| 规则需要复用    | `strings.Replacer`     |
| 输出直接写流    | `Replacer.WriteString` |

> ReplaceAll 是便利函数，Replacer 是引擎。