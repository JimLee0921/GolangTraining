# `sort.IntSlice`

IntSlice 是 `[]int` 的一个轻量包装，为它补齐 `sort.Interface`，并额外提供 Sort / Search 等便利方法，默认为升序排序

```
type IntSlice []int
```

## 方法

### Len()

实现 Interface 接口方法，返回数量

```
func (x IntSlice) Len() int {
    return len(x)
}
```

### Less(i, j)

核心排序规则，默认就是升序排序

```
func (x IntSlice) Less(i, j int) bool {
    return x[i] < x[j]
}
```

### Swap(i, j)

交换元素，很简单的互换：`x[i], x[j] = x[j], x[i]`

```
func (x IntSlice) Swap(i, j int)
```

### Sort

语法糖，用于面向对象调用，等价于 `sort.Sort(x)`

```
func (x IntSlice) Sort()
```

### Search(x)

二分查找，等价于内部调用 `sort.SearchInts(p, x)`，必须已经实现完排序才能使用

```
func (p IntSlice) Search(x int) int
```
