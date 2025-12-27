# slices 函数

slices 作为工具包，主要都是提供对于 slice 通用操作函数，主要分为下面几类

## 极值

不改动 slice，就是获取最大值最小值

- `Min(x)`
- `Max(x)`
- `MinFunc(x, cmp)`
- `MaxFunc(x, cmp)`

### `slices.Min/Max`

Min/Max 用于在一个非空的 slice 中选出最大值和最小值，用于元素本身就有字然大小关系的 slice

```
func Max[S ~[]E, E cmp.Ordered](x S) E
func Min[S ~[]E, E cmp.Ordered](x S) E
```

**参数**

- `x S`：任意 slice 类型，只要底层是 `[]E`
- `E cmp.Ordered`：必须是可排序的内建类型

**返回值**

- 返回 slice 中的一个元素（类型为 `E`） ，值拷贝，不是引用

**注意事项**

1. 空 slice 会引发 panic
2. 返回的是元素值而不是新的对象，修改不会影响原切片
3. 浮点数的 NaN 没法被处理，需要在业务层先进性过滤或使用 `MaxFunc` / `MinFunc` 自定义规则

### `slices.MinFunc/MaxFunc`

`MinFunc` 和 `MaxFunc` 可以自定义比较规则在 slice 中选出最小/最大元素，用于元素没有自然顺序，或想用业务顺序

```
func MaxFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E
func MinFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E
```

**参数**

- `x S`：任意 slice 类型，不要求元素可比较
- `cmp`：比较函数约定（与 `SortFunc` / `CompareFunc` / `BinarySearchFunc` 完全一致）

**返回值**

- 返回 slice 中的某个元素（值拷贝），是按 cmp 意义下的极值

**注意事项**

1. 空 slice 会引发 panic
2. 返回的是值而不是下标位置

## 搜索/包含/谓词查找

- `Contains(s, v) bool`
- `ContainsFunc(s, f) bool`
- `Index(s, v) int`
- `IndexFunc(s, f) int`

> 不修改 slice，`Contains/Index` 要求 `E comparable` 元素可比较，`*Func` 用回调规避约束

### `slices.Contains`

判断 slice 中是否存在等于 v 的元素

```
func Contains[S ~[]E, E comparable](s S, v E) bool
```

**参数**

- `s S`：任意 slice，底层元素类型为 `E`
- `v E`：要查找的目标值，必须为 `E comparable`可比较类型

**返回值**

- bool：找到返回 true，没找到返回 false

**注意事项**

- 不排序
- 不建索引
- 不适合大规模高频查找

### `slices.ContainsFunc`

自定义包含规则，是否存在至少一个元素使 `f(element) == true`

```
func CompareFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, cmp func(E1, E2) int) int
```

**参数**

- `s S`：任意 slice，元素类型不限
- `f func(E) bool`
    - 对 slice 中每个元素执行
    - 返回 true 表示命中

**返回值**

- bool：找到返回 true，没找到返回 false

### `slices.Index`

返回 slice 中 第一个等于 v 的元素下标

```
func Index[S ~[]E, E comparable](s S, v E) int
```

**参数**

- `s S`：任意 slice，元素类型不限
- `f func(E) bool`
    - 对 slice 中每个元素执行
    - 返回 true 表示命中

**返回值**

- `int`：如果找到就是第一个匹配元素的下标，没找到返回 `-1`

**注意事项**

- 不排序
- 不建索引
- 不适合大规模高频查找

### `slices.IndexFunc`

自定义匹配规则，返回第一个满足条件的元素下标

```
func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int
```

**参数**

- `s S`：任意 slice，底层元素类型为 `E`
- `v E`：要查找的目标值，必须为 `E comparable`可比较类型

**返回值**

- `int`：如果找到就是第一个匹配元素的下标，没找到返回 `-1`

## 等价判断与字典序比较

不修改 slice，`Equal/Compare` 用 `comparable / cmp.Ordered`，`*Func` 用自定义比较

1. 相等
    - `Equal(s1, s2) bool`
    - `EqualFunc(s1, s2, eq) bool`

2. 字典序比较（lexicographical）
    - `Compare(s1, s2) int`（返回 -1/0/1）
    - `CompareFunc(s1, s2, cmp) int`

### `slices.Equal`

判断两个 slice 是否完全相同：长度相同且对应位置元素相等（按 `==`）

```
func Equal[S ~[]E, E comparable](s1, s2 S) bool
```

**参数**

