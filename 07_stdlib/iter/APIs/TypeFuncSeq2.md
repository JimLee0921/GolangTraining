# `iter.Seq2`

`iter.Seq2` 表示每次产生一对值的迭代器，通常用于 key-value 或 index-value 的遍历。

可以对标 slice / map

| Go 原生                     | iter                     |
|---------------------------|--------------------------|
| `for k, v := range map`   | `for k, v := range Seq2` |
| `for i, v := range slice` | `for i, v := range Seq2` |

## 定义

```
type Seq2[K, V any] func(yield func(K, V) bool)
```

- `type Seql2`：和 Seq 一样，定义一个迭代器函数类型
- `[K, V any]`：两个泛型参数
    - K 通常为 key/index/name/id
    - V 通常为真正的数据
- `func(yield func(K, V) bool)`：迭代器每次调用 `yield(k, v)` 来调用每一对值
    - 如果 yield 返回 true 就继续传入下一对 `K/V`
    - 如果 yield 返回 false 就提前结束
