# `sort.StringSlice`

StringSlice 是 `[]string` 的排序视图，实现了 `sort.Interface`，默认按字典序（lexicographical order）升序排序。

```
type StringSlice []string
```

> 默认就是按照 UTF-8 字节序逐字节比较

## 方法

### Len

元素数量

```
func (x StringSlice) Len() int {
    return len(x)
}
```

### Less(i, j)

字符串排序规则，内部实现使用 `return x[i] < x[j]` 默认字典序

```
func (x StringSlice) Less(i, j int) bool
```

### Swap(i, j)

元素互换，等价于 `x[i], x[j] = x[j], x[i]`

```
func (x StringSlice) Swap(i, j int)
```

### Sort()

语法糖，等价于：`sort.Sort(sort.StringSlice(strs))`

```
func (x StringSlice) Sort()
```

### Search(x)

二分查找，返回最小的 i，使得 `p[i] >= x`，注意调用前，slice 必须已经按照字典序排好，否则结果是未定义的

```
func (p StringSlice) Search(x string) int
```

> 等价于 `sort.SearchStrings(p, x)`