- `s1, s2 S`：要比较的两个 slice，任意 slice 类型，只要底层是 `[]E`
- `E comparable`：要求元素必须是可比较的

**返回值**

- bool：
    - true：长度相同 + 所有对应元素相等
    - false：任一条件不满足

**注意事项**

1. 顺序敏感，是有序比较，长度不同直接返回 false，遇到第一个不相同的元素时直接返回 false 不再进行比较
2. nil 和空 slice 默认是相同的，因为两者长度都是0且没有任何元素不相等，这一点与 `reflect.DeepEqual` 不同

### `slices.EqualFunc`

Equal 的泛化版本，用自定义的 eq 函数，定义什么叫相等，再逐元素比较两个 slice

```
func EqualFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, eq func(E1, E2) bool) bool
```

**参数**

- `s1, s2 S`：要比较的两个 slice，任意 slice 类型，元素类型也可以不同
- `eq`：自定义比较函数，返回 true 认为两个元素相等，每个位置都会调用一次

**返回值**

- bool：
    - true：长度相同 + 所有对应元素相等
    - false：任一条件不满足

**注意事项**

1. 顺序敏感，是有序比较
2. 不要在 eq 中操作 slice

### `slices.Compare`

按字典序（lexicographical order）比较两个 slice 的大小关系，内部使用的 `cmp.Compare` 函数

```
func Compare[S ~[]E, E cmp.Ordered](s1, s2 S) int
```

**参数**

- `s1, s2 S`：要比较的两个 slice，任意 slice 类型，只要底层是 `[]E`
- `E cmp.Ordered`：元素本身必须有自然大小关系，支持 `< <= >= >`

**返回值**

- int：不保证返回值一定是 -1 / 0 / 1，只保证正负号语义。
    - `< 0`：`s1 < s2 `
    - `= 0`：`s1 == s2 `
    - `> 0`：`s1 > s2`

**注意事项**

1. 顺序敏感，是有序比较
2. 浮点 NaN 问题需要自行处理

### `slices.CompareFunc`

按你提供的比较规则，比较两个 slice 的字典序大小，做的事与 Compare 相同

- 元素不要求有自然顺序
- 比较规则由 cmp 决定

```
func CompareFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, cmp func(E1, E2) int) int
```

**参数**

- `s1, s2`：两个 slice，元素类型可以不同
- `cmp`：比较约定函数，这个约定必须稳定，一致
    - `< 0`：a 小于 b
    - `= 0`：a 等于 b
    - `> 0`：a 大于 b

**返回值**

- int：含义与 Compare 完全一致

**注意事项**

- cmp 必须定义全序
- 仍然是顺序比较
- 不要在 cmp 中操作 slice

## 拼接/重复/分块

结构性构造，`Concat/Repeat` 会分配新底层数组；Chunk 本质是分段视图（通常每块引用原底层数组），所以原 slice 改动可能影响块内容

- `Concat(slices ...S) S`：拼接多个 slice（返回新 slice）
- `Repeat(x, count) S`：重复元素序列 count 次（返回新 slice）

### `slices.Concat`

把多个 slice 按顺序拼接成一个新的 slice，等价于手写的预估长度 + 逐个 append，但由标准库统一实现，语义清晰

```
func Concat[S ~[]E, E any](slices ...S) S
```

**参数**

- `slices ...S`：可变参数：可以传 0 个、1 个或多个 slice，每个参数都是同一种 slice 类型 S（底层元素类型一致）

**返回值**

- `S`：返回拼接后的新 slice，结果长度为所有输入 slice 长度之和

**注意事项**

1. 返回值为新 slice 需要接收
2. 拿到的新结果对老的 slice 不会产生影响

### `slices.Repeat`

把一个 slice 的内容重复 count 次，构造出一个新的 slice

```
func Repeat[S ~[]E, E any](x S, count int) S
```

**参数**

- `x S`：被重复的原始 slice，可以是 nil 或空 slice
- `count int`：重复次数，必须是非负整数

**返回值**

- `S`：返回的一个新的 slice，长度是 `len(x) * count`

**注意事项**

1. `count<0`会导致 panic
2. 结果长度可能溢出或导致巨量分配
3. 返回的是新 slice，不会影响原 slice，必须接收

## 删除/插入/替换

这些逻辑上都会改动结果 slice 的内容与长度，并且在容量不足时会触发重新分配。即便不重新分配，也会在原底层数组上做搬移/覆盖（因此要注意别名引用问题）

