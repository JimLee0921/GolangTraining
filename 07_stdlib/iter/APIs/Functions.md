# iter包函数/方法

一共有 `Pull` 和 `Pull2` 两个方法 分别用于 `Seq` 和 `Seq2`，主要用于把推模型（yield 回调）的 iterator，转换成拉模型（`next()`
）的 iterator，也就是当不想或不能使用 range 时，仍然可以逐个取值并安全停止迭代。

## `iter.Pull`

传入一个 `Seq[V]`，输出：

- `next()`：返回下一个值 V 已经是否还有值
- `stop()`：显性通知迭代器停止迭代

```
func Pull[V any](seq Seq[V]) (
    next func() (V, bool),
    stop func(),
)
```

## `iter.Pull2`

输入一个 `Seq2[K, V]`，输出：

- `next()`：返回下一个 (K, V)，以及是否还有值
- `stop()`：停止迭代

```
func Pull2[K, V any](seq Seq2[K, V]) (
    next func() (K, V, bool),
    stop func(),
)
```

## 存在意义

在某些场景下比如：

- 需要在多个 iterator 之间交替取值
- 必须把 iterator 嵌入到状态机/select/循环重
- 需要取一个处理一个再决定要不要继续时

这时使用 range 就并不适合，需要使用 `Pull/Pull2` 拉取模型

`Pull/Pull2` 的返回值 `next` 和 `stop` 可以更细粒度的控制迭代器

- `next()` 的返回值是 `v, ok := next()`

    - `ok == true`：v 是一个有效的值
    - `ok == false`：iterator 已经结束，后续再调用 `next()`，仍然返回 `ok == false`
- `stop` 用于提前终止信号，等价于在 range 中调用 break，但是在 Pull 模式下必须显示调用
    - 如果没有把 iterator 消耗到结束必须调用 `stop()`
    - 否则一次性使用迭代器可能造成资源泄露，底层 defer 可能永远不执行

注意 `stop()` 不是 break 的替代品，不能用于控制循环语句，唯一职责是在没有自然耗尽 iterator 时，显式通知上游停止，从而触发清理资源

## 使用模板

```
next, stop := iter.Pull(seq)
defer stop()

for {
	v, ok := next()
	if !ok {
		break
	}
	// 使用 v
}

```