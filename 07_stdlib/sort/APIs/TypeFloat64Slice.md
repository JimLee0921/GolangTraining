# `sort.Float64Slice`

`sort.Float64Slice` 是一个专门为 float64 排序而设计的安全排序实现，主要为了正确处理 NaN（Not a Number 非数值）参与排序的情况

```
type Float64Slice []float64
```

Float64Slice 实现了 `sort.Interface` 接口：`Len()`、`Less(i, j int)`、`Swap(i, j int)`，
可以直接 `sort.Sort(sort.Float64Slice(data))` 或更常用的 `sort.Flat64s(data)`

## 对于 Nan 的处理

在 IEEE 754 浮点标准中 `math.Nan()` 有一个特性就是 NaN 与任何值包括它自己比较结果都为 false，而这样如果 data 中有
NaN，在比较时 `Less(i, j)` 和 `Less(j, i)` 可能永远都返回 false，排序算法也就无法建立 Strict Weak Ordering，可能就会导致顺序混乱

`sort.Float64Slice` 明确把 NaN 放在一侧，通常是最后来保证 Less 满足 Strict Weak Ordering

- 所有非 NaN 数值正常比较
- NaN 统一当作最大或最小

## 相关方法

### Len()

返回切片长度，用于满足 `sort.Interface`

```
func (x Float64Slice) Len() int {
    return len(x)
}
```

### Less(i, j)

判断 `x[i]` 是否应该排在 `x[j]` 前面

```
func (x Float64Slice) Less(i, j int) bool
```

实现逻辑为：

```
x[i] < x[j] || (math.IsNaN(x[i]) && !math.IsNaN(x[j]))
```

- `x[i] < x[j]` 正常比较，适用于两个都不是 NaN 的正常浮点排序
- `math.IsNaN(x[i]) && !math.IsNaN(x[j])` 为 NaN 特殊处理
    - 如果 i 是 NaN，而 j 不是 NaN，i 会排在前面
    - NaN 最后被统一放在最后面

### Swap(i, j)

实现 Interface 定义交换规则

```
func (x Float64Slice) Swap(i, j int)
```

等价于 `x[i], x[j] = x[j], x[i]`

### Sort()

语义糖，用于方法调用，提供面向对象式写法，`x.Sort()` 等价于 `sort.Sort(x)`，内部等价于：

```
func (x Float64Slice) Sort() {
    sort.Sort(x)
}
```

### Search(x)

已排序 slice 的二分查找，对接收者执行 `SearchFloat64(p, x)`，前置条件是调用 Search 之前，slice 必须已经按照 Float64Slice
的规则排好序

```
func (p Float64Slice) Search(x float64) int
```

- 搜索普通数字 -> 跳过 NaN 区
- 搜索 NaN -> 返回 NaN 区的起始位置