- `Delete(s, i, j) S`：删区间 [i, j)
- `DeleteFunc(s, del) S`：按条件删除
- `Insert(s, i, v...E) S`：在位置 i 插入
- `Replace(s, i, j, v...E) S`：把 [i, j) 替换为 v
- `Reverse[S ~[]E, E any](s S)`：反转

### `slices.Delete`

删除 slice 中索引区间 `[i, j)` 的元素，删除后剩余元素向前移动，返回一个新的修改后的切片，长度变短（会影响原切片）

```
func Delete[S ~[]E, E any](s S, i, j int) S
```

**参数**

- `s S`：任意 slice
- `i, j int`：删除区间是 `[i, j)`，要求 `0 <= i <= j <= len(s)`，否则直接 panic

**返回值**

- `S`：返回删除后的 slice，可能与原 slice 共享底层数组

**注意事项**

1. 会修改原底层数组，Delete 本质是 copy，不能假设原 slice 保持原样
2. 旧的 slice 尾部元素可能已改变，真正使用场景是直接用 slice 进行接口

### `slices.DeleteFunc`

删除 slice 中所有满足条件的元素并返回新切片（需要接收），删除规则由提供的 del 函数决定

- `del(e) == true` -> 该元素会被删除
- `del(e) == false` -> 该元素会被保留

```
func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S
```

**参数**

- `s S`：任意 slice，元素类型不限
- `del func(E) bool`：删除判断函数
    - 对 slice 中每一个元素都会调用一次
    - 对于 true 表示删掉这个元素

**返回值**

- `S`：返回删除后的 slice，长度可能变短，通常与原 slice 共享底层数组

**注意事项**

1. 会原地覆盖底层数组，真正使用会直接使用原 slice 进行接收

### `slices.Insert`

在 slice 的索引 i 处插入若干元素，原本索引 `>= i` 的元素整体后移，slice 长度增加

**参数**

- `s S`：原 slice
- `i int`：插入位置，要求 `0 ≤ i ≤ len(s)`，否则 panic
- `v ...E`：要插入的元素，可以是 0 个或多个

**返回值**

- `S`：返回插入后的 slice（可能重新分配底层数组）

**注意事项**

1. 原 slice 容量不足可能重新分配，返回的新 slice 和老的 slice 不再共享底层数组
2. 可以插入 0 个元素，等价于原 slice
3. 插入位置如果是 `i == len(s)` 等价于 append

### `slices.Replace`

用新元素 v，替换 slice 中区间 `[i, j)`，等价于先删除，再插入，但只做一次移动

```
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S
```

**参数**

- `s S`：原 slice
- `i, j int`：替换区间 `[i, j)`，要求 `0 <= i <= j <= len(s)`，否则 panic
- `v ...E`：替换用的新元素，可以 少于/等于/多与 `j-i`

**返回值**

- `S`：返回替换后的 slice，长度变化取决于 `len(v)` 与 `(j-i)` 的差

**注意事项**

1. 可能修改底层数组，需要重新接收
2. 最搞笑的区间更新方式，只做一次搬移，更少分配

### `slices.Reverse`

就地反转 slice 中元素的顺序，这是一个纯原地操作：

- 不返回新 slice
- 不改变长度
- 只改变顺序

> 和 Delete / Insert / Replace 不同

```
func Reverse[S ~[]E, E any](s S)
```

**参数**

- `s S`：任意 slice，元素类型不限

**返回值**

无返回值，直接操作原 slice

**注意事项**

1. 直接修改原 slice
2. nil / 空 slice 是安全的

## 去重/压缩

进行去重处理，只处理相邻重复，通常要先排序或保证相同元素聚在一起，会就地覆盖并返回缩短后的 slice

- `Compact(s) S`：移除相邻重复（要求 E comparable）
- `CompactFunc(s, eq) S`：自定义相等

### `slices.Compact`

删除 slice 中相邻的重复元素，只保留每一段连续重复中的第一个，不是全局去重，而是压缩连续重复

```
func Compact[S ~[]E, E comparable](s S) S
```

**参数**

- `s S`：任意 slice，元素类型是 `E`
- `E comparable`：元素必须是可比较类型

**返回值**

- `S`：返回压缩后的 slice，长度可能变短，通常与原 slice 共享底层数组

**注意事项**

1. 会原地覆盖底层数组，返回的 slice 才是可信结果
2. 只处理相邻的重复元素
3. 空 / nil slice 是安全的

### `slices.CompactFunc`

