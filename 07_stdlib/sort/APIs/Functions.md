# sort 函数

sort 包的函数可以分为下面几种

## 排序执行器

Interface 接口的核心入口，直接实现了对 `sort.Interface` 的集合排序

- Sort(data Interface)
- Stable(data Interface)

> 稳定性差异是核心分界：Stable 保留相等元素的原始相对顺序，Sort 不保证

### `sort.Sort`

对 data 进行原地排序，排序规则来自实现 Interface 接口的 `Len` / `Less` / `Swap`，不返回新切片，直接改变底层数据

```
func Sort(data Interface)
```

> 不稳定排序：如果 `Less(i,j)==false` 且 `Less(j,i)==false`（等价元素），最终相对顺序不保证保持输入顺序

### `sort.Stable`

对 data 进行原地稳定排序，同样使用实现的 Interface 接口的 `Len` / `Less` / `Swap`，不返回新切片，直接改变底层数据

```
func Stable(data Interface)
```

> 稳定排序，对等价元素保持原样输出顺序

## 通用 slice 语法糖

这些函数面向任意 slice（any），通过 less 闭包定义规则，更推荐编写形式，不用写 Interface

- `Slice(x any, less func(i, j int) bool)`
- `SliceStable(x any, less func(i, j int) bool)`
- `SliceIsSorted(x any, less func(i, j int) bool) bool`

> 对 `[]struct`、`[]T` 排序/判有序首选使用

### `sort.Slice`

对切片 x 进行原地不稳定排序，规则由 `Less` 决定

```
func Slice(x any, less func(i, j int) bool)
```

- x 必须为切片否则引发 panic
- 不稳定排序，相等的元素顺序可能会与原始顺序颠倒，稳定排序使用SliceStable 函数

### `sort.SliceStable`

对切片 x 做原地稳定排序，规则由 less 决定，稳定排序，与 Stable 一致，等价元素保持原样顺序

```
func SliceStable(x any, less func(i, j int) bool)
```

> x 必须为切片否则引发 panic

### `sort.SliceIsSorted`

判断切片 x 是否已经按 less 定义的顺序排好

```
func SliceIsSorted(x any, less func(i, j int) bool) bool
```

- 验证的是对所有相邻元素，是否满足没有逆序，也就是不存在 `less(i, i-1) == true` 的情况
- 常用于排序前后断言已经在调用二分查找前做前置校验
- 如果 x 不是一个切片，则该函数会引发 panic

## 内置基础类型快捷排序

对 `[]int` / `[]string` / `[]float64` 直接升序排序

- `Ints(x []int)`
- `Strings(x []string)`
- `Float64s(x []float64)`

> Float64s 具备 `NaN-safe` 的排序规则（等价于 Float64Slice）

### `sort.Ints`

对 `[]int` 进行原地升序排序，等价于 `sort.Sort(sort.IntSlice(nums))`

```
func Ints(x []int)
```

> 默认升序，非稳定排序

### `sort.Strings`

对 `[]string` 按字典序升序排序，等价于 `sort.Sort(sort.StringSlice(strs))`

```
func Strings(x []string)
```

### `sort.Float64s`

对 `[]float64` 进行 NaN-safe 的升序排序，内部使用 Float64Slice，等价于 `sort.Sort(sort.Float64Slice(vals))`

```
func Float64s(x []float64)
```

## 是否已排序检查

校验与断言，用于排序前后校验、调试断言、以及二分搜索前置条件检查

1. 针对基础类型
    - `IntsAreSorted(x []int) bool `
    - `StringsAreSorted(x []string) bool`
    - `Float64sAreSorted(x []float64) bool`
2. 针对 Interface
    - `IsSorted(data Interface) bool`
3. 针对任意 slice + 规则
    - `SliceIsSorted(x any, less func(i, j int) bool) bool`

### `sort.IntsAreSorted`

判断 `[]int` 是否已经升序排列，等价于 `sort.IsSorted(sort.InSlice(nums))`

```
func IntsAreSorted(x []int) bool
```

### `sort.StringsAreSorted`

判断 `[]string` 是否按字典序升序

```
func StringsAreSorted(x []string) bool
```

### `sort.Float64sAreSorted`

判断 `[]float64` 是否按 Float64Slice 规则排好（NaN-safe 规则）

```
func Float64sAreSorted(x []float64) bool
```

### `sort.IsSorted`

判断任意实现了 `sort.Interface` 的集合是否有序

```
func IsSorted(data Interface) bool
```

> 使用定义的 Less，如果 `Less(i, i-1)` 为 true 说明乱序

## 二分查找

Search 系列，主要用于定位元素、找插入位置、阈值分界、区间查询，有序数组上的定位/插入点

1. 核心原语（通用）
    - `Search(n int, f func(int) bool) int`：返回最小的 i 使得 `f(i)==true`
2. 针对基础类型的封装（升序前提）
    - `SearchInts(a []int, x int) int`
    - `SearchStrings(a []string, x string) int`
    - `SearchFloat64s(a []float64, x float64) int`（遵循 float 的 NaN 语义/实现约定）

### `sort.Search`

所有的 Search 系列原语，进行二分查找，返回最小的 i 索引（`[0, n)`）使得 `f(i) == true`

```
func Search(n int, f func(int) bool) int
```

> 前置条件为必须已经排好序，例如，给定一个按升序排序的切片数据，调用
`Search(len(data), func(i int) bool { return data[i] >= 23 })` 返回满足 `data[i] >= 23` 的最小索引 i
>
> 如果调用者想要查找 23 是否在切片中，则必须单独测试 `data[i] == 23`
>
> 搜索按降序排列的数据时，应使用 `<=` 运算符而不是 `>=` 运算符。

### `sort.SearchInts`

在已排序的整数切片中查找 x，并返回 Search 函数指定的索引。如果 x 不存在，则返回值为要插入 x 的索引（也可能是 `len(a)`）。int
切片必须按升序排列

```
func SearchInts(a []int, x int) int
```

等价于：

```
sort.Search(len(a), func (i int) bool{
    return a[i] >= x
}
```

### `sort.SearchStrings`

已排序的字符串切片中查找字符串 x，并返回 Search 指定的索引。如果 x 不存在，则返回值为要插入 x 的索引（也可能是 `len(a)`
），切片必须按字典序升序排列。

```
func SearchStrings(a []string, x string) int
```

### `sort.SearchFloat64s`

`SearchFloat64s` 函数在已排序的 float64s 切片中查找元素 x，并返回由`Search`指定的索引。如果 x 不存在，则返回值为要插入 x
的索引（也可能是 `len(a)`）。float64s 切片必须按升序排列

```
func SearchFloat64s(a []float64, x float64) int
```

## 查找+是否命中新式接口

Search 更语义化版本，更接近比较器风格的查找接口（返回位置 + 是否找到），比 Search 更直接表达命中/未命中

- `Find(n int, cmp func(int) int) (i int, found bool)`

### `sort.Find`

Go 后期新增 的接口，让是否找到这件事更直观，和 Search 一样使用二分查找

```
func Find(n int, cmp func(int) int) (i int, found bool)`
```

| 返回值 i | 含义     | 返回值 found | 含义  |
|-------|--------|-----------|-----|
| `< 0` | 当前位置偏小 | false     | 未找到 |
| `= 0` | 命中目标   | true      | 找到了 |
| `> 0` | 当前位置偏大 | false     | 未找到 |


