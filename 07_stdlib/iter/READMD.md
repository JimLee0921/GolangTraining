# iter 包

`iter` 是 Go 提供的迭代器抽象，用于用一种统一、可组合、可中断的方式按顺序进行产出数据，而不需要一次性生成全部结果。
本质上就是用函数+回调来标准化逐个产生值的模式，是一种惰性调用。

## 迭代器

迭代器本身就是一个函数，用于把序列中的下一个元素，一个一个地传给一个回调函数（约定俗成叫做 yield），迭代器函数要么在序列已经结束的时候停止执行，
要么在 yield 回调函数返回 false （提前停止迭代）的时候停止运行，本质上就是可以被 range 的函数

### 迭代器函数定义

iter 包定义了 Seq 和 Seq2 两种标准格式定义为迭代器的简写形式，分别将每个序列元素的一个或两个值传递给 yield

```
type (
	Seq[V any]     func(yield func(V) bool)     
	Seq2[K, V any] func(yield func(K, V) bool)
)
```

- Seq 为单值迭代协议：`for _, v := range iter`
- Seq2 为双值迭代协议：`for k, v := range iter`

> 详情见 [FuncSeq.md](APIs/TypeFuncSeq.md) 和 [FuncSeq2.md](APIs/TypeFuncSeq2.md)

### 实现方式

如果迭代器应该继续处理序列中的下一个元素，yield1 应该返回 true，如果需要停止则返回 false 提前终止迭代，
例如 `maps.Keys` 返回一个迭代器，会产生 map m 的键序列，实现如下：

```
func Keys[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range m {
			if !yield(k) {
				return
			}
		}
	}
}
```

- `Map ~map[K]V` 表示泛型，接收任意底层类型是 `map[K]V` 的类型
- 返回值为 `iter.Seq[K]`，不返回 `[]K` 而是返回遍历 key 的方式
- `return func(yield func(K) bool)`：展开就等价于 `Seq[K]` 约等于 `func(yield func(K) bool)`
    - keys 做的事情不是遍历 map，而是返回一个将来如何遍历 map 的函数

真正的遍历逻辑为：

```
for k := range m {
    if !yield()k {
        return
    }
}
```

1. 内部仍然使用 Go 原生 range map
2. 每拿到一个 k：调用 `yield(k)`
3. 如果 yield 返回 false，立即 return，停止遍历 map

在自定义 `Seq/Seq2` 迭代器时，必须有：

```
if !yield(v){
    return
}

if !yield(k, v){
    return
}
```

也就是在每次消费者调用 yield 后检查返回值，并在返回 false 时立刻停止 return，在消费者消费阶段调用了
`break/return/Pull().Stop()`本质上都是在让 yield 返回 false

### 调用方式

迭代器通常是通过 range 循环来调用的，例如：

```
func PrintAll[V any](seq iter.Seq[V]) {
    for v := range seq {
        fmt.Println(v)
    }
}
```

- `range seq` 不是在遍历容器，而是在调用 seq 并把循环体变成一个 yield 函数，等价概念如下（只是概念等价）：
    ```
    seq(func(v V) bool {
        fmt.Println(v)
        return true
    })
    ```
- PrintAll 不关心数据的来源，只关心能不能给个 V，也就是把算法与数据源解耦

### 迭代器特点

1. 把便利从数据结构中进行解耦
2. 支持惰性和提前终止
3. 是否继续迭代的控制权在于消费者，通过 yield 回调函数进行控制

## 命名规则

约定俗成的迭代器函数和非法应该以它所遍历的序列来命名，名字描述在遍历什么，而不是如何遍历

在集合类型上，迭代器方法通常命名为 All，因为它遍历的是集合中所有的值所组成的序列。

```
func (s *Set[V]) All() iter.Seq[V]
```

- All 返回一个遍历 s 中所有元素的迭代器
- 返回的是 Seq 而不是 slice（惰性遍历而不是一次性取值）

对于一个包含可能多种序列的类型，迭代器的名字可以用来知名返回的是哪一种序列

```
func (c *Country) Cities() iter.Seq[*City]
func (c *Country) Languages() iter.Seq[string]
```

- Cities 返回一个遍历该国家主要城市的迭代器
- Languages 返回一个遍历该国家官方语言的迭代器
- 同一个 Country 可以遍历 cities / languages / people / ....，每个 iterator 对应一个逻辑序列

如果一个迭代器需要额外的配置，构造函数可以接收额外的配置参数，参数属于构造 iterator 的过程，而不是遍历过程

```
func (m *Map[K, V]) Scan(min, max K) iter.Seq2[K, V]
```

- Scan 返回一个遍历键值对的迭代器，其中键满足 `min <= key <= max`
- `min / max` 决定遍历范围，iterator 本身仍然是惰性的，seq2 表示 key-value

```
func Split(s, seq string) iter.Seq[string]
```

`strings.SplitSeq` 函数是一个非常典型的 iter 的使用场景，返回一个迭代器，遍历由 seq 分割的子串，用于流式返回字串

当存在多种可能的遍历顺序时，方法命也可以指明具体的遍历顺序，这在 树/图/双向链表 中非常重要

```
// All returns an iterator over the list from head to tail.
func (l *List[V]) All() iter.Seq[V]
// Backward returns an iterator over the list from tail to head.
func (l *List[V]) Backward() iter.Seq[V]
```

All 表示链表的默认方向，Backward 表示反向遍历链表

```
// Preorder returns an iterator over all nodes of the syntax tree
// beneath (and including) the specified root, in depth-first preorder,
// visiting a parent node before its children.
func Preorder(root node) iter.Seq[Node]
```

这是树的一个复杂遍历规则：深度优先，前序，访问父节点在前，方法名+注释定义了遍历语义

### 一次性迭代器

不是所有的 `Seq / Seq2` 都可以反复遍历，迭代器是有生命周期的，可以区分为两类：

- 一次性迭代器（single-use iterators）
- 可重复使用迭代器（reusable iterators）

大多数迭代器都支持遍历整个序列（Seq/Seq2的正常期望行为）：

- 当被调用时，迭代器会执行开始遍历所需要的初始化操作
- 然后对序列中的连续元素调用 yield 回调函数
- 在返回之前完成清理工作
- 再次调用该迭代器会重新遍历一遍序列

但是有些迭代器只允许对序列进行一次遍历，因为有些 iterator 的数据源天生不可回放，这些一次性迭代器通常来自于无法回退到起点的数据源

- `io.Reader`
- 网络 socket
- HTTP response body
- 实时日志流
- Kafka / MQ consumer
- 数据库 cursor（非缓存）

这些源的共同点是读取过后就没了，如果在提前停止后再次调用该迭代器，可能会继续读取六种的后续内容，单如果在序列已经结束后再次调用，将不会再产生任何值

- 提前停止就是提前在 range 中调用 break 使得 yield 回调函数返回 false
- 自然读完是读到 EOF

> 对于返回一次性迭代器的函数或方法在文档注释中应该特殊说明

## 标准库使用

在 Go 标准库中已经有很多公共 API 设计层开始适配 iterator 迭代器方法，标准库中有几个包比如 maps、slices、strings 提供了基于迭代器的
API，比如：

- `maps.Keys` 返回一个遍历 map 键的迭代器
- `slices.Sorted` 会把一个迭代器中的值收集成一个 slice
- `strings.SplitSeq` 等方法返回的也是一个迭代器