按自定义的相等规则，删除 slice 中相邻的等价元素

```
func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S
```

**参数**

- `s S`：任意 slice，元素类型不限
- `eq`：判断两个元素是否等价，返回 true 会把后一个元素删除

**返回值**

- `S`：返回压缩后的 slice，长度可能变短，同样是原地压缩语义

**注意事项**

1. eq 只会比较相邻元素
2. eq 应是稳定的等价关系，最好满足自反/对称/传递
3. 不要在 eq 中修改 slice

**返回值**

- `S`：返回压缩后的 slice，长度可能变短，通常与原 slice 共享底层数组

**注意事项**

1. 会原地覆盖底层数组，返回的 slice 才是可信结果
2. 只处理相邻的重复元素
3. 空 / nil slice 是安全的

## 排序与有序性检查

`Sort*` 会原地重排（in-place），`IsSorted*` 只是用于判断。`Sort/IsSorted` 需要 `cmp.Ordered`，否则用 `*Func`

1. 排序
    - `Sort`
    - `SortFunc`
    - `SortStableFunc`（稳定）
2. 是否已排序
    - `IsSorted`
    - `IsSortedFunc`
3. `iter.Seq*`相关
    - `Sorted`
    - `SortedFunc`
    - `SortedStableFunc`

### `slices.Sort`

按元素的自然升序，对 slice 进行原地排序，最直接，最常用的排序函数

```
func Sort[S ~[]E, E cmp.Ordered](x S)
```

**参数**

- `x S`：任意 slice，元素类型必须是 `cmp.Ordered`

**返回值**

无返回值，排序结果体现在 `x` 本身

**注意事项**

1. 一定会修改原 slice
2. 不是稳定排序
3. 空 slice / nil slice 安全
4. 对浮点数进行排序时，NaN 值会排在其他值之前

### `slices.SortFunc`

自定义提供比较函数，对 slice 进行原地排序

```
func SortFunc[S ~[]E, E any](x S, cmp func(a, b E) int)
```

**参数**

- `x S`：待排序 slice，元素可以是任意类型
- `cmp`：自定义比较函数
    - `a < b`：cmp 应返回正数
    - `a == b`：cmp 应返回 0
    - `a > b`：cmp 应返回负数

**返回值**

无返回值，直接修改 x

**注意事项**

- cmp 必须定义全序
- 会修改原 slice
- 非稳定排序

### `slices.SortStableFunc`

按自定义的比较函数，对 slice 进行稳定排序，当 `cmp(a, b) == 0`，a 与 b 的相对位置保持不变

```
func SortStableFunc[S ~[]E, E any](x S, cmp func(a, b E) int)
```

> 参数和返回值与 SortFunc 一致，只要区别在于是稳定排序， a == b 不会修改其位置，可能比 SortFunc 稍慢

### `slices.IsSorted`

判断 slice 是否已经按自然升序排列

```
func IsSorted[S ~[]E, E cmp.Ordered](x S) bool
```

**参数**

- `x S`：被检查的 slice，元素必须是 `cmp.Ordered`

**返回值**

- bool：返回 true 表示已升序排列，false 表示未升序

### `slices.IsSortedFunc`

按自定义的比较规则，判断 slice 是否已经按照升序排序

```
func IsSortedFunc[S ~[]E, E any](x S, cmp func(a, b E) int) bool
```

**参数**

- `x S`：被检查 slice，元素可以是任意类型
- `cmp`：排序规则（必须和排序时一致）

**返回值**

和 IsSorted 返回值一致

> cmp 必须定义全序

### `slices.Sorted`

从一个序列 `iter.Seq` 中收集元素，排序后，返回一个新的 slice

```
func SortedFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E
```

**参数**

- `seq iter.Seq[E]`：输入序列，元素必须是 `cmp.Ordered`

**返回值**

- `[]E`：新 slice，原序列与任何原 slice 不受影响

**注意事项**

- 会完整消费序列
- 返回新的 slice 必须接收，不会修改任何原 slice

### `slices.SortedFunc`

按自定义比较规则，对序列排序并生成新 slice

```
func SortedFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E
```

**参数**

- `seq iter.Seq[E]`：输入序列，元素可以是任意类型
- `cmp`：自定义比较规则

**返回值**

- `[]E`：新 slice，原序列与任何原 slice 不受影响

**注意事项**

- cmp 必须是全序，非稳定排序
- 会完整消费序列
- 返回新的 slice 必须接收，不会修改任何原 slice

