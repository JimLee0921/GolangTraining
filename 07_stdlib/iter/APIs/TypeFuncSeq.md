# `iter.Seq`

`iter.Seq` 是 iter 包定义的一种迭代器定义格式，用于表示一个可以被遍历的序列的产生方式，而不是序列本身

> Seq 定义的迭代器遍历的序列是单值序列

## 定义

```
type Seq[V any] func(yield func(V) bool)
```

- `type Seq`：定义一个名为 Seq 的类型
- `[V any]`：`Seq` 是一个泛型类型，类型参数是 V，约束是 any（任何类型都可以）
- `func(yield func(V) bool)`：Seq 的底层就是一个函数类型，这个函数接收一个参数 yield，yield 本身也是一个函数，签名是
  `func(V) bool`

这里的 yield 也就是迭代器调用的回调函数，也就是消费者提供的函数，Seq 类型迭代器会不断调用 yield 回调函数

- 调用：`yield(V)` 也就是把每一个元素交给消费者
- 如果 yield 返回 true 就给 yield 下一个值，如果 yield 返回 false 表示提前迭代结束，保证了消费停止权在消费者手中

