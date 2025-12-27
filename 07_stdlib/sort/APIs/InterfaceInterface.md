# `sort.Interface`

`srot.Interface` 接口是 Go 排序系统的最小抽象边界，主要描述三件事

1. 有多少个元素
2. 两个元素比较时谁在前
3. 如何交换两个元素

## 接口结构

```
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```

接口定义的三个方法分别对应排序算法三个问题：

- `Len()`： 一共要处理多少个元素
- `Less(i, j)`： i 是否应该排在 j 前面
- `Swap(i, j)`： 如果要把 i 和 j 交换，怎么换

### Len()

返回集合长度，排序算法通过下标 `0 ~ Len()-1` 访问元素，`Len()` 在排序过程中不能变化，如果长度变化，行为是未定义的，排序算法假定集合是一个静态视图

### Swap()

排序算法不能直接操作数据，排序算法并不知道数据内部结构，只是简单的用自定义的 `Swap(i, j)` 进行元素交换

- 可以交换 slice 元素
- 可以同时交换多个字段
- 甚至可以在 swap 中维护额外索引（不推荐）

### Less()

`Less(i, j)` 定义的是顺序，不是值的大小

| Less(i, j) | Less(j, i) | 含义     |
|------------|------------|--------|
| true       | false      | i < j  |
| false      | true       | j < i  |
| false      | false      | i == j |
| true       | true       | 非法     |

绝对不能出现 `Less(i, j)` 和 `Less(j, i)` 同时为 true

> `sort.Sort` 不保证稳定，`sort.Stable` 保证稳定
>
> 稳定排序就是如果 `Less(i, j) == false` 且 `Less(j, i) == false` 那么 Stable 会保留原始顺序

**float 和 NaN**

```
NaN < x == false
x < NaN == false
Nan < NaN == false
```

这会导致所有 NaN 彼此相等但又无法形成可靠顺序，需要使用 `sort.Flat64Slice.Less`，内部对 NaN 做了特殊处理

> 如果排序的数据可能含有 NaN，不要自己写比较规则

## `sort.Reverse`

`sort.Reverse` 函数用于在不改动原排序逻辑的前提下，把排序顺序整体反转

- 原本是升序改为降序
- 原本是 A 在前改为 A 在后
- 不需要重写 Less

```
func Reverse(data Interface) Interface
```

- 入参参数是 Interface
- 返回的还是 Interface

> 不是排序函数，而是一个适配器/装饰器，使用先使用 Reverse 进行排序规则反转，再使用 Sort 等进行排序