### `slices.SortedStableFunc`

对序列进行稳定排序，并返回新 slice

```
func SortedStableFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E
```

> 参数和返回值参考 SortedFunc，区别在于稳定排序

## 有序二分查找

必须切片已经有序，不改变 slice，返回的 int 在未命中时通常是可插入位置（保持有序的下标）

- `BinarySearch(x, target) (int, bool)`：x 必须按升序（与 cmp.Ordered 的自然序一致）
- `BinarySearchFunc(x, target, cmp) (int, bool)`：自定义比较，适用于结构体/不同 key

### `slices.BinarySearch`

在一个已经按升序排列的 slice 中，用二分查找定位目标元素

```
func BinarySearch[S ~[]E, E cmp.Ordered](x S, target E) (int, bool)
```

**参数**

- `x S`：被查找的 slice，必须已经按升序排序，排序规则必须与 `cmp.Ordered` 的自然顺序一致
- `target E`：要查找的目标值

**返回值**

- int：第一个返回值 idx：
    - 若找到：目标元素所在的下标
    - 若没找到：目标应该插入的位置
- found：找到返回 true，没找到返回 false

**注意事项**

1. slice 必须有序
2. 排序规则必须一致
3. 有重复元素是不保证返回的是哪一个

### `slices.BinarySearchFunc`

在按自定义规则排序的 slice 中，用自定义的比较函数做二分查找，主要为了解决 slice 元素不是 `cmp.Ordered` 或需要进行自定义业务规则

```
func BinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool)
```

**参数**

- `x S`：被查找的 slice，必须已经按升序排序，排序规则必须与 `cmp.Ordered` 的自然顺序一致
- `target T`：要查找的目标值，类型可以与元素不同
- `cmp`：比较约定
    - `cmp(e, target) < 0`：`e < target`
    - `cmp(e, target) = 0`：`e = target`
    - `cmp(e, target) > 0`：`e > target`

**返回值**

`int / bool` 与 BinarySearch 语义一样

**注意事项**

1. slice 必须有序
2. cmp 必须和排序时使用的规则一致
3. 有重复元素是不保证返回的是哪一个

## 迭代器/序列适配

这类函数把 `[]E` 与迭代器序列(`iter.Seq`/`iter.Seq2`)互相转换加工

1. 从 slice 产生序列
    - `All(s) iter.Seq2[int, E]`：产生 (index, value) 序列，正向
    - `Backward(s) iter.Seq2[int, E]`：产生 (index, value) 序列，反向
    - `Values(s) iter.Seq[E]`：只产生 value 序列
    - `Chunk(s, n) iter.Seq[Slice]`：把 slice 分成块的序列（每块是 Slice）

2. 从序列产生/追加 slice
    - `Collect(seq) []E`：把序列收集成 `[]E`
    - `AppendSeq(s, seq) Slice`：把序列追加到 slice 后面（等价于循环 append）

> `Collect/AppendSeq` 一定会分配/扩容（取决于容量），`All/Backward/Values` 不改原 slice，只是视图式遍历

### `slices.All`

把一个 slice 转换为一个正向遍历的二元序列 `inter.Seq2` (index, value)，迭代器序列形式

```
func All[Slice ~[]E, E any](s Slice) iter.Seq2[int, E]
```

**参数**

- `s Slice`：任意 slice 元素类型不限，不会被修改

**返回值**

- `iter.Seq2[int, E]`：
    - 每一步产出 `(index, value)`
    - index 从 0 递增到 `len(s) - 1`
    - value 是 `s[index]`

> 不分配内存，不复制元素，是惰性便利

### `slices.Backward`

把 slice 转成反向遍历的 `iter.Seq2` (index, value) 迭代器序列，遍历顺序是从 `len(s)-1` 一直到 0

```
func Backward[Slice ~[]E, E any](s Slice) iter.Seq2[int, E]
```

**参数**

- `s Slice`：任意 slice 元素类型不限

**返回值**

- `iter.Seq2[int, E]`：
    - 每一步产出 `(index, value)`
    - index 是原 slice 的真实索引

**注意事项**

1. 返回的 index 不是从 0 开始递增（方便直接回写到 slice）
2. 不创建新的 slice，不复制

### `slices.Values`

把 slice 转成只遍历值的序列，等价于 `for _, v := range s` 的 `iter.Seq` 迭代器版本

```
func Values[Slice ~[]E, E any](s Slice) iter.Seq[E]
```

**参数**

- `s Slice`：任意 slice，元素类型不限

**返回值**

- `iter.Seq[E]`：每一步只产出 value，不暴露索引

> 最轻量的 slice -> seq 方式，适合：map/filter/reduce 风格，不关心索引

### `slices.Chunk`

把一个 slice 按固定大小 n，分成若干块，并以序列形式逐块产出，每一块都是一个 Slice（子 slice）

```
func Chunk[Slice ~[]E, E any](s Slice, n int) iter.Seq[Slice]
```

**参数**

- `s Slice`：原始 slice
- `n int`：每一块的大小，必须 `> 0`，否则会 panic

**返回值**

- `iter.Seq[Slice]`：每一步都产出一个子 slice，最后一块的长度可能 `<n`

**注意事项**

1. 每个 chunk 可能共享底层数组，工程上应把 chunk 当作视图
2. 是惰性的，不会一次性生成所有块，适合大的 slice 的批处理

### `slices.Collect`

把一个迭代器序列 iter.Seq[E] 的所有元素，一次性收集成一个新的切片 `[]E`

```
func Collect[E any](seq iter.Seq[E]) []E
```

**参数**

- `seq iter.Seq[E]`：任意元素类型的序列，可以来自：
    - `slices.Values`
    - `slices.All（取 value）`
    - `slices.Chunk`（元素是子 slice）
    - `Sorted*` 之前的任意序列管道

**返回值**

- `[]E`：返回一个全新的 slice，元素顺序与 seq 产生顺序一致，长度等于序列中元素总数

**注意事项**

1. 回信分配 slice，需要被接收
2. 序列是惰性的，但是 Collect 是贪婪的，会完整消费整个序列

### `slices.AppendSeq`

把一个序列 seq 产生的元素，逐个 append 到已有的 `slice s` 后面

```
func AppendSeq[Slice ~[]E, E any](s Slice, seq iter.Seq[E]) Slice
```

**参数**

- `s Slice`：已存在的 slice，作为 append 的起点
- `seq iter.Seq[E]`：要追加的序列，逐个产出 E

**返回值**

- `Slice`：返回 append 完成后的 slice 必须接收返回值

**注意事项**

1. 容量不够会自动分配，可能重新分配底层数组，返回的新的 slice 与原 slice 不再共享底层数组
2. 返回值一定要显式接收
3. seq 会被完整消费，不能重复使用

## 容量/拷贝/内存形状控制

更偏性能与内存控制，做后端数据管道时非常实用，偏工程化

- `Clone(s) S`：拷贝一份（新底层数组）
- `Grow(s, n) S`：确保还能再 append n 个元素而不频繁扩容（可能重新分配）
- `Clip(s) S`：把 cap 裁到 len（常用于避免意外持有大数组）

### `slices.Clone`

复制一个 slice，生成一份内容相同、但底层数组完全独立的新 slice，解决的问题是：

- slice 别名（alias）
- 原 slice 被修改影响到其他地方

```
func Clone[S ~[]E, E any](s S) S
```

**参数**

- `s S`：任意 slice，可以是 nil / 空 slice

**返回值**

- `S`：新的 slice，len 与 cap 与 s 相同，底层数组一定不同

**注意事项**

1. 这是浅拷贝，slice结构被复制，元素内部引用不会被深拷贝
2. `slices.Clone(nil)` 还为 nil，`slices.Clone([]int{})` 还为 `[]`

### `slices.Grow`

确保 `slice s` 还能至少再 append n 个元素，而不触发扩容，容量保证工具

```
func Grow[S ~[]E, E any](s S, n int) S
```

**参数**

- `s S`：任意 slice
- `n int`：期望未来还可以 append 的元素数量，必须 `n>=0` 否则引发 panic

**返回值**

- `S`：新的 slice，满足 `cap(result) >= len(s) + n`，可能重新分配底层数组

> 是性能工具，不是语义工具

### `slices.Clip`

把 slice 的 capacity 裁剪到和 length 一样大，也就是移除切片中未使用的容量

```
func Clip[S ~[]E, E any](s S) S
```

**参数**

- `s S`：任意 slice

**返回值**

- `S`：返回 slice，元素内容不变，len 不变，cap 变为 len，也就是 ` s[:len(s):len(s)]`，注意结果会保留 s 的nil值

**注意事项**

1. 可能重新分配，如果 `cap(s) > len(s)`，通常会分配新数组，返回的 slice 不再持有原大数组
2. 防止内存泄漏工具，nil / 空 slice 安全
